package captcha

type Captcha struct {
	Id           uint   `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:10;"`                    //id字段
	Uid          uint   `json:"uid" form:"uid" gorm:"column:user_id;comment:用户id;size:10;"`                    //用户id
	UserType     uint   `json:"userType" form:"userType" gorm:"column:user_type;comment:用户类型 1:普通用户 2: 试玩用户;"` //用户类型 1:普通用户 2: 试玩用户
	ChannelType  uint   `json:"channelType" form:"channelType" gorm:"column:channel_type;comment:发送类型：1 手机 2 Email"`
	UniqueId     string `json:"uniqueId" form:"uniqueId" gorm:"column:unique_id;comment:邮箱/手机号"`
	NationalCode string `json:"countryCode" form:"countryCode" gorm:"column:national_code;comment:邮箱/手机号"`
	CaptchaCode  string `json:"captchaCode" form:"captchaCode" gorm:"column:captcha_code;comment:邮箱/手机号"`
	AccessIp     string `json:"accessIp" form:"accessIp" gorm:"column:access_ip;comment:邮箱/手机号"`
	CreatedAt    int64  `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:;"` //createdAt字段
	UpdatedAt    int64  `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at" `          //更新时间
}

// TableName coupon表 Coupon自定义表名 coupon
func (Captcha) TableName() string {
	return "sys_captcha_log"
}
