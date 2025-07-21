package contract

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	ContractAccountRouter
	ContractOrderRouter
	ContractPositionRouter
	ContractEntrustRouter
	ContractLeverageRouter
}

var (
	contractAccountApi  = api.ApiGroupApp.ContractApiGroup.ContractAccountApi
	contractOrderApi    = api.ApiGroupApp.ContractApiGroup.ContractOrderApi
	contractPositionApi = api.ApiGroupApp.ContractApiGroup.ContractPositionApi
	contractEntrustApi  = api.ApiGroupApp.ContractApiGroup.ContractEntrustApi
	contractLeverageApi = api.ApiGroupApp.ContractApiGroup.ContractLeverageApi
)
