package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	. "github.com/shopspring/decimal"
	"time"
)

type ChangeMarginType int

const (
	TransferIn  ChangeMarginType = 1
	TransferOut ChangeMarginType = 2
)

type ContractAccountSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type ChangeMarginReq struct {
	Amount Decimal          `json:"amount" form:"amount"` // 划转金额
	Type   ChangeMarginType `json:"type" form:"type"`     // 划转类型 1-转入 2-转出
}
