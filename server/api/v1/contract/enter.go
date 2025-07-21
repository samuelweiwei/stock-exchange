package contract

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	ContractAccountApi
	ContractOrderApi
	ContractPositionApi
	ContractEntrustApi
	ContractLeverageApi
}

var (
	contractAccountService  = service.ServiceGroupApp.ContractServiceGroup.ContractAccountService
	contractOrderService    = service.ServiceGroupApp.ContractServiceGroup.ContractOrderService
	contractPositionService = service.ServiceGroupApp.ContractServiceGroup.ContractPositionService
	contractEntrustService  = service.ServiceGroupApp.ContractServiceGroup.ContractEntrustService
	contractLeverageService = service.ServiceGroupApp.ContractServiceGroup.ContractLeverageService
)
