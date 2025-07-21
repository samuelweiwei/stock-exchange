package fund

import (
	"errors"
	"fmt"
)

// ActionType 定义资金变动操作类型的枚举
type ActionType string

const (
	Recharge                ActionType = "1"  // 充值
	Withdraw                ActionType = "2"  // 提现
	TransferToContract      ActionType = "3"  // 划转给合约账户
	TransferFromContract    ActionType = "4"  // 合约账户划转进入资金账户
	TradeFollow             ActionType = "5"  // 跟单
	CancelTradeFollow       ActionType = "6"  // 取消跟单
	OperationRefused        ActionType = "7"  // 运营拒绝
	SettleProfit            ActionType = "8"  // 提取盈利
	AutoSettle              ActionType = "9"  // 自动结算
	ProfitSharing           ActionType = "10" // 分润
	SysSend                 ActionType = "11" // 系统赠送
	ApplyOrderFollow        ActionType = "12" // 申请追单
	RefusedApplyOrderFollow ActionType = "13" // 审核拒绝（申请追单）
	StakeEarnProduct        ActionType = "14" // 质押产品
	RedeemEarnProduct       ActionType = "15" // 赎回质押产品
	ApplyWithdraw           ActionType = "16" //申请提现
)

// 打印操作类型的示例
func printActionType(action ActionType) {
	switch action {
	case Recharge:
		fmt.Println("充值操作：快捷充值")
	case Withdraw:
		fmt.Println("充值操作：系统充值")
	case TransferToContract:
		fmt.Println("资金操作：划转给合约账户")
	case TransferFromContract:
		fmt.Println("资金操作：合约账户划转进入资金账户")
	case TradeFollow:
		fmt.Println("资金操作：跟单")
	case SettleProfit:
		fmt.Println("资金操作：盈利结算")
	default:
		fmt.Println("未知操作")
	}
}

// 根据字符串获取对应的 ActionType
func GetActionTypeFromString(value string) (ActionType, error) {
	// 定义一个映射，将字符串映射到 ActionType
	actionTypes := map[string]ActionType{
		"1":  Recharge,
		"2":  Withdraw,
		"3":  TransferToContract,
		"4":  TransferFromContract,
		"5":  TradeFollow,
		"6":  CancelTradeFollow,
		"7":  OperationRefused,
		"8":  SettleProfit,
		"9":  AutoSettle,
		"10": ProfitSharing,
		"11": SysSend,
		"12": ApplyOrderFollow,
		"13": RefusedApplyOrderFollow,
		"14": StakeEarnProduct,
		"15": RedeemEarnProduct,
		"16": ApplyWithdraw,
	}

	// 查找对应的 ActionType
	if action, exists := actionTypes[value]; exists {
		return action, nil
	}

	// 如果没有找到对应的 ActionType，返回错误
	return "", errors.New("无效的 ActionType 值")
}

var (
	BalanceDoesNotEnoughError = errors.New("balance does not enough")
)
