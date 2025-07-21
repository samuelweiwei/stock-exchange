package coupon

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CouponIssuedRouter struct{}

// InitCouponIssuedRouter 初始化 couponIssued表 路由信息
func (s *CouponIssuedRouter) InitCouponIssuedRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	IssuedCouponRouter := Router.Group("coupon").Use(middleware.OperationRecord())
	IssuedCouponRouterWithoutRecord := Router.Group("coupon")
	{
		IssuedCouponRouter.POST("issue", couponIssuedApi.CreateCouponIssued)            // 发放优惠券
		IssuedCouponRouter.GET("issued/list", couponIssuedApi.AdminGetCouponIssuedList) // 分页查询发放的优惠券
		IssuedCouponRouter.PUT("issued/use", couponIssuedApi.UseCoupon)                 // 使用优惠券
		IssuedCouponRouter.GET("issued/options", couponIssuedApi.Options)               // 优惠券下拉列表
	}
	{
		IssuedCouponRouterWithoutRecord.GET("findIssuedCoupon", couponIssuedApi.FindCouponIssued)       // 根据ID获取IssuedCoupon表
		IssuedCouponRouterWithoutRecord.GET("getIssuedCouponList", couponIssuedApi.GetCouponIssuedList) // 获取IssuedCoupon表列表
	}
}

// InitFrontCouponIssuedRouter 初始化 couponIssued表 路由信息
func (s *CouponIssuedRouter) InitFrontCouponIssuedRouter(Router *gin.RouterGroup, PrivateGroupFrontUser *gin.RouterGroup) {
	IssuedCouponRouter := PrivateGroupFrontUser.Group("coupon").Use(middleware.OperationRecord())
	{
		IssuedCouponRouter.GET("front/issued/list", couponIssuedApi.GetCouponIssuedList) // 分页查询发放的优惠券
	}

}
