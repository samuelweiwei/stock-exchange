package order

import (
	"encoding/json"
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/errorx"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/coupon"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/symbol"
	"gorm.io/gorm"
	"io"
	"net/http"
	"net/url"
	"time"
)

func validateApply(req *request.UserFollowOrderApplyReq) error {
	if err := validateApplyAdvisorAndProdRelated(req); err != nil {
		return err
	}

	if req.CouponRecordId > 0 {
		if err := validateApplyCouponRelated(req); err != nil {
			return err
		}
	}

	return nil
}

// validateApplyAdvisorAndProdRelated 用户跟单申请，校验导师和导师产品相关规则
func validateApplyAdvisorAndProdRelated(req *request.UserFollowOrderApplyReq) error {
	var advisorProd order.AdvisorProd
	err := global.GVA_DB.First(&advisorProd, req.AdvisorProdId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errorx.NewWithCode(errorx.AdvisorProdNotFound)
	} else if err != nil {
		return err
	} else if advisorProd.ActiveStatus == order.Inactive {
		return errorx.NewWithCode(errorx.AdvisorProdDisabled)
	}
	if req.FollowAmount < advisorProd.MinAmount {
		return errorx.NewWithCode(errorx.FollowAmountShouldGtMin, "minAmount", advisorProd.MinAmount)
	}
	if req.FollowAmount > advisorProd.MaxAmount {
		return errorx.NewWithCode(errorx.FollowAmountShouldLtMax, "maxAmount", advisorProd.MaxAmount)
	}
	return nil
}

// validateApplyCouponRelated 用户跟单申请，校验使用优惠券相关
func validateApplyCouponRelated(req *request.UserFollowOrderApplyReq) error {
	if req.AutoRenew == order.EnableAutoRenew {
		return errorx.NewWithCode(errorx.CouponNotAllowed)
	}

	var couponIssueRecord coupon.CouponIssued
	err := global.GVA_DB.Model(&couponIssueRecord).
		Where("id = ? and user_id = ?", req.CouponRecordId, req.UserId).
		First(&couponIssueRecord).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errorx.NewWithCode(errorx.CouponNotFound)
	} else if err != nil {
		return err
	}

	var advisorProdCoupon order.AdvisorProdCoupon
	err = global.GVA_DB.Model(&advisorProdCoupon).
		Where("advisor_prod_id = ? and coupon_id = ?", req.AdvisorProdId, couponIssueRecord.CouponId).First(&advisorProdCoupon).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errorx.NewWithCode(errorx.CouponNotAllowedForProd)
	} else if err != nil {
		return err
	}

	if couponIssueRecord.Status == coupon.AlreadyUsed {
		return errorx.NewWithCode(errorx.CouponAlreadyUsed)
	}
	currentUnixTime := time.Now().UnixMilli()
	if currentUnixTime < couponIssueRecord.ValidStart {
		return errorx.NewWithCode(errorx.CouponIsExpired)
	}
	if currentUnixTime > couponIssueRecord.ValidEnd {
		return errorx.NewWithCode(errorx.CouponIsExpired)
	}
	if req.FollowAmount <= *couponIssueRecord.CouponAmount {
		return errorx.NewWithCode(errorx.FollowAmountShouldGtCouponAmount)
	}
	return nil
}

// validateMarketStatus 校验当前交易市场状态
func validateMarketStatus(stockId uint) error {
	var stock symbol.Symbols
	err := global.GVA_DB.First(&stock, stockId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errorx.NewWithCode(errorx.SymbolNotFount)
	} else if err != nil {
		return err
	}

	// 加密货币不校验市场状态
	if symbol.StockType(*stock.Type) == symbol.Crypto {
		return nil
	}

	marketStatus, err := getMarketStatus()
	if err != nil {
		return err
	} else if marketStatus.Exchanges.Nasdaq == closed {
		return errorx.NewWithCode(errorx.MarketNotOpened)
	}
	return nil
}

func getMarketStatus() (*MarketStatus, error) {
	api, err := url.Parse(global.GVA_CONFIG.Polygon.BaseURL + polygonMarketStatusApi)
	if err != nil {
		return nil, err
	}
	param := url.Values{}
	param.Set("apiKey", global.GVA_CONFIG.Polygon.APIKey)
	api.RawQuery = param.Encode()
	resp, err := http.Get(api.String())
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, _ := io.ReadAll(resp.Body)
	var marketStatus MarketStatus
	err = json.Unmarshal(body, &marketStatus)
	return &marketStatus, err
}
