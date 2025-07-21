// 自动生成模板Symbols
package symbol

import (
	"time"
)

type StockType int

const (
	Stock  StockType = 0
	Crypto StockType = 1
	Forex  StockType = 2
)

// symbols表 结构体  Symbols
type Symbols struct {
	Id             *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:主键;size:20;"`                    //主键
	Symbol         string     `json:"symbol" form:"symbol" gorm:"column:symbol;comment:股票代码，通常是股票的唯一标识;size:10;"`      //股票代码，通常是股票的唯一标识
	Corporation    string     `json:"corporation" form:"corporation" gorm:"column:corporation;comment:公司名称;size:255;"` //公司名称
	Industry       string     `json:"industry" form:"industry" gorm:"column:industry;comment:行业;size:255;"`            //行业
	Exchange       string     `json:"exchange" form:"exchange" gorm:"column:exchange;comment:所属交易所;size:50;"`          //所属交易所
	MarketCap      *int       `json:"marketCap" form:"marketCap" gorm:"column:market_cap;comment:市值;size:19;"`         //市值
	ListDate       *time.Time `json:"listDate" form:"listDate" gorm:"column:list_date;comment:上市日期;"`                  //上市日期
	Description    string     `json:"description" form:"description" gorm:"column:description;comment:公司描述;type:text;"`
	SicDescription string     `json:"sicDescription" form:"sicDescription" gorm:"column:sic_description;comment:公司所属行业描述;type:text;"`
	CurrentPrice   *float64   `json:"currentPrice" form:"currentPrice" gorm:"column:current_price;comment:最新价格;size:10;"`     //最新价格
	AverageVolume  *int       `json:"averageVolume" form:"averageVolume" gorm:"column:average_volume;comment:日均交易量;size:19;"` //日均交易量
	ChangeRatio    *float64   `json:"changeRatio" form:"changeRatio" gorm:"column:change_ratio;comment:日涨跌幅度;size:10;"`       //日涨跌幅度
	PeRatio        *float64   `json:"peRatio" form:"peRatio" gorm:"column:pe_ratio;comment:市盈率;size:10;"`                     //市盈率
	CreatedAt      *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:创建时间;"`                      //创建时间
	UpdatedAt      *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;comment:更新时间;"`                      //更新时间
	Icon           string     `json:"icon" form:"icon" gorm:"column:icon;comment:股票图标;size:255;"`                             //股票图标
	Type           *int       `json:"type" form:"type" gorm:"column:type;comment:资产类型(0:股票 1:加密货币 2:外汇);size:4;"`
	Status         *int       `json:"status" form:"status" gorm:"column:status;comment:启用状态(0:关闭 1:启用);default:1;size:4;"`      //启用状态
	TicketSize     *float64   `json:"ticketSize" form:"ticketSize" gorm:"column:ticket_size;comment:价格精度;default:0.01;"`        //价格精度
	TicketNumSize  *float64   `json:"ticketNumSize" form:"ticketNumSize" gorm:"column:ticket_num_size;comment:数量精度;default:1;"` //数量精度
	Sort           *int       `json:"sort" form:"sort" gorm:"column:sort;comment:排序权重;default:0;size:10;"`                      //排序权重
}

// TableName symbols表 Symbols自定义表名 symbols
func (Symbols) TableName() string {
	return "symbols"
}
