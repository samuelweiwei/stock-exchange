package earn

import "github.com/flipped-aurora/gin-vue-admin/server/service"

const (
	earnProductStakeCurrency = "USDT"
)

type ApiGroup struct {
	EarnInterestRatesApi
	EarnProductsApi
	EarnDailyIncomeMoneyLogApi
	EarnSubscribeLogApi
}

var (
	earnInterestRatesService       = service.ServiceGroupApp.EarnServiceGroup.EarnInterestRatesService
	earnProductsService            = service.ServiceGroupApp.EarnServiceGroup.EarnProductsService
	earnDailyIncomeMoneyLogService = service.ServiceGroupApp.EarnServiceGroup.EarnDailyIncomeMoneyLogService
	earnSubscribeLogService        = service.ServiceGroupApp.EarnServiceGroup.EarnSubscribeLogService
	currenciesService              = service.ServiceGroupApp.UserfundServiceGroup.CurrenciesService
)
