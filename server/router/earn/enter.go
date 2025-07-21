package earn

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	EarnInterestRatesRouter
	EarnProductsRouter
	EarnDailyIncomeMoneyLogRouter
	EarnSubscribeLogRouter
}

var (
	earnInterestRatesApi       = api.ApiGroupApp.EarnApiGroup.EarnInterestRatesApi
	earnProductsApi            = api.ApiGroupApp.EarnApiGroup.EarnProductsApi
	earnDailyIncomeMoneyLogApi = api.ApiGroupApp.EarnApiGroup.EarnDailyIncomeMoneyLogApi
	earnSubscribeLogApi        = api.ApiGroupApp.EarnApiGroup.EarnSubscribeLogApi
)
