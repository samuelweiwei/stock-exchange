package utils

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"math"
	"strings"
)

// RoundFloat 将浮点数四舍五入到指定小数位
func RoundFloat(val float64, precision int) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

// RoundDecimal 将 Decimal 四舍五入到指定小数位
func RoundDecimal(val decimal.Decimal, precision int32) decimal.Decimal {
	return val.Round(precision)
}

// TruncateFloat 将浮点数向下取整到指定小数位
func TruncateFloat(val float64, precision int) float64 {
	ratio := math.Pow(10, float64(precision))
	truncated := math.Floor(val * ratio)
	return truncated / ratio
}

// TruncateDecimal 将 decimal.Decimal 向下截取到指定小数位
func TruncateDecimal(val decimal.Decimal, precision int32) decimal.Decimal {
	// 计算 10^precision
	scale := decimal.NewFromFloat(1).Shift(precision)

	// 放大后取整（舍去多余小数部分），然后再缩小
	truncated := val.Mul(scale).Truncate(0) // Truncate(0) 只保留整数部分
	return truncated.Div(scale)
}

// GetDecimalPlaces 获取小数的小数位数
func GetDecimalPlaces(decimal float64) int {
	// 转换为字符串
	str := fmt.Sprintf("%g", decimal) // 使用%g格式去除尾部的0

	// 查找小数点位置
	dotIndex := strings.Index(str, ".")
	if dotIndex == -1 {
		return 0
	}

	return len(str) - dotIndex - 1
}

// GetDecimalPlaces 获取 Decimal 的小数位数
func GetDecimalPlaces2(d decimal.Decimal) int {
	// 将 decimal 转换为字符串
	str := d.String()

	// 找到小数点的位置
	if point := strings.Index(str, "."); point != -1 {
		// 去掉末尾的零
		str = strings.TrimRight(str[point+1:], "0")
		return len(str)
	}
	return 0
}

// FloorFloat 将浮点数向下取整到指定小数位
func FloorFloat(val float64, precision int) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Floor(val*ratio) / ratio
}

// FloorDecimal 将 decimal.Decimal 向下取整到指定小数位
func FloorDecimal(val decimal.Decimal, precision int32) (decimal.Decimal, error) {
	dLength := GetDecimalPlaces2(val)
	// 如果小数位数超过要求，返回错误
	if dLength > int(precision) {
		return decimal.NewFromInt(0), errors.New(fmt.Sprintf("操作金额精度超过限制：当前操作金额为 %s，最多允许%d位小数", val.String(), precision))
	}
	// 计算 10^precision
	scale := decimal.NewFromFloat(1).Shift(precision)
	// 放大后取整，然后再缩小
	return val.Mul(scale).Floor().Div(scale), nil
}

func FloorDecimalNoValidate(val decimal.Decimal, precision int32) decimal.Decimal {

	// 计算 10^precision
	scale := decimal.NewFromFloat(1).Shift(precision)
	// 放大后取整，然后再缩小
	return val.Mul(scale).Floor().Div(scale)
}
