package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

// UserFollowAppendOrderApplyReq 用户申请追单请求
// @Description 用户申请追单请求
type UserFollowAppendOrderApplyReq struct {
	FollowOrderId uint    `json:"-"`
	AppendAmount  float64 `json:"appendAmount"` //跟单金额
	UserId        uint    `json:"-"`
}

// UserFollowAppendOrderPageQueryReq 分页查询用户追单列表请求
// @Description 分页查询永辉追单列表请求
type UserFollowAppendOrderPageQueryReq struct {
	request.PageInfo
	Phone             *string `form:"phone"`             //手机号码
	UserId            *string `form:"userId"`            //用户ID
	UserType          *uint   `form:"userType"`          //用户类型
	FollowOrderId     *string `form:"followOrderId"`     //跟单ID
	AppendOrderStatus *string `form:"appendOrderStatus"` //追单状态
	UserEmail         *string `form:"userEmail"`         //用户邮箱
	AdvisorName       *string `form:"advisorName"`       //分析师名称
	UserPhone         *string `form:"userPhone"`         //用户手机号
	AppendOrderNo     *string `form:"appendOrderNo"`     //追单号
	RootUserIdSearch  *uint   `form:"rootUserId"`        //根用户ID
	ParentUserId      *uint   `form:"parentUserId"`      //上级用户ID
	ProductName       *string `form:"productName"`       //产品名称
	RootUserId        uint    `form:"-"`                 //上级代理用户ID
}
