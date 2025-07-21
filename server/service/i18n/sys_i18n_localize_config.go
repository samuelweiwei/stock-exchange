package i18n

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/i18n"
	i18nReq "github.com/flipped-aurora/gin-vue-admin/server/model/i18n/request"
	"time"
)

type SysI18nLocalizeConfigService struct{}

// CreateSysI18nLocalizeConfig 创建sysI18nLocalizeConfig表记录
// Author [yourname](https://github.com/yourname)
func (sysI18nLocalizeConfigService *SysI18nLocalizeConfigService) CreateSysI18nLocalizeConfig(sysI18nLocalizeConfig *i18n.SysI18nLocalizeConfig) (err error) {
	var (
		now = time.Now()
	)
	if sysI18nLocalizeConfig.CreatedAt == 0 {
		sysI18nLocalizeConfig.CreatedAt = now.UnixMilli()
	}
	if sysI18nLocalizeConfig.UpdatedAt == 0 {
		sysI18nLocalizeConfig.UpdatedAt = now.UnixMilli()
	}
	err = global.GVA_DB.Create(sysI18nLocalizeConfig).Error
	return err
}

// DeleteSysI18nLocalizeConfig 删除sysI18nLocalizeConfig表记录
// Author [yourname](https://github.com/yourname)
func (sysI18nLocalizeConfigService *SysI18nLocalizeConfigService) DeleteSysI18nLocalizeConfig(id string) (err error) {
	err = global.GVA_DB.Delete(&i18n.SysI18nLocalizeConfig{}, "id = ?", id).Error
	return err
}

// DeleteSysI18nLocalizeConfigByIds 批量删除sysI18nLocalizeConfig表记录
// Author [yourname](https://github.com/yourname)
func (sysI18nLocalizeConfigService *SysI18nLocalizeConfigService) DeleteSysI18nLocalizeConfigByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]i18n.SysI18nLocalizeConfig{}, "id in ?", ids).Error
	return err
}

// UpdateSysI18nLocalizeConfig 更新sysI18nLocalizeConfig表记录
// Author [yourname](https://github.com/yourname)
func (sysI18nLocalizeConfigService *SysI18nLocalizeConfigService) UpdateSysI18nLocalizeConfig(sysI18nLocalizeConfig i18n.SysI18nLocalizeConfig) (err error) {
	err = global.GVA_DB.Model(&i18n.SysI18nLocalizeConfig{}).Where("id = ?", sysI18nLocalizeConfig.Id).Updates(&sysI18nLocalizeConfig).Error
	return err
}

// GetSysI18nLocalizeConfig 根据id获取sysI18nLocalizeConfig表记录
// Author [yourname](https://github.com/yourname)
func (sysI18nLocalizeConfigService *SysI18nLocalizeConfigService) GetSysI18nLocalizeConfig(id string) (sysI18nLocalizeConfig i18n.SysI18nLocalizeConfig, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysI18nLocalizeConfig).Error
	return
}

// GetSysI18nLocalizeConfigInfoList 分页获取sysI18nLocalizeConfig表记录
// Author [yourname](https://github.com/yourname)
func (sysI18nLocalizeConfigService *SysI18nLocalizeConfigService) GetSysI18nLocalizeConfigInfoList(info i18nReq.SysI18nLocalizeConfigSearch) (list []*i18n.SysI18nLocalizeConfig, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&i18n.SysI18nLocalizeConfig{})
	var sysI18nLocalizeConfigs []*i18n.SysI18nLocalizeConfig
	// 如果有条件搜索 下方会自动创建搜索语句

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	if info.MessageId != "" {
		db = db.Where("message_id = ?", info.MessageId)
	}
	if info.TagLang != "" {
		db = db.Where("lang_tag = ?", info.TagLang)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Find(&sysI18nLocalizeConfigs).Error
	return sysI18nLocalizeConfigs, total, err
}
func (sysI18nLocalizeConfigService *SysI18nLocalizeConfigService) GetSysI18nLocalizeConfigPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
