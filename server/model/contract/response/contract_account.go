package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/contract"
	. "github.com/shopspring/decimal"
)

type ContractAccountRes struct {
	TotalAmount             Decimal                `json:"totalAmount" form:"totalAmount" `                        //总金额
	Balance                 Decimal                `json:"balance" form:"balance" `                                //账户余额
	RealizedProfitLoss      Decimal                `json:"realizedProfitLoss" form:"realizedProfitLoss"`           //已实现盈亏
	TodayRealizedProfitLoss Decimal                `json:"todayRealizedProfitLoss" form:"todayRealizedProfitLoss"` //今日已实现盈亏
	UnrealizedProfitLoss    Decimal                `json:"unrealizedProfitLoss" form:"unrealizedProfitLoss"`       //未实现盈亏
	AvailableMargin         Decimal                `json:"availableMargin" form:"availableMargin"`                 //可下单保证金
	TransferableAmount      Decimal                `json:"transferableAmount" form:"transferableAmount"`           //可划转金额
	AccountStatus           contract.AccountStatus `json:"accountStatus" form:"accountStatus"`                     //账号状态
}
