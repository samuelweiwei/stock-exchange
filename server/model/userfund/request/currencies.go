package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CurrenciesSearch struct {
	request.PageInfo
	Currency string `json:"currency" form:"currency"`
	CoinType *int   `json:"coinType" form:"coinType"` // 货币类型：1数字货币 2法币
	Source   string `json:"source" form:"source"`
}
