// 自动生成模板FrontendUsers
package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/userfund"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

// frontendUsers表 结构体  FrontendUsers
type FrontendUsers struct {
	global.GVA_MODEL
	UUid                 uuid.UUID                  `json:"uuid" gorm:"column:uuid;comment:用户UUID"`                                                                                       //用户UUID
	Username             string                     `json:"username" form:"username" gorm:"column:username;comment:用户登录名;size:191;"`                                                      //用户登录名
	Password             string                     `json:"password" form:"password" gorm:"column:password;comment:用户登录密码;size:191;"`                                                     //用户登录密码
	PaymentPassword      string                     `json:"paymentPassword" form:"paymentPassword" gorm:"column:payment_password;comment:用户支付密码;size:200;"`                               //用户支付密码
	NickName             string                     `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:用户昵称;size:191;"`                                                      //用户昵称
	HeaderImg            string                     `json:"headerImg" form:"headerImg" gorm:"column:header_img;comment:用户头像;size:191;"`                                                   //用户头像
	Phone                string                     `json:"phone" form:"phone" gorm:"column:phone;comment:用户手机号;size:191;"`                                                               //用户手机号
	Email                string                     `json:"email" form:"email" gorm:"column:email;comment:用户邮箱;size:191;"`                                                                //用户邮箱
	Enable               int                        `json:"enable" form:"enable" gorm:"column:enable;comment:用户是否被冻结 1正常 2冻结;size:19;"`                                                   //用户是否被冻结 1正常 2冻结
	ParentId             uint                       `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:直接上级用户ID;size:20;"`                                                   //直接上级用户ID
	ParentUserId         uint                       `json:"parentUserId" form:"parentId" gorm:"column:parent_id;comment:直接上级用户ID;size:20;"`                                               //直接上级用户ID
	GrandparentId        uint                       `json:"grandparentId" form:"grandparentId" gorm:"column:grandparent_id;comment:上级的上级用户ID;size:20;"`                                   //上级的上级用户ID
	GreatGrandparentId   uint                       `json:"greatGrandparentId" form:"greatGrandparentId" gorm:"column:great_grandparent_id;comment:上级的上级的上级用户ID;size:20;"`                //上级的上级的上级用户ID
	CountryId            uint                       `json:"countryId" form:"countryId" gorm:"column:country_id;comment:用户手机号国家码;size:20;"`                                                //用户手机号国家码
	IdType               int                        `json:"idType" form:"idType" gorm:"column:id_type;comment:证件类型：1-身份证，2-护照;"`                                                          //证件类型：1-身份证，2-护照
	IdImages             string                     `json:"idImages" form:"idImages" gorm:"column:id_images;comment:证件照;size:500;"`                                                       //证件照
	AuthenticationStatus *int                       `json:"authenticationStatus" form:"authenticationStatus" gorm:"column:authentication_status;comment:实名状态：0-未实名，1-待审核，2-未通过， 3-审核通过;"` //实名状态：0-未实名，1-待审核， 3-审核通过
	RealName             string                     `json:"realName" form:"realName" gorm:"column:real_name;comment:真实姓名;size:100;"`                                                      //真实姓名
	IdNumber             string                     `json:"idNumber" form:"idNumber" gorm:"column:id_number;comment:证件号码;size:200;"`                                                      //证件号码
	LastLoginIp          string                     `json:"lastLoginIp" form:"lastLoginIp" gorm:"column:last_login_ip;comment:最近登录ip;size:200;"`                                          //证件号码
	RootUserid           uint                       `json:"rootUserid" form:"rootUserid" gorm:"column:root_userid;comment:根用户id;size:20;"`
	RootUserId           uint                       `json:"rootUserId" form:"rootUserId" gorm:"column:root_userid;comment:根用户id;size:20;"`
	UserFund             *userfund.UserFundAccounts `json:"userFund" form:"userFund" gorm:"foreignKey:UserId;references:ID;comment:用户资金"`
	UserType             uint                       `json:"userType" form:"userType" gorm:"column:user_type;comment:用户类型：1-普通用户，2-试玩用户;"`             //用户类型：1-普通用户，2-试玩用户
	LastLoginTime        uint64                     `json:"lastLoginTime" form:"lastLoginTime" gorm:"column:last_login_time;comment:最近登录时间;size:20;"` //最近登录时间
	CreateAtInt          int64                      `json:"createAtInt" gorm:"-"`
	InviteCode           string                     `json:"inviteCode" gorm:"-"`
}

// TableName frontendUsers表 FrontendUsers自定义表名 frontend_users
func (FrontendUsers) TableName() string {
	return "frontend_users"
}

func (s *FrontendUsers) GetUsername() string {
	return s.Username
}

func (s *FrontendUsers) GetNickname() string {
	return s.NickName
}

func (s *FrontendUsers) GetUUID() uuid.UUID {
	return s.UUid
}

func (s *FrontendUsers) GetUserId() uint {
	return s.ID
}
func (s *FrontendUsers) GetFrontUserId() uint {
	return s.ID
}

func (s *FrontendUsers) GetAuthorityId() uint {
	return 0
}

func (s *FrontendUsers) GetUserInfo() any {
	return *s
}

func (s *FrontendUsers) GetUserType() uint {
	return s.UserType
}

func (u *FrontendUsers) AfterFind(tx *gorm.DB) (err error) {
	// 查询后自动赋值
	u.ParentUserId = u.ParentId
	u.RootUserId = u.RootUserid
	u.CreateAtInt = u.CreatedAt.UnixMilli()
	u.InviteCode = utils.EncryptID(uint64(u.ID))
	return nil
}
