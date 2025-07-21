package order

import (
	"database/sql"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/symbol"
	"github.com/shopspring/decimal"
	"time"
)

type StockOrderStatus uint

// 导师开仓单状态：1-持仓中;5-已完结
const (
	// StockOrderStatusHolding 开单状态: 持仓中
	StockOrderStatusHolding = 1
	// StockOrderStatusFinished 开单状态: 已完结
	StockOrderStatusFinished = 5
)

type AdvisorStockOrder struct {
	global.GVA_MODEL
	AdvisorId uint                `gorm:"column:advisor_id"`
	StockId   uint                `gorm:"column:stock_id"`
	BuyPrice  decimal.Decimal     `gorm:"column:buy_price"`
	BuyTime   time.Time           `gorm:"column:buy_time"`
	SellPrice decimal.NullDecimal `gorm:"column:sell_price"`
	SellTime  sql.NullTime        `gorm:"column:sell_time"`
	Status    StockOrderStatus    `gorm:"column:status"`
	Advisor   Advisor             `gorm:"foreignKey:AdvisorId"`
	Stock     symbol.Symbols      `gorm:"foreignKey:StockId"`
}

func (AdvisorStockOrder) TableName() string {
	return "advisor_stock_order"
}
