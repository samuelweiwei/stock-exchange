package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/order"
	"github.com/guregu/null/v5"
)

// AdvisorStockOrderPageData 导师开单分页数据
// @Description 导师开单分页数据
type AdvisorStockOrderPageData struct {
	StockOrderId uint                   `json:"stockOrderId"` //导师开单记录ID
	StockId      uint                   `json:"stockId"`      //股票ID
	AdvisorName  string                 `json:"advisorName"`  //导师昵称
	StockName    string                 `json:"stockName"`    //股票名称
	BuyPrice     float64                `json:"buyPrice"`     //买入价格
	BuyTime      int64                  `json:"buyTime"`      //买入时间，Unix时间戳
	SellPrice    null.Float             `json:"sellPrice"`    //卖出价格
	SellTime     null.Int64             `json:"sellTime"`     //卖出时间，Unix时间戳
	Status       order.StockOrderStatus `json:"status"`       //导师开仓单状态：1-持仓中;5-已完结
	StatusText   string                 `json:"statusText"`   //导师开仓单状态文案
	CreatedAt    int64                  `json:"createdAt"`    //创建时间，Unix时间戳
	UpdatedAt    int64                  `json:"updatedAt"`    //更新时间，Unix时间戳
}

// AdvisorStockOrderConfirmSummary 导师开仓卖出确认详情
// @Description 导师开仓卖出确认详情
type AdvisorStockOrderConfirmSummary struct {
	StockId           uint                              `json:"stockId"`           //股票ID
	StockName         string                            `json:"stockName"`         //股票名称
	BuyPrice          float64                           `json:"buyPrice"`          //买入价格
	SellPrice         float64                           `json:"sellPrice"`         //卖出价格
	PercentageChange  float64                           `json:"percentageChange"`  //涨跌幅，单位%
	TotalActualReturn float64                           `json:"totalActualReturn"` //跟单实际总盈亏
	TotalTradeReturn  float64                           `json:"totalTradeReturn"`  //跟单交易总盈亏
	TotalStockNum     float64                           `json:"totalStockNum"`     //跟单总股票数量
	FollowOrders      []*AdvisorStockOrderConfirmDetail `json:"followOrders"`      //跟单明细
}

// AdvisorStockOrderConfirmDetail 导师开仓单跟单明细
// @Description 导师开仓单跟单明细
type AdvisorStockOrderConfirmDetail struct {
	UserId                 string  `json:"userId"`                 //会员ID
	StockNum               float64 `json:"stockNum"`               //股票数量
	ActualReturn           float64 `json:"actualReturn"`           //跟单实际盈亏
	TradeReturn            float64 `json:"tradeReturn"`            //跟单交易盈亏
	AdvisorCommission      float64 `json:"advisorCommission"`      //分析师佣金
	PlatformCommission     float64 `json:"platformCommission"`     //平台佣金
	PlatformCommissionRate float64 `json:"platformCommissionRate"` //平台佣金比例
	AdvisorCommissionRate  float64 `json:"advisorCommissionRate"`  //分析师佣金比例
}
