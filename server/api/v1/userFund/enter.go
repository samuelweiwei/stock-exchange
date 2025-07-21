package userFund

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service"
)

type ApiGroup struct {
	UserFundAccountsApi
	RechargeRecordsApi
	RechargeChannelsApi
	UserAccountFlowApi
	WithdrawChannelsApi
	WithdrawRecordsApi
	CurrenciesApi
}

var (
	userFundAccountsService = service.ServiceGroupApp.UserfundServiceGroup.UserFundAccountsService
	rechargeRecordsService  = service.ServiceGroupApp.UserfundServiceGroup.RechargeRecordsService
	rechargeChannelsService = service.ServiceGroupApp.UserfundServiceGroup.RechargeChannelsService
	userAccountFlowService  = service.ServiceGroupApp.UserfundServiceGroup.UserAccountFlowService
	withdrawChannelsService = service.ServiceGroupApp.UserfundServiceGroup.WithdrawChannelsService
	withdrawRecordsService  = service.ServiceGroupApp.UserfundServiceGroup.WithdrawRecordsService
	frontendUsersService    = service.ServiceGroupApp.UserServiceGroup.FrontendUsersService
	currenciesService       = service.ServiceGroupApp.UserfundServiceGroup.CurrenciesService
	systemConfigService     = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
	userFollowOrderService  = service.ServiceGroupApp.OrderServiceGroup.UserFollowOrderService
	contractAccountService  = service.ServiceGroupApp.ContractServiceGroup.ContractAccountService
	contractPositionService = service.ServiceGroupApp.ContractServiceGroup.ContractPositionService
)
