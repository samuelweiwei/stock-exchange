package earn

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EarnInterestRatesRouter struct {}

// InitEarnInterestRatesRouter 初始化 earnInterestRates表 路由信息
func (s *EarnInterestRatesRouter) InitEarnInterestRatesRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	earnInterestRatesRouter := Router.Group("earnInterestRates").Use(middleware.OperationRecord())
	earnInterestRatesRouterWithoutRecord := Router.Group("earnInterestRates")
	earnInterestRatesRouterWithoutAuth := PublicRouter.Group("earnInterestRates")
	{
		earnInterestRatesRouter.POST("createEarnInterestRates", earnInterestRatesApi.CreateEarnInterestRates)   // 新建earnInterestRates表
		earnInterestRatesRouter.DELETE("deleteEarnInterestRates", earnInterestRatesApi.DeleteEarnInterestRates) // 删除earnInterestRates表
		earnInterestRatesRouter.DELETE("deleteEarnInterestRatesByIds", earnInterestRatesApi.DeleteEarnInterestRatesByIds) // 批量删除earnInterestRates表
		earnInterestRatesRouter.PUT("updateEarnInterestRates", earnInterestRatesApi.UpdateEarnInterestRates)    // 更新earnInterestRates表
	}
	{
		earnInterestRatesRouterWithoutRecord.GET("findEarnInterestRates", earnInterestRatesApi.FindEarnInterestRates)        // 根据ID获取earnInterestRates表
		earnInterestRatesRouterWithoutRecord.GET("getEarnInterestRatesList", earnInterestRatesApi.GetEarnInterestRatesList)  // 获取earnInterestRates表列表
	}
	{
	    earnInterestRatesRouterWithoutAuth.GET("getEarnInterestRatesPublic", earnInterestRatesApi.GetEarnInterestRatesPublic)  // earnInterestRates表开放接口
	}
}
