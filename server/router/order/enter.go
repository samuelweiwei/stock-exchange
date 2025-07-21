package order

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	AdvisorRouter
	AdvisorProdRouter
	AdvisorStockOrderRouter
	UserFollowOrderRouter
	UserFollowAppendOrderRouter
	SystemOrderMsgRouter
	UserProfitShareRouter
}

var (
	advisorApi               = api.ApiGroupApp.OrderApiGroup.AdvisorApi
	advisorProdApi           = api.ApiGroupApp.OrderApiGroup.AdvisorProdApi
	advisorStockOrderApi     = api.ApiGroupApp.OrderApiGroup.AdvisorStockOrderApi
	userFollowOrderAdminApi  = api.ApiGroupApp.OrderApiGroup.UserFollowOrderAdminApi
	userFollowOrderH5Api     = api.ApiGroupApp.OrderApiGroup.UserFollowOrderH5Api
	userFollowAppendOrderApi = api.ApiGroupApp.OrderApiGroup.UserFollowAppendOrderApi
	systemOrderMsgApi        = api.ApiGroupApp.OrderApiGroup.SystemOrderMsgApi
	userProfitShareApi       = api.ApiGroupApp.OrderApiGroup.UserProfitShareApi
)
