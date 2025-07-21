package contract

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ContractPositionRouter struct{}

func (s *ContractPositionRouter) InitContractPositionRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup, FrontUserRouter *gin.RouterGroup) {
	contractPositionRouter := Router.Group("contractPosition").Use(middleware.OperationRecord())
	contractPositionRouterWithoutRecord := Router.Group("contractPosition")
	contractPositionRouterWithoutAuth := PublicRouter.Group("contractPosition")
	contractPositionRouterFrontUser := FrontUserRouter.Group("contractPosition")
	{
		contractPositionRouter.POST("createContractPosition", contractPositionApi.CreateContractPosition)
		contractPositionRouter.DELETE("deleteContractPosition", contractPositionApi.DeleteContractPosition)
		contractPositionRouter.DELETE("deleteContractPositionByIds", contractPositionApi.DeleteContractPositionByIds)
		contractPositionRouter.PUT("updateContractPosition", contractPositionApi.UpdateContractPosition)
	}
	{
		contractPositionRouterWithoutRecord.GET("findContractPosition", contractPositionApi.FindContractPosition)
	}
	{
		contractPositionRouterWithoutAuth.GET("getContractPositionPublic", contractPositionApi.GetContractPositionPublic)
	}
	{
		contractPositionRouterFrontUser.GET("getContractPositionList", contractPositionApi.GetContractPositionList)
	}
}
