package symbol

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SymbolsHotRouter struct{}

// InitSymbolsHotRouter 初始化 symbolsHot表 路由信息
func (s *SymbolsHotRouter) InitSymbolsHotRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	symbolsHotRouter := Router.Group("symbolsHot").Use(middleware.OperationRecord())
	symbolsHotRouterWithoutRecord := Router.Group("symbolsHot")
	symbolsHotRouterWithoutAuth := PublicRouter.Group("symbolsHot")
	{
		symbolsHotRouter.POST("create", symbolsHotApi.CreateSymbolsHot)              // 新建symbolsHot表
		symbolsHotRouter.DELETE("delete", symbolsHotApi.DeleteSymbolsHot)            // 删除symbolsHot表
		symbolsHotRouter.DELETE("deleteSByIds", symbolsHotApi.DeleteSymbolsHotByIds) // 批量删除symbolsHot表
		symbolsHotRouter.PUT("update", symbolsHotApi.UpdateSymbolsHot)               // 更新symbolsHot表
	}
	{
		symbolsHotRouterWithoutRecord.GET("detail/:id", symbolsHotApi.FindSymbolsHot) // 根据ID获取symbolsHot表
		symbolsHotRouterWithoutRecord.GET("list", symbolsHotApi.GetSymbolsHotList)    // 获取symbolsHot表列表
	}
	{
		symbolsHotRouterWithoutAuth.GET("public/list", symbolsHotApi.GetSymbolsHotPublic) // symbolsHot表开放接口
	}
}
