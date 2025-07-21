package earn

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EarnProductsRouter struct{}

// InitEarnProductsRouter 初始化 earnProducts表 路由信息
func (s *EarnProductsRouter) InitEarnProductsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	earnProductsRouter := Router.Group("earn/products").Use(middleware.OperationRecord())
	earnProductsRouterWithoutRecord := Router.Group("earn/products")
	earnProductsRouterWithoutAuth := PublicRouter.Group("earn/products")
	{
		earnProductsRouter.POST("create", earnProductsApi.CreateEarnProducts)                         // 新建earnProducts表
		earnProductsRouter.DELETE("delete", earnProductsApi.DeleteEarnProducts)                       // 删除earnProducts表
		earnProductsRouter.DELETE("deleteEarnProductsByIds", earnProductsApi.DeleteEarnProductsByIds) // 批量删除earnProducts表
		earnProductsRouter.PUT("edit", earnProductsApi.UpdateEarnProducts)                            // 更新earnProducts表
	}
	{
		earnProductsRouterWithoutRecord.GET("findEarnProducts", earnProductsApi.FindEarnProducts) // 根据ID获取earnProducts表
		earnProductsRouterWithoutRecord.GET("list", earnProductsApi.GetEarnProductsList)          // 获取earnProducts表列表
	}
	{
		earnProductsRouterWithoutAuth.GET("getEarnProductsPublic", earnProductsApi.GetEarnProductsPublic) // earnProducts表开放接口
	}
}

// InitFrontCouponIssuedRouter
func (s *EarnProductsRouter) InitFrontEarnProductRouter(Router *gin.RouterGroup, PrivateGroupFrontUser *gin.RouterGroup) {
	earnProductsRouter := PrivateGroupFrontUser.Group("/front/earn/product/").Use(middleware.OperationRecord())
	{
		earnProductsRouter.GET("list", earnProductsApi.GetFrontEarnProductsList)
	}
}
