package coupon

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/coupon"
	couponReq "github.com/flipped-aurora/gin-vue-admin/server/model/coupon/request"
	couponRes "github.com/flipped-aurora/gin-vue-admin/server/model/coupon/response"

	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	"strings"
	"time"
)

type CouponIssuedService struct{}

// CreateCouponIssued 赠送优惠券
// Author [yourname](https://github.com/yourname)
func (couponIssuedService *CouponIssuedService) CreateCouponIssued(req *couponReq.IssueCoupon) (res []string, err error) {
	var (
		c            coupon.Coupon
		users        []*user.FrontendUsers
		IssuedCoupon *coupon.CouponIssued
		userIds      = couponIssuedService.getUsers(req.UserIdList)
		ok           bool
	)

	if err = global.GVA_DB.Where("id in(?)", userIds).Find(&users).Error; err != nil {
		return res, errors.New("UserDoesNotExist")
	}

	if err = global.GVA_DB.Where("id = ?", req.CouponId).First(&c).Error; err != nil {
		return res, errors.New("CouponDoesNotExist")
	}
	if *c.ActiveStatus == false {
		return res, errors.New("CouponDoesNotActive")
	}

	if res, ok = couponIssuedService.isAllUserExist(userIds, users); !ok {
		return res, errors.New("UserDoesNotExist")
	}

	for _, u := range users {
		IssuedCoupon = couponIssuedService.buildIssuedCoupon(&c, u)
		if err = global.GVA_DB.Create(IssuedCoupon).Error; err != nil {
			return res, nil
		}
	}

	return res, err
}

func (couponIssuedService *CouponIssuedService) isAllUserExist(reqUserId []string, users []*user.FrontendUsers) ([]string, bool) {
	var (
		res []string
		m   = make(map[string]*user.FrontendUsers)
	)
	for _, v := range users {
		m[fmt.Sprintf("%v", v.ID)] = v
	}
	for _, v := range reqUserId {
		if _, ok := m[v]; !ok {
			res = append(res, v)
		}
	}
	return res, len(res) == 0
}

func (couponIssuedService *CouponIssuedService) getUsers(uidList string) []string {
	return strings.Split(strings.ReplaceAll(uidList, "，", ","), ",")
}

func (couponIssuedService *CouponIssuedService) buildIssuedCoupon(c *coupon.Coupon, u *user.FrontendUsers) *coupon.CouponIssued {
	var (
		n     = time.Now()
		id    = uint64(u.ID)
		start int64
		end   = n.AddDate(0, 0, c.ValidDays).UnixMilli()
	)
	if c.ValidDays != 0 {
		start = n.UnixMilli()
		end = n.AddDate(0, 0, c.ValidDays).UnixMilli()
	} else {
		start = c.ValidStart
		end = c.ValidEnd
	}

	return &coupon.CouponIssued{
		UserId:       id,
		CouponName:   c.Name,
		CouponId:     c.Id,
		CouponAmount: c.Amount,
		ValidStart:   start,
		ValidEnd:     end,
		CreatedAt:    n.UnixMilli(),
		UpdatedAt:    n.UnixMilli(),
		Type:         c.Type,
		Status:       coupon.NotUsed,
		UserType:     u.UserType,
	}
}

// DeleteCouponIssued 删除couponIssued表记录
// Author [yourname](https://github.com/yourname)
func (couponIssuedService *CouponIssuedService) DeleteCouponIssued(id string) (err error) {
	err = global.GVA_DB.Delete(&coupon.CouponIssued{}, "id = ?", id).Error
	return err
}

// DeleteCouponIssuedByIds 批量删除couponIssued表记录
// Author [yourname](https://github.com/yourname)
func (couponIssuedService *CouponIssuedService) DeleteCouponIssuedByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]coupon.CouponIssued{}, "id in ?", ids).Error
	return err
}

// UpdateCouponIssued 更新couponIssued表记录
// Author [yourname](https://github.com/yourname)
func (couponIssuedService *CouponIssuedService) UpdateCouponIssued(couponIssued coupon.CouponIssued) (err error) {
	err = global.GVA_DB.Model(&coupon.CouponIssued{}).Where("id = ?", couponIssued.Id).Updates(&couponIssued).Error
	return err
}

