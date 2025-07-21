package earn

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EarnDailyIncomeMoneyLogRouter struct{}

// InitEarnDailyIncomeMoneyLogRouter 初始化 earnDailyIncomeMoneyLog表 路由信息
func (s *EarnDailyIncomeMoneyLogRouter) InitEarnDailyIncomeMoneyLogRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	earnDailyIncomeMoneyLogRouter := Router.Group("/earn/product/daily/income").Use(middleware.OperationRecord())
	earnDailyIncomeMoneyLogRouterWithoutRecord := Router.Group("/earn/product/daily/income")
	earnDailyIncomeMoneyLogRouterWithoutAuth := PublicRouter.Group("/earn/product/daily/income")
	{
		earnDailyIncomeMoneyLogRouter.POST("createEarnDailyIncomeMoneyLog", earnDailyIncomeMoneyLogApi.CreateEarnDailyIncomeMoneyLog)             // 新建earnDailyIncomeMoneyLog表
		earnDailyIncomeMoneyLogRouter.DELETE("deleteEarnDailyIncomeMoneyLog", earnDailyIncomeMoneyLogApi.DeleteEarnDailyIncomeMoneyLog)           // 删除earnDailyIncomeMoneyLog表
		earnDailyIncomeMoneyLogRouter.DELETE("deleteEarnDailyIncomeMoneyLogByIds", earnDailyIncomeMoneyLogApi.DeleteEarnDailyIncomeMoneyLogByIds) // 批量删除earnDailyIncomeMoneyLog表
		earnDailyIncomeMoneyLogRouter.PUT("updateEarnDailyIncomeMoneyLog", earnDailyIncomeMoneyLogApi.UpdateEarnDailyIncomeMoneyLog)              // 更新earnDailyIncomeMoneyLog表
	}
	{
		earnDailyIncomeMoneyLogRouterWithoutRecord.GET("findEarnDailyIncomeMoneyLog", earnDailyIncomeMoneyLogApi.FindEarnDailyIncomeMoneyLog) // 根据ID获取earnDailyIncomeMoneyLog表
		earnDailyIncomeMoneyLogRouterWithoutRecord.GET("summary", earnDailyIncomeMoneyLogApi.GetEarnProductDailyIncomeMoneyLogSummary)        // 获取earnDailyIncomeMoneyLog表列表
	}
	{
		earnDailyIncomeMoneyLogRouterWithoutAuth.GET("getEarnDailyIncomeMoneyLogPublic", earnDailyIncomeMoneyLogApi.GetEarnDailyIncomeMoneyLogPublic) // earnDailyIncomeMoneyLog表开放接口
	}
	{
		earnDailyIncomeMoneyLogRouterWithoutRecord.GET("list", earnDailyIncomeMoneyLogApi.GetEarnProductDailyIncomeMoneyLogSummary)
	}
}

// InitEarnDailyIncomeMoneyLogRouter 初始化 earnDailyIncomeMoneyLog表 路由信息
func (s *EarnDailyIncomeMoneyLogRouter) InitFrontEarnDailyIncomeMoneyLogRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	earnDailyIncomeMoneyLogRouter := Router.Group("/front/earn/product/daily/income").Use(middleware.OperationRecord())
	{
		earnDailyIncomeMoneyLogRouter.GET("detail", earnDailyIncomeMoneyLogApi.GetEarnProductDailyIncomeMoneyLogDetail)
	}
}
