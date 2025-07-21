package request

import "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"

type CaptchaListSearch struct {
	request.PageInfo
	UserId      int    `json:"userId" form:"userId"`           //用户id
	UniqueId    string `json:"uniqueId" form:"uniqueId"`       // 邮件/手机号
	CountryCode string `json:"countryCode" form:"countryCode"` // 国家或地区码
	UserType    int    `json:"userType" form:"userType"`       //用户类型 1:普通用户 2:试玩用户
	CaptchaCode string `json:"captchaCode" form:"captchaCode"` // 验证码
	ChannelType int    `json:"channelType" form:"channelType"` // 发送渠道  1:手机 2:邮件
	AccessIp    string `json:"accessIp" form:"accessIp"`       // 发送请求IP
}
