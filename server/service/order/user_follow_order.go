package order

import (
	"context"
	"database/sql"
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/constants"
	"github.com/flipped-aurora/gin-vue-admin/server/enums/fund"
	"github.com/flipped-aurora/gin-vue-admin/server/errorx"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/coupon"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/symbol"
	"github.com/flipped-aurora/gin-vue-admin/server/model/userfund"
	. "github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/guregu/null/v5"
	. "github.com/shopspring/decimal"
	"gorm.io/gorm"
	"strconv"
	"sync"
	"time"
)

const (
	FollowOrderNoPrefix     = "FQ"
	FollowOrderNoTimeFormat = "20060102150405"
	FollowOrderRandMaxPos   = 4
)

type UserFollowOrderService struct {
}

func (s *UserFollowOrderService) Apply(req *request.UserFollowOrderApplyReq) error {
	if err := validateApply(req); err != nil {
		return err
	}

	platformCommissionRate, _ := global.GVA_REDIS.Get(context.Background(), constants.RedisKeyPlatformCommissionRate).Float64()

	couponAmount := NewFromFloat(0)
	if req.CouponRecordId > 0 {
		var couponIssued coupon.CouponIssued
		global.GVA_DB.First(&couponIssued, req.CouponRecordId)
		couponAmount = couponAmount.Add(NewFromFloat(*couponIssued.CouponAmount))
	}

	return global.GVA_DB.Transaction(func(tx *gorm.DB) (txErr error) {
		var advisorProd order.AdvisorProd
		tx.First(&advisorProd, req.AdvisorProdId)

		followOrder := &order.UserFollowOrder{
			GVA_MODEL: global.GVA_MODEL{
				ID: uint(global.Snowflake.Generate()),
			},
			FollowOrderNo:          GenerateOderNo(FollowOrderNoPrefix, FollowOrderRandMaxPos, FollowOrderNoTimeFormat),
			UserId:                 req.UserId,
			AdvisorProdId:          req.AdvisorProdId,
			AutoRenew:              req.AutoRenew,
			AdvisorCommissionRate:  NullDecimal{Decimal: advisorProd.CommissionRate, Valid: true},
			PlatformCommissionRate: NullDecimal{Decimal: NewFromFloat(platformCommissionRate), Valid: true},
			FollowOrderStatus:      order.FollowOrderStatusAuditing,
			StockStatus:            order.FollowOrderStockStatusUnOpened,
			FollowAmount:           NewFromFloat(req.FollowAmount),
			FollowAvailableAmount:  NewFromFloat(req.FollowAmount),
			RetrievableAmount:      NewFromFloat(0),
			ApplyTime:              sql.NullTime{Time: time.Now(), Valid: true},
		}

		if req.CouponRecordId > 0 {
			followOrder.CouponRecordId = sql.NullInt64{Int64: int64(req.CouponRecordId), Valid: true}
		}

		txErr = tx.Create(&followOrder).Error
		if txErr != nil {
			return txErr
		}

		//更新优惠券使用记录
		txErr = tx.Model(&coupon.CouponIssued{}).
			Where("id = ? and user_id = ?", req.CouponRecordId, req.UserId).
			Updates(map[string]interface{}{
				"status":     coupon.AlreadyUsed,
				"updated_at": time.Now().UnixMilli(),
			}).Error
		if txErr != nil {
			return txErr
		}

		//扣除资金账户
		actualPaidAmount := NewFromFloat(req.FollowAmount).Sub(couponAmount).RoundCeil(2).InexactFloat64()
		txErr = userFundService.UpdateUserFundAccountsAndNewFlow(int(req.UserId), fund.TradeFollow, actualPaidAmount, strconv.Itoa(int(followOrder.ID)))
		if txErr != nil {
			return txErr
		}

		//新增订单系统站内信
		tx.Create(&order.SystemOrderMsg{
			Type:       order.MsgTypeNewFollowOrder,
			OrderId:    followOrder.ID,
			UserId:     followOrder.UserId,
			ReadStatus: order.MsgReadStatusUnRead,
		})

		return nil
	})
}

