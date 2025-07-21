package report

import "github.com/gin-gonic/gin"

type PlatformReportRouter struct{}

func (r *PlatformReportRouter) InitPlatformReportRouter(PrivateRouter *gin.RouterGroup, PublicGroup *gin.RouterGroup) {
	privateRouterWithoutRecord := PrivateRouter.Group("report")
	{
		privateRouterWithoutRecord.GET("summary", platformReportApi.GetPlatformReportSummary)        //查询平台固定报表
		privateRouterWithoutRecord.GET("daily/page", platformReportApi.PageQueryPlatformReportDaily) //分页查询每日报表
		privateRouterWithoutRecord.GET("dynamic", platformReportApi.GetPlatformDynamicReport)        //查询平台动态报表
	}
}
