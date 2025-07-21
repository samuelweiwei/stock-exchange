package contract

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/contract"
	contractReq "github.com/flipped-aurora/gin-vue-admin/server/model/contract/request"
	contractRes "github.com/flipped-aurora/gin-vue-admin/server/model/contract/response"
	. "github.com/shopspring/decimal"
	"gorm.io/gorm"
	"strconv"
)

type ContractEntrustService struct{}

// CreateContractEntrust 创建contractEntrust表记录
// Author [yourname](https://github.com/yourname)
func (contractEntrustService *ContractEntrustService) CreateContractEntrust(contractEntrustReq *contractReq.ContractEntrustReq, userId uint) (err error) {
	// 验证下数据有效性
	if contractEntrustReq.TriggerType != contract.TakeProfit && contractEntrustReq.TriggerType != contract.StopLoss {
		err = errors.New("触发类型不正确")
		return
	}
	// 先看看有没有持仓
	var contractPosition contract.ContractPosition
	err = global.GVA_DB.Where("id = ?", contractEntrustReq.PositionId).First(&contractPosition).Error
	if err != nil {
		return err
	}
	operationType := contract.CloseLong
	if contractPosition.PositionType == contract.Short {
		operationType = contract.CloseShort
	}
	// 委托数据
	contractEntrust := contract.ContractEntrust{
		GVA_MODEL: global.GVA_MODEL{
			ID: uint(global.Snowflake.Generate()),
		},
		UserId:        userId,
		PositionId:    contractPosition.ID,
		StockId:       contractPosition.StockId,
		StockName:     contractPosition.StockName,
		TriggerType:   contractEntrustReq.TriggerType,
		TriggerPrice:  contractEntrustReq.TriggerPrice.InexactFloat64(),
		OperationType: operationType,
		Quantity:      contractEntrustReq.Quantity.InexactFloat64(),
		EntrustStatus: contract.Untriggered,
		CreatedBy:     userId,
	}

	err = global.GVA_DB.Create(&contractEntrust).Error
	return err
}

// DeleteContractEntrust 删除contractEntrust表记录
// Author [yourname](https://github.com/yourname)
func (contractEntrustService *ContractEntrustService) DeleteContractEntrust(ID string, userId uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var contractEntrust contract.ContractEntrust
		err := tx.Where("id = ? and user_id = ? and entrust_status = ?", ID, userId, contract.Untriggered).First(&contractEntrust).Error
		if err != nil {
			return err
		}
		if contractEntrust.TriggerType == contract.Limit {
			// 限价单要删掉order
			var contractOrder contract.ContractOrder
			err = tx.Where("id = ?", contractEntrust.OrderId).First(&contractOrder).Error
			if err != nil {
				return err
			}
			contractOrder.OrderStatus = contract.Cancelled
			contractOrder.UpdatedBy = userId
			err = tx.Model(&contract.ContractOrder{}).Where("id = ?", contractOrder.ID).Select("OrderStatus", "UpdatedBy").Updates(&contractOrder).Error
			if err != nil {
				return err
			}
			// 还要更新一下账户金额
			var contractAccount contract.ContractAccount
			err = tx.Where("user_id = ?", userId).First(&contractAccount).Error
			if err != nil {
				return err
			}
			contractAccount.AvailableMargin = contractAccount.AvailableMargin.Add(contractEntrust.Margin)
			contractAccount.FrozenMargin = contractAccount.FrozenMargin.Sub(contractEntrust.Margin)
			contractAccount.UpdatedBy = userId
			err = tx.Model(&contract.ContractAccount{}).Where("id = ?", contractAccount.ID).Select("AvailableMargin", "FrozenMargin", "UpdatedBy").Updates(contractAccount).Error
			if err != nil {
				return err
			}
		}
		// 更新委托表
		contractEntrust.EntrustStatus = contract.Deleted
		contractEntrust.DeletedBy = userId
		if err = tx.Model(&contract.ContractEntrust{}).Where("id = ?", ID).Select("EntrustStatus", "DeletedBy").Updates(&contractEntrust).Error; err != nil {
			return err
		}
		if err = tx.Delete(&contract.ContractEntrust{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteAllContractEntrust 一键撤销委托
func (contractEntrustService *ContractEntrustService) DeleteAllContractEntrust(userId uint) (err error) {
	// 更新委托信息
	var entrusts []contract.ContractEntrust
	err = global.GVA_DB.Model(&contract.ContractEntrust{}).Where("user_id = ? AND entrust_status = ?", userId, contract.Untriggered).Find(&entrusts).Error
	if err != nil {
		return
	}
	for _, entrust := range entrusts {
		err = contractEntrustService.DeleteContractEntrust(strconv.Itoa(int(entrust.ID)), userId)
		if err != nil {
			return
		}
	}
	return err
}

// DeleteContractEntrustByIds 批量删除contractEntrust表记录
// Author [yourname](https://github.com/yourname)
func (contractEntrustService *ContractEntrustService) DeleteContractEntrustByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&contract.ContractEntrust{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&contract.ContractEntrust{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateContractEntrust 更新contractEntrust表记录
// Author [yourname](https://github.com/yourname)
func (contractEntrustService *ContractEntrustService) UpdateContractEntrust(contractEntrust contract.ContractEntrust) (err error) {
	err = global.GVA_DB.Model(&contract.ContractEntrust{}).Where("id = ?", contractEntrust.ID).Updates(&contractEntrust).Error
	return err
}

// GetContractEntrust 根据ID获取contractEntrust表记录
// Author [yourname](https://github.com/yourname)
func (contractEntrustService *ContractEntrustService) GetContractEntrust(ID string) (contractEntrust contract.ContractEntrust, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&contractEntrust).Error
	return
}

// GetContractEntrustInfoList 分页获取contractEntrust表记录
// Author [yourname](https://github.com/yourname)
func (contractEntrustService *ContractEntrustService) GetContractEntrustInfoList(info contractReq.ContractEntrustSearch, userId uint) (resList []contractRes.ContractEntrustRes, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&contract.ContractEntrust{}).Where("user_id = ? AND entrust_status = ?", userId, contract.Untriggered)
	var contractEntrusts []contract.ContractEntrust
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Keyword != "" {
		db = db.Where("stock_name LIKE ?", "%"+info.Keyword+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&contractEntrusts).Error
	if err != nil {
		return
	}

	for _, entrust := range contractEntrusts {
		resList = append(resList, contractRes.ContractEntrustRes{
			ID:            strconv.Itoa(int(entrust.ID)),
			PositionId:    strconv.Itoa(int(entrust.PositionId)),
			OrderId:       strconv.Itoa(int(entrust.OrderId)),
			StockId:       strconv.Itoa(int(entrust.StockId)),
			StockName:     entrust.StockName,
			TriggerType:   entrust.TriggerType,
			TriggerPrice:  NewFromFloat(entrust.TriggerPrice),
			Margin:        entrust.Margin,
			OperationType: entrust.OperationType,
			Quantity:      NewFromFloat(entrust.Quantity),
			EntrustStatus: entrust.EntrustStatus,
		})
	}
	return resList, total, err
}
func (contractEntrustService *ContractEntrustService) GetContractEntrustPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
