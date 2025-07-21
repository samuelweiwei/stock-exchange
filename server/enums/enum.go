package enums

// 定义充值类型枚举
const (
	RechargeTypeSystem = "1" // 系统充值
	RechargeTypeQuick  = "2" // 快捷充值

	WithDrawTypeSystem = "1" // 系统提现
	WithDrawTypeQuick  = "2" // 快捷提现

	UserAction_SUBMIT = "SUBMIT"

	UN_CHECK = "1" //待审核
	CHECKED  = "2" //审核

	IS_LOCK = "1" //已锁定
	UNLOCK  = "0" //未锁定

	//订单状态
	PENDING = "1" // 支付中
	SUCCESS = "2" // 支付成功
	FAILED  = "3" // 支付失败
	OUTTIME = "4" //支付超时

	//提现订单状态
	WITHDRAWING      = "1" // 发起提现，等待审核
	WITHDRAW_CHECKED = "2" // 审核完成
	WITHDRAWED       = "3" //提现完成

	CHECKING = "1" //审核中
	REFUSED  = "2" //拒绝
	PASSED   = "3" //审核通过
	INVALID  = "4" //审核无效（目前提现失败使用）

	WithdrawCommissionQuota = 1
	WithdrawCommissionRate  = 2
)

func GetRechargeStatusString(status string) string {
	switch status {
	case PENDING:
		return "支付中"
	case SUCCESS:
		return "支付成功"
	case FAILED:
		return "支付失败"
	case OUTTIME:
		return "支付超时"
	}
	return "无此状态：" + status
}

func GetWithdrawStatusString(status string) string {
	switch status {
	case WITHDRAWING:
		return "等待审核"
	case WITHDRAW_CHECKED:
		return "审核完成"
	case WITHDRAWED:
		return "提现完成"
	}
	return "无此状态：" + status
}
