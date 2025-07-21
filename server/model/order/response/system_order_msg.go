package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/order"

// SystemOrderMsgPageData 系统订单站内信分页数据
// @Description 系统订单站内信分页数据
type SystemOrderMsgPageData struct {
	Id             uint                `json:"id"`             //记录ID
	Type           order.MsgType       `json:"type"`           //站内信类型
	TypeText       string              `json:"typeText"`       //站内信类型文案
	OrderId        uint                `json:"orderId"`        //关联订单ID
	UserId         uint                `json:"userId"`         //客户ID
	Phone          string              `json:"phone"`          //客户手机号
	ReadStatus     order.MsgReadStatus `json:"readStatus"`     //已读状态
	ReadStatusText string              `json:"readStatusText"` //已读状态文案
}
