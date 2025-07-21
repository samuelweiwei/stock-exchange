package earn

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/earn"
	earnReq "github.com/flipped-aurora/gin-vue-admin/server/model/earn/request"
	"time"
)

type EarnInterestRatesService struct{}

// CreateEarnInterestRates 创建earnInterestRates表记录
// Author [yourname](https://github.com/yourname)
func (earnInterestRatesService *EarnInterestRatesService) CreateEarnInterestRates(earnInterestRates *earn.EarnInterestRates) (err error) {
	err = global.GVA_DB.Create(earnInterestRates).Error
	return err
}

// BatchCreateEarnInterestRates 创建earnInterestRates表记录
// Author [yourname](https://github.com/yourname)
func (earnInterestRatesService *EarnInterestRatesService) BatchCreateEarnInterestRates(earnInterestRates []*earn.EarnInterestRates) (err error) {
	err = global.GVA_DB.CreateInBatches(earnInterestRates, 10000).Error
	return err
}

// DeleteEarnInterestRates 删除earnInterestRates表记录
// Author [yourname](https://github.com/yourname)
func (earnInterestRatesService *EarnInterestRatesService) DeleteEarnInterestRates(id string) (err error) {
	err = global.GVA_DB.Delete(&earn.EarnInterestRates{}, "id = ?", id).Error
	return err
}

// DeleteEarnInterestRatesByIds 批量删除earnInterestRates表记录
// Author [yourname](https://github.com/yourname)
func (earnInterestRatesService *EarnInterestRatesService) DeleteEarnInterestRatesByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]earn.EarnInterestRates{}, "id in ?", ids).Error
	return err
}

// UpdateEarnInterestRates 更新earnInterestRates表记录
// Author [yourname](https://github.com/yourname)
func (earnInterestRatesService *EarnInterestRatesService) RandomEarnInterestRate(earnInterestRates earn.EarnInterestRates) (err error) {
	err = global.GVA_DB.Model(&earn.EarnInterestRates{}).Where("id = ?", earnInterestRates.Id).Updates(&earnInterestRates).Error
	return err
}

// UpdateEarnInterestRates 更新earnInterestRates表记录
// Author [yourname](https://github.com/yourname)
func (earnInterestRatesService *EarnInterestRatesService) UpdateEarnInterestRates(earnInterestRates earn.EarnInterestRates) (err error) {
	err = global.GVA_DB.Model(&earn.EarnInterestRates{}).Where("id = ?", earnInterestRates.Id).Updates(&earnInterestRates).Error
	return err
}

// GetEarnInterestRates 根据id获取earnInterestRates表记录
// Author [yourname](https://github.com/yourname)
func (earnInterestRatesService *EarnInterestRatesService) GetEarnInterestRates(id string) (earnInterestRates earn.EarnInterestRates, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&earnInterestRates).Error
	return
}

// GetPeriodEarnInterestRates 根据id获取earnInterestRates表记录
// Author [yourname](https://github.com/yourname)
func (earnInterestRatesService *EarnInterestRatesService) GetPeriodEarnInterestRates(productId uint, period time.Time) (earnInterestRates *earn.EarnInterestRates, err error) {
	err = global.GVA_DB.Where("product_id = ? and period = ?", productId, period).First(&earnInterestRates).Error
	return
}

// GetCurrentPeriodEarnInterestRates 根据id获取earnInterestRates表记录
// Author [yourname](https://github.com/yourname)
func (earnInterestRatesService *EarnInterestRatesService) GetCurrentPeriodEarnInterestRates(productId []uint, period time.Time) (earnInterestRates []*earn.EarnInterestRates, err error) {
	err = global.GVA_DB.Where("product_id in (?) and period = ?", productId, period).Find(&earnInterestRates).Error
	return
}

// GetEarnInterestRatesInfoList 分页获取earnInterestRates表记录
// Author [yourname](https://github.com/yourname)
func (earnInterestRatesService *EarnInterestRatesService) GetEarnInterestRatesInfoList(info earnReq.EarnInterestRatesSearch) (list []earn.EarnInterestRates, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&earn.EarnInterestRates{})
	var earnInterestRatess []earn.EarnInterestRates
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&earnInterestRatess).Error
	return earnInterestRatess, total, err
}
func (earnInterestRatesService *EarnInterestRatesService) GetEarnInterestRatesPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
