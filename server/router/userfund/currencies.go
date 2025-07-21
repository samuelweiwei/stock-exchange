package userfund

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CurrenciesRouter struct{}

// InitCurrenciesRouter 初始化 currencies表 路由信息
func (s *CurrenciesRouter) InitCurrenciesRouter(Router *gin.RouterGroup, privateGroupFrontUser *gin.RouterGroup) {
	currenciesRouter := Router.Group("currencies").Use(middleware.OperationRecord())
	currenciesRouterWithoutRecord := Router.Group("currencies")
	currenciesRouteFrontendAuth := privateGroupFrontUser.Group("currencies")
	{
		currenciesRouter.POST("create", currenciesApi.CreateCurrencies)             // 新建currencies表
		currenciesRouter.DELETE("delete", currenciesApi.DeleteCurrencies)           // 删除currencies表
		currenciesRouter.DELETE("deleteByIds", currenciesApi.DeleteCurrenciesByIds) // 批量删除currencies表
		currenciesRouter.PUT("update", currenciesApi.UpdateCurrencies)              // 更新currencies表
	}
	{
		currenciesRouterWithoutRecord.GET("detail/:id", currenciesApi.FindCurrencies) // 根据ID获取currencies表
		currenciesRouterWithoutRecord.GET("list", currenciesApi.GetCurrenciesList)    // 获取currencies表列表
	}
	{
		currenciesRouteFrontendAuth.GET("frontend/list", currenciesApi.GetCurrenciesPublic) // currencies表开放接口
	}
}
