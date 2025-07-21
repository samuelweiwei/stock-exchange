package userfund

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WithdrawRecordsRouter struct{}

// InitWithdrawRecordsRouter 初始化 withdrawRecords表 路由信息
func (s *WithdrawRecordsRouter) InitWithdrawRecordsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup, PrivateGroupFrontUser *gin.RouterGroup) {
	withdrawRecordsRouter := Router.Group("withdrawRecords").Use(middleware.OperationRecord())
	withdrawRecordsRouterWithoutRecord := Router.Group("withdrawRecords")
	withdrawRecordsRouterWithoutAuth := PublicRouter.Group("withdrawRecords")
	privateGroupFrontUser := PrivateGroupFrontUser.Group("withdrawRecords")

	{
		withdrawRecordsRouter.POST("createWithdrawRecords", withdrawRecordsApi.CreateWithdrawRecords)             // 新建withdrawRecords表
		withdrawRecordsRouter.DELETE("deleteWithdrawRecords", withdrawRecordsApi.DeleteWithdrawRecords)           // 删除withdrawRecords表
		withdrawRecordsRouter.DELETE("deleteWithdrawRecordsByIds", withdrawRecordsApi.DeleteWithdrawRecordsByIds) // 批量删除withdrawRecords表
		withdrawRecordsRouter.PUT("updateWithdrawRecords", withdrawRecordsApi.UpdateWithdrawRecords)              // 更新withdrawRecords表
		withdrawRecordsRouter.POST("reviewRecord", withdrawRecordsApi.ReviewRecord)                               //审核充值订单
		withdrawRecordsRouter.POST("lock", withdrawRecordsApi.LockRecord)
		withdrawRecordsRouter.GET("getWithdrawRecordsList", withdrawRecordsApi.GetWithdrawRecordsList) // 获取withdrawRecords表列表
	}
	{
		withdrawRecordsRouterWithoutRecord.GET("findWithdrawRecords", withdrawRecordsApi.FindWithdrawRecords) // 根据ID获取withdrawRecords表

	}
	{
		withdrawRecordsRouterWithoutAuth.GET("getWithdrawRecordsPublic", withdrawRecordsApi.GetWithdrawRecordsPublic) // withdrawRecords表开放接口
		withdrawRecordsRouterWithoutAuth.POST("withdrawNotify/:paymentId", withdrawRecordsApi.WithdrawNotify)
	}
	{
		privateGroupFrontUser.GET("getUserRecordsList", withdrawRecordsApi.GetUserRecordsList) // 根据用户ID获取提现记录
		privateGroupFrontUser.GET("getUserRecordDetail/:id", withdrawRecordsApi.GetUserRecordDetail)

	}
}
