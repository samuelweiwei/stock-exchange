package order

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AdvisorProdRouter struct{}

func (r *AdvisorProdRouter) InitAdvisorProdRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup, FrontPrivateGroup *gin.RouterGroup) {
	advisorProdRouter := Router.Group("advisor/prod").Use(middleware.OperationRecord()).Use(middleware.ErrorHandler())
	advisorProdRouterWithoutRecord := Router.Group("advisor/prod").Use(middleware.OperationRecord()).Use(middleware.ErrorHandler())
	{
		advisorProdRouter.POST("", advisorProdApi.CreateProd)
		advisorProdRouter.PUT(":id", advisorProdApi.UpdateProd)
		advisorProdRouter.DELETE(":id", advisorProdApi.DeleteProd)

		advisorProdRouterWithoutRecord.GET("page", advisorProdApi.PageQuery)
		advisorProdRouterWithoutRecord.GET(":advisorProdId", advisorProdApi.GetAdvisorProdById)
	}

	advisorProdFrontRouterWithoutRecord := FrontPrivateGroup.Group("").Use(middleware.ErrorHandler())
	{
		advisorProdFrontRouterWithoutRecord.GET("advisor/prods/:advisorId", advisorProdApi.ListByAdvisorId)
	}
}
