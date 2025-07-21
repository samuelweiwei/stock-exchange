package contract

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/contract"
	contractReq "github.com/flipped-aurora/gin-vue-admin/server/model/contract/request"
	contractRes "github.com/flipped-aurora/gin-vue-admin/server/model/contract/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service/symbol"
	. "github.com/shopspring/decimal"
	"gorm.io/gorm"
	"strconv"
)

type ContractPositionService struct{}

// CreateContractPosition 创建contractPosition表记录
// Author [yourname](https://github.com/yourname)
func (contractPositionService *ContractPositionService) CreateContractPosition(contractPosition *contract.ContractPosition) (err error) {
	err = global.GVA_DB.Create(contractPosition).Error
	return err
}

// DeleteContractPosition 删除contractPosition表记录
// Author [yourname](https://github.com/yourname)
func (contractPositionService *ContractPositionService) DeleteContractPosition(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&contract.ContractPosition{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&contract.ContractPosition{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteContractPositionByIds 批量删除contractPosition表记录
// Author [yourname](https://github.com/yourname)
func (contractPositionService *ContractPositionService) DeleteContractPositionByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&contract.ContractPosition{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&contract.ContractPosition{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateContractPosition 更新contractPosition表记录
// Author [yourname](https://github.com/yourname)
func (contractPositionService *ContractPositionService) UpdateContractPosition(contractPosition contract.ContractPosition) (err error) {
	err = global.GVA_DB.Model(&contract.ContractPosition{}).Where("id = ?", contractPosition.ID).Updates(&contractPosition).Error
	return err
}

// GetContractPosition 根据ID获取contractPosition表记录
// Author [yourname](https://github.com/yourname)
func (contractPositionService *ContractPositionService) GetContractPosition(ID string) (contractPosition contract.ContractPosition, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&contractPosition).Error
	return
}

// GetContractPositionInfoList 分页获取contractPosition表记录
// Author [yourname](https://github.com/yourname)
func (contractPositionService *ContractPositionService) GetContractPositionInfoList(info contractReq.ContractPositionSearch, userId uint) (resList []contractRes.ContractPositionRes, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&contract.ContractPosition{}).Where("user_id = ? AND position_status = ?", userId, contract.Unclosed)
	var contractPositions []contract.ContractPosition
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

	err = db.Find(&contractPositions).Error
	if err != nil {
		return
	}

	for _, position := range contractPositions {
		closePrice, unrealizedProfitLoss, err := contractPositionService.GetUnrealizedProfitLossSingle(position.StockId, NewFromFloat(position.Quantity), NewFromFloat(position.OpenPrice), position.PositionType)
		if err != nil {
			return resList, total, err
		}
		//positionType := "买"
		//if position.PositionType == contract.Short {
		//	positionType = "卖"
		//}
		resList = append(resList, contractRes.ContractPositionRes{
			ID:                   strconv.Itoa(int(position.ID)),
			StockId:              strconv.Itoa(int(position.StockId)),
			StockName:            position.StockName,
			PositionTime:         position.PositionTime.UnixMilli(),
			Quantity:             NewFromFloat(position.Quantity),
			OpenPrice:            NewFromFloat(position.OpenPrice),
			CurPrice:             closePrice,
			ROI:                  unrealizedProfitLoss.Div(position.Margin).Mul(NewFromInt(100)).RoundCeil(2).String() + "%",
			SafetyFactor:         "0.00%",
			LeverageRatio:        position.LeverageRatio,
			Margin:               position.Margin,
			PositionAmount:       NewFromFloat(position.PositionAmount),
			ForceClosePrice:      NewFromFloat(position.ForceClosePrice),
			PositionType:         position.PositionType,
			UnrealizedProfitLoss: unrealizedProfitLoss,
		})
	}

	return resList, total, err
}
func (contractPositionService *ContractPositionService) GetContractPositionPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// GetUnrealizedProfitLoss 计算账号未实现盈亏
func (contractPositionService *ContractPositionService) GetUnrealizedProfitLoss(userId uint) (unrealizedProfitLoss Decimal, err error) {
	var contractPositions []contract.ContractPosition
	err = global.GVA_DB.Model(&contract.ContractPosition{}).Where("user_id = ? AND position_status = ?", userId, contract.Unclosed).Find(&contractPositions).Error
	if err != nil {
		return
	}
	for _, position := range contractPositions {
		_, single, err := contractPositionService.GetUnrealizedProfitLossSingle(position.StockId, NewFromFloat(position.Quantity), NewFromFloat(position.OpenPrice), position.PositionType)
		if err != nil {
			return Zero, err
		}
		unrealizedProfitLoss = unrealizedProfitLoss.Add(single)
	}
	return
}

// GetUnrealizedProfitLossSingle 计算单个持仓未实现盈亏
func (contractPositionService *ContractPositionService) GetUnrealizedProfitLossSingle(stockId uint, quantity Decimal, openPrice Decimal, positionType contract.PositionType) (closePrice Decimal, unrealizedProfitLoss Decimal, err error) {
	// 在redis里查找最新的股票价格
	symbolsService := symbol.SymbolsService{}
	stockPrice, err := symbolsService.GetSymbolPriceById(fmt.Sprintf("%d", stockId))
	if err != nil {
		return Zero, Zero, err
	}
	closePrice = NewFromFloat(stockPrice.(float64))
	switch positionType {
	case contract.Long:
		unrealizedProfitLoss = closePrice.Sub(openPrice).Mul(quantity).RoundFloor(2)
	case contract.Short:
		unrealizedProfitLoss = closePrice.Sub(openPrice).Mul(quantity).RoundCeil(2).Neg()
	}
	return
}

// GetForceClosePrice 计算强平价格
func (contractPositionService *ContractPositionService) GetForceClosePrice(openPrice Decimal, quantity Decimal, totalMargin Decimal, unrealizedProfitLoss Decimal, positionType contract.PositionType, priceSize int32) (forceClosePrice float64) {
	buffer := Zero // 预留一定的金额，TODO
	if quantity.Equal(Zero) {
		return
	}
	switch positionType {
	case contract.Long:
		forceClosePrice = openPrice.Sub(totalMargin.Add(unrealizedProfitLoss).Sub(buffer).Div(quantity).RoundFloor(priceSize)).InexactFloat64()
	case contract.Short:
		forceClosePrice = openPrice.Add(totalMargin.Add(unrealizedProfitLoss).Sub(buffer).Div(quantity).RoundFloor(priceSize)).InexactFloat64()
	}
	return forceClosePrice
}

// GetNewPositionInfo 追单后的开仓价格和杠杆倍数
func (contractPositionService *ContractPositionService) GetNewPositionInfo(oldOpenPrice Decimal, oldQuantity Decimal, oldMargin Decimal, newOpenPrice Decimal, newQuantity Decimal, leverageRatio int, positionType contract.PositionType, priceSize int32, quantitySize int32) (openPrice float64, quantity float64, margin Decimal) {
	positionAmount := oldOpenPrice.Mul(oldQuantity).Add(newOpenPrice.Mul(newQuantity))
	switch positionType {
	case contract.Long:
		openPrice = positionAmount.Div(oldQuantity.Add(newQuantity)).RoundCeil(priceSize).InexactFloat64()
	case contract.Short:
		openPrice = positionAmount.Div(oldQuantity.Add(newQuantity)).RoundFloor(priceSize).InexactFloat64()
	}
	quantity = positionAmount.Div(NewFromFloat(openPrice)).RoundFloor(quantitySize).InexactFloat64()
	margin = NewFromFloat(quantity).Mul(NewFromFloat(openPrice)).Div(NewFromInt(int64(leverageRatio))).RoundCeil(2)
	return openPrice, quantity, margin.Sub(oldMargin)
}
