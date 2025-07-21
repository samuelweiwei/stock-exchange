package earn

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EarnSubscribeLogRouter struct{}

// InitEarnSubscribeLogRouter 初始化 earnSubscribeLog表 路由信息
func (s *EarnSubscribeLogRouter) InitEarnSubscribeLogRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	earnSubscribeLogRouter := Router.Group("earn/subscribe/log").Use(middleware.OperationRecord())
	earnSubscribeLogRouterWithoutRecord := Router.Group("earn/subscribe/log")
	earnSubscribeLogRouterWithoutAuth := PublicRouter.Group("earn/subscribe/log")
	{
		earnSubscribeLogRouter.POST("create", earnSubscribeLogApi.CreateEarnSubscribeLog)                             // 新建earnSubscribeLog表
		earnSubscribeLogRouter.DELETE("delete", earnSubscribeLogApi.DeleteEarnSubscribeLog)                           // 删除earnSubscribeLog表
		earnSubscribeLogRouter.DELETE("deleteEarnSubscribeLogByIds", earnSubscribeLogApi.DeleteEarnSubscribeLogByIds) // 批量删除earnSubscribeLog表
		earnSubscribeLogRouter.PUT("edit", earnSubscribeLogApi.UpdateEarnSubscribeLog)                                // 更新earnSubscribeLog表
	}
	{
		//earnSubscribeLogRouterWithoutRecord.GET("find", earnSubscribeLogApi.FindEarnSubscribeLog)    // 根据ID获取earnSubscribeLog表
		earnSubscribeLogRouterWithoutRecord.GET("list", earnSubscribeLogApi.GetEarnSubscribeLogList) // 获取earnSubscribeLog表列表
	}
	{
		earnSubscribeLogRouterWithoutAuth.GET("getEarnSubscribeLogPublic", earnSubscribeLogApi.GetEarnSubscribeLogPublic) // earnSubscribeLog表开放接口
	}
}

func (s *EarnSubscribeLogRouter) InitFrontEarnSubscribeLogRouter(Router *gin.RouterGroup, PrivateGroupFrontUser *gin.RouterGroup) {
	earnProductsRouter := PrivateGroupFrontUser.Group("/front/earn/product/subscription").Use(middleware.OperationRecord())
	{
		earnProductsRouter.GET("list", earnSubscribeLogApi.GetFrontEarnSubscribeLogList)
		earnProductsRouter.POST("stake", earnSubscribeLogApi.Stake)
		earnProductsRouter.POST("redeem", earnSubscribeLogApi.Redeem)
		earnProductsRouter.GET("summary", earnSubscribeLogApi.GetFrontEarnSubscribeSummary)
	}
}
