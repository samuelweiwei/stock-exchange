package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysRouter struct{}

func (s *SysRouter) InitSystemRouter(Router *gin.RouterGroup, PublicGroup *gin.RouterGroup, FrontGroup *gin.RouterGroup) {
	sysRouter := Router.Group("system").Use(middleware.OperationRecord()).Use(middleware.ErrorHandler())
	sysRouterWithoutRecord := Router.Group("system").Use(middleware.ErrorHandler())
	sysPublicRouter := PublicGroup.Group("system").Use(middleware.ErrorHandler())
	{
		sysRouter.POST("setSystemConfig", systemApi.SetSystemConfig) // 设置配置文件内容
		sysRouter.POST("reloadSystem", systemApi.ReloadSystem)       // 重启服务
		sysRouter.POST("config", systemConfigApi.SaveSystemConfig)   //保存平台基本配置
	}
	{
		sysRouterWithoutRecord.POST("getSystemConfig", systemApi.GetSystemConfig) // 获取配置文件内容
		sysRouterWithoutRecord.POST("getServerInfo", systemApi.GetServerInfo)     // 获取服务器信息
		sysRouterWithoutRecord.GET("config", systemConfigApi.GetSystemConfig)     //获取平台基本配置
	}
	{
		sysPublicRouter.GET("platform-commission-rate", systemConfigApi.GetPlatformCommissionRate) //获取平台佣金比例
		sysPublicRouter.GET("getDomainInfo", systemConfigApi.GetDomainInfo)                        //前台-获取站点信息
	}

}
