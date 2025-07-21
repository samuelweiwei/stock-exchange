package order

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
)

type AutoRenewStatus int

const (
	DisableAutoRenew AutoRenewStatus = iota
	EnableAutoRenew
)

type AdvisorProd struct {
	global.GVA_MODEL
	AdvisorId          uint                `gorm:"column:advisor_id"`
	ProductName        string              `gorm:"column:product_name"`
	AmountStep         float64             `gorm:"column:amount_step"`
	MaxAmount          float64             `gorm:"column:max_amount"`
	MinAmount          float64             `gorm:"column:min_amount"`
	FollowPeriod       int                 `gorm:"column:follow_period"`
	CommissionRate     decimal.Decimal     `gorm:"column:commission_rate"`
	AutoRenew          AutoRenewStatus     `gorm:"column:auto_renew"`
	ActiveStatus       ActiveStatus        `gorm:"column:active_status"`
	Advisor            Advisor             `gorm:"foreignKey:AdvisorId"`
	AdvisorProdCoupons []AdvisorProdCoupon `gorm:"foreignKey:AdvisorProdId"`
}

type AdvisorProdCoupon struct {
	global.GVA_MODEL
	AdvisorProdId uint `gorm:"column:advisor_prod_id"`
	CouponId      uint `gorm:"column:coupon_id"`
}

func (AdvisorProd) TableName() string {
	return "advisor_prod"
}

func (AdvisorProdCoupon) TableName() string {
	return "advisor_prod_coupon"
}
