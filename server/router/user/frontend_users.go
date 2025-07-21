package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FrontendUsersRouter struct{}

// InitFrontendUsersRouter 初始化 frontendUsers表 路由信息
func (s *FrontendUsersRouter) InitFrontendUsersRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup, PrivateGroupFrontUser *gin.RouterGroup) {
	frontendUsersRouter := Router.Group("users").Use(middleware.OperationRecord())
	frontendUsersRouterWithoutRecord := Router.Group("users")
	frontendUsersRouterWithoutAuth := PublicRouter.Group("users")
	privateGroupFrontUser := PrivateGroupFrontUser.Group("users")
	{
		frontendUsersRouter.POST("createFrontendUsers", frontendUsersApi.CreateFrontendUsers)             // 新建frontendUsers表
		frontendUsersRouter.DELETE("deleteFrontendUsers", frontendUsersApi.DeleteFrontendUsers)           // 删除frontendUsers表
		frontendUsersRouter.DELETE("deleteFrontendUsersByIds", frontendUsersApi.DeleteFrontendUsersByIds) // 批量删除frontendUsers表
		frontendUsersRouter.PUT("updateFrontendUsers", frontendUsersApi.UpdateFrontendUsers)              // 更新frontendUsers表
		frontendUsersRouter.POST("realNameAuthentication", frontendUsersApi.RealNameAuthentication)       // 实名认证审核
		frontendUsersRouter.POST("changeParent", frontendUsersApi.ChangeParent)                           // 修改上级代理
		frontendUsersRouter.POST("updateUserPassword", frontendUsersApi.UpdateUserPassword)               // 修改用户密码
	}
	{
		frontendUsersRouterWithoutRecord.GET("findFrontendUsers", frontendUsersApi.FindFrontendUsers)       // 根据ID获取frontendUsers表
		frontendUsersRouterWithoutRecord.GET("getFrontendUsersList", frontendUsersApi.GetFrontendUsersList) // 获取frontendUsers表列表

	}
	// public 接口
	{
		frontendUsersRouterWithoutAuth.POST("login", frontendUsersApi.UserLogin)
		frontendUsersRouterWithoutAuth.POST("register", frontendUsersApi.UserRegister)
		frontendUsersRouterWithoutAuth.POST("resetLoginPassword", frontendUsersApi.ResetLoginPassword)
		PublicRouter.POST("/base/sendCaptcha", frontendUsersApi.SendCaptcha)

	}

	// 前台用户接口
	{
		privateGroupFrontUser.POST("bindEmail", frontendUsersApi.BindEmail)                         // 绑定邮箱
		privateGroupFrontUser.POST("bindPhone", frontendUsersApi.BindPhone)                         // 绑定手机号
		privateGroupFrontUser.POST("changePaymentPassword", frontendUsersApi.ChangePaymentPassword) // 用户修改支付密码
		privateGroupFrontUser.POST("changeLoginPassword", frontendUsersApi.ChangeLoginPassword)     // 用户修改登录密码
		privateGroupFrontUser.GET("getAncestors", frontendUsersApi.GetAncestors)                    // 获取用户三级代理 [1,2,3] , 1为上级，2为上上级，3为上上上级
		privateGroupFrontUser.GET("getFrontUserInfo", frontendUsersApi.GetFrontendUsers)            // 获取用户基本信息
		privateGroupFrontUser.POST("userIdentity", frontendUsersApi.UserIdentity)                   // 身份认证
		privateGroupFrontUser.GET("getSubUserList", frontendUsersApi.GetSubUserList)                // 获取下级
		privateGroupFrontUser.GET("getTeamCount", frontendUsersApi.GetTeamCount)                    // 获取团队总人数
		privateGroupFrontUser.POST("updateUserInfo", frontendUsersApi.UpdateUserInfo)               // 修改用户基本信息
		privateGroupFrontUser.GET("teamCount", frontendUsersApi.TeamCount)                          // 团队人数
	}
}
