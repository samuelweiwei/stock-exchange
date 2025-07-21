package userfund

type ServiceGroup struct {
	UserFundAccountsService
	RechargeRecordsService
	RechargeChannelsService
	UserAccountFlowService
	WithdrawChannelsService
	WithdrawRecordsService
	CurrenciesService
}
