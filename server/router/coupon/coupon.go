package coupon

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CouponRouter struct{}

// InitCouponRouter 初始化 coupon表 路由信息
func (s *CouponRouter) InitCouponRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	cpRouter := Router.Group("coupon").Use(middleware.OperationRecord())
	cpRouterWithoutRecord := Router.Group("coupon")
	{
		cpRouter.POST("create", cpApi.CreateCoupon)   // 新建coupon表
		cpRouter.DELETE("delete", cpApi.DeleteCoupon) // 删除coupon表
		cpRouter.PUT("update", cpApi.UpdateCoupon)    // 更新coupon表
		cpRouter.GET("list", cpApi.GetCouponList)     // 分页获取coupon表列表
		cpRouter.GET("options", cpApi.Options)        // 获取coupon表列表
	}
	{
		cpRouterWithoutRecord.GET("find", cpApi.FindCoupon) // 根据ID获取coupon表
	}

}
