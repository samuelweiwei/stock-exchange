package contract

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/enums/fund"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/contract"
	contractReq "github.com/flipped-aurora/gin-vue-admin/server/model/contract/request"
	symbolModel "github.com/flipped-aurora/gin-vue-admin/server/model/symbol"
	"github.com/flipped-aurora/gin-vue-admin/server/service/userfund"
	. "github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ContractAccountService struct{}

// CreateContractAccount 创建contractAccount表记录
// Author [yourname](https://github.com/yourname)
func (contractAccountService *ContractAccountService) CreateContractAccount(contractAccount *contract.ContractAccount) (err error) {
	err = global.GVA_DB.Create(contractAccount).Error
	return err
}

// DeleteContractAccount 删除contractAccount表记录
// Author [yourname](https://github.com/yourname)
func (contractAccountService *ContractAccountService) DeleteContractAccount(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&contract.ContractAccount{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&contract.ContractAccount{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteContractAccountByIds 批量删除contractAccount表记录
// Author [yourname](https://github.com/yourname)
func (contractAccountService *ContractAccountService) DeleteContractAccountByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&contract.ContractAccount{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&contract.ContractAccount{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateContractAccount 更新contractAccount表记录
// Author [yourname](https://github.com/yourname)
func (contractAccountService *ContractAccountService) UpdateContractAccount(contractAccount contract.ContractAccount) (err error) {
	err = global.GVA_DB.Model(&contract.ContractAccount{}).Where("id = ?", contractAccount.ID).Updates(&contractAccount).Error
	return err
}

// GetContractAccount 根据ID获取contractAccount表记录
// Author [yourname](https://github.com/yourname)
func (contractAccountService *ContractAccountService) GetContractAccount(ID string) (contractAccount contract.ContractAccount, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&contractAccount).Error
	return
}

// GetContractAccount 根据ID获取contractAccount表记录
// Author [yourname](https://github.com/yourname)
func (contractAccountService *ContractAccountService) GetContractAccountByUserId(userId uint) (contractAccount contract.ContractAccount, err error) {
	err = global.GVA_DB.Where("user_id = ?", userId).First(&contractAccount).Error
	return
}

// GetContractAccountInfoList 分页获取contractAccount表记录
// Author [yourname](https://github.com/yourname)
func (contractAccountService *ContractAccountService) GetContractAccountInfoList(info contractReq.ContractAccountSearch) (list []contract.ContractAccount, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&contract.ContractAccount{})
	var contractAccounts []contract.ContractAccount
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

	err = db.Find(&contractAccounts).Error
	return contractAccounts, total, err
}
func (contractAccountService *ContractAccountService) GetContractAccountPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// GetTransferableAmount 获取可划转金额
// Author [yourname](https://github.com/yourname)
func (contractAccountService *ContractAccountService) GetTransferableAmount(totalMargin Decimal, availableMargin Decimal, unrealizedProfitLoss Decimal) (amount Decimal) {
	// 当前有持仓时,可划转金额,只能划转80%的金额
	if totalMargin.Equal(availableMargin) {
		amount = availableMargin
	} else {
		amount = availableMargin.Add(unrealizedProfitLoss).Mul(NewFromFloat(0.8)).RoundFloor(2)
	}
	return amount
}

// ChangeAccountMargin 合约账户转入或转出
func (contractAccountService *ContractAccountService) ChangeAccountMargin(changeMargin contractReq.ChangeMarginReq, userId uint) (err error) {
	// 验证下数据有效性
	if changeMargin.Type != contractReq.TransferIn && changeMargin.Type != contractReq.TransferOut {
		err = errors.New("划转类型不正确")
		return
	}
	// 获取账户信息
	account, err := contractAccountService.GetContractAccountByUserId(userId)
	if err != nil {
		return
	}
	// 计算未实现盈亏
	contractPositionService := ContractPositionService{}
	unrealizedProfitLoss, err := contractPositionService.GetUnrealizedProfitLoss(userId)
	if err != nil {
		return
	}
	// 转出时判断下
	amount := changeMargin.Amount
	actionType := fund.TransferToContract
	if changeMargin.Type == contractReq.TransferOut {
		// 最大划转金额
		maxAmount := contractAccountService.GetTransferableAmount(account.TotalMargin, account.AvailableMargin, unrealizedProfitLoss)
		if amount.GreaterThan(maxAmount) {
			err = errors.New(fmt.Sprintf("最大可划转保证金为%.2f", maxAmount.InexactFloat64()))
			return
		}
		amount = amount.Neg()
		actionType = fund.TransferFromContract
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
	// 更新账户字段
	account.TotalMargin = account.TotalMargin.Add(amount)
	account.AvailableMargin = account.AvailableMargin.Add(amount)
	account.UpdatedBy = userId
	err = tx.Model(&contract.ContractAccount{}).Where("id = ?", account.ID).Select("TotalMargin", "AvailableMargin", "UpdatedBy").Updates(&account).Error
	if err != nil {
		return
	}
	// 更新持仓信息
	var positions []contract.ContractPosition
	err = tx.Model(&contract.ContractPosition{}).Where("user_id = ? AND position_status = ?", userId, contract.Unclosed).Find(&positions).Error
	if err != nil {
		return
	}
	for _, position := range positions {
		// 查一下股票对应的精度
		var stock symbolModel.Symbols
		err = tx.Where("id = ?", position.StockId).First(&stock).Error
		if err != nil {
			return err
		}
		ticketSize := NewFromFloat(*stock.TicketSize)
		//ticketNumSize := NewFromFloat(*stock.TicketNumSize)
		// 获得小数点位数
		priceSize := int32(decimalPlaces(ticketSize))
		//quantitySize := int32(decimalPlaces(ticketNumSize))
		_, single, err := contractPositionService.GetUnrealizedProfitLossSingle(position.StockId, NewFromFloat(position.Quantity), NewFromFloat(position.OpenPrice), position.PositionType)
		if err != nil {
			return err
		}
		position.ForceClosePrice = contractPositionService.GetForceClosePrice(NewFromFloat(position.OpenPrice), NewFromFloat(position.Quantity), account.TotalMargin, unrealizedProfitLoss.Sub(single), position.PositionType, priceSize)
		position.UpdatedBy = userId
		err = tx.Model(&contract.ContractPosition{}).Where("id = ?", position.ID).Select("ForceClosePrice", "UpdatedBy").Updates(&position).Error
		if err != nil {
			return err
		}
	}
	// 加入资金流水
	err = userfund.NewUserFundAccountService(tx, true).UpdateUserFundAccountsAndNewFlow(int(userId), actionType, changeMargin.Amount.InexactFloat64(), "")
	if err != nil {
		return
	}
	return
}