func (s *UserFollowOrderService) PageQueryMyFollowOrder(req *request.MyFollowOrderPageReq) ([]*response.MyFollowOrderPageData, int64, error) {
	db := global.GVA_DB.Model(&order.UserFollowOrder{})
	if req.FollowOrderStatus != nil {
		db = db.Where("follow_order_status = ?", req.FollowOrderStatus)
	}
	if req.CreatedAtStart != nil {
		db = db.Where("created_at >= ?", time.UnixMilli(*req.CreatedAtStart))
	}
	if req.CreatedAtEnd != nil {
		db = db.Where("created_at <= ?", time.UnixMilli(*req.CreatedAtEnd))
	}
	if req.AdvisorId != nil {
		db = db.Where("advisor_id = ?", req.AdvisorId)
	}
	db = db.Where("user_id = ?", req.UserId)

	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	} else if total == 0 {
		return nil, 0, nil
	}

	var orderList []*order.UserFollowOrder
	err := req.Paginate()(db).
		Preload("AdvisorProd").Preload("AdvisorProd.Advisor").Preload("UserFollowStockDetails").
		Order("created_at desc").
		Find(&orderList).
		Error
	if err != nil {
		return nil, 0, err
	}

	list := make([]*response.MyFollowOrderPageData, len(orderList))
	for i, v := range orderList {
		list[i] = &response.MyFollowOrderPageData{
			FollowOrderId:         strconv.Itoa(int(v.ID)),
			FollowOrderNo:         v.FollowOrderNo,
			AdvisorName:           v.AdvisorProd.Advisor.NickName,
			ProductName:           v.AdvisorProd.ProductName,
			FollowAmount:          v.FollowAmount.InexactFloat64(),
			AdvisorCommissionRate: v.AdvisorProd.CommissionRate.InexactFloat64(),
			FollowOrderStatus:     v.FollowOrderStatus,
			RetrievableAmount:     v.RetrievableAmount.InexactFloat64(),
		}
		if v.PeriodStart.Valid {
			list[i].PeriodStart = null.IntFrom(v.PeriodStart.Time.UnixMilli())
		}
		if v.PeriodEnd.Valid {
			list[i].PeriodEnd = null.IntFrom(v.PeriodEnd.Time.UnixMilli())
		}
		if v.ApplyTime.Valid {
			list[i].ApplyTime = null.IntFrom(v.ApplyTime.Time.UnixMilli())
		}
		if v.EndTime.Valid {
			list[i].FinishTime = null.IntFrom(v.EndTime.Time.UnixMilli())
		}

		if len(v.UserFollowStockDetails) > 0 {
			totalTradeReturn, totalPlatformCommission := NewFromFloat(0), NewFromFloat(0)
			for _, d := range v.UserFollowStockDetails {
				if d.SellPrice.Valid {
					d.UserFollowOrder = *v
					d.CalculateProfit()
					totalTradeReturn = totalTradeReturn.Add(d.TradeReturn)
					if d.PlatformCommission.Valid {
						totalPlatformCommission = totalPlatformCommission.Add(d.PlatformCommission.Decimal)
					}
				}
			}
			list[i].PlatformCommission = totalPlatformCommission.InexactFloat64()
			list[i].ReturnAmount = totalTradeReturn.InexactFloat64()
		}
	}

	return list, total, nil
}

