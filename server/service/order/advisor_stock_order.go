package order

import (
	"context"
	"database/sql"
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/constants"
	"github.com/flipped-aurora/gin-vue-admin/server/enums/fund"
	"github.com/flipped-aurora/gin-vue-admin/server/errorx"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/symbol"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	. "github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/guregu/null/v5"
	. "github.com/shopspring/decimal"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"time"
)

const (
	polygonMarketStatusApi = "/v1/marketstatus/now"
	closed                 = "closed"
)

type MarketStatus struct {
	Exchanges Exchanges `json:"exchanges"`
}
type Exchanges struct {
	Nasdaq string `json:"nasdaq"`
	Nyse   string `json:"nyse"`
	Otc    string `json:"otc"`
}

type AdvisorStockOrderService struct{}

func (service *AdvisorStockOrderService) CreateStockOrder(req *request.AdvisorStockOrderCreateReq) error {
	err := validateCreateReq(req)
	if err != nil {
		return err
	}

	var stock symbol.Symbols
	err = global.GVA_DB.First(&stock, req.StockId).Error
	if err != nil {
		return err
	}

	advisorStockOrder := &order.AdvisorStockOrder{
		AdvisorId: req.AdvisorId,
		StockId:   req.StockId,
		BuyPrice:  NewFromFloat(req.BuyPrice).RoundFloor(int32(GetScale(*stock.TicketSize))),
		BuyTime:   time.UnixMilli(req.BuyTime),
		Status:    order.StockOrderStatusHolding,
	}

	return global.GVA_DB.Create(advisorStockOrder).Error
}

func (service *AdvisorStockOrderService) PageQuery(req *request.AdvisorStockOrderPageQueryReq) (
	[]*response.AdvisorStockOrderPageData, int64, error) {
	db := global.GVA_DB.Model(&order.AdvisorStockOrder{}).Joins("Advisor")

	if req.AdvisorName != nil && len(*req.AdvisorName) > 0 {
		db = db.Where("Advisor.nick_name = ?", *req.AdvisorName)
	}
	if req.StockOrderStatus != nil {
		db = db.Where("status = ?", *req.StockOrderStatus)
	}
	if req.CreatedAtStart != nil && *req.CreatedAtStart > 0 {
		db = db.Where("advisor_stock_order.created_at >= ?", time.UnixMilli(*req.CreatedAtStart))
	}
	if req.CreatedAtEnd != nil && *req.CreatedAtEnd > 0 {
		db = db.Where("advisor_stock_order.created_at <= ?", time.UnixMilli(*req.CreatedAtEnd))
	}
	if req.UpdatedAtStart != nil && *req.UpdatedAtStart > 0 {
		db = db.Where("advisor_stock_order.updated_at >= ?", time.UnixMilli(*req.UpdatedAtStart))
	}
	if req.UpdatedAtEnd != nil && *req.UpdatedAtEnd > 0 {
		db = db.Where("advisor_stock_order.updated_at <= ?", time.UnixMilli(*req.UpdatedAtEnd))
	}

	var total int64

	if err := db.Count(&total).Error; err != nil || total == 0 {
		return nil, 0, err
	}

	var stockOrders []*order.AdvisorStockOrder
	err := req.Paginate()(db).Preload("Advisor").Preload("Stock").Order("advisor_stock_order.updated_at desc").Find(&stockOrders).Error
	if err != nil {
		return nil, 0, err
	}

	list := make([]*response.AdvisorStockOrderPageData, len(stockOrders))
	for i, item := range stockOrders {
		pageData := &response.AdvisorStockOrderPageData{
			StockOrderId: item.ID,
			StockId:      item.StockId,
			StockName:    item.Stock.Symbol,
			AdvisorName:  item.Advisor.NickName,
			BuyPrice:     item.BuyPrice.InexactFloat64(),
			BuyTime:      item.BuyTime.UnixMilli(),
			Status:       item.Status,
			CreatedAt:    item.CreatedAt.UnixMilli(),
			UpdatedAt:    item.UpdatedAt.UnixMilli(),
		}
		if item.SellPrice.Valid {
			pageData.SellPrice = null.FloatFrom(item.SellPrice.Decimal.InexactFloat64())
		}
		if item.SellTime.Valid {
			pageData.SellTime = null.IntFrom(item.SellTime.Time.UnixMilli())
		}
		list[i] = pageData
	}
	return list, total, nil
}

