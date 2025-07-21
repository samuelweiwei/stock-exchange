package symbol

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	SymbolsRouter
	SymbolsCustomRouter
	SymbolsHotRouter
}

var (
	symbolsApi       = api.ApiGroupApp.SymbolApiGroup.SymbolsApi
	symbolsCustomApi = api.ApiGroupApp.SymbolApiGroup.SymbolsCustomApi
	symbolsHotApi    = api.ApiGroupApp.SymbolApiGroup.SymbolsHotApi
)
