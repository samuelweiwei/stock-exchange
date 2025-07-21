// 自动生成模板Countries
package user
import (
	"time"
)

// countries表 结构体  Countries
type Countries struct {
    Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:10;"`  //id字段 
    CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:;"`  //createdAt字段 
    UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;comment:;"`  //updatedAt字段 
    DeletedAt  *time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;comment:;"`  //deletedAt字段 
    Name  string `json:"name" form:"name" gorm:"column:name;comment:名称;size:200;"`  //名称 
    NameI18nKey  string `json:"nameI18nKey" form:"nameI18nKey" gorm:"column:name_i18n_key;comment:名称多语言key;size:200;"`  //名称多语言key 
    PhotoSide  string `json:"photoSide" form:"photoSide" gorm:"column:photo_side;comment:状态: 1 正面, 2 反面, 3 正面、反面;size:50;"`  //状态: 1 正面, 2 反面, 3 正面、反面 
    PhoneCode  string `json:"phoneCode" form:"phoneCode" gorm:"column:phone_code;comment:手机区号;size:50;"`  //手机区号 
    Status  *bool `json:"status" form:"status" gorm:"column:status;comment:状态: 1 开启, 0 未开启;"`  //状态: 1 开启, 0 未开启 
}


// TableName countries表 Countries自定义表名 countries
func (Countries) TableName() string {
    return "countries"
}

