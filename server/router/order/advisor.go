package order

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AdvisorRouter struct{}

// InitAdvisorRouter 初始化 advisor表 路由信息
func (s *AdvisorRouter) InitAdvisorRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	advisorRouterWithRecord := Router.Group("advisor").Use(middleware.OperationRecord()).Use(middleware.ErrorHandler())
	advisorPublicRouter := PublicRouter.Group("advisor").Use(middleware.ErrorHandler())
	// 需要鉴权路由
	{
		advisorRouterWithRecord.POST("", advisorApi.CreateAdvisor)                //创建导师
		advisorRouterWithRecord.PUT(":id", advisorApi.UpdateAdvisor)              //更新导师
		advisorRouterWithRecord.PUT("status/:id", advisorApi.UpdateAdvisorStatus) // 更新导师状态
	}

	{
		advisorPublicRouter.GET("page", advisorApi.PageQuery)             //分页查询导师
		advisorPublicRouter.GET(":id", advisorApi.GetAdvisor)             //查询导师详情
		advisorPublicRouter.GET("options", advisorApi.ListAdvisorOptions) //查询导师选项列表
		advisorPublicRouter.GET("front/page", advisorApi.PageQueryFront)  //前端分页查询导师
	}
}
