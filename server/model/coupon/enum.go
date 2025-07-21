package coupon

import (
	"encoding/json"
	"fmt"
)

type Type int

func (t *Type) String() string {
	return fmt.Sprint(*t)
}
func (t *Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(*t))
}
func (t *Type) UnmarshalJSON(b []byte) error {
	var (
		err error
		n   int
	)
	err = json.Unmarshal(b, &n)
	*t = Type(n)
	return err
}

const (
	UnKnown Type = iota
	Register
	RealName
	Manual
)

type IssuedCouponStatus int

const (
	NotUsed IssuedCouponStatus = iota
	AlreadyUsed
	Expired
)

func (t *IssuedCouponStatus) String() string {
	return fmt.Sprint(*t)
}
func (t *IssuedCouponStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(*t))
}
func (t *IssuedCouponStatus) UnmarshalJSON(b []byte) error {
	var (
		err error
		n   int
	)
	err = json.Unmarshal(b, &n)
	*t = IssuedCouponStatus(n)
	return err
}