func (service *AdvisorStockOrderService) GetSellConfirmSummary(req *request.AdvisorStockOrderSellConfirmReq) (
	*response.AdvisorStockOrderConfirmSummary, error) {
	var stockOrder order.AdvisorStockOrder
	err := global.GVA_DB.Preload("Stock").First(&stockOrder, req.StockOrderId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errorx.NewWithCode(errorx.AdvisorStockOrderNotFound)
	} else if err != nil {
		return nil, err
	}

	var userFollowStockDetails []*order.UserFollowStockDetail
	err = global.GVA_DB.Preload("UserFollowOrder").
		Model(&order.UserFollowStockDetail{}).
		Where("stock_order_id = ?", req.StockOrderId).
		Find(&userFollowStockDetails).
		Error

	if len(userFollowStockDetails) == 0 {
		return &response.AdvisorStockOrderConfirmSummary{
			StockId:           stockOrder.StockId,
			StockName:         stockOrder.Stock.Symbol,
			BuyPrice:          stockOrder.BuyPrice.InexactFloat64(),
			SellPrice:         req.SellPrice,
			PercentageChange:  0,
			TotalStockNum:     0,
			TotalActualReturn: 0,
			TotalTradeReturn:  0,
			FollowOrders:      make([]*response.AdvisorStockOrderConfirmDetail, 0),
		}, nil
	}

	summary := &response.AdvisorStockOrderConfirmSummary{
		StockId:      stockOrder.StockId,
		StockName:    stockOrder.Stock.Symbol,
		BuyPrice:     stockOrder.BuyPrice.InexactFloat64(),
		SellPrice:    req.SellPrice,
		FollowOrders: make([]*response.AdvisorStockOrderConfirmDetail, len(userFollowStockDetails)),
	}

	totalStockNum, totalTradeReturn, totalActualReturn := NewFromFloat(0), NewFromFloat(0), NewFromFloat(0)
	for i, v := range userFollowStockDetails {
		v.SellPrice = NullDecimal{Decimal: NewFromFloat(req.SellPrice), Valid: true}
		v.CalculateProfit()
		confirmDetail := &response.AdvisorStockOrderConfirmDetail{
			UserId:   strconv.Itoa(int(v.UserId)),
			StockNum: v.StockNum.InexactFloat64(),
		}
		if v.UserFollowOrder.PlatformCommissionRate.Valid {
			confirmDetail.PlatformCommissionRate = v.UserFollowOrder.PlatformCommissionRate.Decimal.InexactFloat64()
		}
		if v.UserFollowOrder.AdvisorCommissionRate.Valid {
			confirmDetail.AdvisorCommissionRate = v.UserFollowOrder.AdvisorCommissionRate.Decimal.InexactFloat64()
		}
		confirmDetail.TradeReturn = v.TradeReturn.InexactFloat64()
		confirmDetail.ActualReturn = v.ActualReturn.InexactFloat64()
		if v.PlatformCommission.Valid {
			confirmDetail.PlatformCommission = v.PlatformCommission.Decimal.InexactFloat64()
		}
		if v.AdvisorCommission.Valid {
			confirmDetail.AdvisorCommission = v.AdvisorCommission.Decimal.InexactFloat64()
		}

		summary.FollowOrders[i] = confirmDetail

		totalStockNum = totalStockNum.Add(v.StockNum)
		totalTradeReturn = totalTradeReturn.Add(v.TradeReturn)
		totalActualReturn = totalActualReturn.Add(v.ActualReturn)
	}

	summary.TotalActualReturn = totalActualReturn.InexactFloat64()
	summary.TotalTradeReturn = totalTradeReturn.InexactFloat64()
	summary.TotalStockNum = totalStockNum.InexactFloat64()
	return summary, nil
}

