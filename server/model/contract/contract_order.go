// 自动生成模板ContractOrder
package contract

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	. "github.com/shopspring/decimal"
	"time"
)

type OrderType int
type OperationType int
type OrderStatus int

// 市价单-1，限价单-2，平仓单-3
const (
	MarketOrder OrderType = 1
	LimitOrder  OrderType = 2
	CloseOrder  OrderType = 3
)

// 开多-1，开空-2，平多-3，平空-4
const (
	OpenLong   OperationType = 1
	OpenShort  OperationType = 2
	CloseLong  OperationType = 3
	CloseShort OperationType = 4
)

// 待成交-1，部分成交-2，完全成交-3，已取消-99
const (
	Pending         OrderStatus = 1
	PartiallyFilled OrderStatus = 2
	FullyFilled     OrderStatus = 3
	Cancelled       OrderStatus = 99
)

// contractOrder表 结构体  ContractOrder
type ContractOrder struct {
	global.GVA_MODEL
	OrderNumber        string        `json:"orderNumber" form:"orderNumber" gorm:"column:order_number;comment:订单号;size:50;"`                         //订单号
	OrderTime          time.Time     `json:"orderTime" form:"orderTime" gorm:"column:order_time;comment:订单时间;"`                                      //订单时间
	UserId             uint          `json:"userId" form:"userId" gorm:"column:user_id;comment:关联用户表的用户 ID;size:20;"`                                //关联用户表的用户 ID
	StockId            uint          `json:"stockId" form:"stockId" gorm:"column:stock_id;comment:股票id;size:20;"`                                    //股票id
	StockName          string        `json:"stockName" form:"stockName" gorm:"column:stock_name;comment:股票名称;size:255;"`                             //股票名称
	OrderType          OrderType     `json:"orderType" form:"orderType" gorm:"column:order_type;comment:下单类型;size:10;"`                              //下单类型
	OpenPrice          float64       `json:"openPrice" form:"openPrice" gorm:"column:open_price;comment:开仓价格;size:10;"`                              //开仓价格
	ClosePrice         float64       `json:"closePrice" form:"closePrice" gorm:"column:close_price;comment:平仓价格;size:10;"`                           //平仓价格
	OperationType      OperationType `json:"operationType" form:"operationType" gorm:"column:operation_type;comment:操作类型;size:10;"`                  //操作类型
	Quantity           float64       `json:"quantity" form:"quantity" gorm:"column:quantity;comment:数量;size:10;"`                                    //数量
	OrderStatus        OrderStatus   `json:"orderStatus" form:"orderStatus" gorm:"column:order_status;comment:订单状态;size:10;"`                        //订单状态
	Fee                Decimal       `json:"fee" form:"fee" gorm:"column:fee;comment:手续费;size:10;"`                                                  //手续费
	RealizedProfitLoss Decimal       `json:"realizedProfitLoss" form:"realizedProfitLoss" gorm:"column:realized_profit_loss;comment:已实现盈亏;size:10;"` //已实现盈亏
	CreatedBy          uint          `gorm:"column:created_by;comment:创建者"`
	UpdatedBy          uint          `gorm:"column:updated_by;comment:更新者"`
	DeletedBy          uint          `gorm:"column:deleted_by;comment:删除者"`
}

// TableName contractOrder表 ContractOrder自定义表名 contract_order
func (ContractOrder) TableName() string {
	return "contract_order"
}
