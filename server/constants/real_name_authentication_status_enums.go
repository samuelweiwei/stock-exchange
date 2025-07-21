package constants

// RealNameAuthenticationStatus 用户实名审核状态
type RealNameAuthenticationStatus uint8

// 定义枚举值
const (
	AuthenticationStatusInvalid RealNameAuthenticationStatus = iota // 0
	AuthenticationStatusPending                                     // 1
	AuthenticationStatusRejected
	AuthenticationStatusPassed
)

func ParseRealNameAuthenticationStatus(input int) RealNameAuthenticationStatus {
	switch input {
	case 1:
		return AuthenticationStatusPending
	case 2:
		return AuthenticationStatusRejected
	case 3:
		return AuthenticationStatusPassed
	default:
		return AuthenticationStatusInvalid
	}
}
