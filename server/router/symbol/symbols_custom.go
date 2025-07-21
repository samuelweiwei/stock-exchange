package symbol

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SymbolsCustomRouter struct{}

// InitSymbolsCustomRouter 初始化 symbolsCustom表 路由信息
func (s *SymbolsCustomRouter) InitSymbolsCustomRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup, PrivateGroupFrontUser *gin.RouterGroup) {
	symbolsCustomRouter := Router.Group("symbolsCustom").Use(middleware.OperationRecord())
	symbolsCustomRouterWithoutRecord := Router.Group("symbolsCustom")
	//symbolsCustomRouterWithoutAuth := PublicRouter.Group("symbolsCustom")
	symbolsCustomFrontRecord := PrivateGroupFrontUser.Group("symbolsCustom")

	{
		symbolsCustomRouter.POST("create", symbolsCustomApi.CreateSymbolsCustom)                          // 新建symbolsCustom表
		symbolsCustomRouter.DELETE("delete", symbolsCustomApi.DeleteSymbolsCustom)                        // 删除symbolsCustom表
		symbolsCustomRouter.DELETE("deleteByIds", symbolsCustomApi.DeleteSymbolsCustomByIds)              // 批量删除symbolsCustom表
		symbolsCustomRouter.PUT("update", symbolsCustomApi.UpdateSymbolsCustom)                           // 更新symbolsCustom表
		symbolsCustomRouter.POST("initUserSymbols/:userId", symbolsCustomApi.InitUserDefaultSymbolsAdmin) // 初始化用户默认自选股

	}
	{
		symbolsCustomRouterWithoutRecord.GET("detail/:id", symbolsCustomApi.FindSymbolsCustom) // 根据ID获取symbolsCustom表
		symbolsCustomRouterWithoutRecord.GET("list", symbolsCustomApi.GetSymbolsCustomList)    // 获取symbolsCustom表列表
	}
	{
		symbolsCustomFrontRecord.GET("/frontend/list", symbolsCustomApi.GetSymbolsCustomPublic)             // 获取用户自定义交易对
		symbolsCustomFrontRecord.POST("/frontend/create", symbolsCustomApi.CreateSymbolsCustomPublic)       // 新增用户自定义交易对
		symbolsCustomFrontRecord.DELETE("/frontend/delete", symbolsCustomApi.DeleteSymbolsCustomPublic)     // 删除用户自定义交易对
		symbolsCustomFrontRecord.POST("/frontend/initUserSymbols", symbolsCustomApi.InitUserDefaultSymbols) // 初始化用户默认自选股
	}
}
