package order

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/errorx"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/coupon"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order"
	orderReq "github.com/flipped-aurora/gin-vue-admin/server/model/order/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order/response"
	. "github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type AdvisorProdService struct {
}

// CreateProduct 创建导师产品
func (advisorProdService *AdvisorProdService) CreateProduct(req *orderReq.AdvisorProdCreateReq) error {
	if err := validate(req); err != nil {
		return err
	}
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		advisorProd := &order.AdvisorProd{
			AdvisorId:      req.AdvisorId,
			ProductName:    req.ProductName,
			AmountStep:     req.AmountStep,
			MaxAmount:      req.MaxAmount,
			MinAmount:      req.MinAmount,
			FollowPeriod:   req.FollowPeriod,
			CommissionRate: NewFromFloat(req.CommissionRate),
			AutoRenew:      req.AutoRenew,
			ActiveStatus:   req.ActiveStatus,
		}

		err := tx.Create(advisorProd).Error
		if err != nil {
			return err
		}

		if len(req.CouponIdList) > 0 {
			prodCouponList := make([]*order.AdvisorProdCoupon, len(req.CouponIdList))
			for i := 0; i < len(req.CouponIdList); i++ {
				prodCouponList[i] = &order.AdvisorProdCoupon{
					AdvisorProdId: advisorProd.ID,
					CouponId:      req.CouponIdList[i],
				}
			}
			err = tx.Create(&prodCouponList).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (advisorProdService *AdvisorProdService) UpdateProduct(req *orderReq.AdvisorProdUpdateReq) error {
	if err := validate(&req.AdvisorProdCreateReq); err != nil {
		return err
	}
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		advisorProd := &order.AdvisorProd{
			GVA_MODEL: global.GVA_MODEL{
				ID: req.Id,
			},
			AdvisorId:      req.AdvisorId,
			ProductName:    req.ProductName,
			AmountStep:     req.AmountStep,
			MaxAmount:      req.MaxAmount,
			MinAmount:      req.MinAmount,
			FollowPeriod:   req.FollowPeriod,
			CommissionRate: NewFromFloat(req.CommissionRate),
			AutoRenew:      req.AutoRenew,
			ActiveStatus:   req.ActiveStatus,
		}

		err := tx.Omit("CreatedAt").Save(advisorProd).Error
		if err != nil {
			return err
		}

		err = tx.Where("advisor_prod_id", req.Id).Delete(&order.AdvisorProdCoupon{}).Error
		if err != nil {
			return err
		}
		if len(req.CouponIdList) > 0 {
			prodCouponList := make([]*order.AdvisorProdCoupon, len(req.CouponIdList))
			for i := 0; i < len(req.CouponIdList); i++ {
				prodCouponList[i] = &order.AdvisorProdCoupon{
					AdvisorProdId: advisorProd.ID,
					CouponId:      req.CouponIdList[i],
				}
			}
			err = tx.Create(&prodCouponList).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (advisorProdService *AdvisorProdService) DeleteProduct(productId uint) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Delete(&order.AdvisorProd{}, productId).Error
		if err != nil {
			return err
		}

		err = tx.Where("advisor_prod_id", productId).Delete(&order.AdvisorProdCoupon{}).Error
		if err != nil {
			return err
		}
		return nil
	})
}

func (advisorProdService *AdvisorProdService) PageQuery(req *orderReq.AdvisorProdPageQueryReq) ([]*response.AdvisorProdPageData, int64, error) {
	db := global.GVA_DB.Model(&order.AdvisorProd{})
	var total int64
	if err := db.Count(&total).Error; err != nil || total == 0 {
		return nil, 0, err
	}

	var prodList []*order.AdvisorProd
	err := db.Scopes(req.Paginate()).Preload("Advisor").Order("advisor_prod.updated_at desc").Find(&prodList).Error
	if err != nil {
		return nil, 0, err
	}

	pageDataList := make([]*response.AdvisorProdPageData, len(prodList))
	for i, item := range prodList {
		pageDataList[i] = &response.AdvisorProdPageData{
			AdvisorProdId:  item.ID,
			AdvisorName:    item.Advisor.NickName,
			ProductName:    item.ProductName,
			FollowPeriod:   item.FollowPeriod,
			CommissionRate: item.CommissionRate.InexactFloat64(),
			MaxAmount:      item.MaxAmount,
			MinAmount:      item.MinAmount,
			AmountStep:     item.AmountStep,
			AutoRenew:      item.AutoRenew,
			ActiveStatus:   item.ActiveStatus,
		}
	}
	return pageDataList, total, nil
}

func (advisorProdService *AdvisorProdService) GetAdvisorProdById(advisorProdId uint) (*response.AdvisorProdDetail, error) {
	var advisorProd order.AdvisorProd
	err := global.GVA_DB.Preload("AdvisorProdCoupons").First(&advisorProd, advisorProdId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errorx.NewWithCode(errorx.AdvisorProdNotFound)
	} else if err != nil {
		return nil, err
	}

	resp := &response.AdvisorProdDetail{
		AdvisorProdId:  advisorProd.ID,
		AdvisorId:      advisorProd.AdvisorId,
		ProductName:    advisorProd.ProductName,
		FollowPeriod:   advisorProd.FollowPeriod,
		CommissionRate: advisorProd.CommissionRate.InexactFloat64(),
		MaxAmount:      advisorProd.MaxAmount,
		MinAmount:      advisorProd.MinAmount,
		AmountStep:     advisorProd.AmountStep,
		AutoRenew:      advisorProd.AutoRenew,
		ActiveStatus:   advisorProd.ActiveStatus,
	}
	if len(advisorProd.AdvisorProdCoupons) > 0 {
		couponIds := make([]uint, len(advisorProd.AdvisorProdCoupons))
		for i, v := range advisorProd.AdvisorProdCoupons {
			couponIds[i] = v.CouponId
		}
		resp.CouponIdList = couponIds
	}
	return resp, nil
}

func (advisorProdService *AdvisorProdService) ListByAdvisorId(advisorId uint, userId uint) ([]*response.AdvisorProdListData, error) {
	var advisorProds []*order.AdvisorProd
	err := global.GVA_DB.Preload("AdvisorProdCoupons").Where("advisor_id = ? and active_status = ?", advisorId, order.Active).Find(&advisorProds).Error
	if err != nil {
		return nil, err
	}

	couponIds := make([]uint, 0)
	for _, v := range advisorProds {
		if len(v.AdvisorProdCoupons) > 0 {
			for _, c := range v.AdvisorProdCoupons {
				couponIds = append(couponIds, c.CouponId)
			}
		}
	}
	couponMap := make(map[uint][]coupon.CouponIssued)
	if len(couponIds) > 0 {
		var couponIssueList []coupon.CouponIssued
		err2 := global.GVA_DB.Model(&coupon.CouponIssued{}).Where("user_id = ? and coupon_id in (?)", userId, couponIds).Find(&couponIssueList).Error
		if err2 == nil {
			for _, v := range couponIssueList {
				couponIssues, ok := couponMap[uint(*v.CouponId)]
				if !ok {
					couponIssues = make([]coupon.CouponIssued, 0)
				}
				if v.ValidStart <= time.Now().UnixMilli() && v.ValidEnd > time.Now().UnixMilli() && v.Status == coupon.NotUsed {
					couponIssues = append(couponIssues, v)
					couponMap[uint(*v.CouponId)] = couponIssues
				}
			}
		}
	}

	list := make([]*response.AdvisorProdListData, len(advisorProds))
	for i, item := range advisorProds {
		list[i] = &response.AdvisorProdListData{
			AdvisorProdId:  item.ID,
			ProductName:    item.ProductName,
			FollowPeriod:   item.FollowPeriod,
			CommissionRate: item.CommissionRate.InexactFloat64(),
			MaxAmount:      item.MaxAmount,
			MinAmount:      item.MinAmount,
			AutoRenew:      item.AutoRenew,
		}

		if len(item.AdvisorProdCoupons) > 0 {
			usableCoupons := make([]*response.AdvisorProdUsableCoupon, 0)
			for _, c := range item.AdvisorProdCoupons {
				couponIssues, ok := couponMap[c.CouponId]
				if ok {
					for _, d := range couponIssues {
						usableCoupons = append(usableCoupons, &response.AdvisorProdUsableCoupon{
							CouponRecordId: uint(*d.Id),
							CouponName:     d.CouponName,
						})
					}
				}
			}
			list[i].UsableCoupons = usableCoupons
		}
	}

	return list, err
}

func validate(req *orderReq.AdvisorProdCreateReq) error {
	err := global.GVA_DB.First(&order.Advisor{}, req.AdvisorId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errorx.NewWithCode(errorx.AdvisorNotFound)
	} else {
		return err
	}
}
