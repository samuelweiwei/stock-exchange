package request

import "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"

// SystemOrderMsgPageQueryReq 订单站内信列表分页查询请求
// @Description 订单站内信列表分页查询请求
type SystemOrderMsgPageQueryReq struct {
	request.PageInfo
}

// SystemOrderMsgSetReadStatusReq 系统订单站内信已读状态设置请求
// @Description 系统订单站内信已读状态设置请求
type SystemOrderMsgSetReadStatusReq struct {
	MsgIds []int64 `json:"msgIds"` //消息ID集合
}
