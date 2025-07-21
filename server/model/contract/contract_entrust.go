// 自动生成模板ContractEntrust
package contract

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	. "github.com/shopspring/decimal"
)

type TriggerType int
type EntrustStatus int

// 限价-1，止盈-2，止损-3
const (
	Limit      TriggerType = 1
	TakeProfit TriggerType = 2
	StopLoss   TriggerType = 3
)

// 未触发-1，已触发-2,已取消-99
const (
	Untriggered EntrustStatus = 1
	Triggered   EntrustStatus = 2
	Deleted     EntrustStatus = 99
)

// contractEntrust表 结构体  ContractEntrust
type ContractEntrust struct {
	global.GVA_MODEL
	UserId        uint          `json:"userId" form:"userId" gorm:"column:user_id;comment:关联用户表的用户 ID;size:20;"`               //关联用户表的用户 ID
	PositionId    uint          `json:"positionId" form:"positionId" gorm:"column:position_id;comment:关联持仓表的持仓 ID;size:20;"`   //关联持仓表的持仓 ID
	OrderId       uint          `json:"orderId" form:"orderId" gorm:"column:order_id;comment:关联订单表的订单 ID;size:20;"`            //关联订单表的订单 ID
	StockId       uint          `json:"stockId" form:"stockId" gorm:"column:stock_id;comment:股票id;size:20;"`                   //股票id
	StockName     string        `json:"stockName" form:"stockName" gorm:"column:stock_name;comment:股票名称;size:255;"`            //股票名称
	TriggerType   TriggerType   `json:"triggerType" form:"triggerType" gorm:"column:trigger_type;comment:触发类型;size:10;"`       //触发类型 限价-1，止盈-2，止损-3
	TriggerPrice  float64       `json:"triggerPrice" form:"triggerPrice" gorm:"column:trigger_price;comment:触发价格;size:10;"`    //触发价格
	Margin        Decimal       `json:"margin" form:"margin" gorm:"column:margin;comment:保证金;size:10;"`                        //保证金
	OperationType OperationType `json:"operationType" form:"operationType" gorm:"column:operation_type;comment:操作类型;size:10;"` //操作类型 开多-1，开空-2，平多-3，平空-4
	Quantity      float64       `json:"quantity" form:"quantity" gorm:"column:quantity;comment:数量;size:10;"`                   //数量
	EntrustStatus EntrustStatus `json:"entrustStatus" form:"entrustStatus" gorm:"column:entrust_status;comment:委托状态;size:10;"` //委托状态 未触发-1，已触发-2,已取消-99
	CreatedBy     uint          `gorm:"column:created_by;comment:创建者"`
	UpdatedBy     uint          `gorm:"column:updated_by;comment:更新者"`
	DeletedBy     uint          `gorm:"column:deleted_by;comment:删除者"`
}

// TableName contractEntrust表 ContractEntrust自定义表名 contract_entrust
func (ContractEntrust) TableName() string {
	return "contract_entrust"
}
