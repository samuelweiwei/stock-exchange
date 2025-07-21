// 自动生成模板SymbolsCustom
package symbol

import (
	"time"
)

// symbolsCustom表 结构体  SymbolsCustom
type SymbolsCustom struct {
	Id        *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:主键;size:20;"`           //主键
	UserId    *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID;size:19;"`       //用户ID
	SymbolId  *int       `json:"symbolId" form:"symbolId" gorm:"column:symbol_id;comment:股票ID;size:19;"` //股票ID
	CreatedAt *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:创建时间;"`      //创建时间
	UpdatedAt *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;comment:更新时间;"`      //更新时间
}

// TableName symbolsCustom表 SymbolsCustom自定义表名 symbols_custom
func (SymbolsCustom) TableName() string {
	return "symbols_custom"
}