func (s *UserFollowOrderService) GetMyFollowOrderDetail(followOrderId uint, userId uint) (*response.MyFollowOrderDetail, error) {
	var myFollowOrder order.UserFollowOrder
	err := global.GVA_DB.Preload("AdvisorProd").Preload("AdvisorProd.Advisor").Model(&myFollowOrder).
		Where("user_id = ? and id = ?", userId, followOrderId).
		First(&myFollowOrder).
		Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errorx.NewWithCode(errorx.FollowOrderNotFound)
	} else if err != nil {
		return nil, err
	}

	var stockOrderDetails []*order.UserFollowStockDetail
	err = global.GVA_DB.Preload("UserFollowOrder").Model(&order.UserFollowStockDetail{}).
		Where("follow_order_id = ?", followOrderId).
		Find(&stockOrderDetails).
		Error
	if err != nil {
		return nil, err
	}

	resp := &response.MyFollowOrderDetail{
		FollowOrderId:     strconv.Itoa(int(myFollowOrder.ID)),
		FollowOrderNo:     myFollowOrder.FollowOrderNo,
		FollowOrderStatus: myFollowOrder.FollowOrderStatus,
		StepAmount:        myFollowOrder.AdvisorProd.AmountStep,
		RetrievableAmount: myFollowOrder.RetrievableAmount.InexactFloat64(),
		FollowAmount:      myFollowOrder.FollowAmount.InexactFloat64(),
		AutoRenew:         myFollowOrder.AutoRenew,
		ProductName:       myFollowOrder.AdvisorProd.ProductName,
		AdvisorName:       myFollowOrder.AdvisorProd.Advisor.NickName,
		AdvisorAvatarUrl:  myFollowOrder.AdvisorProd.Advisor.AvatarUrl,
		Details:           make([]*response.MyFollowOrderStockDetail, len(stockOrderDetails)),
	}

	if myFollowOrder.CouponRecordId.Valid {
		var couponIssued coupon.CouponIssued
		if global.GVA_DB.First(&couponIssued, myFollowOrder.CouponRecordId.Int64).Error == nil {
			resp.CouponName = &couponIssued.CouponName
			resp.CouponAmount = couponIssued.CouponAmount
		}
	}

	if len(stockOrderDetails) > 0 {
		stockIds := make([]uint, len(stockOrderDetails))
		for _, v := range stockOrderDetails {
			stockIds = append(stockIds, v.StockId)
		}
		stockMap := make(map[uint]symbol.Symbols)
		var stockList []symbol.Symbols
		global.GVA_DB.Model(&symbol.Symbols{}).Where("id in (?)", stockIds).Find(&stockList)
		for _, v := range stockList {
			stockMap[uint(*v.Id)] = v
		}

		totalTradeReturnDecimal, totalActualReturnDecimal := NewFromFloat(0), NewFromFloat(0)
		for i, v := range stockOrderDetails {
			resp.Details[i] = &response.MyFollowOrderStockDetail{
				StockNum: v.StockNum.InexactFloat64(),
				BuyPrice: v.BuyPrice.InexactFloat64(),
				BuyTime:  v.BuyTime.UnixMilli(),
			}
			if v.SellPrice.Valid {
				resp.Details[i].SellPrice = null.FloatFrom(v.SellPrice.Decimal.InexactFloat64())
				v.CalculateProfit()
				resp.Details[i].TradeReturn = v.TradeReturn.InexactFloat64()
				resp.Details[i].ActualReturn = v.ActualReturn.InexactFloat64()
				if v.AdvisorCommission.Valid {
					resp.Details[i].AdvisorCommission = v.AdvisorCommission.Decimal.InexactFloat64()
				}
				if v.PlatformCommission.Valid {
					resp.Details[i].PlatformCommission = v.PlatformCommission.Decimal.InexactFloat64()
				}

				totalTradeReturnDecimal = totalTradeReturnDecimal.Add(v.TradeReturn)
				totalActualReturnDecimal = totalActualReturnDecimal.Add(v.ActualReturn)
			}
			if v.SellTime.Valid {
				resp.Details[i].SellTime = null.IntFrom(v.SellTime.Time.UnixMilli())
			}
			if stock, ok := stockMap[v.StockId]; ok {
				resp.Details[i].StockName = stock.Symbol
			}
		}
		resp.TotalTradeReturn = totalTradeReturnDecimal.InexactFloat64()
		resp.TotalActualReturn = totalActualReturnDecimal.InexactFloat64()
	}

	return resp, nil
}

