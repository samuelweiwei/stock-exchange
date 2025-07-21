// 自动生成模板FrontendUserLoginLog
package user

// frontendUserLoginLog表 结构体  FrontendUserLoginLog
type FrontendUserLoginLog struct {
	Id          *int   `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:10;"`                                      //id字段
	Uid         uint   `json:"uid" form:"uid" gorm:"primarykey;column:uid;comment:用户id;size:10;"`                               //用户id
	IsSameIp    int    `json:"isSameIp" form:"isSameIp" gorm:"primarykey;column:is_same_ip;comment:是否不同用户同ip 0:否 1:是;size:10;"` //是否不同用户同ip 0:否 1:是
	LoginIp     string `json:"loginIp" form:"loginIp" gorm:"primarykey;column:login_ip;comment:最近登录ip;size:128;"`               //最近登录ip
	LoginRegion string `json:"loginRegion" form:"loginRegion" gorm:"primarykey;column:login_region;comment:最近登录地区;size:128;"`   //最近登录地区
	LoginTime   int64  `json:"loginTime" form:"loginTime" gorm:"primarykey;column:login_time;comment:最近登录时间戳（毫秒）;size:20;"`     //最近登录时间戳（毫秒）
	UserAgent   string `json:"userAgent" form:"userAgent" gorm:"primarykey;column:user_agent;comment:登录设备;size:1024;"`          //登录设备
}

// TableName frontendUserLoginLog表 FrontendUserLoginLog自定义表名 frontend_user_login_log
func (FrontendUserLoginLog) TableName() string {
	return "frontend_user_login_log"
}
