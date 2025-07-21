package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	userReq "github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
)

type CountriesService struct{}

// CreateCountries 创建countries表记录
// Author [yourname](https://github.com/yourname)
func (countriesService *CountriesService) CreateCountries(countries *user.Countries) (err error) {
	err = global.GVA_DB.Create(countries).Error
	return err
}

// DeleteCountries 删除countries表记录
// Author [yourname](https://github.com/yourname)
func (countriesService *CountriesService) DeleteCountries(id string) (err error) {
	err = global.GVA_DB.Delete(&user.Countries{}, "id = ?", id).Error
	return err
}

// DeleteCountriesByIds 批量删除countries表记录
// Author [yourname](https://github.com/yourname)
func (countriesService *CountriesService) DeleteCountriesByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]user.Countries{}, "id in ?", ids).Error
	return err
}

// UpdateCountries 更新countries表记录
// Author [yourname](https://github.com/yourname)
func (countriesService *CountriesService) UpdateCountries(countries user.Countries) (err error) {
	err = global.GVA_DB.Model(&user.Countries{}).Where("id = ?", countries.Id).Updates(&countries).Error
	return err
}

// GetCountries 根据id获取countries表记录
// Author [yourname](https://github.com/yourname)
func (countriesService *CountriesService) GetCountries(id string) (countries user.Countries, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&countries).Error
	return
}

// GetCountriesInfoList 分页获取countries表记录
// Author [yourname](https://github.com/yourname)
func (countriesService *CountriesService) GetCountriesInfoList(info userReq.CountriesSearch) (list []user.Countries, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&user.Countries{})
	var countriess []user.Countries
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}

	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}

	if info.PhoneCode != "" {
		db = db.Where("phone_code LIKE ?", "%"+info.PhoneCode+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Find(&countriess).Error
	return countriess, total, err
}
func (countriesService *CountriesService) GetCountriesPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
