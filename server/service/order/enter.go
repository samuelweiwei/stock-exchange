package order

import "github.com/flipped-aurora/gin-vue-admin/server/service/userfund"

type ServiceGroup struct {
	AdvisorService
	AdvisorProdService
	AdvisorStockOrderService
	UserFollowOrderService
	UserFollowAppendOrderService
	SystemOrderMsgService
	UserProfitShareService
}

var userFundService = userfund.UserFundAccountsService{}