func (service *AdvisorStockOrderService) SubmitSellConfirm(req *request.AdvisorStockOrderSellConfirmSubmitReq) error {
	var stockOrder order.AdvisorStockOrder
	err := global.GVA_DB.First(&stockOrder, req.StockOrderId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errorx.NewWithCode(errorx.AdvisorStockOrderNotFound)
	} else if err != nil {
		return err
	}

	if err = validateMarketStatus(stockOrder.StockId); err != nil {
		return err
	}

	if stockOrder.Status == order.StockOrderStatusFinished {
		return errorx.NewWithCode(errorx.AdvisorStockOrderAlreadyFinished)
	}

	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		stockOrder.Status = order.StockOrderStatusFinished
		stockOrder.SellPrice = NullDecimal{Decimal: NewFromFloat(req.SellPrice), Valid: true}
		stockOrder.SellTime = sql.NullTime{Time: time.UnixMilli(req.SellTime), Valid: true}
		txErr := tx.Model(&stockOrder).Select("Status", "SellPrice", "SellTime").Updates(stockOrder).Error
		if txErr != nil {
			return txErr
		}

		var userFollowStockDetails []*order.UserFollowStockDetail
		txErr = tx.Preload("UserFollowOrder").Where("stock_order_id = ?", req.StockOrderId).Find(&userFollowStockDetails).Error
		if txErr != nil {
			return txErr
		}

		if len(userFollowStockDetails) > 0 {
			updateFollowOrderMap := make(map[uint]*order.UserFollowOrder)
			for _, v := range userFollowStockDetails {
				v.SellPrice = stockOrder.SellPrice
				v.SellTime = sql.NullTime{Time: time.UnixMilli(req.SellTime), Valid: true}
				v.CalculateProfit()

				txErr = tx.Model(v).Select("SellPrice", "SellTime", "AdvisorCommission", "PlatformCommission").Updates(*v).Error
				if txErr != nil {
					return txErr
				}

				updateFollowOrder, ok := updateFollowOrderMap[v.FollowOrderId]
				if !ok {
					updateFollowOrder = &order.UserFollowOrder{
						GVA_MODEL: global.GVA_MODEL{
							ID: v.FollowOrderId,
						},
						UserId:            v.UserId,
						FollowOrderStatus: v.UserFollowOrder.FollowOrderStatus,
					}
					if v.UserFollowOrder.StockStatus == order.FollowOrderStockStatusOpened {
						if v.UserFollowOrder.FollowOrderStatus == order.FollowOrderStatusFinished {
							updateFollowOrder.StockStatus = order.FollowOrderStockStatusExpired
						} else {
							updateFollowOrder.StockStatus = order.FollowOrderStockStatusUnOpened
						}
					} else {
						updateFollowOrder.StockStatus = v.UserFollowOrder.StockStatus
					}
					updateFollowOrderMap[v.FollowOrderId] = updateFollowOrder
				}

				if v.ActualReturn.InexactFloat64() > 0 {
					updateFollowOrder.FollowAvailableAmount = updateFollowOrder.FollowAvailableAmount.Add(v.BuyPrice.Mul(v.StockNum)).RoundFloor(2)
					updateFollowOrder.RetrievableAmount = updateFollowOrder.RetrievableAmount.Add(v.ActualReturn)
				} else {
					updateFollowOrder.FollowAvailableAmount = updateFollowOrder.FollowAvailableAmount.Add(v.SellPrice.Decimal.Mul(v.StockNum)).RoundFloor(2)
				}
			}

			for _, v := range updateFollowOrderMap {
				if v.FollowOrderStatus == order.FollowOrderStatusFinished {
					settleAmount := v.FollowAvailableAmount.Add(v.RetrievableAmount).RoundFloor(2).InexactFloat64()
					txErr = userFundService.UpdateUserFundAccountsAndNewFlow(int(v.UserId), fund.AutoSettle, settleAmount, strconv.Itoa(int(v.ID)))
					if txErr != nil {
						return txErr
					}
				} else {
					txErr = tx.Model(v).
						Where("id = ? and user_id = ?", v.ID, v.UserId).
						Updates(map[string]interface{}{
							"follow_available_amount": gorm.Expr("follow_available_amount + ?", v.FollowAvailableAmount),
							"retrievable_amount":      gorm.Expr("retrievable_amount + ?", v.RetrievableAmount),
							"stock_status":            v.StockStatus,
						}).Error
					if txErr != nil {
						return txErr
					}
				}

				if v.RetrievableAmount.InexactFloat64() > 0 {
					shareProfit(v.UserId, v.ID, v.RetrievableAmount)
				}
			}
		}
		return nil
	})
}

