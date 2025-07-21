package utils

import (
	decimal2 "github.com/govalues/decimal"
	"github.com/shopspring/decimal"
	"strconv"
)

var HundredPercent = decimal.NewFromInt(100)

// ToFloat64Fix 将decimal转换为float64类型，并且保留places位数
func ToFloat64Fix(d decimal.Decimal, places int32) float64 {
	f, _ := strconv.ParseFloat(d.StringFixed(places), 64)
	return f
}

// GetScale 获取小数位数
func GetScale(f float64) int {
	d, _ := decimal2.NewFromFloat64(f)
	return d.Scale()
}
