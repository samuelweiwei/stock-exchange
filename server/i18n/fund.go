package i18n

var (
	//FUND_UPDATE_FAILED = "FundUpdateFailed"
	FundUpdateFailed          = "FundUpdateFailed"
	FundQueryFailed           = "FundQueryFailed"
	FloatNotPatchError        = "FloatNotPatchError"
	RechargeMoutMinFaild      = "RechargeMoutMinFaild"
	CurrencyQueryFailed       = "CurrencyQueryFailed"       // 货币信息获取失败
	CalculateUsdtFaild        = "CalculateUsdtFaild"        // 充值资金转换USDT失败
	RechargeRecordCreateFaild = "RechargeRecordCreateFaild" // 充值订单创建失败
	PayPwdVerifyFaild         = "PayPwdVerifyFaild"         // 支付密码验证失败
	PayPwdMismatchError       = "PayPwdMismatchError"       //支付密码错误
	FundAmountNotEnough       = "FundAmountNotEnough"       // 资金账户余额不足
	WithdrawAmoutCheckFaild   = "WithdrawAmoutCheckFaild"   // 提现数量不满足提现条件
	ThirdRequestError         = "ThirdRequestError"         // 请求第三方通道报错
	WithdrawRecordCreateFaild = "WithdrawRecordCreateFaild" // 提现订单创建失败
	AccountFlowQueryFailed    = "AccountFlowQueryFailed"    // 流水查询失败
	AccountFlowSaveFailed     = "AccountFlowSaveFailed"     // 流水新增失败
	NotifyBindError           = "NotifyBindError"           //  回调解析失败
	PayNotifyUpdateFundError  = "PayNotifyUpdateFundError"  // 支付回调更新资金失败
	ActionAuthCheckFailed     = "ActionAuthCheckFailed"
)
