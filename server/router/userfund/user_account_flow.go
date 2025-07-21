package userfund

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserAccountFlowRouter struct{}

// InitUserAccountFlowRouter 初始化 userAccountFlow表 路由信息
func (s *UserAccountFlowRouter) InitUserAccountFlowRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup, PrivateGroupFrontUser *gin.RouterGroup) {
	userAccountFlowRouter := Router.Group("userAccountFlow").Use(middleware.OperationRecord())
	userAccountFlowRouterWithoutRecord := Router.Group("userAccountFlow")
	userAccountFlowRouterWithoutAuth := PublicRouter.Group("userAccountFlow")
	privateGroupFrontUser := PrivateGroupFrontUser.Group("userAccountFlow")
	{
		userAccountFlowRouter.POST("createUserAccountFlow", userAccountFlowApi.CreateUserAccountFlow)             // 新建userAccountFlow表
		userAccountFlowRouter.DELETE("deleteUserAccountFlow", userAccountFlowApi.DeleteUserAccountFlow)           // 删除userAccountFlow表
		userAccountFlowRouter.DELETE("deleteUserAccountFlowByIds", userAccountFlowApi.DeleteUserAccountFlowByIds) // 批量删除userAccountFlow表
		userAccountFlowRouter.PUT("updateUserAccountFlow", userAccountFlowApi.UpdateUserAccountFlow)              // 更新userAccountFlow表
	}
	{
		userAccountFlowRouterWithoutRecord.GET("findUserAccountFlow", userAccountFlowApi.FindUserAccountFlow)       // 根据ID获取userAccountFlow表
		userAccountFlowRouterWithoutRecord.GET("getUserAccountFlowList", userAccountFlowApi.GetUserAccountFlowList) // 获取userAccountFlow表列表
	}
	{
		userAccountFlowRouterWithoutAuth.GET("getUserAccountFlowPublic", userAccountFlowApi.GetUserAccountFlowPublic) // userAccountFlow表开放接口
	}
	{
		//privateGroupFrontUser.POST("findById", userFundAccountsApi.Recharge)
		privateGroupFrontUser.GET("getFlowList", userAccountFlowApi.GetAccountFlowList)        // 获取userAccountFlow表列表
		privateGroupFrontUser.GET("getFundChangeTypes", userAccountFlowApi.GetFundChangeTypes) // 获取所有资金变动类型的国际化值

	}
}
