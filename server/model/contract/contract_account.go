// 自动生成模板ContractAccount
package contract

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	. "github.com/shopspring/decimal"
)

type AccountStatus int

// 未审核-1，已审核-2
const (
	Unreviewed AccountStatus = 1
	Reviewed   AccountStatus = 2
)

// contractAccount表 结构体  ContractAccount
type ContractAccount struct {
	global.GVA_MODEL
	UserId             uint          `json:"userId" form:"userId" gorm:"column:user_id;comment:关联用户表的用户 ID;size:20;"`                                //关联用户表的用户 ID
	TotalMargin        Decimal       `json:"totalMargin" form:"totalMargin" gorm:"column:total_margin;comment:总保证金;size:10;"`                        //总保证金
	AvailableMargin    Decimal       `json:"availableMargin" form:"availableMargin" gorm:"column:available_margin;comment:可用保证金;size:10;"`           //可用保证金
	FrozenMargin       Decimal       `json:"frozenMargin" form:"frozenMargin" gorm:"column:frozen_margin;comment:冻结保证金;size:10;"`                    //冻结保证金
	UsedMargin         Decimal       `json:"usedMargin" form:"usedMargin" gorm:"column:used_margin;comment:已用保证金;size:10;"`                          //已用保证金
	RealizedProfitLoss Decimal       `json:"realizedProfitLoss" form:"realizedProfitLoss" gorm:"column:realized_profit_loss;comment:已实现盈亏;size:10;"` //已实现盈亏
	AccountStatus      AccountStatus `json:"accountStatus" form:"accountStatus" gorm:"column:account_status;comment:账号状态;size:10;"`                  //账号状态
	CreatedBy          uint          `gorm:"column:created_by;comment:创建者"`
	UpdatedBy          uint          `gorm:"column:updated_by;comment:更新者"`
	DeletedBy          uint          `gorm:"column:deleted_by;comment:删除者"`
}

// TableName contractAccount表 ContractAccount自定义表名 contract_account
func (ContractAccount) TableName() string {
	return "contract_account"
}

func (status AccountStatus) Text() string {
	switch status {
	case Unreviewed:
		return "未审核"
	case Reviewed:
		return "已审核"
	default:
		return "未知"
	}
}
