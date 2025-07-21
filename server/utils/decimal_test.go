package utils

import (
	"github.com/shopspring/decimal"
	"testing"
)

func TestDecimal(t *testing.T) {
	s1 := GetScale(0.01)
	if s1 != 2 {
		t.Errorf("GetScale returned %d,expect %d", s1, 2)
	}
	s2 := GetScale(1.0000)
	if s2 != 0 {
		t.Errorf("GetScale returned %d, expect %d", s2, 0)
	}

	d := decimal.Decimal{}
	println(d.InexactFloat64() == 0)
}
