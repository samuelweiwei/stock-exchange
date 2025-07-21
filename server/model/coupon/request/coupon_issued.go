package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/coupon"
)

type CouponIssuedSearch struct {
	request.PageInfo
	UserId       uint64 `json:"userId" form:"userId"`
	UserType     int    `json:"userType" form:"userType"`
	Email        string `json:"email" form:"email"`
	SuperiorId   uint   `json:"rootUserId" form:"rootUserId"`
	Status       string `json:"status" form:"status"`
	Name         string `json:"name" form:"name"`
	Phone        string `json:"phone" form:"phone"`
	ParentId     uint   `json:"parentUserId" form:"parentUserId"`
	UseStartTime int64  `json:"useStartTime" form:"useStartTime"`
	UseEndTime   int64  `json:"useEndTime" form:"useEndTime"`
}

type IssuedCouponSearch struct {
	request.PageInfo
	Type coupon.Type `json:"type" form:"type" binding:"required"`
}

type IssueCoupon struct {
	CouponId   int64  `json:"couponId" form:"couponId" binding:"required"`
	UserIdList string `json:"userIdList" form:"userIdList" binding:"required"`
}

type UseIssuedCoupon struct {
	IssuedCouponId uint64 `json:"issuedCouponId" form:"couponId" binding:"required"`
	UserId         uint64 `json:"userId" form:"userId" binding:"required"`
}
