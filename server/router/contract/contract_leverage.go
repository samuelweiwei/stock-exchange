package contract

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ContractLeverageRouter struct{}

// InitContractLeverageRouter 初始化 contractLeverage表 路由信息
func (s *ContractLeverageRouter) InitContractLeverageRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup, FrontUserRouter *gin.RouterGroup) {
	contractLeverageRouter := Router.Group("contractLeverage").Use(middleware.OperationRecord())
	contractLeverageRouterWithoutRecord := Router.Group("contractLeverage")
	contractLeverageRouterWithoutAuth := PublicRouter.Group("contractLeverage")
	contractLeverageRouterFrontUser := FrontUserRouter.Group("contractLeverage")
	{
		contractLeverageRouter.POST("createContractLeverage", contractLeverageApi.CreateContractLeverage)             // 新建contractLeverage表
		contractLeverageRouter.DELETE("deleteContractLeverage", contractLeverageApi.DeleteContractLeverage)           // 删除contractLeverage表
		contractLeverageRouter.DELETE("deleteContractLeverageByIds", contractLeverageApi.DeleteContractLeverageByIds) // 批量删除contractLeverage表
		//contractLeverageRouter.PUT("updateContractLeverage", contractLeverageApi.UpdateContractLeverage)              // 更新contractLeverage表
	}
	{
		//contractLeverageRouterWithoutRecord.GET("findContractLeverage", contractLeverageApi.FindContractLeverage)       // 根据ID获取contractLeverage表
		contractLeverageRouterWithoutRecord.GET("getContractLeverageList", contractLeverageApi.GetContractLeverageList) // 获取contractLeverage表列表
	}
	{
		contractLeverageRouterWithoutAuth.GET("getContractLeveragePublic", contractLeverageApi.GetContractLeveragePublic) // contractLeverage表开放接口
	}
	{
		contractLeverageRouterFrontUser.GET("findContractLeverage", contractLeverageApi.FindContractLeverage)      // 根据ID获取contractLeverage表
		contractLeverageRouterFrontUser.POST("updateContractLeverage", contractLeverageApi.UpdateContractLeverage) // 更新contractLeverage表
	}
}
