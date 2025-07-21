package contract

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ContractOrderRouter struct{}

func (s *ContractOrderRouter) InitContractOrderRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup, FrontUserRouter *gin.RouterGroup) {
	contractOrderRouter := Router.Group("contractOrder").Use(middleware.OperationRecord())
	contractOrderRouterWithoutRecord := Router.Group("contractOrder")
	contractOrderRouterWithoutAuth := PublicRouter.Group("contractOrder")
	contractOrderRouterFrontUser := FrontUserRouter.Group("contractOrder")
	{
		//contractOrderRouter.POST("createContractOrder", contractOrderApi.CreateContractOrder)
		contractOrderRouter.DELETE("deleteContractOrder", contractOrderApi.DeleteContractOrder)
		contractOrderRouter.DELETE("deleteContractOrderByIds", contractOrderApi.DeleteContractOrderByIds)
		contractOrderRouter.PUT("updateContractOrder", contractOrderApi.UpdateContractOrder)
		//contractOrderRouter.POST("closeContractOrder", contractOrderApi.CloseContractOrder)
	}
	{
		contractOrderRouterWithoutRecord.GET("findContractOrder", contractOrderApi.FindContractOrder)
		//contractOrderRouterWithoutRecord.GET("getContractOrderList", contractOrderApi.GetContractOrderList)
	}
	{
		contractOrderRouterWithoutAuth.GET("getContractOrderPublic", contractOrderApi.GetContractOrderPublic)
	}
	{
		contractOrderRouterFrontUser.POST("createContractOrder", contractOrderApi.CreateContractOrder)     // 新建contractOrder表
		contractOrderRouterFrontUser.POST("closeContractOrder", contractOrderApi.CloseContractOrder)       // 平仓
		contractOrderRouterFrontUser.POST("closeAllContractOrder", contractOrderApi.CloseAllContractOrder) // 平仓
		contractOrderRouterFrontUser.GET("getContractOrderList", contractOrderApi.GetContractOrderList)    // 成交列表
	}
}
