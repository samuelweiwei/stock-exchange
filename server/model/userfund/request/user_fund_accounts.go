package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type UserFundAccountsSearch struct {
	UserId     int64   `json:"userId"` // 账户ID
	UserIds    string  `json:"userIds"`
	SendAmount float64 `json:"sendAmount"`
	UserType   int     `json:"userType" form:"userType"`
	request.PageInfo
}
