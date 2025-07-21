package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CountriesSearch struct {
	Status    *int   `json:"status" form:"status"`
	Name      string `json:"name" form:"name"`           //名称
	PhoneCode string `json:"phoneCode" form:"phoneCode"` //手机区号
	request.PageInfo
}
