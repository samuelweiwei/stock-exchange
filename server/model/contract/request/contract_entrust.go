package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/contract"
	. "github.com/shopspring/decimal"
	"time"
)

type ContractEntrustSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type ContractEntrustReq struct {
	PositionId   string               `json:"positionId" form:"positionId"`      //持仓ID
	TriggerPrice Decimal              `json:"triggerPrice" form:"triggerPrice" ` //触发价格
	Quantity     Decimal              `json:"quantity" form:"quantity" `         //数量
	TriggerType  contract.TriggerType `json:"triggerType" form:"triggerType" `   //杠杆倍数
}
