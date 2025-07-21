package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type UserAccountFlowSearch struct {
	UserId          int    `json:"userId" form:"userId"`
	TransactionType string `json:"transactionType" form:"transactionType"`
	StartTime       int64  `json:"startTime" form:"startTime"`       // 开始时间戳(毫秒)
	EndTime         int64  `json:"endTime" form:"endTime"`           // 结束时间戳(毫秒)
	PhoneNumber     string `json:"phoneNumber" form:"phoneNumber"`   // 手机号
	Email           string `json:"email" form:"email"`               // 邮箱
	UserName        string `json:"username" form:"username"`         // 用户名
	UserType        int    `json:"userType" form:"userType"`         // 用户类型
	OrderId         string `json:"orderId" form:"orderId"`           //订单号
	RootUserId      int    `json:"rootUserId" form:"rootUserId"`     //根用户
	ParentId        int    `json:"parentUserId" form:"parentUserId"` //上级用户
	request.PageInfo
}
