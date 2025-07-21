package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/shopspring/decimal"
)

type EarnDailyIncomeMoneyLogSearch struct {
	Uid              string `p:"uid" dc:"会员ID" form:"uid" json:"Uid,omitempty"`
	UserType         uint   `json:"userType,omitempty" form:"userType" dc:"用户类型 1:普通用户 2: 试玩用户"`
	SuperiorId       string `p:"superiorId" form:"rootUserId" dc:"最上级推荐id" json:"SuperiorId,omitempty"`
	ParentId         string `p:"parentId"  form:"parentUserId" dc:"上线推荐人id" json:"ParentId,omitempty"`
	Phone            string `p:"phone" form:"phone" dc:"用户手机号" json:"Phone,omitempty"`
	Email            string `p:"email" form:"email"  dc:"邮箱" json:"Email,omitempty"`
	BeginTime        uint   `p:"beginTime" form:"beginTime"  dc:"起始时间戳(秒)" json:"BeginTime,omitempty"`
	EndTime          uint   `p:"endTime" form:"endTime" dc:"结束时间戳(秒)" json:"EndTime,omitempty"`
	request.PageInfo `json:"Request.PageInfo"`
}

type EarnProductsIncome struct {
	Phone         string          `json:"phone"`
	Email         string          `json:"email"`
	Uid           uint            `json:"uid"`
	UserType      uint            `json:"userType"`
	StakeNum      uint            `json:"stakeNum"`
	StakeAmount   decimal.Decimal `json:"stakeAmount"`
	StakeEarnings decimal.Decimal `json:"stakeEarnings"`
	ParentId      uint            `json:"parentUserId"` // 推荐人ID
	SuperiorId    uint            `json:"rootUserId"`   // 一级推荐人ID
}

type EarnProductsIncomeSummaryRes struct {
	List             []*EarnProductsIncome `json:"list"`
	TotalStakeNum    int64                 `json:"totalStakeNum" dc:"总质押笔数"`
	TotalStakeAmount decimal.Decimal       `json:"totalStakeAmount" dc:"总质押金额"`
	TotalStakeProfit decimal.Decimal       `json:"totalStakeProfit" dc:"盈亏总额"`
}

type EarnDailyIncomeMoneyLogDetailSearch struct {
	SubscriptionId uint `p:"subscriptionId" dc:"理财订阅id" form:"subscriptionId"`
	request.PageInfo
}
