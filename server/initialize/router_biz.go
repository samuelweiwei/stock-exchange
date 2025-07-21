package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
)

func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}
func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]
	privateGroupFrontUser := routers[2]
	holder(publicGroup, privateGroup)
	{
		orderRouter := router.RouterGroupApp.Order
		orderRouter.InitAdvisorRouter(privateGroup, publicGroup)
		orderRouter.InitAdvisorProdRouter(privateGroup, publicGroup, privateGroupFrontUser)
		orderRouter.InitAdvisorStockOrderRouter(privateGroup, publicGroup)
		orderRouter.InitUserFollowOrderAdminRouter(privateGroup, publicGroup)
		orderRouter.InitUserFollowOrderH5Router(privateGroupFrontUser, publicGroup)
		orderRouter.InitSystemOrderMsgRouter(privateGroup, publicGroup)
		orderRouter.InitUserFollowAppendOrderAdminRouter(privateGroup, publicGroup)
		orderRouter.InitUserFollowAppendOrderH5Router(privateGroupFrontUser, publicGroup)
		orderRouter.InitUserProfitShareH5Router(privateGroupFrontUser, publicGroup)
	}
	{
		contractRouter := router.RouterGroupApp.Contract
		contractRouter.InitContractAccountRouter(privateGroup, publicGroup, privateGroupFrontUser)
		contractRouter.InitContractOrderRouter(privateGroup, publicGroup, privateGroupFrontUser)
		contractRouter.InitContractPositionRouter(privateGroup, publicGroup, privateGroupFrontUser)
		contractRouter.InitContractEntrustRouter(privateGroup, publicGroup, privateGroupFrontUser)
		contractRouter.InitContractLeverageRouter(privateGroup, publicGroup, privateGroupFrontUser)
	}
	{
		userfundRouter := router.RouterGroupApp.Userfund
		userfundRouter.InitUserFundAccountsRouter(privateGroup, publicGroup, privateGroupFrontUser)
		userfundRouter.InitRechargeRecordsRouter(privateGroup, publicGroup, privateGroupFrontUser)
		userfundRouter.InitRechargeChannelsRouter(privateGroup, publicGroup, privateGroupFrontUser)
		userfundRouter.InitUserAccountFlowRouter(privateGroup, publicGroup, privateGroupFrontUser)
		userfundRouter.InitWithdrawChannelsRouter(privateGroup, publicGroup, privateGroupFrontUser)
		userfundRouter.InitWithdrawRecordsRouter(privateGroup, publicGroup, privateGroupFrontUser)
		userfundRouter.InitCurrenciesRouter(privateGroup, privateGroupFrontUser)
	}
	{
		userRouter := router.RouterGroupApp.User
		userRouter.InitCountriesRouter(privateGroup, publicGroup)
		userRouter.InitFrontendUsersRouter(privateGroup, publicGroup, privateGroupFrontUser)
		userRouter.InitFrontendUserLoginLogRouter(privateGroup)
	}
	{
		couponRouter := router.RouterGroupApp.Coupon
		couponRouter.InitCouponRouter(privateGroup, publicGroup)
		couponRouter.InitCouponIssuedRouter(privateGroup, publicGroup)
		couponRouter.InitFrontCouponIssuedRouter(privateGroup, privateGroupFrontUser)
	}
	{
		i18nRouter := router.RouterGroupApp.I18n
		i18nRouter.InitSysI18nLocalizeConfigRouter(privateGroup, publicGroup)
	}
	{
		symbolRouter := router.RouterGroupApp.Symbol
		symbolRouter.InitSymbolsRouter(privateGroup, publicGroup)
		symbolRouter.InitSymbolsCustomRouter(privateGroup, publicGroup, privateGroupFrontUser)
		symbolRouter.InitSymbolsHotRouter(privateGroup, publicGroup)
	}
	{
		reportRouter := router.RouterGroupApp.Report
		reportRouter.InitPlatformReportRouter(privateGroup, publicGroup)
	}
	{
		earnRouter := router.RouterGroupApp.Earn
		earnRouter.InitEarnInterestRatesRouter(privateGroup, publicGroup)
		earnRouter.InitEarnProductsRouter(privateGroup, publicGroup)
		earnRouter.InitEarnDailyIncomeMoneyLogRouter(privateGroup, publicGroup)
		earnRouter.InitEarnSubscribeLogRouter(privateGroup, publicGroup)
		earnRouter.InitFrontEarnProductRouter(privateGroup, privateGroupFrontUser)
		earnRouter.InitFrontEarnDailyIncomeMoneyLogRouter(privateGroup, privateGroupFrontUser)
		earnRouter.InitFrontEarnSubscribeLogRouter(privateGroup, privateGroupFrontUser)
	}
	{
		settingManageRouter := router.RouterGroupApp.SettingManage
		settingManageRouter.InitServiceLinkRouter(privateGroup, publicGroup)
	}
}
