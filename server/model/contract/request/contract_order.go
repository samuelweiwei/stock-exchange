package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/contract"
	. "github.com/shopspring/decimal"
	"time"
)

type ContractOrderSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type ContractOrderReq struct {
	StockId       uint                   `json:"stockId" form:"stockId"`              //股票id
	StockName     string                 `json:"stockName" form:"stockName"`          //股票名称
	OrderType     contract.OrderType     `json:"orderType" form:"orderType"`          //下单类型
	OpenPrice     Decimal                `json:"openPrice" form:"openPrice" `         //开仓价格
	OperationType contract.OperationType `json:"operationType" form:"operationType" ` //操作类型
	Margin        Decimal                `json:"margin" form:"margin" `               //保证金
	LeverageRatio int                    `json:"leverageRatio" form:"leverageRatio" ` //杠杆倍数
}

type ContractCloseReq struct {
	PositionId string  `json:"positionId" form:"positionId"` //持仓ID
	Quantity   Decimal `json:"quantity" form:"quantity" `    //数量
}
