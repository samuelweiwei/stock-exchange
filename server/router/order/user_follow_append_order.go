package order

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserFollowAppendOrderRouter struct{}

func (router *UserFollowAppendOrderRouter) InitUserFollowAppendOrderAdminRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	adminRouterWithoutRecord := Router.Group("user/append/order").Use(middleware.ErrorHandler())
	adminRouterWithRecord := Router.Group("user/append/order").Use(middleware.OperationRecord()).Use(middleware.ErrorHandler())
	{
		adminRouterWithoutRecord.GET("page", userFollowAppendOrderApi.PageQuery)               //分页查询用户追单列表
		adminRouterWithRecord.POST("approve/:appendOrderId", userFollowAppendOrderApi.Approve) //审核通过用户追单
		adminRouterWithRecord.POST("reject/:appendOrderId", userFollowAppendOrderApi.Reject)   //审核驳回用户追单
	}
}

func (router *UserFollowAppendOrderRouter) InitUserFollowAppendOrderH5Router(PrivateRouter *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	h5RouterWithoutRecord := PrivateRouter.Group("").Use(middleware.ErrorHandler())
	h5RouterWithRecord := PrivateRouter.Group("").Use(middleware.OperationRecord()).Use(middleware.ErrorHandler())
	{
		h5RouterWithoutRecord.GET("user/append/orders/my/:followOrderId", userFollowAppendOrderApi.QueryMyAppendOrderList) //查询我的追单列表
		h5RouterWithRecord.POST("user/append/order/:followOrderId", userFollowAppendOrderApi.Apply)                        //申请追单
	}
}
