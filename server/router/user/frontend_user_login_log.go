package user

import (
	"github.com/gin-gonic/gin"
)

type FrontendUserLoginLogRouter struct{}

// InitFrontendUserLoginLogRouter 初始化 frontendUserLoginLog表 路由信息
func (s *FrontendUserLoginLogRouter) InitFrontendUserLoginLogRouter(Router *gin.RouterGroup) {
	frontendUserLoginLogRouterWithoutRecord := Router.Group("frontendUserLoginLog")
	{
		frontendUserLoginLogRouterWithoutRecord.GET("getFrontendUserLoginLogList", frontendUserLoginLogApi.GetFrontendUserLoginLogList) // 获取frontendUserLoginLog表列表
	}
}
