package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/contract"
	. "github.com/shopspring/decimal"
)

type ContractEntrustRes struct {
	ID            string                 `json:"ID" form:"ID"`                        // 主键ID
	PositionId    string                 `json:"positionId" form:"positionId"`        //关联持仓表的持仓 ID
	OrderId       string                 `json:"orderId" form:"orderId" `             //关联订单表的订单 ID
	StockId       string                 `json:"stockId" form:"stockId"`              //股票id
	StockName     string                 `json:"stockName" form:"stockName"`          //股票名称
	TriggerType   contract.TriggerType   `json:"triggerType" form:"triggerType"`      //触发类型 限价-1，止盈-2，止损-3
	TriggerPrice  Decimal                `json:"triggerPrice" form:"triggerPrice" `   //触发价格
	Margin        Decimal                `json:"margin" form:"margin"`                //保证金
	OperationType contract.OperationType `json:"operationType" form:"operationType" ` //操作类型 开多-1，开空-2，平多-3，平空-4
	Quantity      Decimal                `json:"quantity" form:"quantity" `           //数量
	EntrustStatus contract.EntrustStatus `json:"entrustStatus" form:"entrustStatus" ` //委托状态 未触发-1，已触发-2,已取消-99
}
