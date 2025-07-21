package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/order"

// MyFollowAppendOrderListData 我的追单列表数据
// @Description 我的追单列表数据
type MyFollowAppendOrderListData struct {
	AppendOrderNo         string                  `json:"appendOrderNo"`         //追单号
	AppendAmount          float64                 `json:"appendAmount"`          //追单金额
	AppendOrderStatus     order.AppendOrderStatus `json:"appendOrderStatus"`     //追单状态
	AppendOrderStatusText string                  `json:"appendOrderStatusText"` //追单状态文案
	CreatedAt             int64                   `json:"createdAt"`             //创建时间，Unix时间戳
}

// UserFollowAppendOrderPageData 用户追单分页数据
// @Description 用户追单分页数据
type UserFollowAppendOrderPageData struct {
	AppendOrderId         string                  `json:"appendOrderId"`         //追单ID
	AppendOrderNo         string                  `json:"appendOrderNo"`         //追单号
	FollowOrderId         string                  `json:"followOrderId"`         //跟单ID
	UserId                string                  `json:"userId"`                //客户ID
	UserType              uint                    `json:"userType"`              //客户类型
	UserTypeText          string                  `json:"userTypeText"`          //客户类型文案
	Phone                 string                  `json:"phone"`                 //客户手机号
	AppendAmount          float64                 `json:"appendAmount"`          //追单金额
	ProductName           string                  `json:"productName"`           //产品名称
	CreatedAt             int64                   `json:"createdAt"`             //创建时间
	AppendOrderStatus     order.AppendOrderStatus `json:"appendOrderStatus"`     //追单状态
	AppendOrderStatusText string                  `json:"appendOrderStatusText"` //追单状态文案
	UserEmail             string                  `json:"userEmail"`             //客户邮箱
	AdvisorName           string                  `json:"advisorName"`           //分析师名称
	RootUserId            uint                    `json:"rootUserId"`            //根用户ID
	ParentUserId          uint                    `json:"parentUserId"`          //上级用户ID
}
