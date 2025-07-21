package constants

type UserType uint8

// 定义枚举值
const (
	UserTypeInvalid UserType = iota // 0
	UserTypeAdmin                   // 1
	UserTypeFront
)
