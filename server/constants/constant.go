package constants

var (
	// InviteCodeKey 邀请码秘钥
	InviteCodeKey = "2TPN4fApRAeCzWCy"
	SMSAppKey     = "fBz91C"
	SMSAppSecret  = "UQN6Se"
	SMSAppCode    = "1000"
)

var (
	// InitCustomTradingPairsUrl 初始化用户默认的自定义交易对（市值前十）接口地址
	InitCustomTradingPairsUrl      = "/api/symbolsCustom/frontend/initUserSymbols"
	InitCustomTradingPairsAdminUrl = "/api/symbolsCustom/frontend/initUserSymbols/%d"
)
