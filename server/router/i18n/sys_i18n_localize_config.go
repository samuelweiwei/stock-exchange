package i18n

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysI18nLocalizeConfigRouter struct{}

// InitSysI18nLocalizeConfigRouter 初始化 sysI18nLocalizeConfig表 路由信息
func (s *SysI18nLocalizeConfigRouter) InitSysI18nLocalizeConfigRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	sysI18nLocalizeConfigRouter := Router.Group("i18n/localize").Use(middleware.OperationRecord())
	sysI18nLocalizeConfigRouterWithoutRecord := Router.Group("i18n/localize")
	sysI18nLocalizeConfigRouterWithoutAuth := PublicRouter.Group("i18n/localize")
	{
		sysI18nLocalizeConfigRouter.POST("create", sysI18nLocalizeConfigApi.CreateSysI18nLocalizeConfig)               // 新建sysI18nLocalizeConfig表
		sysI18nLocalizeConfigRouter.DELETE("delete", sysI18nLocalizeConfigApi.DeleteSysI18nLocalizeConfig)             // 删除sysI18nLocalizeConfig表
		sysI18nLocalizeConfigRouter.DELETE("/batch/delete", sysI18nLocalizeConfigApi.DeleteSysI18nLocalizeConfigByIds) // 批量删除sysI18nLocalizeConfig表
		sysI18nLocalizeConfigRouter.PUT("update", sysI18nLocalizeConfigApi.UpdateSysI18nLocalizeConfig)                // 更新sysI18nLocalizeConfig表
	}
	{
		sysI18nLocalizeConfigRouterWithoutRecord.GET("find", sysI18nLocalizeConfigApi.FindSysI18nLocalizeConfig)    // 根据ID获取sysI18nLocalizeConfig表
		sysI18nLocalizeConfigRouterWithoutRecord.GET("list", sysI18nLocalizeConfigApi.GetSysI18nLocalizeConfigList) // 获取sysI18nLocalizeConfig表列表
	}
	{
		sysI18nLocalizeConfigRouterWithoutAuth.GET("getSysI18nLocalizeConfigPublic", sysI18nLocalizeConfigApi.GetSysI18nLocalizeConfigPublic) // sysI18nLocalizeConfig表开放接口
	}
}
