package order

import (
	"database/sql"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	. "github.com/flipped-aurora/gin-vue-admin/server/utils"
	. "github.com/shopspring/decimal"
	"time"
)

type FollowOrderStatus uint
type FollowOrderStockStatus uint

// 用户跟单状态: 0-审核中;3-已驳回;5-跟单中;8-已撤回;10-已结束
const (
	FollowOrderStatusAuditing  = 0
	FollowOrderStatusRejected  = 3
	FollowOrderStatusFollowing = 5
	FollowOrderStatusCancelled = 8
	FollowOrderStatusFinished  = 10
)

// 持仓状态： 0-未开单；1-持仓中；3-已过期
const (
	FollowOrderStockStatusUnOpened = 0
	FollowOrderStockStatusOpened   = 1
	FollowOrderStockStatusExpired  = 3
)

type UserFollowOrder struct {
	global.GVA_MODEL
	FollowOrderNo          string                  `gorm:"column:follow_order_no"`
	UserId                 uint                    `gorm:"column:user_id"`
	AdvisorProdId          uint                    `gorm:"column:advisor_prod_id"`
	AutoRenew              AutoRenewStatus         `gorm:"column:auto_renew"`
	PeriodStart            sql.NullTime            `gorm:"column:period_start"`
	PeriodEnd              sql.NullTime            `gorm:"column:period_end"`
	AdvisorCommissionRate  NullDecimal             `gorm:"column:advisor_commission_rate"`
	PlatformCommissionRate NullDecimal             `gorm:"column:platform_commission_rate"`
	FollowOrderStatus      FollowOrderStatus       `gorm:"column:follow_order_status"`
	StockStatus            FollowOrderStockStatus  `gorm:"column:stock_status"`
	FollowAmount           Decimal                 `gorm:"column:follow_amount"`
	FollowAvailableAmount  Decimal                 `gorm:"column:follow_available_amount"`
	RetrievableAmount      Decimal                 `gorm:"column:retrievable_amount"`
	CouponRecordId         sql.NullInt64           `gorm:"column:coupon_record_id"`
	ApplyTime              sql.NullTime            `gorm:"column:apply_time"`
	EndTime                sql.NullTime            `gorm:"column:end_time"`
	AdvisorProd            AdvisorProd             `gorm:"foreignKey:AdvisorProdId"`
	UserFollowStockDetails []UserFollowStockDetail `gorm:"foreignKey:FollowOrderId"`
	User                   user.FrontendUsers      `gorm:"foreignKey:UserId"`
}

type UserFollowStockDetail struct {
	global.GVA_MODEL
	UserId             uint            `gorm:"column:user_id"`
	StockOrderId       uint            `gorm:"column:stock_order_id"`
	FollowOrderId      uint            `gorm:"column:follow_order_id"`
	StockId            uint            `gorm:"column:stock_id"`
	StockNum           Decimal         `gorm:"column:stock_num"`
	BuyPrice           Decimal         `gorm:"column:buy_price"`
	BuyTime            time.Time       `gorm:"column:buy_time"`
	SellPrice          NullDecimal     `gorm:"column:sell_price"`
	SellTime           sql.NullTime    `gorm:"column:sell_time"`
	AdvisorCommission  NullDecimal     `gorm:"column:advisor_commission"`
	PlatformCommission NullDecimal     `gorm:"column:platform_commission"`
	UserFollowOrder    UserFollowOrder `gorm:"foreignKey:FollowOrderId"`
	TradeReturn        Decimal         `gorm:"-:all"`
	ActualReturn       Decimal         `gorm:"-:all"`
}

type UserProfitShare struct {
	global.GVA_MODEL
	FromUserId    uint    `gorm:"column:from_user_id"`
	ToUserId      uint    `gorm:"column:to_user_id"`
	FollowOrderId uint    `gorm:"column:follow_order_id"`
	Amount        Decimal `gorm:"column:amount"`
	ShareRate     Decimal `gorm:"column:share_rate"`
}

func (UserFollowOrder) TableName() string {
	return "user_follow_order"
}
func (UserFollowStockDetail) TableName() string {
	return "user_follow_stock_detail"
}
func (UserProfitShare) TableName() string {
	return "user_profit_share"
}

// CalculateProfit 计算利润相关数据
func (detail *UserFollowStockDetail) CalculateProfit() {
	if !detail.SellPrice.Valid {
		return
	}

	advisorCommissionRate := detail.UserFollowOrder.AdvisorCommissionRate.Decimal
	platformCommissionRate := detail.UserFollowOrder.PlatformCommissionRate.Decimal

	if detail.SellPrice.Decimal.Cmp(detail.BuyPrice) >= 0 {
		detail.TradeReturn = detail.SellPrice.Decimal.Sub(detail.BuyPrice).Mul(detail.StockNum).RoundFloor(2)
		detail.AdvisorCommission = NullDecimal{Decimal: detail.TradeReturn.Mul(advisorCommissionRate).Div(HundredPercent).RoundCeil(2), Valid: true}
		detail.PlatformCommission = NullDecimal{Decimal: detail.TradeReturn.Mul(platformCommissionRate).Div(HundredPercent).RoundCeil(2), Valid: true}
		detail.ActualReturn = detail.TradeReturn.Sub(detail.AdvisorCommission.Decimal).Sub(detail.PlatformCommission.Decimal)
	} else {
		detail.TradeReturn = detail.SellPrice.Decimal.Sub(detail.BuyPrice).Mul(detail.StockNum).RoundFloor(2)
		detail.ActualReturn = detail.TradeReturn
	}
}
