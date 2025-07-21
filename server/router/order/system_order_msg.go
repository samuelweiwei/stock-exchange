package order

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SystemOrderMsgRouter struct{}

func (router *SystemOrderMsgRouter) InitSystemOrderMsgRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	systemOrderMsgRouterWithoutRecord := Router.Group("system/order/msg").Use(middleware.ErrorHandler())
	{
		systemOrderMsgRouterWithoutRecord.GET("unread/count", systemOrderMsgApi.GetUnReadCount)
		systemOrderMsgRouterWithoutRecord.GET("page", systemOrderMsgApi.PageQuery)
	}

	systemOrderMsgRouterWithRecord := Router.Group("system/order/msg").Use(middleware.OperationRecord()).Use(middleware.ErrorHandler())
	{
		systemOrderMsgRouterWithRecord.POST("read", systemOrderMsgApi.SetReadStatus)
	}
}
