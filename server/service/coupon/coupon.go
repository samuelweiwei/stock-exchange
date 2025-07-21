package coupon

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/coupon"
	couponReq "github.com/flipped-aurora/gin-vue-admin/server/model/coupon/request"
	"github.com/go-playground/validator/v10"
	"time"
)

type CouponService struct{}

// CreateCoupon 创建coupon表记录
// Author [yourname](https://github.com/yourname)
func (cpService *CouponService) CreateCoupon(cp *coupon.Coupon) (err error) {
	if err = validator.New().Struct(cp); err != nil {
		return err
	}
	cp.CreatedAt = time.Now().UnixMilli()
	cp.UpdatedXAt = time.Now().UnixMilli()
	err = global.GVA_DB.Create(cp).Error
	return err
}

// DeleteCoupon 删除coupon表记录
// Author [yourname](https://github.com/yourname)
func (cpService *CouponService) DeleteCoupon(id string) (err error) {
	err = global.GVA_DB.Delete(&coupon.Coupon{}, "id = ?", id).Error
	return err
}

// DeleteCouponByIds 批量删除coupon表记录
// Author [yourname](https://github.com/yourname)
func (cpService *CouponService) DeleteCouponByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]coupon.Coupon{}, "id in ?", ids).Error
	return err
}

// UpdateCoupon 更新coupon表记录
// Author [yourname](https://github.com/yourname)
func (cpService *CouponService) UpdateCoupon(cp coupon.Coupon) (err error) {
	cp.UpdatedXAt = time.Now().UnixMilli()
	err = global.GVA_DB.Model(&coupon.Coupon{}).Where("id = ?", cp.Id).Save(&cp).Error
	return err
}

// GetCoupon 根据id获取coupon表记录
// Author [yourname](https://github.com/yourname)
func (cpService *CouponService) GetCoupon(id string) (cp coupon.Coupon, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&cp).Error
	return
}

// GetCouponInfoList 分页获取coupon表记录
// Author [yourname](https://github.com/yourname)
func (cpService *CouponService) GetCouponInfoList(info couponReq.CouponSearch) (list []coupon.Coupon, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&coupon.Coupon{}).Order("id desc")

	if coupon.Type(info.Type) == coupon.Manual {
		db = db.Where("type = ?", info.Type)
	}
	var cps []coupon.Coupon
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&cps).Error
	for i, item := range cps {
		if item.ValidDays != 0 {
			cps[i].Period = append(cps[i].Period, int64(item.ValidDays))
		} else {
			cps[i].Period = append(cps[i].Period, item.ValidStart, item.ValidEnd)
		}
	}
	return cps, total, err
}
func (cpService *CouponService) GetCouponPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
