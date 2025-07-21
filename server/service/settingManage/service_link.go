package settingManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/settingManage"
	settingManageReq "github.com/flipped-aurora/gin-vue-admin/server/model/settingManage/request"
)

type ServiceLinkService struct{}

// CreateServiceLink 创建serviceLink表记录
// Author [yourname](https://github.com/yourname)
func (serviceLinkService *ServiceLinkService) CreateServiceLink(serviceLink *settingManageReq.ServiceLinkCreate) (err error) {
	err = global.GVA_DB.Create(serviceLink).Error
	return err
}

// DeleteServiceLink 删除serviceLink表记录
// Author [yourname](https://github.com/yourname)
func (serviceLinkService *ServiceLinkService) DeleteServiceLink(id string) (err error) {
	err = global.GVA_DB.Delete(&settingManage.ServiceLink{}, "id = ?", id).Error
	return err
}

// DeleteServiceLinkByIds 批量删除serviceLink表记录
// Author [yourname](https://github.com/yourname)
func (serviceLinkService *ServiceLinkService) DeleteServiceLinkByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]settingManage.ServiceLink{}, "id in ?", ids).Error
	return err
}

// UpdateServiceLink 更新serviceLink表记录
// Author [yourname](https://github.com/yourname)
func (serviceLinkService *ServiceLinkService) UpdateServiceLink(serviceLink settingManageReq.ServiceLinkUpdate) (err error) {
	err = global.GVA_DB.Model(&settingManage.ServiceLink{}).Where("id = ?", serviceLink.Id).Updates(&serviceLink).Error
	return err
}

// GetServiceLink 根据id获取serviceLink表记录
// Author [yourname](https://github.com/yourname)
func (serviceLinkService *ServiceLinkService) GetServiceLink(id string) (serviceLink settingManage.ServiceLink, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&serviceLink).Error
	return
}

// GetServiceLinkInfoList 分页获取serviceLink表记录
// Author [yourname](https://github.com/yourname)
func (serviceLinkService *ServiceLinkService) GetServiceLinkInfoList(info settingManageReq.ServiceLinkSearch) (list []settingManage.ServiceLink, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&settingManage.ServiceLink{})
	if info.Status != nil {
		db.Where(" `status` = ? ", &info.Status)
	}
	var serviceLinks []settingManage.ServiceLink
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&serviceLinks).Error
	return serviceLinks, total, err
}
func (serviceLinkService *ServiceLinkService) GetServiceLinkPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
