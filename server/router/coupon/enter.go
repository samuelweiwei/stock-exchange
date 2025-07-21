package coupon

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	CouponRouter
	CouponIssuedRouter
}

var (
	cpApi           = api.ApiGroupApp.CouponApiGroup.CouponApi
	couponIssuedApi = api.ApiGroupApp.CouponApiGroup.CouponIssuedApi
)