func shareProfit(userId uint, followOrderId uint, profit Decimal) {
	var frontUser user.FrontendUsers
	if err := global.GVA_DB.First(&frontUser, userId).Error; err != nil {
		global.GVA_LOG.Error("query front user error", zap.Error(err))
		return
	}
	if frontUser.UserType == 2 {
		return
	}
	if frontUser.ParentId > 0 {
		firstGradeShareRate, _ := global.GVA_REDIS.Get(context.Background(), constants.RedisKeyFirstGradeShareRate).Float64()
		firstGradeShareRateDecimal := NewFromFloat(firstGradeShareRate)
		shareAmount := profit.Mul(NewFromFloat(0.1)).Mul(firstGradeShareRateDecimal).RoundFloor(2)
		if shareAmount.InexactFloat64() > 0 {
			userShareProfit := &order.UserProfitShare{
				FromUserId:    userId,
				ToUserId:      frontUser.ParentId,
				FollowOrderId: followOrderId,
				Amount:        shareAmount,
				ShareRate:     firstGradeShareRateDecimal,
			}
			if err := global.GVA_DB.Create(userShareProfit).Error; err != nil {
				global.GVA_LOG.Error("create userShareProfit error", zap.Error(err))
			}

			err := userFundService.UpdateUserFundAccountsAndNewFlow(int(frontUser.ParentId), fund.ProfitSharing, shareAmount.InexactFloat64(), strconv.Itoa(int(userShareProfit.ID)))
			if err != nil {
				global.GVA_LOG.Error("Update to user account error", zap.Int("userId", int(frontUser.ParentId)), zap.Float64("amount", shareAmount.InexactFloat64()), zap.Error(err))
			}
		}
	}
	if frontUser.GrandparentId > 0 {
		secondGradeShareRate, _ := global.GVA_REDIS.Get(context.Background(), constants.RedisKeySecondGradeShareRate).Float64()
		secondGradeShareRateDecimal := NewFromFloat(secondGradeShareRate)
		shareAmount := profit.Mul(NewFromFloat(0.1)).Mul(secondGradeShareRateDecimal).RoundFloor(2)
		if shareAmount.InexactFloat64() > 0 {
			userShareProfit := &order.UserProfitShare{
				FromUserId:    userId,
				ToUserId:      frontUser.GrandparentId,
				FollowOrderId: followOrderId,
				Amount:        shareAmount,
				ShareRate:     secondGradeShareRateDecimal,
			}
			if err := global.GVA_DB.Create(userShareProfit).Error; err != nil {
				global.GVA_LOG.Error("create userShareProfit error", zap.Error(err))
			}

			err := userFundService.UpdateUserFundAccountsAndNewFlow(int(frontUser.GrandparentId), fund.ProfitSharing, shareAmount.InexactFloat64(), strconv.Itoa(int(userShareProfit.ID)))
			if err != nil {
				global.GVA_LOG.Error("Update to user account error", zap.Int("userId", int(frontUser.GrandparentId)), zap.Float64("amount", shareAmount.InexactFloat64()), zap.Error(err))
			}
		}
	}

	if frontUser.GreatGrandparentId > 0 {
		thirdGradeShareRate, _ := global.GVA_REDIS.Get(context.Background(), constants.RedisKeyThirdGradeShareRate).Float64()
		thirdGradeShareRateDecimal := NewFromFloat(thirdGradeShareRate)
		shareAmount := profit.Mul(NewFromFloat(0.1)).Mul(thirdGradeShareRateDecimal).RoundFloor(2)
		if shareAmount.InexactFloat64() > 0 {
			userShareProfit := &order.UserProfitShare{
				FromUserId:    userId,
				ToUserId:      frontUser.GreatGrandparentId,
				FollowOrderId: followOrderId,
				Amount:        shareAmount,
				ShareRate:     thirdGradeShareRateDecimal,
			}
			if err := global.GVA_DB.Create(userShareProfit).Error; err != nil {
				global.GVA_LOG.Error("create userShareProfit error", zap.Error(err))
			}

			err := userFundService.UpdateUserFundAccountsAndNewFlow(int(frontUser.GreatGrandparentId), fund.ProfitSharing, shareAmount.InexactFloat64(), strconv.Itoa(int(userShareProfit.ID)))
			if err != nil {
				global.GVA_LOG.Error("Update to user account error", zap.Int("userId", int(frontUser.GreatGrandparentId)), zap.Float64("amount", shareAmount.InexactFloat64()), zap.Error(err))
			}
		}
	}
}

