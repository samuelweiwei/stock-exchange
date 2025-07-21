package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/contract"
	. "github.com/shopspring/decimal"
)

type ContractOrderRes struct {
	ID                 string                 `json:"ID" form:"ID"`                                  // 主键ID
	OrderNumber        string                 `json:"orderNumber" form:"orderNumber"`                //订单号
	OrderTime          int64                  `json:"orderTime" form:"orderTime"`                    //订单时间
	StockId            string                 `json:"stockId" form:"stockId" `                       //股票id
	StockName          string                 `json:"stockName" form:"stockName" `                   //股票名称
	OrderType          contract.OrderType     `json:"orderType" form:"orderType" `                   //下单类型
	OpenPrice          Decimal                `json:"openPrice" form:"openPrice"`                    //开仓价格
	ClosePrice         Decimal                `json:"closePrice" form:"closePrice" `                 //平仓价格
	OperationType      contract.OperationType `json:"operationType" form:"operationType" `           //操作类型
	Quantity           Decimal                `json:"quantity" form:"quantity" `                     //数量
	OrderStatus        contract.OrderStatus   `json:"orderStatus" form:"orderStatus"`                //订单状态
	Fee                Decimal                `json:"fee" form:"fee"`                                //手续费
	RealizedProfitLoss Decimal                `json:"realizedProfitLoss" form:"realizedProfitLoss" ` //已实现盈亏
}
