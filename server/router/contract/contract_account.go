package contract

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ContractAccountRouter struct{}

func (s *ContractAccountRouter) InitContractAccountRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup, FrontUserRouter *gin.RouterGroup) {
	contractAccountRouter := Router.Group("contractAccount").Use(middleware.OperationRecord())
	contractAccountRouterWithoutRecord := Router.Group("contractAccount")
	contractAccountRouterWithoutAuth := PublicRouter.Group("contractAccount")
	contractAccountRouterFrontUser := FrontUserRouter.Group("contractAccount")
	{
		contractAccountRouter.DELETE("deleteContractAccount", contractAccountApi.DeleteContractAccount)
		contractAccountRouter.DELETE("deleteContractAccountByIds", contractAccountApi.DeleteContractAccountByIds)
		contractAccountRouter.PUT("updateContractAccount", contractAccountApi.UpdateContractAccount)
	}
	{
		contractAccountRouterWithoutRecord.GET("getContractAccountList", contractAccountApi.GetContractAccountList)
	}
	{
		contractAccountRouterWithoutAuth.GET("getContractAccountPublic", contractAccountApi.GetContractAccountPublic)
	}
	{
		contractAccountRouterFrontUser.POST("createContractAccount", contractAccountApi.CreateContractAccount)
		contractAccountRouterFrontUser.GET("findContractAccount", contractAccountApi.FindContractAccount)
		contractAccountRouterFrontUser.POST("changeAccountMargin", contractAccountApi.ChangeAccountMargin)
	}
}
