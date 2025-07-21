// 自动生成模板UserAccountFlow
package userfund

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	. "github.com/shopspring/decimal"
	"time"
)

// userAccountFlow表 结构体  UserAccountFlow
type UserAccountFlow struct {
	global.GVA_MODEL
	UserId int `json:"userId" form:"userId" gorm:"column:user_id;comment:;size:10;"` //userId字段

	ParentId int `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:父级用户ID;"` //会员ID，标识提现会员
	RootId   int `json:"rootId" form:"rootId" gorm:"column:root_id;comment:root_id;"`      //根用户ID

	TransactionType     string  `json:"transactionType" form:"transactionType" gorm:"column:transaction_type;comment:1 充值 2 提现 3 划转到合约账户 4 合约账户划转进来;"` //1 充值 2 提现 3 划转到合约账户 4 合约账户划转进来
	Amount              Decimal `json:"amount" form:"amount" gorm:"column:amount;comment:;size:10;"`
	BalanceBefore       Decimal `json:"balanceBefore" form:"balanceBefore" gorm:"column:balance_before;comment:;size:10;"`                    //balanceBefore字段
	BalanceAfter        Decimal `json:"balanceAfter" form:"balanceAfter" gorm:"column:balance_after;comment:;size:10;"`                       //balanceAfter字段
	FrozenBalanceBefore Decimal `json:"frozenBalanceBefore" form:"frozenBalanceBefore" gorm:"column:frozen_balance_before;comment:;size:10;"` //frozenbalanceBefore字段
	FrozenBalanceAfter  Decimal `json:"frozenBalanceAfter" form:"frozenBalanceAfter" gorm:"column:frozen_balance_after;comment:;size:10;"`    //frozenbalanceAfter字段
	AvaBalanceBefore    Decimal `json:"avaBalanceBefore" form:"avaBalanceBefore" gorm:"column:ava_balance_before;comment:;size:10;"`          //frozenbalanceBefore字段
	AvaBalanceAfter     Decimal `json:"avaBalanceAfter" form:"avaBalanceAfter" gorm:"column:ava_balance_after;comment:;size:10;"`             //frozenbalanceAfter字段
	TotalCommission     Decimal `json:"totalCommission" form:"totalCommission" gorm:"column:total_commission;comment:提现总共手续费;"`               //手续费

	TransactionDate time.Time `json:"transactionDate" form:"transactionDate" gorm:"column:transaction_date;comment:;"` //transactionDate字段
	Description     string    `json:"description" form:"description" gorm:"column:description;comment:;size:255;"`     //description字段
	OrderId         string    `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单id;size:19;"`             //订单id
	UserType        uint      `json:"userType" form:"userType" gorm:"column:user_type;comment:用户类型：1-普通用户，2-试玩用户;"`    //用户类型：1-普通用户，2-试玩用户

	PhoneNumber string `json:"phoneNumber" gorm:"-"`
	Email       string `json:"email" gorm:"-"`
	Username    string `json:"username" gorm:"-"`
}

// userAccountFlow表 结构体  UserAccountFlow
type UserAccountFlowUnion struct {
	global.GVA_MODEL
	UserId              int     `json:"userId" form:"userId" gorm:"column:user_id;comment:;size:10;"`                                                  //userId字段
	TransactionType     string  `json:"transactionType" form:"transactionType" gorm:"column:transaction_type;comment:1 充值 2 提现 3 划转到合约账户 4 合约账户划转进来;"` //1 充值 2 提现 3 划转到合约账户 4 合约账户划转进来
	Amount              Decimal `json:"amount" form:"amount" gorm:"column:amount;comment:;size:10;"`
	BalanceBefore       Decimal `json:"balanceBefore" form:"balanceBefore" gorm:"column:balance_before;comment:;size:10;"`                    //balanceBefore字段
	BalanceAfter        Decimal `json:"balanceAfter" form:"balanceAfter" gorm:"column:balance_after;comment:;size:10;"`                       //balanceAfter字段
	FrozenBalanceBefore Decimal `json:"frozenBalanceBefore" form:"frozenBalanceBefore" gorm:"column:frozen_balance_before;comment:;size:10;"` //frozenbalanceBefore字段
	FrozenBalanceAfter  Decimal `json:"frozenBalanceAfter" form:"frozenBalanceAfter" gorm:"column:frozen_balance_after;comment:;size:10;"`    //frozenbalanceAfter字段
	AvaBalanceBefore    Decimal `json:"avaBalanceBefore" form:"avaBalanceBefore" gorm:"column:ava_balance_before;comment:;size:10;"`          //frozenbalanceBefore字段
	AvaBalanceAfter     Decimal `json:"avaBalanceAfter" form:"avaBalanceAfter" gorm:"column:ava_balance_after;comment:;size:10;"`             //frozenbalanceAfter字段
	TotalCommission     Decimal `json:"totalCommission" form:"totalCommission" gorm:"column:total_commission;comment:提现总共手续费;"`               //手续费

	TransactionDate time.Time `json:"transactionDate" form:"transactionDate" gorm:"column:transaction_date;comment:;"` //transactionDate字段
	Description     string    `json:"description" form:"description" gorm:"column:description;comment:;size:255;"`     //description字段
	OrderId         string    `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单id;size:19;"`             //订单id

	ParentId int `json:"parentUserId" form:"parentUserId" gorm:"column:parent_id;comment:父级用户ID;"` //会员ID，标识提现会员
	RootId   int `json:"rootUserId" form:"rootUserId" gorm:"column:root_id;comment:root_id;"`      //根用户ID

	PhoneNumber         string `json:"phoneNumber" gorm:"column:phone_number;comment:手机号;size:19;"`
	Email               string `json:"email" gorm:"column:email;comment:邮箱;size:19;"`
	Username            string `json:"username" gorm:"column:username;comment:姓名;size:19;"`
	UserType            uint   `json:"userType" form:"userType" gorm:"column:user_type;comment:用户类型：1-普通用户，2-试玩用户;"`
	TransactionDateInt  int64  `json:"transactionDateInt" gorm:"-"`
	TransactionTypeI18n string `json:"TransactionTypeI18n" gorm:"-"`
}

// TableName userAccountFlow表 UserAccountFlow自定义表名 user_account_flow
func (UserAccountFlow) TableName() string {
	return "user_account_flow"
}
