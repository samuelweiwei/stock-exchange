package order

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	AdvisorApi
	AdvisorProdApi
	AdvisorStockOrderApi
	UserFollowOrderAdminApi
	UserFollowOrderH5Api
	UserFollowAppendOrderApi
	SystemOrderMsgApi
	UserProfitShareApi
}

var (
	advisorService               = service.ServiceGroupApp.OrderServiceGroup.AdvisorService
	advisorProdService           = service.ServiceGroupApp.OrderServiceGroup.AdvisorProdService
	advisorStockOrderService     = service.ServiceGroupApp.OrderServiceGroup.AdvisorStockOrderService
	userFollowOrderService       = service.ServiceGroupApp.OrderServiceGroup.UserFollowOrderService
	userFollowAppendOrderService = service.ServiceGroupApp.OrderServiceGroup.UserFollowAppendOrderService
	systemOrderMsgService        = service.ServiceGroupApp.OrderServiceGroup.SystemOrderMsgService
	userProfitShareService       = service.ServiceGroupApp.OrderServiceGroup.UserProfitShareService
)
