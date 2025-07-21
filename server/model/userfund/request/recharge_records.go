package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type RechargeRecordsSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	RechargeType   string     `json:"rechargeType" form:"rechargeType" `
	OrderId        string     `json:"orderId" form:"orderId" `          //订单号查询
	MemberId       int        `json:"memberId" form:"memberId" `        //会员ID
	MemberPhone    string     `json:"memberPhone" form:"memberPhone" `  //订单号查询
	MemberEmail    string     `json:"memberEmail" form:"memberEmail" `  //订单号查询
	Currency       string     `json:"currency" form:"currency" `        //币种
	IsLock         string     `json:"is_lock" form:"isLock"`            //渠道类型，如银行、第三方、钱包
	OrderStatus    string     `json:"orderStatus" form:"orderStatus"`   //订单状态1 支付中 2 成功 3 失败 4 超时
	UserAction     string     `json:"userAction" form:"userAction"`     //用户行为 1 请求提现 2 撤销
	ReviewStatus   string     `json:"reviewStatus" form:"reviewStatus"` //审核状态
	FromAddress    string     `json:"fromAddress" form:"fromAddress"`
	ToAddress      string     `json:"toAddress" form:"toAddress"`
	ParentId       int        `json:"parentUserId" form:"parentUserId"` //会员ID，标识提现会员
	RootId         int        `json:"rootUserId" form:"rootUserId" `    //根用户ID

	StartRechargeTime int64 `json:"startRechargeTime" form:"startRechargeTime"` // 开始时间戳(毫秒)
	EndRechargeTime   int64 `json:"endRechargeTime" form:"endRechargeTime"`     // 结束时间戳(毫秒)

	StartUpdateTime int64 `json:"startUpdateTime" form:"startUpdateTime"` // 开始时间戳(毫秒)
	EndUpdateTime   int64 `json:"endUpdateTime" form:"endUpdateTime"`     // 结束时间戳(毫秒)
	UserType        int   `json:"userType" form:"userType"`
	request.PageInfo
}

// UserRechargeRecordsSearch 用户充值记录查询结构体
type UserRechargeRecordsSearch struct {
	Page     int `form:"page" json:"page" binding:"required"`         // 页码
	PageSize int `form:"pageSize" json:"pageSize" binding:"required"` // 每页大小
}
