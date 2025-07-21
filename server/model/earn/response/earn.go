package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/earn"
	"github.com/shopspring/decimal"
)

type PurchasedEarnProduct struct {
	Subscription      *earn.EarnSubscribeLog `json:"purchased,omitempty"`
	TotalEarned       decimal.Decimal        `json:"totalEarned"`
	TodayInterestRate decimal.Decimal        `json:"todayInterestRate"`
	EarnProduct       *earn.EarnProducts     `json:"earnProduct"`
}

type PurchasedEarnProductsListRes struct {
	AllProductsEarned decimal.Decimal         `json:"allProductsEarned,omitempty"`
	List              []*PurchasedEarnProduct `json:"list" `
}
