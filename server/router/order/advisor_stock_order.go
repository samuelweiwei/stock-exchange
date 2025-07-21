package order

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AdvisorStockOrderRouter struct {
}

func (router *AdvisorStockOrderRouter) InitAdvisorStockOrderRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	advisorStockOrderRouterWithRecord := Router.Group("advisor/stock/order").Use(middleware.OperationRecord()).Use(middleware.ErrorHandler())
	{
		advisorStockOrderRouterWithRecord.POST("", advisorStockOrderApi.CreateStockOrder)
		advisorStockOrderRouterWithRecord.POST("auto-follow/:stockOrderId", advisorStockOrderApi.AutoFollow)
		advisorStockOrderRouterWithRecord.POST("confirm/:stockOrderId", advisorStockOrderApi.SubmitSellConfirm)
	}

	advisorStockOrderRouterWithoutRecord := Router.Group("advisor/stock/order").Use(middleware.ErrorHandler())
	{
		advisorStockOrderRouterWithoutRecord.GET("/confirm/:stockOrderId", advisorStockOrderApi.GetSellConfirmDetail)
		advisorStockOrderRouterWithoutRecord.GET("page", advisorStockOrderApi.PageQuery)
	}

}
