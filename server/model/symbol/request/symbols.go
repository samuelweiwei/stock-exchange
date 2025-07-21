package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type SymbolsSearch struct {
	request.PageInfo
	Symbol string `json:"symbol" form:"symbol"` // 交易对名称
	Type   *int   `json:"type" form:"type"`     // 交易对类型：0-股票 1-加密货币 2-外汇
}

// WebSocketSubscribeReq WebSocket订阅请求结构
type WebSocketSubscribeReq struct {
	Req     string   `json:"req"`
	Symbols []string `json:"symbols"`
	Type    int      `json:"type"` // 0:股票 1:加密货币 2:外汇
}

// SymbolsPublicSearch 公开列表的搜索条件
type SymbolsPublicSearch struct {
	OrderBy  string `form:"orderBy" json:"orderBy"`   // 排序字段
	Order    string `form:"order" json:"order"`       // 排序方式
	Page     int    `form:"page" json:"page"`         // 页码
	PageSize int    `form:"pageSize" json:"pageSize"` // 每页大小
}

// Paginate 分页方法
func (s *SymbolsPublicSearch) Paginate() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (s.Page - 1) * s.PageSize
		return db.Offset(offset).Limit(s.PageSize)
	}
}

// TickerSearchReq 查询可添加的ticker列表请求参数
type TickerSearchReq struct {
	Type   int    `json:"type" form:"type" binding:"required,oneof=0 1 2"` // 0:股票 1:加密货币 2:外汇
	Search string `json:"search" form:"search"`                            // 搜索关键字
}

// AddSymbolRequest 定义添加交易对的请求结构
type AddSymbolRequest struct {
	Symbol string `json:"symbol" form:"symbol"`
	Type   int    `json:"type" form:"type"`
}

// SymbolNameRequest 根据symbol查询的请求结构
type SymbolNameRequest struct {
	Symbol string `json:"symbol" binding:"required"` // 交易对代码
}