func (s *UserFollowOrderService) RetrieveFollowOrder(req *request.RetrieveFollowOrderReq) error {
	var followOrder order.UserFollowOrder
	err := global.GVA_DB.Where("id = ? and user_id = ?", req.FollowOrderId, req.UserId).First(&followOrder).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errorx.NewWithCode(errorx.FollowOrderNotFound)
	} else if err != nil {
		return err
	}

	if req.RetrieveAmount > followOrder.RetrievableAmount.InexactFloat64() {
		return errorx.NewWithCode(errorx.InsufficientRetrievableAmount)
	}

	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		txErr := tx.Model(&followOrder).
			Update("retrievable_amount", gorm.Expr("retrievable_amount - ?", req.RetrieveAmount)).
			Error
		if txErr != nil {
			return txErr
		}

		txErr = userFundService.UpdateUserFundAccountsAndNewFlow(int(followOrder.UserId), fund.SettleProfit, req.RetrieveAmount, strconv.Itoa(int(followOrder.ID)))
		if txErr != nil {
			return txErr
		}
		return nil
	})
}

func (s *UserFollowOrderService) PageQuery(req *request.UserFollowOrderPageQueryReq) ([]*response.UserFollowOrderPageData, int64, error) {
	db := global.GVA_DB.Model(&order.UserFollowOrder{}).Joins("AdvisorProd").Joins("AdvisorProd.Advisor").Joins("User")
	if req.FollowOrderId != nil {
		followOrderId, _ := strconv.Atoi(*req.FollowOrderId)
		db = db.Where("user_follow_order.id = ?", followOrderId)
	}
	if req.ProductName != nil {
		db = db.Where("AdvisorProd.product_name like ?", "%"+*req.ProductName+"%")
	}
	if req.FollowOrderStatus != nil {
		db = db.Where("follow_order_status = ?", *req.FollowOrderStatus)
	}
	if req.StockStatus != nil {
		db = db.Where("stock_status = ?", *req.StockStatus)
	}
	if req.AdvisorId != nil {
		db = db.Where("AdvisorProd.advisor_id = ?", *req.AdvisorId)
	}
	if req.UserType != nil && *req.UserType > 0 {
		db = db.Where("User.user_type=?", *req.UserType)
	}
	if req.RootUserId > 0 {
		db = db.Where("User.root_userid = ?", req.RootUserId)
	}
	if req.UserId != nil {
		db = db.Where("user_follow_order.user_id = ?", *req.UserId)
	}
	if req.UserPhone != nil && len(*req.UserPhone) > 0 {
		db = db.Where("User.phone = ?", *req.UserPhone)
	}
	if req.UserEmail != nil && len(*req.UserEmail) > 0 {
		db = db.Where("User.email = ?", *req.UserEmail)
	}
	if req.ApplyTimeStart != nil {
		db = db.Where("apply_time >= ?", time.UnixMilli(*req.ApplyTimeStart))
	}
	if req.ApplyTimeEnd != nil {
		db = db.Where("apply_time <= ?", time.UnixMilli(*req.ApplyTimeEnd))
	}
	if req.RootUserIdSearch != nil {
		db = db.Where("User.root_userid = ?", *req.RootUserIdSearch)
	}
	if req.ParentUserId != nil {
		db = db.Where("User.parent_id = ?", *req.ParentUserId)
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	} else if total == 0 {
		return make([]*response.UserFollowOrderPageData, 0), 0, nil
	}

	var followOrderList []*order.UserFollowOrder
	err := req.Paginate()(db).Order("user_follow_order.created_at desc").Find(&followOrderList).Error
	if err != nil {
		return nil, 0, err
	}

	resultList := make([]*response.UserFollowOrderPageData, len(followOrderList))
	for i, item := range followOrderList {
		resultList[i] = &response.UserFollowOrderPageData{
			FollowOrderId:          strconv.Itoa(int(item.ID)),
			AdvisorName:            item.AdvisorProd.Advisor.NickName,
			UserId:                 strconv.Itoa(int(item.UserId)),
			UserName:               item.User.NickName,
			Phone:                  item.User.Phone,
			UserType:               item.User.UserType,
			ProductName:            item.AdvisorProd.ProductName,
			AutoRenew:              item.AutoRenew,
			AdvisorCommissionRate:  item.AdvisorCommissionRate.Decimal.InexactFloat64(),
			FollowOrderStatus:      item.FollowOrderStatus,
			StockStatus:            item.StockStatus,
			CouponRecordId:         null.IntFrom(item.CouponRecordId.Int64),
			Amount:                 item.FollowAmount.InexactFloat64(),
			UserEmail:              item.User.Email,
			PlatformCommissionRate: item.PlatformCommissionRate.Decimal.InexactFloat64(),
			ApplyTime:              item.ApplyTime.Time.UnixMilli(),
			RootUserId:             item.User.RootUserid,
			ParentUserId:           item.User.ParentId,
		}
	}

	return resultList, total, nil
}

