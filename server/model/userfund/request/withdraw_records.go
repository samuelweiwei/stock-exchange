package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type WithdrawRecordsSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	WithdrawType   string     `json:"withdrawType" form:"withdrawType" `
	OrderId        string     `json:"orderId" form:"orderId" `   //订单号查询
	MemberId       int        `json:"memberId" form:"memberId" ` //会员ID

	ParentId int `json:"parentUserId" form:"parentUserId"` //会员ID，标识提现会员
	RootId   int `json:"rootUserId" form:"rootUserId" `    //根用户ID

	UserType     int    `json:"userType" form:"userType"`
	MemberPhone  string `json:"memberPhone" form:"memberPhone" `  //手机号查询
	MemberEmail  string `json:"memberEmail" form:"memberEmail" `  //邮箱查询
	Currency     string `json:"currency" form:"currency" `        //币种
	IsLock       string `json:"is_lock" form:"isLock"`            //是否锁定
	OrderStatus  string `json:"orderStatus" form:"orderStatus"`   //订单状态1 发起  2 完成
	UserAction   string `json:"userAction" form:"userAction"`     //用户行为 1 请求提现 2 撤销
	ReviewStatus string `json:"reviewStatus" form:"reviewStatus"` //审核状态 1 审核中 2 通过 3 拒绝
	FromAddress  string `json:"fromAddress" form:"fromAddress"`
	ToAddress    string `json:"toAddress" form:"toAddress"`

	StartWithdrawTime int64 `json:"startWithdrawTime" form:"startWithdrawTime"` // 开始时间戳(毫秒)
	EndWithdrawTime   int64 `json:"endWithdrawTime" form:"endWithdrawTime"`     // 结束时间戳(毫秒)

	StartUpdateTime int64 `json:"startUpdateTime" form:"startUpdateTime"` // 开始时间戳(毫秒)
	EndUpdateTime   int64 `json:"endUpdateTime" form:"endUpdateTime"`     // 结束时间戳(毫秒)

	FrontMemberId int `json:"frontMemberId" form:"frontMemberId" ` //root user id
	request.PageInfo
}

// UserWithdrawRecordsSearch 用户提现记录查询结构体
type UserWithdrawRecordsSearch struct {
	Page     int `form:"page" json:"page" binding:"required"`         // 页码
	PageSize int `form:"pageSize" json:"pageSize" binding:"required"` // 每页大小
}
