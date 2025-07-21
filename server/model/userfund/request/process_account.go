package request

import "github.com/flipped-aurora/gin-vue-admin/server/enums/fund"

// ActionType 是操作类型的枚举
type ActionType string

// AccountOperation 表示账户操作的结构体
type AccountOperation struct {
	UserId  int64           `json:"userId"`  // 账户ID
	Amount  float64         `json:"amount"`  // 操作金额
	OrderId string          `json:"orderId"` //orderId
	Action  fund.ActionType `json:"action"`  // 操作类型（枚举）
}
