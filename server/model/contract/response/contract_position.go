package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/contract"
	. "github.com/shopspring/decimal"
)

type ContractPositionRes struct {
	ID                   string                `json:"ID" form:"ID"`                                         // 主键ID
	StockId              string                `json:"stock_id" form:"stock_id"`                             // 股票id
	StockName            string                `json:"stock_name" form:"stock_name"`                         // 股票名称
	PositionTime         int64                 `json:"position_time" form:"position_time"`                   // 持仓时间
	Quantity             Decimal               `json:"quantity" form:"quantity"`                             // 持仓数量
	OpenPrice            Decimal               `json:"open_price" form:"open_price"`                         // 开仓价格
	CurPrice             Decimal               `json:"cur_price" form:"cur_price"`                           // 当前价格
	LeverageRatio        int                   `json:"leverage_ratio" form:"leverage_ratio"`                 // 杠杆倍数
	ROI                  string                `json:"ROI" form:"ROI"`                                       // 投资回报率
	SafetyFactor         string                `json:"safety_factor" form:"safety_factor"`                   // 安全系数
	Margin               Decimal               `json:"margin" form:"margin"`                                 // 保证金
	PositionAmount       Decimal               `json:"position_amount" form:"position_amount"`               // 持仓金额
	ForceClosePrice      Decimal               `json:"force_close_price" form:"force_close_price"`           // 强平价格
	PositionType         contract.PositionType `json:"position_type" form:"position_type"`                   // 持仓类型  多单-1，空单-2
	UnrealizedProfitLoss Decimal               `json:"unrealized_profit_loss" form:"unrealized_profit_loss"` // 未实现盈亏
}
