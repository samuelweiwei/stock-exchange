package request

import "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"

// MyUserProfitShareRecordPageQueryReq 我的用户分成记录分页查询接口
//
// @Description 我的用户分成记录分页查询接口
type MyUserProfitShareRecordPageQueryReq struct {
	request.PageInfo
	UserId uint `json:"-"` //用户ID
}
