package contract

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ContractEntrustRouter struct{}

// InitContractEntrustRouter 初始化 contractEntrust表 路由信息
func (s *ContractEntrustRouter) InitContractEntrustRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup, FrontUserRouter *gin.RouterGroup) {
	contractEntrustRouter := Router.Group("contractEntrust").Use(middleware.OperationRecord())
	contractEntrustRouterWithoutRecord := Router.Group("contractEntrust")
	contractEntrustRouterWithoutAuth := PublicRouter.Group("contractEntrust")
	contractEntrustRouterFrontUser := FrontUserRouter.Group("contractEntrust")
	{
		//contractEntrustRouter.POST("createContractEntrust", contractEntrustApi.CreateContractEntrust)             // 新建contractEntrust表
		//contractEntrustRouter.DELETE("deleteContractEntrust", contractEntrustApi.DeleteContractEntrust)           // 删除contractEntrust表
		contractEntrustRouter.DELETE("deleteContractEntrustByIds", contractEntrustApi.DeleteContractEntrustByIds) // 批量删除contractEntrust表
		contractEntrustRouter.PUT("updateContractEntrust", contractEntrustApi.UpdateContractEntrust)              // 更新contractEntrust表
	}
	{
		contractEntrustRouterWithoutRecord.GET("findContractEntrust", contractEntrustApi.FindContractEntrust) // 根据ID获取contractEntrust表
		//contractEntrustRouterWithoutRecord.GET("getContractEntrustList", contractEntrustApi.GetContractEntrustList) // 获取contractEntrust表列表
	}
	{
		contractEntrustRouterWithoutAuth.GET("getContractEntrustPublic", contractEntrustApi.GetContractEntrustPublic) // contractEntrust表开放接口
	}
	{
		contractEntrustRouterFrontUser.POST("createContractEntrust", contractEntrustApi.CreateContractEntrust)       // 新建contractEntrust表
		contractEntrustRouterFrontUser.POST("deleteContractEntrust", contractEntrustApi.DeleteContractEntrust)       // 删除contractEntrust表
		contractEntrustRouterFrontUser.POST("deleteAllContractEntrust", contractEntrustApi.DeleteAllContractEntrust) // 删除contractEntrust表
		contractEntrustRouterFrontUser.GET("getContractEntrustList", contractEntrustApi.GetContractEntrustList)      // 获取contractEntrust表列表
	}
}
