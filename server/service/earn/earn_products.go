package earn

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/earn"
	earnReq "github.com/flipped-aurora/gin-vue-admin/server/model/earn/request"
	"time"
)

type EarnProductsService struct{}

// CreateEarnProducts 创建earnProducts表记录
// Author [yourname](https://github.com/yourname)
func (earnProductsService *EarnProductsService) CreateEarnProducts(earnProducts *earn.EarnProducts) (err error) {
	earnProducts.UpdatedAt = time.Now()
	earnProducts.CreatedAt = time.Now().UnixMilli()
	earnProducts.Stock = -1
	err = global.GVA_DB.Create(earnProducts).Error
	return err
}

// DeleteEarnProducts 删除earnProducts表记录
// Author [yourname](https://github.com/yourname)
func (earnProductsService *EarnProductsService) DeleteEarnProducts(id string) (err error) {
	err = global.GVA_DB.Delete(&earn.EarnProducts{}, "id = ?", id).Error
	return err
}

// DeleteEarnProductsByIds 批量删除earnProducts表记录
// Author [yourname](https://github.com/yourname)
func (earnProductsService *EarnProductsService) DeleteEarnProductsByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]earn.EarnProducts{}, "id in ?", ids).Error
	return err
}

// UpdateEarnProducts 更新earnProducts表记录
// Author [yourname](https://github.com/yourname)
func (earnProductsService *EarnProductsService) UpdateEarnProducts(earnProducts earn.EarnProducts) (err error) {
	earnProducts.UpdatedAt = time.Now()
	err = global.GVA_DB.Model(&earn.EarnProducts{}).Where("id = ?", earnProducts.Id).Save(&earnProducts).Error
	return err
}

// GetEarnProducts 根据id获取earnProducts表记录
// Author [yourname](https://github.com/yourname)
func (earnProductsService *EarnProductsService) GetEarnProducts(id string) (earnProducts earn.EarnProducts, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&earnProducts).Error
	return
}

// GetEarnProducts 根据id获取earnProducts表记录
// Author [yourname](https://github.com/yourname)
func (earnProductsService *EarnProductsService) GetEarnProductsList(id []uint) (earnProducts []*earn.EarnProducts, err error) {
	err = global.GVA_DB.Where("id in(?)", id).Find(&earnProducts).Error
	return
}

// GetEarnProductsInfoList 分页获取earnProducts表记录
// Author [yourname](https://github.com/yourname)
func (earnProductsService *EarnProductsService) GetEarnProductsInfoList(info earnReq.EarnProductsSearch) (list []earn.EarnProducts, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&earn.EarnProducts{})
	var earnProductss []earn.EarnProducts
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Order("id desc").Find(&earnProductss).Error
	return earnProductss, total, err
}

// GetEarnProductsInfoList 分页获取earnProducts表记录
// Author [yourname](https://github.com/yourname)
func (earnProductsService *EarnProductsService) GetFrontEarnProductsInfoList(info earnReq.EarnProductsSearch) (list []earn.EarnProducts, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&earn.EarnProducts{})
	var earnProductss []earn.EarnProducts
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Where("status = ?", earn.Available).Order("id desc").Find(&earnProductss).Error
	return earnProductss, total, err
}
func (earnProductsService *EarnProductsService) GetEarnProductsPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