func (s *UserFollowOrderService) Approve(followOrderId uint) error {
	var followOrder order.UserFollowOrder
	err := global.GVA_DB.Preload("AdvisorProd").Where("id = ?", followOrderId).First(&followOrder).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errorx.NewWithCode(errorx.FollowOrderNotFound)
	} else if err != nil {
		return err
	}

	if followOrder.FollowOrderStatus != order.FollowOrderStatusAuditing {
		return errorx.NewWithCode(errorx.FollowOrderApprovedFailed)
	}

	followOrder.FollowOrderStatus = order.FollowOrderStatusFollowing
	followOrder.PeriodStart = sql.NullTime{Time: time.Now(), Valid: true}
	followOrder.PeriodEnd = sql.NullTime{Time: global.Calendar.WorkdaysFrom(time.Now(), followOrder.AdvisorProd.FollowPeriod), Valid: true}

	err = global.GVA_DB.Model(&followOrder).
		Select("FollowOrderStatus", "PeriodStart", "PeriodEnd").
		Updates(followOrder).
		Error

	return err
}

func (s *UserFollowOrderService) ApproveAll() error {
	var followOrders []order.UserFollowOrder
	err := global.GVA_DB.Preload("AdvisorProd").
		Where("follow_order_status = ?", order.FollowOrderStatusAuditing).
		Find(&followOrders).Error
	if err != nil {
		return err
	}
	if len(followOrders) == 0 {
		return nil
	}

	var wait sync.WaitGroup
	wait.Add(len(followOrders))

	for _, v := range followOrders {
		go func() {
			defer func() { wait.Done() }()

			v.FollowOrderStatus = order.FollowOrderStatusFollowing
			v.PeriodStart = sql.NullTime{Time: time.Now(), Valid: true}
			v.PeriodEnd = sql.NullTime{Time: global.Calendar.WorkdaysFrom(time.Now(), v.AdvisorProd.FollowPeriod), Valid: true}

			global.GVA_DB.Model(&v).
				Select("FollowOrderStatus", "PeriodStart", "PeriodEnd").
				Where("follow_order_status = ?", order.FollowOrderStatusAuditing).
				Updates(v)
		}()
	}

	wait.Wait()
	return nil
}

func (s *UserFollowOrderService) Reject(followOrderId uint) error {
	var followOrder order.UserFollowOrder
	err := global.GVA_DB.Where("id = ?", followOrderId).First(&followOrder).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errorx.NewWithCode(errorx.FollowOrderNotFound)
	} else if err != nil {
		return err
	}

	if followOrder.FollowOrderStatus != order.FollowOrderStatusAuditing {
		return errorx.NewWithCode(errorx.FollowOrderRejectFailed)
	}

	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		followOrder.FollowOrderStatus = order.FollowOrderStatusRejected
		txErr := tx.Model(&followOrder).Select("FollowOrderStatus").Updates(followOrder).Error
		if txErr != nil {
			return txErr
		}

		retrieveAmount := followOrder.FollowAvailableAmount
		if followOrder.CouponRecordId.Valid {
			var couponIssued coupon.CouponIssued
			txErr = global.GVA_DB.First(&couponIssued, followOrder.CouponRecordId.Int64).Error
			if txErr != nil {
				return txErr
			}

			retrieveAmount = retrieveAmount.Sub(NewFromFloat(*couponIssued.CouponAmount))
			txErr = global.GVA_DB.Model(&couponIssued).Updates(map[string]interface{}{
				"status": coupon.NotUsed,
			}).Error
			if txErr != nil {
				return txErr
			}
		}

		txErr = userFundService.UpdateUserFundAccountsAndNewFlow(int(followOrder.UserId), fund.OperationRefused, retrieveAmount.InexactFloat64(), strconv.Itoa(int(followOrderId)))
		return txErr
	})
}

