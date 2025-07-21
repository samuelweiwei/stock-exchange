package userfund

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	UserFundAccountsRouter
	RechargeRecordsRouter
	RechargeChannelsRouter
	UserAccountFlowRouter
	WithdrawChannelsRouter
	WithdrawRecordsRouter
	CurrenciesRouter
}

var (
	userFundAccountsApi = api.ApiGroupApp.UserfundApiGroup.UserFundAccountsApi
	rechargeRecordsApi  = api.ApiGroupApp.UserfundApiGroup.RechargeRecordsApi
	rechargeChannelsApi = api.ApiGroupApp.UserfundApiGroup.RechargeChannelsApi
	userAccountFlowApi  = api.ApiGroupApp.UserfundApiGroup.UserAccountFlowApi
	withdrawChannelsApi = api.ApiGroupApp.UserfundApiGroup.WithdrawChannelsApi
	withdrawRecordsApi  = api.ApiGroupApp.UserfundApiGroup.WithdrawRecordsApi
	currenciesApi       = api.ApiGroupApp.UserfundApiGroup.CurrenciesApi
)
