package userfund

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RechargeRecordsRouter struct{}

// InitRechargeRecordsRouter 初始化 rechargeRecords表 路由信息
func (s *RechargeRecordsRouter) InitRechargeRecordsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup, PrivateGroupFrontUser *gin.RouterGroup) {
	rechargeRecordsRouter := Router.Group("rechargeRecords").Use(middleware.OperationRecord())
	rechargeRecordsRouterWithoutRecord := Router.Group("rechargeRecords")
	rechargeRecordsRouterWithoutAuth := PublicRouter.Group("rechargeRecords")
	privateGroupFrontUser := PrivateGroupFrontUser.Group("rechargeRecords")

	{
		rechargeRecordsRouter.POST("createRechargeRecords", rechargeRecordsApi.CreateRechargeRecords)             // 新建rechargeRecords表
		rechargeRecordsRouter.DELETE("deleteRechargeRecords", rechargeRecordsApi.DeleteRechargeRecords)           // 删除rechargeRecords表
		rechargeRecordsRouter.DELETE("deleteRechargeRecordsByIds", rechargeRecordsApi.DeleteRechargeRecordsByIds) // 批量删除rechargeRecords表
		rechargeRecordsRouter.PUT("updateRechargeRecords", rechargeRecordsApi.UpdateRechargeRecords)              // 更新rechargeRecords表
		rechargeRecordsRouter.POST("reviewRecord", rechargeRecordsApi.ReviewRecord)                               //审核充值订单
		rechargeRecordsRouter.POST("lock", rechargeRecordsApi.LockRecord)                                         //锁定/解锁
		rechargeRecordsRouter.GET("getRechargeRecordsList", rechargeRecordsApi.GetRechargeRecordsList)            // 获取rechargeRecords表列表
	}
	{
		rechargeRecordsRouterWithoutRecord.GET("findRechargeRecords", rechargeRecordsApi.FindRechargeRecords) // 根据ID获取rechargeRecords表
	}
	{
		rechargeRecordsRouterWithoutAuth.GET("getRechargeRecordsPublic", rechargeRecordsApi.GetRechargeRecordsPublic) // rechargeRecords表开放接口
		rechargeRecordsRouterWithoutAuth.POST("payNotify/:paymentId", rechargeRecordsApi.PayNotify)
	}
	{
		privateGroupFrontUser.GET("getUserRecordsList", rechargeRecordsApi.GetUserRecordsList) // 根据用户ID获取充值记录
		privateGroupFrontUser.GET("getUserRecordDetail/:id", rechargeRecordsApi.GetUserRecordDetail)
	}
}