func (s *UserFollowOrderService) Cancel(followOrderId uint, userId uint) error {
	var followOrder order.UserFollowOrder
	err := global.GVA_DB.
		Where("id = ? and user_id = ?", followOrderId, userId).
		First(&followOrder).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errorx.NewWithCode(errorx.FollowOrderNotFound)
	} else if err != nil {
		return err
	}

	if followOrder.FollowOrderStatus != order.FollowOrderStatusAuditing {
		return errorx.NewWithCode(errorx.FollowOrderCancelFailed)
	}

	return global.GVA_DB.Transaction(func(tx *gorm.DB) (txErr error) {
		followOrder.FollowOrderStatus = order.FollowOrderStatusCancelled
		txErr = tx.Model(&followOrder).Update("FollowOrderStatus", order.FollowOrderStatusCancelled).Error
		if txErr != nil {
			return txErr
		}

		retrieveAmount := followOrder.FollowAvailableAmount
		if followOrder.CouponRecordId.Valid {
			var couponIssued coupon.CouponIssued
			txErr = global.GVA_DB.First(&couponIssued, followOrder.CouponRecordId.Int64).Error
			if txErr != nil {
				return txErr
			}

			retrieveAmount = retrieveAmount.Sub(NewFromFloat(*couponIssued.CouponAmount))
			txErr = global.GVA_DB.Model(&couponIssued).Updates(map[string]interface{}{
				"status": coupon.NotUsed,
			}).Error
			if txErr != nil {
				return txErr
			}
		}

		txErr = userFundService.UpdateUserFundAccountsAndNewFlow(int(userId), fund.CancelTradeFollow, retrieveAmount.InexactFloat64(), strconv.Itoa(int(followOrderId)))
		return
	})
}

func (s *UserFollowOrderService) SubmitFollowConfirm(req *request.UserFollowOrderConfirmSubmitReq) error {
	if req.StockNum <= 0 {
		return errorx.NewWithCode(errorx.StockNumShouldGtZero)
	}

	var followOrder order.UserFollowOrder
	err := global.GVA_DB.
		Where("id = ? and user_id = ?", req.FollowOrderId, req.UserId).
		First(&followOrder).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errorx.NewWithCode(errorx.FollowOrderNotFound)
	} else if err != nil {
		return err
	}

	var advisorStockOrder order.AdvisorStockOrder
	err = global.GVA_DB.First(&advisorStockOrder, req.StockOrderId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errorx.NewWithCode(errorx.AdvisorStockOrderNotFound)
	} else if err != nil {
		return err
	}

	if advisorStockOrder.Status != order.StockOrderStatusHolding {
		return errorx.NewWithCode(errorx.AdvisorStockOrderAlreadyFinished)
	}

	var stock symbol.Symbols
	if err = global.GVA_DB.First(&stock, advisorStockOrder.StockId).Error; err != nil {
		return err
	}

	stockNum := NewFromFloat(req.StockNum).RoundFloor(int32(GetScale(*stock.TicketNumSize)))
	if !stockNum.IsPositive() {
		return errorx.NewWithCode(errorx.StockNumShouldGtZero)
	}

	buyTotal := advisorStockOrder.BuyPrice.Mul(stockNum).RoundFloor(2)
	if buyTotal.Cmp(followOrder.FollowAvailableAmount) > 0 {
		return errorx.NewWithCode(errorx.InsufficientFollowAvailableAmount)
	}

	if followOrder.PeriodEnd.Valid && followOrder.PeriodEnd.Time.Before(time.Now()) {
		return errorx.NewWithCode(errorx.FollowPeriodFinished)
	}

	if err = validateMarketStatus(advisorStockOrder.StockId); err != nil {
		return err
	}

	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		txErr := tx.Model(&followOrder).Updates(map[string]interface{}{
			"follow_available_amount": gorm.Expr("follow_available_amount - ?", buyTotal),
			"stock_status":            order.FollowOrderStockStatusOpened,
		}).Error
		if txErr != nil {
			return txErr
		}

		userFollowStockDetail := &order.UserFollowStockDetail{
			GVA_MODEL: global.GVA_MODEL{
				ID: uint(global.Snowflake.Generate()),
			},
			UserId:        followOrder.UserId,
			StockOrderId:  advisorStockOrder.ID,
			FollowOrderId: followOrder.ID,
			StockId:       advisorStockOrder.StockId,
			StockNum:      stockNum,
			BuyPrice:      advisorStockOrder.BuyPrice,
			BuyTime:       advisorStockOrder.BuyTime,
		}
		if txErr = tx.Create(userFollowStockDetail).Error; txErr != nil {
			return txErr
		}

		return nil
	})
}

