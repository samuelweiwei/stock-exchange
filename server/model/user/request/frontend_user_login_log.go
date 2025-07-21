package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type FrontendUserLoginLogSearch struct {
	request.PageInfo
	Uid            *uint   `json:"uid" form:"uid"`                       //用户id
	IsSameIp       *int    `json:"isSameIp" form:"isSameIp" `            //是否不同用户同ip 0:否 1:是
	LoginIp        *string `json:"loginIp" form:"loginIp"`               //最近登录ip
	LoginRegion    *string `json:"loginRegion" form:"loginRegion"`       //最近登录地区
	StartLoginTime *int64  `json:"startLoginTime" form:"startLoginTime"` //开始登录时间戳（毫秒）
	EndLoginTime   *int64  `json:"endLoginTime" form:"endLoginTime"`     //结束登录时间戳（毫秒）
	UserAgent      *string `json:"userAgent" form:"userAgent"`           //登录设备
	Phone          *string `json:"phone" form:"phone"`                   // 手机号
	Email          *string `json:"email" form:"email"`                   // 邮箱
	UserName       *string `json:"username" form:"username"`             // 用户名

}