// UpdateCouponIssued 更新couponIssued表记录
// Author [yourname](https://github.com/yourname)
func (couponIssuedService *CouponIssuedService) Use(couponIssued *couponReq.UseIssuedCoupon) (err error) {
	var (
		status = coupon.AlreadyUsed
	)
	err = global.GVA_DB.Model(&coupon.CouponIssued{}).Select("status").
		Where("id = ? and user_id = ?", couponIssued.IssuedCouponId, couponIssued.UserId).
		Updates(map[string]interface{}{
			"status":    status,
			"update_at": time.Now().UnixMilli(),
		}).Error
	return err
}

// GetCouponIssued 根据id获取couponIssued表记录
// Author [yourname](https://github.com/yourname)
func (couponIssuedService *CouponIssuedService) GetCouponIssued(id string) (couponIssued coupon.CouponIssued, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&couponIssued).Error
	return
}

// AdminGetCouponIssuedInfoList 分页获取couponIssued表记录
// Author [yourname](https://github.com/yourname)
func (couponIssuedService *CouponIssuedService) AdminGetCouponIssuedInfoList(req couponReq.CouponIssuedSearch) (list []*couponRes.CouponIssuedRes, total int64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&coupon.CouponIssued{}).Order("coupon_issued.id desc")
	// 如果有条件搜索 下方会自动创建搜索语句
	db = db.Model(coupon.CouponIssued{}).Select("" +
		"coupon_issued.id, " +
		"coupon_issued.user_id, " +
		"coupon_issued.coupon_id, " +
		"coupon_issued.coupon_name, " +
		"coupon_issued.coupon_amount, " +
		"coupon_issued.phone_number, " +
		"coupon_issued.type, " +
		"coupon_issued.user_type, " +
		"coupon_issued.valid_start, " +
		"coupon_issued.valid_end, " +
		"coupon_issued.created_at, " +
		"coupon_issued.updated_at, " +
		"coupon_issued.status, " +
		"frontend_users.email," +
		"frontend_users.root_userid as root_user_id," +
		"frontend_users.parent_id").Joins("left join frontend_users on coupon_issued.user_id = frontend_users.id ")
	if req.Status != "" {
		db = db.Where("coupon_issued.status = ?", req.Status)
	}
	if req.SuperiorId != 0 {
		db = db.Where("frontend_users.root_userid = ?", req.SuperiorId)
	}
	if req.UserType != 0 {
		db = db.Where("coupon_issued.user_type = ?", req.UserType)
	}
	if req.Email != "" {
		db = db.Where("frontend_users.email = ?", req.Email)
	}

	if req.UserId != 0 {
		db = db.Where("coupon_issued.user_id = ?", req.UserId)
	}

	if req.Name != "" {
		db = db.Where("coupon_issued.name = ?", req.Name)
	}

	if req.Phone != "" {
		db = db.Where("frontend_users.phone = ?", req.Phone)
	}

	if req.ParentId != 0 {
		db = db.Where("frontend_users.parent_id = ?", req.ParentId)
	}

	if req.UseStartTime != 0 {
		db.Where("coupon_issued.status = 1 and coupon_issued.updated_at >= ? and coupon_issued.updated_at < ? ",
			req.UseStartTime, req.UseEndTime)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&list).Error
	for i, cp := range list {
		n := time.Now().UnixMilli()
		c := list[i].CreatedAt
		u := list[i].UpdatedAt
		if cp.Status == coupon.NotUsed && cp.ValidEnd < n {
			list[i].Status = coupon.Expired
		}
		if cp.Status == coupon.AlreadyUsed {
			list[i].UpdateAtTick = u
		}
		list[i].CreatedAtTick = c
	}
	return list, total, err
}

// GetCouponIssuedInfoList 分页获取couponIssued表记录
// Author [yourname](https://github.com/yourname)
func (couponIssuedService *CouponIssuedService) GetCouponIssuedInfoList(info couponReq.CouponIssuedSearch) (list []*coupon.CouponIssued, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&coupon.CouponIssued{}).Order("id desc")
	if info.UserId != 0 {
		n := time.Now().UnixMilli()
		db = db.Where("user_id = ? and valid_start <= ? and valid_end > ? and status = ?", info.UserId, n, n, coupon.NotUsed)
	}
	var couponIssueds []*coupon.CouponIssued
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&couponIssueds).Error
	for i, cp := range couponIssueds {
		n := time.Now().UnixMilli()
		if cp.Status == coupon.NotUsed && cp.ValidEnd < n {
			couponIssueds[i].Status = coupon.Expired
		}
	}
	return couponIssueds, total, err
}

func (couponIssuedService *CouponIssuedService) GetCouponIssuedPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
