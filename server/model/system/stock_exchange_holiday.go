package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

type StockExchangeHoliday struct {
	global.GVA_MODEL
	Name     string    `gorm:"column:name"`     //节日名称
	Exchange string    `gorm:"column:exchange"` //交易所名称
	Date     time.Time `gorm:"column:date"`     //节日日期
	Status   string    `gorm:"column:status"`   //开市状态
}

func (StockExchangeHoliday) TableName() string {
	return "stock_exchange_holiday"
}
