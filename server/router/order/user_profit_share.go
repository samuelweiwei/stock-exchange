package order

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserProfitShareRouter struct {
}

func (r *UserProfitShareRouter) InitUserProfitShareH5Router(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	routerWithoutRecord := Router.Group("user/profit/share").Use(middleware.ErrorHandler())
	{
		routerWithoutRecord.GET("my", userProfitShareApi.QueryMyProfitShareRecord)
	}
}
