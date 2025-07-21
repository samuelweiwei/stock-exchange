package constants

// Role 定义用户角色类型
type NotificationChannel uint8

// 定义枚举值
const (
	ChannelInvalid      NotificationChannel = iota // 0
	ChannelPhoneMessage                            // 1
	ChannelEmail
)

// Value 使用 Value 方法返回对应的角色值
func (r NotificationChannel) Value() string {
	switch r {
	case ChannelPhoneMessage:
		return "phoneMessage"
	case ChannelEmail:
		return "email"
	default:
		return ""
	}
}

// ParseNotificationChannel 将 int 解析为对应的 Role 枚举
func ParseNotificationChannel(input string) NotificationChannel {
	switch input {
	case "email":
		return ChannelEmail
	case "phoneMessage":
		return ChannelPhoneMessage
	default:
		return ChannelInvalid
	}
}
