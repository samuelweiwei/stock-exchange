package order

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/enums/fund"
	"github.com/flipped-aurora/gin-vue-admin/server/errorx"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order/response"
	. "github.com/shopspring/decimal"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
)

const (
	AppendOrderNoPrefix   = "AQ"
	AppendOrderRandMaxPos = 4
	AppendOrderTimeFormat = "20060102150405"
)

type UserFollowAppendOrderService struct{}

func (service *UserFollowAppendOrderService) QueryMyFollowAppendOrderList(followOrderId uint, userId uint) ([]*response.MyFollowAppendOrderListData, error) {
	var appendOrderList []order.UserFollowAppendOrder
	err := global.GVA_DB.Model(&order.UserFollowAppendOrder{}).
		Where("follow_order_id = ? and user_id = ?", followOrderId, userId).
		Find(&appendOrderList).
		Error

	if err != nil {
		return nil, err
	}

	resultList := make([]*response.MyFollowAppendOrderListData, len(appendOrderList))
	for i, item := range appendOrderList {
		resultList[i] = &response.MyFollowAppendOrderListData{
			AppendOrderNo:     item.AppendOrderNo,
			AppendAmount:      item.AppendAmount,
			AppendOrderStatus: item.AppendOrderStatus,
			CreatedAt:         item.CreatedAt.UnixMilli(),
		}
	}

	return resultList, nil
}

func (service *UserFollowAppendOrderService) PageQuery(req *request.UserFollowAppendOrderPageQueryReq) (
	[]*response.UserFollowAppendOrderPageData, int64, error) {
	db := global.GVA_DB.Model(&order.UserFollowAppendOrder{}).Joins("User").Joins("UserFollowOrder").Joins("UserFollowOrder.AdvisorProd").Joins("UserFollowOrder.AdvisorProd.Advisor")

	if req.FollowOrderId != nil && len(*req.FollowOrderId) > 0 {
		followOrderId, _ := strconv.Atoi(*req.FollowOrderId)
		db = db.Where("follow_order_id = ?", followOrderId)
	}
	if req.UserId != nil && len(*req.UserId) > 0 {
		userId, _ := strconv.Atoi(*req.UserId)
		db = db.Where("user_follow_append_order.user_id = ?", userId)
	}
	if req.Phone != nil && len(*req.Phone) > 0 {
		db = db.Where("User.phone = ?", *req.Phone)
	}
	if req.AppendOrderStatus != nil && len(*req.AppendOrderStatus) > 0 {
		status, _ := strconv.Atoi(*req.AppendOrderStatus)
		db = db.Where("append_order_status = ?", status)
	}
	if req.RootUserId > 0 {
		db = db.Where("User.root_userid = ?", req.RootUserId)
	}
	if req.UserType != nil && *req.UserType > 0 {
		db = db.Where("User.user_type = ?", *req.UserType)
	}
	if req.UserEmail != nil && len(*req.UserEmail) > 0 {
		db = db.Where("User.email = ?", *req.UserEmail)
	}
	if req.AdvisorName != nil && len(*req.AdvisorName) > 0 {
		db = db.Where("UserFollowOrder__AdvisorProd__Advisor.nick_name = ?", *req.AdvisorName)
	}
	if req.UserPhone != nil && len(*req.UserPhone) > 0 {
		db = db.Where("User.phone = ?", *req.UserPhone)
	}
	if req.AppendOrderNo != nil && len(*req.AppendOrderNo) > 0 {
		db = db.Where("append_order_no = ?", *req.AppendOrderNo)
	}
	if req.RootUserIdSearch != nil {
		db = db.Where("User.root_userid = ?", *req.RootUserIdSearch)
	}
	if req.ParentUserId != nil {
		db = db.Where("User.parent_id = ?", *req.ParentUserId)
	}
	if req.ProductName != nil && len(*req.ProductName) > 0 {
		db = db.Where("UserFollowOrder__AdvisorProd.product_name like ?", "%"+*req.ProductName+"%")
	}

	var total int64
	err := db.Count(&total).Error
	if err != nil || total == 0 {
		return nil, 0, err
	}

	var appendOrderList []order.UserFollowAppendOrder
	err = db.Scopes(req.Paginate()).
		Order("user_follow_append_order.created_at desc").
		Find(&appendOrderList).
		Error
	if err != nil {
		return nil, 0, err
	}

	resultList := make([]*response.UserFollowAppendOrderPageData, len(appendOrderList))
	for i, item := range appendOrderList {
		resultList[i] = &response.UserFollowAppendOrderPageData{
			AppendOrderId:     strconv.Itoa(int(item.ID)),
			AppendOrderNo:     item.AppendOrderNo,
			FollowOrderId:     strconv.Itoa(int(item.FollowOrderId)),
			UserId:            strconv.Itoa(int(item.UserId)),
			UserType:          item.User.UserType,
			Phone:             item.User.Phone,
			AppendAmount:      item.AppendAmount,
			ProductName:       item.UserFollowOrder.AdvisorProd.ProductName,
			CreatedAt:         item.CreatedAt.UnixMilli(),
			AppendOrderStatus: item.AppendOrderStatus,
			UserEmail:         item.User.Email,
			AdvisorName:       item.UserFollowOrder.AdvisorProd.Advisor.NickName,
			RootUserId:        item.User.RootUserid,
			ParentUserId:      item.User.ParentId,
		}
	}

	return resultList, total, nil
}

