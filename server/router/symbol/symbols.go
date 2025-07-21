package symbol

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SymbolsRouter struct{}

// InitSymbolsRouter 初始化 symbols表 路由信息
func (s *SymbolsRouter) InitSymbolsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	symbolsRouter := Router.Group("symbols").Use(middleware.OperationRecord())
	symbolsRouterWithoutRecord := Router.Group("symbols")
	symbolsRouterWithoutAuth := PublicRouter.Group("symbols")
	var symbolsApi = v1.ApiGroupApp.SymbolApiGroup.SymbolsApi
	{
		symbolsRouter.POST("create", symbolsApi.CreateSymbols)             // 新建symbols表
		symbolsRouter.DELETE("delete", symbolsApi.DeleteSymbols)           // 删除symbols表
		symbolsRouter.DELETE("deleteByIds", symbolsApi.DeleteSymbolsByIds) // 批量删除symbols表
		symbolsRouter.PUT("update", symbolsApi.UpdateSymbols)              // 更新symbols表
		symbolsRouter.POST("addByName", symbolsApi.AddSymbolsByName)       // 需要授权的版本
		symbolsRouter.GET("search", symbolsApi.SearchTickers)              // 查询可添加的ticker列表
		symbolsRouter.POST("validate", symbolsApi.ValidateSymbol)          // 验证symbol是否可用
	}
	{
		symbolsRouterWithoutRecord.GET("detail/:id", symbolsApi.FindSymbols)           // 根据ID获取symbols表
		symbolsRouterWithoutRecord.GET("list", symbolsApi.GetSymbolsList)              // 获取symbols表列表
		symbolsRouterWithoutRecord.GET("price/:symbol", symbolsApi.GetSymbolPrice)     // 获取单个symbol的实时价格
		symbolsRouterWithoutRecord.GET("getAllSimple", symbolsApi.GetAllSymbolsSimple) // 获取所有股票的简化信息
	}
	{
		symbolsRouterWithoutAuth.GET("ws", symbolsApi.SubscribeSymbolPrices)            // WebSocket订阅价格
		symbolsRouterWithoutAuth.GET("public/list", symbolsApi.GetSymbolsPublic)        // symbols表开放接口
		symbolsRouterWithoutAuth.GET("public/price/:id", symbolsApi.GetSymbolPriceById) // 根据ID获取价格
		symbolsRouterWithoutAuth.POST("kline", symbolsApi.GetKlineData)                 // 获取K线历史数据
		symbolsRouterWithoutAuth.GET("public/detail/:id", symbolsApi.FindSymbols)       // 根据ID获取symbols表
		symbolsRouterWithoutAuth.POST("agg", symbolsApi.GetSymbolAggData)               // 获取聚合数据
		symbolsRouterWithoutAuth.GET("agg/ws", symbolsApi.SubscribeSymbolAggData)       // WebSocket订阅K线数据
		symbolsRouterWithoutAuth.GET("front/detail/:id", symbolsApi.FrontSymbolsDetail) // 根据ID获取symbols表
		symbolsRouterWithoutAuth.GET("unified/ws", symbolsApi.UnifiedWebSocket)         // 统一的WebSocket接口
		symbolsRouterWithoutAuth.POST("front/symbol", symbolsApi.GetSymbolBySymbolName) // 根据symbol获取详细信息
	}
}