func (s *UserFollowOrderService) GetFollowConfirmDetail(followOrderId uint) (*response.UserFollowConfirmDetail, error) {
	var followOrder order.UserFollowOrder
	err := global.GVA_DB.Preload("AdvisorProd").
		Model(&followOrder).
		Where("id = ?", followOrderId).
		First(&followOrder).
		Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errorx.NewWithCode(errorx.FollowOrderNotFound)
	} else if err != nil {
		return nil, err
	}

	var advisorStockOrder order.AdvisorStockOrder
	err = global.GVA_DB.
		Model(&advisorStockOrder).
		Where("advisor_id = ? and status = ?", followOrder.AdvisorProd.AdvisorId, order.StockOrderStatusHolding).
		Order("created_at desc").
		Limit(1).
		First(&advisorStockOrder).
		Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errorx.NewWithCode(errorx.AdvisorHasNotOpenStockOrder)
	} else if err != nil {
		return nil, err
	}

	var stock symbol.Symbols
	err = global.GVA_DB.Find(&stock, advisorStockOrder.StockId).Error
	if err != nil {
		return nil, err
	}

	confirmDetail := &response.UserFollowConfirmDetail{
		FollowOrderId:   strconv.Itoa(int(followOrderId)),
		StockOrderId:    strconv.Itoa(int(advisorStockOrder.ID)),
		UserId:          strconv.Itoa(int(followOrder.UserId)),
		BuyPrice:        advisorStockOrder.BuyPrice.InexactFloat64(),
		AvailableAmount: followOrder.FollowAvailableAmount.InexactFloat64(),
		MaxStockNum:     followOrder.FollowAvailableAmount.Div(advisorStockOrder.BuyPrice).RoundFloor(int32(GetScale(*stock.TicketNumSize))).InexactFloat64(),
		StockName:       stock.Symbol,
	}

	return confirmDetail, nil
}

func (s *UserFollowOrderService) CancelAutoRenew(followOrderId uint, userId uint) error {
	var followOrder order.UserFollowOrder
	err := global.GVA_DB.Model(&followOrder).
		Where("id = ? and user_id = ?", followOrderId, userId).
		First(&followOrder).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errorx.NewWithCode(errorx.FollowOrderNotFound)
	} else if err != nil {
		return err
	}

	err = global.GVA_DB.Model(&followOrder).Update("auto_renew", order.DisableAutoRenew).Error
	return err
}

