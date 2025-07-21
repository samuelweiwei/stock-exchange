// 自动生成模板SymbolsHot
package symbol

import (
	"time"
)

// symbolsHot表 结构体  SymbolsHot
type SymbolsHot struct {
	Id        *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:主键;size:20;"`                //主键
	SymbolId  *int       `json:"symbolId" form:"symbolId" gorm:"column:symbol_id;comment:股票ID;size:19;"`      //股票ID
	CreatedAt *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:创建时间;"`           //创建时间
	UpdatedAt *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;comment:更新时间;"`           //更新时间
	Sort      *int       `json:"sort" form:"sort" gorm:"column:sort;comment:排序权重(值越大越靠前);default:0;size:10;"` //排序权重
}

// TableName symbolsHot表 SymbolsHot自定义表名 symbols_hot
func (SymbolsHot) TableName() string {
	return "symbols_hot"
}
