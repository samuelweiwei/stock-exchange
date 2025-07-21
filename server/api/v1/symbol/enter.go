package symbol

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	SymbolsApi
	SymbolsCustomApi
	SymbolsHotApi
}

var (
	symbolsService       = service.ServiceGroupApp.SymbolServiceGroup.SymbolsService
	symbolsCustomService = service.ServiceGroupApp.SymbolServiceGroup.SymbolsCustomService
	symbolsHotService    = service.ServiceGroupApp.SymbolServiceGroup.SymbolsHotService
)
