package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/order"

// AdvisorDetail 导师详情
// @Description 导师详情
type AdvisorDetail struct {
	Id                 uint               `json:"id"`                 //导师ID
	NickName           string             `json:"nickName"`           //导师昵称
	Duty               string             `json:"duty"`               //导师职务
	Intro              string             `json:"intro"`              //导师简介
	AvatarUrl          string             `json:"avatarUrl"`          //导师头像地址
	Exp                *int               `json:"exp"`                //导师从业时间，单位年
	SevenDayReturn     *float64           `json:"sevenDayReturn"`     //导师7日盈亏
	SevenDayReturnRate *float64           `json:"sevenDayReturnRate"` //导师7日回报率
	ActiveStatus       order.ActiveStatus `json:"activeStatus"`       //导师启用状态： 0-禁用；1-启用
	ActiveStatusText   string             `json:"activeStatusText"`   //导师启用状态文案
}

// AdvisorPageData 导师分页数据
// @Description 导师分页数据
type AdvisorPageData struct {
	AdvisorDetail
}

// AdvisorOption 导师选项
// @Description 导师选项
type AdvisorOption struct {
	Id       uint   `json:"id"`       //导师ID
	NickName string `json:"nickName"` //导师昵称
}
