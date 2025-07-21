package settingManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ServiceLinkRouter struct{}

// InitServiceLinkRouter 初始化 serviceLink表 路由信息
func (s *ServiceLinkRouter) InitServiceLinkRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	serviceLinkRouter := Router.Group("serviceLink").Use(middleware.OperationRecord())
	serviceLinkRouterWithoutAuthAndRecord := PublicRouter.Group("serviceLink")
	{
		serviceLinkRouter.POST("createServiceLink", serviceLinkApi.CreateServiceLink)   // 新建serviceLink表
		serviceLinkRouter.DELETE("deleteServiceLink", serviceLinkApi.DeleteServiceLink) // 删除serviceLink表
		serviceLinkRouter.PUT("updateServiceLink", serviceLinkApi.UpdateServiceLink)    // 更新serviceLink表
		serviceLinkRouter.GET("getServiceLinkList", serviceLinkApi.GetServiceLinkList)  // 获取serviceLink表列表
	}
	{
		serviceLinkRouterWithoutAuthAndRecord.GET("getFrontServiceLinkList", serviceLinkApi.GetFrontServiceLinkList) // 获取serviceLink表列表
	}
}
