// 自动生成模板UserFundAccounts
package userfund

import (
	. "github.com/shopspring/decimal"
	"time"
)

// userFundAccounts表 结构体  UserFundAccounts
type UserFundAccounts struct {
	Id                *int        `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:10;"`                         //id字段
	UserId            int         `json:"userId" form:"userId" gorm:"column:user_id;comment:;"`                               //userId字段
	AssetType         string      `json:"assetType" form:"assetType" gorm:"column:asset_type;comment:;size:50;"`              //assetType字段
	Balance           Decimal     `json:"balance" form:"balance" gorm:"column:balance;comment:;"`                             //balance字段
	FrozenBalance     Decimal     `json:"frozenBalance" form:"frozenBalance" gorm:"column:frozen_balance;comment:;"`          //frozenBalance字段
	AvailableBalance  Decimal     `json:"availableBalance" form:"availableBalance" gorm:"column:available_balance;comment:;"` //availableBalance字段
	STATUS            string      `json:"STATUS" form:"STATUS" gorm:"column:STATUS;comment:;size:1;"`                         //STATUS字段
	CreatedAt         *time.Time  `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:;"`                      //createdAt字段
	UpdatedAt         *time.Time  `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;comment:;"`                      //updatedAt字段
	DeletedAt         *time.Time  `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;comment:;"`                      //deletedAt字段
	FirstChargeAmount NullDecimal `json:"firstChargeAmount" form:"firstChargeAmount" gorm:"column:first_charge_amount;comment:;"`
	FirstChargeTime   *time.Time  `json:"firstChargeTime" form:"firstChargeTime" gorm:"column:first_charge_time;comment:;"`
	UserType          uint        `json:"userType" form:"userType" gorm:"column:user_type;comment:用户类型：1-普通用户，2-试玩用户;"` //用户类型：1-普通用户，2-试玩用户
}

type RechargeRequest struct {
	RechargeChannelId int     `json:"rechargeChannelId" binding:"required"` // 充值渠道ID，必填
	RechargeAmount    Decimal `json:"rechargeAmount" binding:"required"`    // 充值金额，保留两位小数
	FromAddress       string  `json:"fromAddress" binding:"omitempty"`
	JumpUrl           string  `json:"jumpUrl" binding:"omitempty"`
	Remark            string  `json:"remark" binding:"omitempty"` // 备注，非必填
}

type WithdrawRequest struct {
	WithdrawChannelId int     `json:"withdrawChannelId" binding:"required"` // 提现渠道ID，必填
	WithdrawAmount    Decimal `json:"withdrawAmount" binding:"required"`    // 提现金额，保留两位小数
	ToAddress         string  `json:"toAddress" binding:"required"`
	Remark            string  `json:"remark" binding:"omitempty"` // 备注，非必填
	PaymentPassword   string  `json:"paymentPassword" binding:"required"`
}

// TableName userFundAccounts表 UserFundAccounts自定义表名 user_fund_accounts
func (UserFundAccounts) TableName() string {
	return "user_fund_accounts"
}
