package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/shopspring/decimal"
)

type SubscriptionStatus uint
type RedeemStatus uint

const (
	UnKnown SubscriptionStatus = iota
	Staking
	Redeemed
)

const (
	RedeemUnKnown RedeemStatus = iota
	RedeemInAdvance
	RedeemNormal
)

type EarnSubscribeLogSearch struct {
	Uid             uint               `json:"uid"  form:"uid"`
	RedeemInAdvance RedeemStatus       `json:"redeemInAdvance"  form:"redeemInAdvance"`
	Status          SubscriptionStatus `json:"status"  form:"status"`
	UserType        int                `json:"userType"  form:"userType"`
	request.PageInfo
}

type EarnProductStake struct {
	ProductId uint            `json:"productId"`
	Amount    decimal.Decimal `json:"amount"`
}

type EarnProductRedeem struct {
	SubscriptionID uint `json:"subscriptionId"`
}

type PurchasedEarnProductsSearch struct {
	request.PageInfo
}
