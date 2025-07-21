// 自动生成模板ContractLeverage
package contract

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// contractLeverage表 结构体  ContractLeverage
type ContractLeverage struct {
	global.GVA_MODEL
	UserId        uint   `json:"userId" form:"userId" gorm:"column:user_id;comment:关联用户表的用户 ID;size:20;"`               //关联用户表的用户 ID
	StockId       uint   `json:"stockId" form:"stockId" gorm:"column:stock_id;comment:股票id;size:20;"`                   //股票id
	StockName     string `json:"stockName" form:"stockName" gorm:"column:stock_name;comment:股票名称;size:255;"`            //股票名称
	LeverageRatio int    `json:"leverageRatio" form:"leverageRatio" gorm:"column:leverage_ratio;comment:杠杆倍数;size:10;"` //杠杆倍数
	CreatedBy     uint   `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建人ID;size:20;"`            //创建人ID
	UpdatedBy     uint   `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新人ID;size:20;"`            //更新人ID
	DeletedBy     uint   `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:删除人ID;size:20;"`            //删除人ID
}

// TableName contractLeverage表 ContractLeverage自定义表名 contract_leverage
func (ContractLeverage) TableName() string {
	return "contract_leverage"
}
