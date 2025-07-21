// 自动生成模板ContractPosition
package contract

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	. "github.com/shopspring/decimal"
	"time"
)

type PositionType int
type PositionStatus int

// 多单-1，空单-2
const (
	Long  PositionType = 1
	Short PositionType = 2
)

// 未平仓-1，已平仓-2
const (
	Unclosed PositionStatus = 1
	Closed   PositionStatus = 2
)

// contractPosition表 结构体  ContractPosition
type ContractPosition struct {
	global.GVA_MODEL
	UserId          uint           `json:"userId" form:"userId" gorm:"column:user_id;comment:所在用户表的用户 ID;size:20;"`                      //所在用户表的用户 ID
	StockId         uint           `json:"stockId" form:"stockId" gorm:"column:stock_id;comment:股票id;size:20;"`                          //股票id
	StockName       string         `json:"stockName" form:"stockName" gorm:"column:stock_name;comment:股票名称;size:255;"`                   //股票名称
	PositionTime    time.Time      `json:"positionTime" form:"positionTime" gorm:"column:position_time;comment:持仓时间;"`                   //持仓时间
	Quantity        float64        `json:"quantity" form:"quantity" gorm:"column:quantity;comment:持仓数量;size:10;"`                        //持仓数量
	OpenPrice       float64        `json:"openPrice" form:"openPrice" gorm:"column:open_price;comment:开仓价格;size:10;"`                    //开仓价格
	LeverageRatio   int            `json:"leverageRatio" form:"leverageRatio" gorm:"column:leverage_ratio;comment:杠杆倍数;size:10;"`        //杠杆倍数
	Margin          Decimal        `json:"margin" form:"margin" gorm:"column:margin;comment:保证金;size:10;"`                               //保证金
	PositionAmount  float64        `json:"positionAmount" form:"positionAmount" gorm:"column:position_amount;comment:持仓金额;size:10;"`     //持仓金额
	ForceClosePrice float64        `json:"forceClosePrice" form:"forceClosePrice" gorm:"column:force_close_price;comment:强平价格;size:10;"` //强平价格
	PositionType    PositionType   `json:"positionType" form:"positionType" gorm:"column:position_type;comment:持仓类型;size:10;"`           //持仓类型
	PositionStatus  PositionStatus `json:"positionStatus" form:"positionStatus" gorm:"column:position_status;comment:持仓状态;size:10;"`     //持仓状态
	CreatedBy       uint           `gorm:"column:created_by;comment:创建者"`
	UpdatedBy       uint           `gorm:"column:updated_by;comment:更新者"`
	DeletedBy       uint           `gorm:"column:deleted_by;comment:删除者"`
}

// TableName contractPosition表 ContractPosition自定义表名 contract_position
func (ContractPosition) TableName() string {
	return "contract_position"
}