func validateCreateReq(req *request.AdvisorStockOrderCreateReq) error {
	err := global.GVA_DB.First(&order.Advisor{}, req.AdvisorId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errorx.NewWithCode(errorx.AdvisorNotFound)
	}

	if err = checkHasHoldingStockOrder(req.AdvisorId); err != nil {
		return err
	}

	if err = validateMarketStatus(req.StockId); err != nil {
		return err
	}

	return nil
}

func (service *AdvisorStockOrderService) AutoFollow(req *request.AdvisorStockOrderAutoFollowReq) error {
	var stockOrder order.AdvisorStockOrder
	err := global.GVA_DB.First(&stockOrder, req.StockOrderId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errorx.NewWithCode(errorx.AdvisorStockOrderNotFound)
	} else if err != nil {
		return err
	}

	if err = validateMarketStatus(stockOrder.StockId); err != nil {
		return err
	}

	var followOrders []order.UserFollowOrder
	err = global.GVA_DB.Joins("AdvisorProd").
		Where("AdvisorProd.advisor_id = ?", stockOrder.AdvisorId).
		Where("user_follow_order.follow_order_status = ? and user_follow_order.stock_status = ?", order.FollowOrderStatusFollowing, order.FollowOrderStockStatusUnOpened).
		Find(&followOrders).Error
	if err != nil {
		return err
	}

	var stock symbol.Symbols
	err = global.GVA_DB.First(&stock, stockOrder.StockId).Error
	if err != nil {
		return err
	}

	return global.GVA_DB.Transaction(func(tx *gorm.DB) (txErr error) {
		for _, followOrder := range followOrders {
			buyAmount := followOrder.FollowAvailableAmount.Mul(NewFromFloat(req.PositionRatio)).Div(HundredPercent)
			maxStockNum := buyAmount.Div(stockOrder.BuyPrice).RoundFloor(int32(GetScale(*stock.TicketNumSize)))
			if maxStockNum.IsPositive() {
				followOrderStockDetail := order.UserFollowStockDetail{
					GVA_MODEL: global.GVA_MODEL{
						ID: uint(global.Snowflake.Generate()),
					},
					UserId:        followOrder.UserId,
					StockOrderId:  stockOrder.ID,
					FollowOrderId: followOrder.ID,
					StockId:       stockOrder.StockId,
					StockNum:      maxStockNum,
					BuyPrice:      stockOrder.BuyPrice,
					BuyTime:       stockOrder.BuyTime,
				}
				txErr = tx.Create(&followOrderStockDetail).Error
				if txErr != nil {
					return txErr
				}

				txErr = tx.Model(&followOrder).Updates(map[string]interface{}{
					"StockStatus":             order.FollowOrderStockStatusOpened,
					"follow_available_amount": gorm.Expr("follow_available_amount - ?", maxStockNum.Mul(stockOrder.BuyPrice).RoundFloor(2)),
				}).Error
				if txErr != nil {
					return txErr
				}
			}
		}
		return nil
	})
}

func checkHasHoldingStockOrder(advisorId uint) error {
	var count int64
	err := global.GVA_DB.Model(&order.AdvisorStockOrder{}).
		Where("advisor_id = ? and status = ? ", advisorId, order.StockOrderStatusHolding).
		Count(&count).Error
	if err != nil {
		return err
	}

	if count > 0 {
		return errorx.NewWithCode(errorx.AdvisorHasOpenedStockOrder)
	}
	return nil
}
