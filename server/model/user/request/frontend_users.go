package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	"time"
)

type FrontendUsersSearch struct {
	user.FrontendUsers
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	UserIDList     string     `json:"userIDList" form:"userIDList"`
	UserType       int        `json:"userType" form:"userType"`
	RootId         uint       `json:"rootId" form:"rootId"`
	request.PageInfo
}

type SubUserReq struct {
	TeamOwner uint `json:"teamOwner" form:"teamOwner"` // 团队所属 userID
	UserID    uint `json:"userID" form:"userID"`       // 团队成员的userID
	request.PageInfo
}

type TeamReq struct {
	TeamOwner uint `json:"teamOwner" form:"teamOwner"` // 团队所属 userID
}

// UserLogin login structure
type UserLogin struct {
	UserName  string `json:"username"`  // 邮箱
	Email     string `json:"email"`     // 邮箱
	CountryId uint   `json:"countryId"` // 手机号国家码
	Phone     string `json:"phone"`     // 手机号
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
	Password  string `json:"password"`  // 密码
}

// BindEmailReq structure
type BindEmailReq struct {
	ID        uint   `json:"-"`                             // 从 JWT 中提取 user id，避免越权
	Email     string `json:"email" description:"邮箱"`        // 邮箱
	Captcha   string `json:"captcha" description:"验证码"`     // 验证码
	CaptchaId string `json:"captchaId" description:"验证码ID"` // 验证码ID
}

// BindPhoneReq structure
type BindPhoneReq struct {
	CountryId uint   `json:"countryId"`                     // 手机号国家码
	Phone     string `json:"phone" description:"手机号"`       // 手机号
	Captcha   string `json:"captcha" description:"验证码"`     // 验证码
	CaptchaId string `json:"captchaId" description:"验证码ID"` // 验证码ID
}

type RealNameAuthenticationReq struct {
	ID     uint `json:"id"`     // user id
	Status int  `json:"status"` // 2： 审核拒绝 3： 审核通过
}

type ChangeParentReq struct {
	ID       uint `json:"id"`       // user id
	ParentID uint `json:"parentId"` // 上级代理 ID
}

type GetAncestorsReq struct {
	ID uint `json:"id" form:"id"` // user id
}

type UpdateUserInfoReq struct {
	NickName  string `json:"nickName" example:"昵称"`  // 昵称
	HeaderImg string `json:"headerImg" example:"头像"` // 头像
}

type UpdateUserPassword struct {
	UserID          uint   `json:"userId"`                         // userID
	Password        string `json:"passWord" example:"密码"`          // 登录密码
	PaymentPassword string `json:"paymentPassword" example:"交易密码"` // 交易密码
}

// UserRegister User register structure
type UserRegister struct {
	CountryId       uint   `json:"countryId" swaggertype:"integer" example:"1"`
	Phone           string `json:"phone" example:"手机号码"`
	Email           string `json:"email" example:"电子邮箱"`
	Password        string `json:"passWord" example:"密码"`
	InviteCode      string `json:"inviteCode" example:"xxxxx"`
	IdType          *int   `json:"idType" example:"1"` // 1. 身份证 2. 护照
	RealName        string `json:"realName" example:"真实姓名"`
	IdNumber        string `json:"idNumber" example:"证件号码"`
	IdImages        string `json:"idImages" example:"证件图片"` // 身份证或者护照的图片，用逗号隔开
	PaymentPassword string `json:"paymentPassword" example:"交易密码"`
	Captcha         string `json:"captcha"`   // 验证码
	CaptchaId       string `json:"captchaId"` // 验证码ID
}
type UserIdentityReq struct {
	CountryId uint   `json:"countryId" swaggertype:"integer" example:"1"` // 国家
	IdType    *int   `json:"idType" example:"1"`                          // 1. 身份证 2. 护照
	RealName  string `json:"realName" example:"真实姓名"`                     // 真实姓名
	IdNumber  string `json:"idNumber" example:"证件号码"`                     // 证件号码
	IdImages  string `json:"idImages" example:"证件图片"`                     // 身份证或者护照的图片，用逗号隔开
}

type ChangePasswordReq struct {
	ID          uint   `json:"-"`           // 从 JWT 中提取 user id，避免越权
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
	Captcha     string `json:"captcha"`
	CaptchaId   string `json:"captchaId"`
	CountryId   uint   `json:"countryId" swaggertype:"integer" example:"1"`
	Phone       string `json:"phone" example:"手机号码"`
	Email       string `json:"email" example:"电子邮箱"`
}
