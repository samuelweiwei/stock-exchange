package contract

import (
	. "github.com/shopspring/decimal"
	"strings"
)

type ServiceGroup struct {
	ContractAccountService
	ContractOrderService
	ContractPositionService
	ContractEntrustService
	ContractLeverageService
}

func decimalPlaces(num Decimal) int {
	str := num.String()
	index := strings.Index(str, ".")
	if index == -1 {
		return 0
	}
	return len(str) - index - 1
}