func (service *UserFollowAppendOrderService) Apply(req *request.UserFollowAppendOrderApplyReq) error {
	var followOrder order.UserFollowOrder
	err := global.GVA_DB.Preload("UserFollowStockDetails").Preload("AdvisorProd").
		Where("id = ? and user_id = ?", req.FollowOrderId, req.UserId).
		First(&followOrder).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errorx.NewWithCode(errorx.FollowOrderNotFound)
	} else if err != nil {
		return err
	}

	// 当前可用金额 + 持仓中股票金额
	currentPrincipal := followOrder.FollowAvailableAmount
	for _, v := range followOrder.UserFollowStockDetails {
		if !v.SellPrice.Valid {
			currentPrincipal = currentPrincipal.Add(v.BuyPrice.Mul(v.StockNum))
		}
	}
	if currentPrincipal.Add(NewFromFloat(req.AppendAmount)).Compare(NewFromFloat(followOrder.AdvisorProd.MaxAmount)) > 0 {
		return errorx.NewWithCode(errorx.IllegalFollowAmount, "maxAmount", followOrder.AdvisorProd.MaxAmount, "currentPrincipal", currentPrincipal.InexactFloat64())
	}

	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		userFollowAppendOrder := &order.UserFollowAppendOrder{
			GVA_MODEL: global.GVA_MODEL{
				ID: uint(global.Snowflake.Generate()),
			},
			FollowOrderId:     req.FollowOrderId,
			UserId:            req.UserId,
			AppendAmount:      req.AppendAmount,
			AppendOrderStatus: order.AppendOrderStatusAuditing,
			AppendOrderNo:     GenerateOderNo(AppendOrderNoPrefix, AppendOrderRandMaxPos, AppendOrderTimeFormat),
		}

		txErr := tx.Create(userFollowAppendOrder).Error
		if txErr != nil {
			return txErr
		}

		txErr = userFundService.UpdateUserFundAccountsAndNewFlow(int(followOrder.UserId), fund.ApplyOrderFollow, req.AppendAmount, strconv.Itoa(int(userFollowAppendOrder.ID)))
		if txErr != nil {
			return txErr
		}

		txErr = global.GVA_DB.Create(&order.SystemOrderMsg{
			Type:       order.MsgTypeNewAppendOrder,
			OrderId:    userFollowAppendOrder.ID,
			UserId:     userFollowAppendOrder.UserId,
			ReadStatus: order.MsgReadStatusUnRead,
		}).Error

		if txErr != nil {
			global.GVA_LOG.Error("Add Append Order Msg Error", zap.Error(err))
		}
		return nil
	})
}

func (service *UserFollowAppendOrderService) Approve(appendOrderId uint) error {
	var appendOrder order.UserFollowAppendOrder
	err := global.GVA_DB.Preload("UserFollowOrder").Preload("UserFollowOrder.AdvisorProd").Preload("UserFollowOrder.UserFollowStockDetails").
		First(&appendOrder, appendOrderId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errorx.NewWithCode(errorx.AppendOrderNotFound)
	} else if err != nil {
		return err
	}

	// 当前可用金额 + 持仓中股票金额
	currentPrincipal := appendOrder.UserFollowOrder.FollowAvailableAmount
	for _, v := range appendOrder.UserFollowOrder.UserFollowStockDetails {
		if !v.SellPrice.Valid {
			currentPrincipal = currentPrincipal.Add(v.BuyPrice.Mul(v.StockNum))
		}
	}
	if currentPrincipal.Add(NewFromFloat(appendOrder.AppendAmount)).Compare(NewFromFloat(appendOrder.UserFollowOrder.AdvisorProd.MaxAmount)) > 0 {
		return errorx.NewWithCode(errorx.IllegalFollowAmount, "maxAmount", appendOrder.UserFollowOrder.AdvisorProd.MaxAmount, "currentPrincipal", currentPrincipal.InexactFloat64())
	}

	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		txErr := tx.Model(&appendOrder).Updates(map[string]interface{}{
			"append_order_status": order.AppendOrderStatusApproved,
		}).Error
		if txErr != nil {
			return txErr
		}

		txErr = tx.Model(&order.UserFollowOrder{}).Where("id = ?", appendOrder.FollowOrderId).
			Updates(map[string]interface{}{
				"follow_amount":           gorm.Expr("follow_amount + ?", appendOrder.AppendAmount),
				"follow_available_amount": gorm.Expr("follow_available_amount + ?", appendOrder.AppendAmount),
			}).Error

		if txErr != nil {
			return txErr
		}

		return nil
	})
}

func (service *UserFollowAppendOrderService) Reject(appendOrderId uint) error {
	var appendOrder order.UserFollowAppendOrder
	err := global.GVA_DB.First(&appendOrder, appendOrderId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errorx.NewWithCode(errorx.AppendOrderNotFound)
	} else if err != nil {
		return err
	}

	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		txErr := tx.Model(&appendOrder).Updates(map[string]interface{}{
			"append_order_status": order.AppendOrderStatusRejected,
		}).Error
		if txErr != nil {
			return txErr
		}

		txErr = userFundService.UpdateUserFundAccountsAndNewFlow(int(appendOrder.UserId), fund.RefusedApplyOrderFollow, appendOrder.AppendAmount, strconv.Itoa(int(appendOrder.ID)))
		if txErr != nil {
			return txErr
		}

		return nil
	})
}
