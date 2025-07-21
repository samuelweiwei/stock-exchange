package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/coupon"

type CouponInfo struct {
	Id uint `json:"id"`
}
type CouponIssuedRes struct {
	*coupon.CouponIssued
	UserType   uint   `json:"userType" gorm:"user_type,omitempty"`
	Email      string `json:"email" gorm:"email,omitempty"`
	RootUserId uint   `json:"rootUserId" gorm:"root_user_id"`
	//SuperiorId uint   `json:"superiorId" gorm:"-"`
	ParentId uint `json:"parentUserId" gorm:"parent_id,omitempty"`
}
