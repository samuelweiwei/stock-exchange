package task

import "github.com/flipped-aurora/gin-vue-admin/server/service"

var (
	earnProductService             = service.ServiceGroupApp.EarnServiceGroup.EarnProductsService
	earnInterestRateService        = service.ServiceGroupApp.EarnServiceGroup.EarnInterestRatesService
	earnProductDailyIncomeService  = service.ServiceGroupApp.EarnServiceGroup.EarnDailyIncomeMoneyLogService
	earnProductSubscriptionService = service.ServiceGroupApp.EarnServiceGroup.EarnSubscribeLogService
)
