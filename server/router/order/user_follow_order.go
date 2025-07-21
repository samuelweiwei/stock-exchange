package order

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserFollowOrderRouter struct{}

func (router *UserFollowOrderRouter) InitUserFollowOrderAdminRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	adminRouterWithoutRecord := Router.Group("user/follow/order").Use(middleware.ErrorHandler())
	adminRouterWithRecord := Router.Group("user/follow/order").Use(middleware.OperationRecord()).Use(middleware.ErrorHandler())
	{
		adminRouterWithoutRecord.GET("page", userFollowOrderAdminApi.PageQuery)                                //分页查询用户跟单列表
		adminRouterWithoutRecord.GET("confirm/:followOrderId", userFollowOrderAdminApi.GetFollowConfirmDetail) //查询跟单确认详情

		adminRouterWithRecord.POST("approve/:followOrderId", userFollowOrderAdminApi.Approve)             //审核通过
		adminRouterWithRecord.POST("approve-all", userFollowOrderAdminApi.ApproveAll)                     //一键审核
		adminRouterWithRecord.POST("reject/:followOrderId", userFollowOrderAdminApi.Reject)               //审核驳回
		adminRouterWithRecord.POST("confirm/:followOrderId", userFollowOrderAdminApi.SubmitFollowConfirm) //提交跟单确认
	}
}

func (router *UserFollowOrderRouter) InitUserFollowOrderH5Router(PrivateRouter *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	h5RouterWithoutRecord := PrivateRouter.Group("/user/follow/order").Use(middleware.ErrorHandler())
	h5RouterWithRecord := PrivateRouter.Group("/user/follow/order").Use(middleware.OperationRecord()).Use(middleware.ErrorHandler())
	{
		h5RouterWithRecord.POST("apply", userFollowOrderH5Api.Apply)                                 //申请跟单
		h5RouterWithRecord.POST("cancel/:followOrderId", userFollowOrderH5Api.Cancel)                //撤销跟单
		h5RouterWithRecord.POST("retrieve/:followOrderId", userFollowOrderH5Api.RetrieveFollowOrder) //提盈
		h5RouterWithRecord.POST("cancel-renew/:followOrderId", userFollowOrderH5Api.CancelAutoRenew) //取消自动续期

		h5RouterWithoutRecord.GET("page/my", userFollowOrderH5Api.PageQueryMyFollowOrder)                          //分页查询我的跟单列表
		h5RouterWithoutRecord.GET("my/:followOrderId", userFollowOrderH5Api.GetMyFollowOrderDetail)                //查询我的跟单详情
		h5RouterWithoutRecord.GET("profit/my", userFollowOrderH5Api.GetUserProfitSummary)                          //查询我的跟单收益报表
		h5RouterWithoutRecord.GET("total-amount/my", userFollowOrderH5Api.GetTotalAmount)                          //查询我的跟单总金额
		h5RouterWithoutRecord.GET("retrieve/records/:followOrderId", userFollowOrderH5Api.PageQueryRetrieveRecord) //提盈记录分页查询接口
	}
}
