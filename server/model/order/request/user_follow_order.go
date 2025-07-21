package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order"
)

// UserFollowOrderApplyReq 用户跟单申请请求
// @Description 用户跟单申请请求
type UserFollowOrderApplyReq struct {
	UserId         uint                  `json:"-"`
	AdvisorProdId  uint                  `json:"advisorProdId"`  //导师产品ID
	FollowAmount   float64               `json:"followAmount"`   //跟单金额
	AutoRenew      order.AutoRenewStatus `json:"autoRenew"`      //是否自动续期：0-否；1-是
	CouponRecordId uint                  `json:"couponRecordId"` //使用优惠券发放记录ID
}

// MyFollowOrderPageReq 我的跟单列表分页查询参数
// @Description 我的跟单列表分页查询参数
type MyFollowOrderPageReq struct {
	request.PageInfo
	FollowOrderStatus *order.FollowOrderStockStatus `form:"followOrderStatus"` //跟单状态：0-审核中;3-已驳回;5-跟单中;8-已撤回;10-已结束
	CreatedAtStart    *int64                        `form:"createdAt"`         //创建时间搜索开始值，Unix时间戳
	CreatedAtEnd      *int64                        `form:"createdAtEnd"`      //创建时间搜索结束值，Unix时间戳
	AdvisorId         *uint                         `form:"advisorId"`         //导师ID
	UserId            uint                          `json:"-"`
}

// RetrieveFollowOrderReq 用户提取盈利请求
// @Description 用户提取盈利请求
type RetrieveFollowOrderReq struct {
	UserId         uint    `json:"-"`
	FollowOrderId  uint    `json:"-"`
	RetrieveAmount float64 `json:"retrieveAmount"` //提取金额
}

// UserFollowOrderPageQueryReq 用户跟单列表分页查询参数
// @Description 用户跟单列表分页查询参数
type UserFollowOrderPageQueryReq struct {
	request.PageInfo
	FollowOrderId     *string                       `form:"followOrderId"`     //跟单ID
	ProductName       *string                       `form:"productName"`       //产品名称
	FollowOrderStatus *order.FollowOrderStatus      `form:"followOrderStatus"` //跟单状态
	StockStatus       *order.FollowOrderStockStatus `form:"stockOrderStatus"`  //跟单持仓状态
	AdvisorId         *uint                         `form:"advisorId"`         //导师ID
	UserType          *uint                         `form:"userType"`          //用户类型: 1-普通用户;2-试玩用户
	UserId            *uint                         `form:"userId"`            //客户ID
	UserPhone         *string                       `form:"userPhone"`         //客户手机号
	UserEmail         *string                       `form:"userEmail"`         //客户邮箱
	ApplyTimeStart    *int64                        `form:"applyTimeStart"`    //申请时间开始值
	ApplyTimeEnd      *int64                        `form:"applyTimeEnd"`      //申请时间结束值
	RootUserIdSearch  *uint                         `form:"rootUserId"`        //根用户ID
	ParentUserId      *uint                         `form:"parentUserId"`      //上级用户ID
	RootUserId        uint                          `form:"-"`                 //代理用户ID
}

// UserFollowOrderConfirmSubmitReq 用户跟单确认页提交请求
// @Description 用户跟单确认页提交请求
type UserFollowOrderConfirmSubmitReq struct {
	StockOrderId  string  `json:"stockOrderId"` //导师开单记录ID
	UserId        string  `json:"userId"`       //用户ID
	StockNum      float64 `json:"stockNum"`     //股票数量
	FollowOrderId uint    `json:"-"`
}

// UserFollowOrderRetrieveRecordPageReq 用户跟单提现记录分页查询参数
// @Description 用户跟单提现记录分页查询参数
type UserFollowOrderRetrieveRecordPageReq struct {
	request.PageInfo
	FollowOrderId uint `json:"-"` //跟单订单ID
	UserId        uint `json:"-"`
}
