package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type SymbolsCustomSearch struct {
	request.PageInfo
}

// CreateSymbolsCustomRequest 创建自定义交易对请求结构
type CreateSymbolsCustomRequest struct {
	SymbolId int `json:"symbolId" binding:"required"`
}
