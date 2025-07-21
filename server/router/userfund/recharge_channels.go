package userfund

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RechargeChannelsRouter struct{}

func (s *RechargeChannelsRouter) InitRechargeChannelsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup, PrivateGroupFrontUser *gin.RouterGroup) {
	rechargeChannelsRouter := Router.Group("rechargeChannels").Use(middleware.OperationRecord())
	rechargeChannelsRouterWithoutRecord := Router.Group("rechargeChannels")
	rechargeChannelsRouterWithoutAuth := PublicRouter.Group("rechargeChannels")
	privateGroupFrontUser := PrivateGroupFrontUser.Group("rechargeChannels")
	{
		rechargeChannelsRouter.POST("createRechargeChannels", rechargeChannelsApi.CreateRechargeChannels)
		rechargeChannelsRouter.DELETE("deleteRechargeChannels", rechargeChannelsApi.DeleteRechargeChannels)
		rechargeChannelsRouter.DELETE("deleteRechargeChannelsByIds", rechargeChannelsApi.DeleteRechargeChannelsByIds)
		rechargeChannelsRouter.PUT("updateRechargeChannels", rechargeChannelsApi.UpdateRechargeChannels)
	}
	{
		rechargeChannelsRouterWithoutRecord.GET("findRechargeChannels", rechargeChannelsApi.FindRechargeChannels)
		rechargeChannelsRouterWithoutRecord.GET("getRechargeChannelsList", rechargeChannelsApi.GetRechargeChannelsList)
	}
	{
		rechargeChannelsRouterWithoutAuth.GET("getRechargeChannelsPublic", rechargeChannelsApi.GetRechargeChannelsPublic)
	}
	{
		privateGroupFrontUser.GET("getRechargeChannels", rechargeChannelsApi.GetRechargeChannelsList2)
	}
}
