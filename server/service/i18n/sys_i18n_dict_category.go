package i18n

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/i18n"
    i18nReq "github.com/flipped-aurora/gin-vue-admin/server/model/i18n/request"
)

type SysI18nDictCategoryService struct {}
// CreateSysI18nDictCategory 创建sysI18nDictCategory表记录
// Author [yourname](https://github.com/yourname)
func (sysI18nDictCategoryService *SysI18nDictCategoryService) CreateSysI18nDictCategory(sysI18nDictCategory *i18n.SysI18nDictCategory) (err error) {
	err = global.GVA_DB.Create(sysI18nDictCategory).Error
	return err
}

// DeleteSysI18nDictCategory 删除sysI18nDictCategory表记录
// Author [yourname](https://github.com/yourname)
func (sysI18nDictCategoryService *SysI18nDictCategoryService)DeleteSysI18nDictCategory(id string) (err error) {
	err = global.GVA_DB.Delete(&i18n.SysI18nDictCategory{},"id = ?",id).Error
	return err
}

// DeleteSysI18nDictCategoryByIds 批量删除sysI18nDictCategory表记录
// Author [yourname](https://github.com/yourname)
func (sysI18nDictCategoryService *SysI18nDictCategoryService)DeleteSysI18nDictCategoryByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]i18n.SysI18nDictCategory{},"id in ?",ids).Error
	return err
}

// UpdateSysI18nDictCategory 更新sysI18nDictCategory表记录
// Author [yourname](https://github.com/yourname)
func (sysI18nDictCategoryService *SysI18nDictCategoryService)UpdateSysI18nDictCategory(sysI18nDictCategory i18n.SysI18nDictCategory) (err error) {
	err = global.GVA_DB.Model(&i18n.SysI18nDictCategory{}).Where("id = ?",sysI18nDictCategory.Id).Updates(&sysI18nDictCategory).Error
	return err
}

// GetSysI18nDictCategory 根据id获取sysI18nDictCategory表记录
// Author [yourname](https://github.com/yourname)
func (sysI18nDictCategoryService *SysI18nDictCategoryService)GetSysI18nDictCategory(id string) (sysI18nDictCategory i18n.SysI18nDictCategory, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysI18nDictCategory).Error
	return
}

// GetSysI18nDictCategoryInfoList 分页获取sysI18nDictCategory表记录
// Author [yourname](https://github.com/yourname)
func (sysI18nDictCategoryService *SysI18nDictCategoryService)GetSysI18nDictCategoryInfoList(info i18nReq.SysI18nDictCategorySearch) (list []i18n.SysI18nDictCategory, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&i18n.SysI18nDictCategory{})
    var sysI18nDictCategorys []i18n.SysI18nDictCategory
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&sysI18nDictCategorys).Error
	return  sysI18nDictCategorys, total, err
}
func (sysI18nDictCategoryService *SysI18nDictCategoryService)GetSysI18nDictCategoryPublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
