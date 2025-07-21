// 自动生成模板SysI18nDictCategory
package i18n
import (
	"time"
)

// sysI18nDictCategory表 结构体  SysI18nDictCategory
type SysI18nDictCategory struct {
    Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;comment:分类ID;size:20;"`  //分类ID 
    ParentCategoryId  *int `json:"parentCategoryId" form:"parentCategoryId" gorm:"column:parent_category_id;comment:父分类ID;size:10;"`  //父分类ID 
    CategoryName  string `json:"categoryName" form:"categoryName" gorm:"column:category_name;comment:分类名称;size:255;"`  //分类名称 
    CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:;"`  //createdAt字段 
    UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;comment:;"`  //updatedAt字段 
    DeletedAt  *time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;comment:;"`  //deletedAt字段 
}


// TableName sysI18nDictCategory表 SysI18nDictCategory自定义表名 sys_i18n_dict_category
func (SysI18nDictCategory) TableName() string {
    return "sys_i18n_dict_category"
}

