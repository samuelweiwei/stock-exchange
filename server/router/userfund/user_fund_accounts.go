package userfund

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserFundAccountsRouter struct{}

func (s *UserFundAccountsRouter) InitUserFundAccountsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup, PrivateGroupFrontUser *gin.RouterGroup) {
	userFundAccountsRouter := Router.Group("userFundAccounts").Use(middleware.OperationRecord())
	userFundAccountsRouterWithoutRecord := Router.Group("userFundAccounts")
	userFundAccountsRouterWithoutAuth := PublicRouter.Group("userFundAccounts")
	privateGroupFrontUser := PrivateGroupFrontUser.Group("userFundAccounts")
	{
		userFundAccountsRouter.POST("createUserFundAccounts", userFundAccountsApi.CreateUserFundAccounts)
		userFundAccountsRouter.DELETE("deleteUserFundAccounts", userFundAccountsApi.DeleteUserFundAccounts)
		userFundAccountsRouter.DELETE("deleteUserFundAccountsByIds", userFundAccountsApi.DeleteUserFundAccountsByIds)
		userFundAccountsRouter.PUT("updateUserFundAccounts", userFundAccountsApi.UpdateUserFundAccounts)
		userFundAccountsRouter.GET("getUserFundStatic/:userId", userFundAccountsApi.GetUserFundStatic)
		//userFundAccountsRouterWithoutAuth.POST("batchSendFund", userFundAccountsApi.BatchSendFund)

	}
	{
		userFundAccountsRouterWithoutRecord.GET("findUserFundAccounts", userFundAccountsApi.FindUserFundAccounts)
		userFundAccountsRouterWithoutRecord.GET("getUserFundAccountsList", userFundAccountsApi.GetUserFundAccountsList)

	}
	{
		//userFundAccountsRouterWithoutAuth.POST("getUserFundStatic", userFundAccountsApi.GetUserFundStatic)
		userFundAccountsRouterWithoutAuth.POST("batchSendFund", userFundAccountsApi.BatchSendFund)
	}
	{
		userFundAccountsRouterWithoutAuth.GET("getUserFundAccountsPublic", userFundAccountsApi.GetUserFundAccountsPublic)
	}
	{
		privateGroupFrontUser.POST("recharge", userFundAccountsApi.Recharge)
		privateGroupFrontUser.POST("withdraw", userFundAccountsApi.Withdraw)
		privateGroupFrontUser.POST("processAccountForFrontUser", userFundAccountsApi.ProcessAccount)
		//查询用户资金信息
		privateGroupFrontUser.GET("getUserAccountInfo", userFundAccountsApi.GetUserAccountInfo)
		privateGroupFrontUser.POST("processAccount", userFundAccountsApi.ProcessAccount)
	}
}
