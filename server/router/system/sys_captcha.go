package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysCaptchaRouter struct{}

func (s *SysParamsRouter) InitCaptchaRouter(Router *gin.RouterGroup) {
	sysParamsRouter := Router.Group("admin/captcha").Use(middleware.OperationRecord())
	{
		sysParamsRouter.GET("log/list", captchaLogApi.List) // 新建参数
	}
}
