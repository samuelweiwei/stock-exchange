// 自动生成模板SysI18nDict
package i18n
import (
	"time"
)

// sysI18nDict表 结构体  SysI18nDict
type SysI18nDict struct {
    Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`  //id字段 
    CategoryId  *int `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:;size:10;"`  //categoryId字段 
    LangTag  string `json:"langTag" form:"langTag" gorm:"column:lang_tag;comment:语言ID;size:10;"`  //语言ID 
    LangKey  string `json:"langKey" form:"langKey" gorm:"column:lang_key;comment:国际化字符键;size:100;"`  //国际化字符键 
    LangValue  string `json:"langValue" form:"langValue" gorm:"column:lang_value;comment:国际化字符值;size:255;"`  //国际化字符值 
    CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:;"`  //createdAt字段 
    UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;comment:;"`  //updatedAt字段 
    DeletedAt  *time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;comment:;"`  //deletedAt字段 
}


// TableName sysI18nDict表 SysI18nDict自定义表名 sys_i18n_dict
func (SysI18nDict) TableName() string {
    return "sys_i18n_dict"
}

