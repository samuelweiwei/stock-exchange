package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order"
)

// AdvisorStockOrderCreateReq 创建导师开单请求
// @Description 创建导师开单
type AdvisorStockOrderCreateReq struct {
	StockId   uint    `json:"stockId"`   //股票ID
	AdvisorId uint    `json:"advisorId"` //导师ID
	BuyPrice  float64 `json:"buyPrice"`  //买入价格
	BuyTime   int64   `json:"buyTime"`   //买入时间，Unix时间戳
}

// AdvisorStockOrderPageQueryReq 导师开单分页查询请求
// @Description 导师开单分页查询请求
type AdvisorStockOrderPageQueryReq struct {
	request.PageInfo
	AdvisorName      *string                 `form:"advisorName"`      //分析师昵称
	StockOrderStatus *order.StockOrderStatus `form:"stockOrderStatus"` //持仓状态：1-持仓中;5-已完结
	CreatedAtStart   *int64                  `form:"createdAtStart"`   //创建时间开始值
	CreatedAtEnd     *int64                  `form:"createdAtEnd"`     //创建时间结束值
	UpdatedAtStart   *int64                  `form:"updatedAtStart"`   //更新时间开始值
	UpdatedAtEnd     *int64                  `form:"updatedAtEnd"`     //更新时间结束值
}

// AdvisorStockOrderSellConfirmReq 查询导师开单卖出确认详情请求
// @Description 查询导师开单卖出确认详情请求
type AdvisorStockOrderSellConfirmReq struct {
	StockOrderId uint    `json:"-"`
	SellPrice    float64 `form:"sellPrice"` //卖出价格
}

// AdvisorStockOrderSellConfirmSubmitReq 导师开单卖出确认详情提交
// @Description 导师开单卖出确认详情提交
type AdvisorStockOrderSellConfirmSubmitReq struct {
	StockOrderId uint    `json:"-"`
	SellPrice    float64 `json:"sellPrice"` //卖出价格
	SellTime     int64   `json:"sellTime"`  //卖出时间，Unix时间戳
}

// AdvisorStockOrderAutoFollowReq 导师开单一键跟单请求
// @Description 导师开单一键跟单请求
type AdvisorStockOrderAutoFollowReq struct {
	StockOrderId  uint    `json:"-"`
	PositionRatio float64 `json:"positionRatio"` //仓位比例
}
