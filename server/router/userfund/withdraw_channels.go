package userfund

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WithdrawChannelsRouter struct{}

// InitWithdrawChannelsRouter 初始化 withdrawChannels表 路由信息
func (s *WithdrawChannelsRouter) InitWithdrawChannelsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup, PrivateGroupFrontUser *gin.RouterGroup) {
	withdrawChannelsRouter := Router.Group("withdrawChannels").Use(middleware.OperationRecord())
	withdrawChannelsRouterWithoutRecord := Router.Group("withdrawChannels")
	withdrawChannelsRouterWithoutAuth := PublicRouter.Group("withdrawChannels")
	privateGroupFrontUser := PrivateGroupFrontUser.Group("withdrawChannels")
	{
		withdrawChannelsRouter.POST("createWithdrawChannels", withdrawChannelsApi.CreateWithdrawChannels)             // 新建withdrawChannels表
		withdrawChannelsRouter.DELETE("deleteWithdrawChannels", withdrawChannelsApi.DeleteWithdrawChannels)           // 删除withdrawChannels表
		withdrawChannelsRouter.DELETE("deleteWithdrawChannelsByIds", withdrawChannelsApi.DeleteWithdrawChannelsByIds) // 批量删除withdrawChannels表
		withdrawChannelsRouter.PUT("updateWithdrawChannels", withdrawChannelsApi.UpdateWithdrawChannels)              // 更新withdrawChannels表
	}
	{
		withdrawChannelsRouterWithoutRecord.GET("findWithdrawChannels", withdrawChannelsApi.FindWithdrawChannels)       // 根据ID获取withdrawChannels表
		withdrawChannelsRouterWithoutRecord.GET("getWithdrawChannelsList", withdrawChannelsApi.GetWithdrawChannelsList) // 获取withdrawChannels表列表
	}
	{
		withdrawChannelsRouterWithoutAuth.GET("getWithdrawChannelsPublic", withdrawChannelsApi.GetWithdrawChannelsPublic) // withdrawChannels表开放接口
	}
	{
		privateGroupFrontUser.GET("getWithdrawChannels", withdrawChannelsApi.GetWithdrawChannelsList2)
	}
}
