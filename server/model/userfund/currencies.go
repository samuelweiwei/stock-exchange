// 自动生成模板Currencies
package userfund

import (
	. "github.com/shopspring/decimal"
	"time"
)

// currencies表 结构体  Currencies
type Currencies struct {
	Id            *int    `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:10;"`                               //id字段
	Currency      string  `json:"currency" form:"currency" gorm:"column:currency;comment:货币名称;size:255;"`                   //货币名称
	Icon          string  `json:"icon" form:"icon" gorm:"column:icon;comment:货币图标;size:255;"`                               //货币图标
	PriceUsdt     Decimal `json:"priceUsdt" form:"priceUsdt" gorm:"column:price_usdt;comment:货币价格;size:10;"`                //货币价格
	CoinType      *int    `json:"coinType" form:"coinType" gorm:"column:coin_type;comment:1:数字货币。2:法币;size:10;"`            //1:数字货币。2:法币
	TicketSize    Decimal `json:"ticketSize" form:"ticketSize" gorm:"column:ticket_size;comment:价格进度;size:10,8;"`           //价格进度
	TicketNumSize Decimal `json:"ticketNumSize" form:"ticketNumSize" gorm:"column:ticket_num_size;comment:数量精度;size:10,8;"` //数量精度
	//MinNum         Decimal    `json:"minNum" form:"minNum" gorm:"column:min_num;comment:可充值/提现最小数量;size:20,8;"`                 //可充值/提现最小数量
	MinWithdrawNum Decimal    `json:"minWithdrawNum" form:"minWithdrawNum" gorm:"column:min_withdraw_num;comment:可提现最小数量;size:20,8;"`
	MinRechargeNum Decimal    `json:"minRechargeNum" form:"minRechargeNum" gorm:"column:min_recharge_num;comment:可充值最小数量;size:20,8;"`
	CreatedAt      *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:;"` //createdAt字段
	UpdatedAt      *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;comment:;"` //updatedAt字段
	DeletedAt      *time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;comment:;"` //deletedAt字段

	CreateAtInt  int64 `json:"createAtInt" gorm:"-"`
	UpdatedAtInt int64 `json:"updateAtInt" gorm:"-"`
}

// TableName currencies表 Currencies自定义表名 currencies
func (Currencies) TableName() string {
	return "currencies"
}
