package contract

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/contract"
	contractReq "github.com/flipped-aurora/gin-vue-admin/server/model/contract/request"
	. "github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ContractLeverageService struct{}

// CreateContractLeverage 创建contractLeverage表记录
// Author [yourname](https://github.com/yourname)
func (contractLeverageService *ContractLeverageService) CreateContractLeverage(contractLeverage *contract.ContractLeverage) (err error) {
	err = global.GVA_DB.Create(contractLeverage).Error
	return err
}

// DeleteContractLeverage 删除contractLeverage表记录
// Author [yourname](https://github.com/yourname)
func (contractLeverageService *ContractLeverageService) DeleteContractLeverage(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&contract.ContractLeverage{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&contract.ContractLeverage{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteContractLeverageByIds 批量删除contractLeverage表记录
// Author [yourname](https://github.com/yourname)
func (contractLeverageService *ContractLeverageService) DeleteContractLeverageByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&contract.ContractLeverage{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&contract.ContractLeverage{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateContractLeverage 更新contractLeverage表记录
// Author [yourname](https://github.com/yourname)
func (contractLeverageService *ContractLeverageService) UpdateContractLeverage(contractLeverage *contract.ContractLeverage, userId uint) (err error) {
	var oldLeverage contract.ContractLeverage
	err = global.GVA_DB.Where("user_id = ? AND stock_id = ?", userId, contractLeverage.StockId).First(&oldLeverage).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			leverage := contract.ContractLeverage{
				GVA_MODEL: global.GVA_MODEL{
					ID: uint(global.Snowflake.Generate()),
				},
				UserId:        userId,
				StockId:       contractLeverage.StockId,
				StockName:     contractLeverage.StockName,
				LeverageRatio: contractLeverage.LeverageRatio,
				CreatedBy:     userId,
			}
			err = global.GVA_DB.Create(&leverage).Error
			return err
		} else {
			return err
		}
	}
	// 开始事务
	tx := global.GVA_DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		// Make sure to rollback when panic, Block error or Commit error
		if err != nil {
			tx.Rollback()
		}
	}()
	// 更新杠杆表
	oldLeverage.LeverageRatio = contractLeverage.LeverageRatio
	oldLeverage.UpdatedBy = userId
	err = tx.Model(&contract.ContractLeverage{}).Where("id = ?", oldLeverage.ID).Select("LeverageRatio", "UpdatedBy").Updates(&oldLeverage).Error
	if err != nil {
		return err
	}
	// 更新持仓信息
	var positions []contract.ContractPosition
	err = tx.Model(&contract.ContractPosition{}).Where("user_id = ? AND stock_id = ?", userId, contractLeverage.StockId).Find(&positions).Error
	if err != nil {
		return
	}
	var diffUsedMargin Decimal
	for _, position := range positions {
		newMargin := NewFromFloat(position.Quantity).Mul(NewFromFloat(position.OpenPrice)).Div(NewFromInt(int64(contractLeverage.LeverageRatio))).RoundCeil(2)
		diffUsedMargin = newMargin.Sub(position.Margin)
		position.LeverageRatio = contractLeverage.LeverageRatio
		position.Margin = newMargin
		position.UpdatedBy = userId
		err = tx.Model(&contract.ContractPosition{}).Where("id = ?", position.ID).Select("LeverageRatio", "Margin", "UpdatedBy").Updates(&position).Error
		if err != nil {
			return err
		}
	}
	// 更新委托信息
	var entrusts []contract.ContractEntrust
	err = tx.Model(&contract.ContractEntrust{}).Where("user_id = ? AND stock_id = ? AND trigger_type = ? AND entrust_status = ?", userId, contractLeverage.StockId, contract.Limit, contract.Untriggered).Find(&entrusts).Error
	if err != nil {
		return
	}
	var diffFrozenMargin Decimal
	for _, entrust := range entrusts {
		newMargin := NewFromFloat(entrust.Quantity).Mul(NewFromFloat(entrust.TriggerPrice)).Div(NewFromInt(int64(contractLeverage.LeverageRatio))).RoundCeil(2)
		diffFrozenMargin = newMargin.Sub(entrust.Margin)
		entrust.Margin = newMargin
		entrust.UpdatedBy = userId
		err = tx.Model(&contract.ContractEntrust{}).Where("id = ?", entrust.ID).Select("Margin", "UpdatedBy").Updates(&entrust).Error
		if err != nil {
			return err
		}
	}
	// 更新账户表
	contractAccountService := ContractAccountService{}
	account, err := contractAccountService.GetContractAccountByUserId(userId)
	if err != nil {
		return
	}
	account.AvailableMargin = account.AvailableMargin.Sub(diffUsedMargin).Sub(diffFrozenMargin)
	account.UsedMargin = account.UsedMargin.Add(diffUsedMargin)
	account.FrozenMargin = account.FrozenMargin.Sub(diffFrozenMargin)
	account.UpdatedBy = userId
	err = tx.Model(&contract.ContractAccount{}).Where("id = ?", account.ID).Select("AvailableMargin", "UsedMargin", "FrozenMargin", "UpdatedBy").Updates(&account).Error
	if err != nil {
		return
	}
	// 提交事务
	err = tx.Commit().Error
	return err
}

// GetContractLeverage 根据ID获取contractLeverage表记录
// Author [yourname](https://github.com/yourname)
func (contractLeverageService *ContractLeverageService) GetContractLeverage(stockId string, userId uint) (contractLeverage contract.ContractLeverage, err error) {
	err = global.GVA_DB.Where("user_id = ? AND stock_id = ?", userId, stockId).First(&contractLeverage).Error
	return
}

// GetContractLeverageInfoList 分页获取contractLeverage表记录
// Author [yourname](https://github.com/yourname)
func (contractLeverageService *ContractLeverageService) GetContractLeverageInfoList(info contractReq.ContractLeverageSearch) (list []contract.ContractLeverage, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&contract.ContractLeverage{})
	var contractLeverages []contract.ContractLeverage
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&contractLeverages).Error
	return contractLeverages, total, err
}
func (contractLeverageService *ContractLeverageService) GetContractLeveragePublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
