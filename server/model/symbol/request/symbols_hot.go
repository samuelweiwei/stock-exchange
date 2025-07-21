package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

// SymbolsHotSearch 用于热门股票查询的结构体
type SymbolsHotSearch struct {
	request.PageInfo
	Symbol string `json:"symbol" form:"symbol"` // 交易对名称
	Type   string `json:"type" form:"type"`     // 交易对类型
}
