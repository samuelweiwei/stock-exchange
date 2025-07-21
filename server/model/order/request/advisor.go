package request

import (
	common "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order"
)

// AdvisorCreateReq
// @Description  导师创建请求
type AdvisorCreateReq struct {
	NickName           string   `json:"nickName"`           //导师昵称
	Duty               string   `json:"duty"`               //导师职务
	Intro              string   `json:"intro"`              //导师简介
	AvatarUrl          string   `json:"avatarUrl"`          //导师头像
	Exp                *int     `json:"exp"`                //导师从业时间，单位年
	SevenDayReturn     *float64 `json:"sevenDayReturn"`     //导师7日盈亏
	SevenDayReturnRate *float64 `json:"sevenDayReturnRate"` //导师7日回报率，单位%
}

// AdvisorUpdateReq 导师信息更新请求
// @Description	导师信息更新请求
type AdvisorUpdateReq struct {
	Id                 int      `json:"-"`
	NickName           string   `json:"nickName"`           //导师昵称
	Duty               string   `json:"duty"`               //导师职务
	Intro              string   `json:"intro"`              //导师简介
	AvatarUrl          string   `json:"avatarUrl"`          //导师头像
	Exp                *int     `json:"exp"`                //导师从业时间，单位年
	SevenDayReturn     *float64 `json:"sevenDayReturn"`     //导师7日盈亏
	SevenDayReturnRate *float64 `json:"sevenDayReturnRate"` //导师7日回报率,单位%
}

// AdvisorStatusUpdateReq 导师启用状态更新请求
// @Description 导师启用状态更新请求
type AdvisorStatusUpdateReq struct {
	Id           int                `json:"-"`
	ActiveStatus order.ActiveStatus `json:"activeStatus"` //启用状态：0-禁用；1-启用
}

// AdvisorPageQueryReq 导师列表分页查询请求
// @Description 导师列表分页查询请求
type AdvisorPageQueryReq struct {
	common.PageInfo
	ActiveStatus *order.ActiveStatus `form:"activeStatus"` //启用状态
}

// AdvisorListQueryReq 导师列表查询请求
//
// @Description 导师列表查询请求
type AdvisorListQueryReq struct {
	ActiveStatus *order.ActiveStatus `form:"activeStatus"` //启用状态
}
