package coupon

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	CouponApi
	CouponIssuedApi
}

var (
	cpService           = service.ServiceGroupApp.CouponServiceGroup.CouponService
	couponIssuedService = service.ServiceGroupApp.CouponServiceGroup.CouponIssuedService
)