func (s *UserFollowOrderService) GetUserProfitSummary(userId uint) *response.UserFollowOrderProfitSummary {
	resp := &response.UserFollowOrderProfitSummary{}
	w := sync.WaitGroup{}
	w.Add(3)
	go func() {
		defer w.Done()
		var followStockDetails []order.UserFollowStockDetail
		global.GVA_DB.Model(&order.UserFollowStockDetail{}).
			Where("user_id = ? and sell_price is not null", userId).
			Find(&followStockDetails)

		totalProfit := NewFromFloat(0)
		for _, detail := range followStockDetails {
			totalProfit = totalProfit.Add(detail.SellPrice.Decimal.Sub(detail.BuyPrice).Mul(detail.StockNum).RoundFloor(2))
			if detail.AdvisorCommission.Valid {
				totalProfit = totalProfit.Sub(detail.AdvisorCommission.Decimal)
			}
			if detail.PlatformCommission.Valid {
				totalProfit = totalProfit.Sub(detail.PlatformCommission.Decimal)
			}
		}
		resp.TotalProfit = totalProfit.InexactFloat64()
	}()

	go func() {
		defer w.Done()
		var followStockDetails []order.UserFollowStockDetail
		global.GVA_DB.Model(&order.UserFollowStockDetail{}).
			Where("user_id = ? and sell_time = cast(? as date)", userId, time.Now()).
			Find(&followStockDetails)

		totalProfit := NewFromFloat(0)
		for _, detail := range followStockDetails {
			totalProfit = totalProfit.Add(detail.SellPrice.Decimal.Sub(detail.BuyPrice).Mul(detail.StockNum).RoundFloor(2))
			if detail.AdvisorCommission.Valid {
				totalProfit = totalProfit.Sub(detail.AdvisorCommission.Decimal)
			}
			if detail.PlatformCommission.Valid {
				totalProfit = totalProfit.Sub(detail.PlatformCommission.Decimal)
			}
		}
		resp.TodayProfit = totalProfit.InexactFloat64()
	}()

	go func() {
		defer w.Done()
		var userShareProfit float64
		global.GVA_DB.Model(&order.UserProfitShare{}).
			Select("sum(amount)").Where("to_user_id = ?", userId).Find(&userShareProfit)
		resp.UserShareProfit = NewFromFloat(userShareProfit).InexactFloat64()
	}()

	w.Wait()
	return resp
}

func (s *UserFollowOrderService) GetTotalAmount(userId uint) float64 {
	var (
		totalAvailableAmount float64
		totalHoldingAmount   float64
		totalAppendAmount    float64
		w                    sync.WaitGroup
	)
	w.Add(3)

	go func() {
		defer w.Done()
		global.GVA_DB.Model(&order.UserFollowStockDetail{}).
			Select("sum(buy_price * stock_num)").
			Where("user_id = ? and sell_price is null", userId).
			First(&totalHoldingAmount)
	}()

	go func() {
		defer w.Done()
		global.GVA_DB.Model(&order.UserFollowOrder{}).
			Select("sum(follow_available_amount + retrievable_amount)").
			Where("user_id = ? and follow_order_status in (?)", userId, []order.FollowOrderStatus{order.FollowOrderStatusAuditing, order.FollowOrderStatusFollowing}).
			First(&totalAvailableAmount)
	}()

	go func() {
		defer w.Done()
		global.GVA_DB.Model(&order.UserFollowAppendOrder{}).
			Select("sum(append_amount)").
			Where("user_id = ? and append_order_status = ?", userId, order.AppendOrderStatusAuditing).First(&totalAppendAmount)
	}()

	w.Wait()
	return NewFromFloat(totalAvailableAmount).Add(NewFromFloat(totalAppendAmount)).Add(NewFromFloat(totalHoldingAmount)).RoundFloor(2).InexactFloat64()
}

func (s *UserFollowOrderService) PageQueryRetrieveRecord(req *request.UserFollowOrderRetrieveRecordPageReq) (
	[]response.UserFollowOrderRetrieveRecord, int64, error) {
	db := global.GVA_DB.Model(&userfund.UserAccountFlow{}).Where("user_id = ? and order_id = ? and transaction_type = ? ", req.UserId, strconv.Itoa(int(req.FollowOrderId)), fund.SettleProfit)
	var total int64
	if err := db.Count(&total).Error; err != nil || total == 0 {
		return nil, 0, err
	}

	var flows []userfund.UserAccountFlow
	if err := db.Scopes(req.Paginate()).Order("transaction_date desc").Find(&flows).Error; err != nil {
		return nil, 0, err
	}

	records := make([]response.UserFollowOrderRetrieveRecord, len(flows))
	for i, v := range flows {
		records[i] = response.UserFollowOrderRetrieveRecord{
			Amount:       v.Amount.InexactFloat64(),
			RetrieveTime: v.TransactionDate.UnixMilli(),
		}
	}

	return records, total, nil
}
