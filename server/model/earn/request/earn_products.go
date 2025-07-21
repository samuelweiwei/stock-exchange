package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type EarnProductsSearch struct {
	Uid uint `json:"uid" form:"uid"`
	request.PageInfo
}
