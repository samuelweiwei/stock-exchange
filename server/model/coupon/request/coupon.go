package request

import "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"

type CouponSearch struct {
	Type int `json:"type" form:"type"`
	request.PageInfo
}
