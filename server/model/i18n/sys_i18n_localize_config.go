// 自动生成模板SysI18nLocalizeConfig
package i18n

// sysI18nLocalizeConfig表 结构体  SysI18nLocalizeConfig
type SysI18nLocalizeConfig struct {
	Id           *int   `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:10;"`               //id字段
	LangTag      string `json:"langTag" form:"langTag" gorm:"column:lang_tag;comment:;size:48;"`          //langTag字段
	MessageId    string `json:"messageId" form:"messageId" gorm:"column:message_id;comment:;size:100;"`   //messageId字段
	TemplateData string `json:"templateData" form:"templateData" gorm:"column:template_data;comment:;"`   //templateData字段
	CategoryId   *int   `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:;size:10;"` //categoryId字段
	ErrorCode    *int   `json:"errorCode" form:"errorCode" gorm:"column:error_code;comment:错误码;size:19;"` //错误码
	CreatedAt    int64  `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:;"`            //createdAt字段
	UpdatedAt    int64  `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;comment:;"`            //updatedAt字段
}

// TableName sysI18nLocalizeConfig表 SysI18nLocalizeConfig自定义表名 sys_i18n_localize_config
func (SysI18nLocalizeConfig) TableName() string {
	return "sys_i18n_localize_config"
}
