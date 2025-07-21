package contract

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/contract"
	contractReq "github.com/flipped-aurora/gin-vue-admin/server/model/contract/request"
	contractRes "github.com/flipped-aurora/gin-vue-admin/server/model/contract/response"
	symbolModel "github.com/flipped-aurora/gin-vue-admin/server/model/symbol"
	"github.com/flipped-aurora/gin-vue-admin/server/service/symbol"
	. "github.com/shopspring/decimal"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type ContractOrderService struct{}

// CreateContractOrder 创建contractOrder表记录
// Author [yourname](https://github.com/yourname)
func (contractOrderService *ContractOrderService) CreateContractOrder(contractOrderReq *contractReq.ContractOrderReq, contractAccount *contract.ContractAccount, userId uint) (err error) {
	// 验证下数据有效性
	if contractOrderReq.OperationType != contract.OpenLong && contractOrderReq.OperationType != contract.OpenShort {
		err = errors.New("操作类型不正确")
		return
	}
	if contractOrderReq.OrderType != contract.MarketOrder && contractOrderReq.OrderType != contract.LimitOrder {
		err = errors.New("下单类型不正确")
		return
	}
	// 计算未实现盈亏
	contractPositionService := ContractPositionService{}
	unrealizedProfitLoss, err := contractPositionService.GetUnrealizedProfitLoss(userId)
	if err != nil {
		return
	}
	// 持仓类型
	positionType := contract.Long
	if contractOrderReq.OperationType == contract.OpenShort {
		positionType = contract.Short
	}
	// 看看有没有历史持仓
	var oldPosition contract.ContractPosition
	err = global.GVA_DB.Where("user_id = ? AND stock_id = ? AND position_type = ?", userId, contractOrderReq.StockId, positionType).First(&oldPosition).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	// 开仓只能用保证金和未实现盈亏的80%
	availableMargin := contractAccount.AvailableMargin.Add(unrealizedProfitLoss).Mul(NewFromFloat(0.8))
	if contractOrderReq.Margin.GreaterThan(availableMargin) {
		err = errors.New(fmt.Sprintf("最大可用保证金为%.2f", availableMargin.InexactFloat64()))
		return
	}
	// 查一下股票对应的精度
	var stock symbolModel.Symbols
	err = global.GVA_DB.Where("id = ?", contractOrderReq.StockId).First(&stock).Error
	if err != nil {
		return err
	}
	ticketSize := NewFromFloat(*stock.TicketSize)
	ticketNumSize := NewFromFloat(*stock.TicketNumSize)
	// 获得小数点位数
	priceSize := int32(decimalPlaces(ticketSize))
	quantitySize := int32(decimalPlaces(ticketNumSize))
	// 开仓价，市价单开多时加0.01，开空时减0.01
	openPrice := contractOrderReq.OpenPrice
	if contractOrderReq.OrderType == contract.MarketOrder {
		symbolsService := symbol.SymbolsService{}
		stockPrice, err := symbolsService.GetSymbolPriceById(fmt.Sprintf("%d", contractOrderReq.StockId))
		if err != nil {
			return err
		}
		if contractOrderReq.OperationType == contract.OpenLong {
			openPrice = NewFromFloat(stockPrice.(float64)).Add(ticketSize)
		} else if contractOrderReq.OperationType == contract.OpenShort {
			openPrice = NewFromFloat(stockPrice.(float64)).Sub(ticketSize)
		}
	}
	// 预估持仓金额
	positionAmount := contractOrderReq.Margin.Mul(NewFromInt(int64(contractOrderReq.LeverageRatio)))
	// 持仓数量，向下取整
	quantity := positionAmount.Div(openPrice).RoundFloor(quantitySize)
	if !quantity.IsPositive() {
		err = errors.New("持仓数量必须大于0")
		return
	}
	// 保证金根据数量变一下
	margin := quantity.Mul(openPrice).Div(NewFromInt(int64(contractOrderReq.LeverageRatio))).RoundCeil(2)
	// 占用的保证金要加上修改旧的持仓倍率导致的变化
	occMargin := margin
	// 预先生成订单Id
	orderId := uint(global.Snowflake.Generate())
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
	// 存入杠杆
	var contractLeverage contract.ContractLeverage
	err = tx.Where("user_id = ? AND stock_id = ?", userId, contractOrderReq.StockId).First(&contractLeverage).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			leverage := contract.ContractLeverage{
				GVA_MODEL: global.GVA_MODEL{
					ID: uint(global.Snowflake.Generate()),
				},
				UserId:        userId,
				StockId:       contractOrderReq.StockId,
				StockName:     contractOrderReq.StockName,
				LeverageRatio: contractOrderReq.LeverageRatio,
				CreatedBy:     userId,
			}
			err = tx.Create(&leverage).Error
			if err != nil {
				return
			}
		} else {
			return
		}
	} else if contractLeverage.LeverageRatio != contractOrderReq.LeverageRatio {
		err = errors.New("杠杆不一致")
		return
	}
	// 订单数据
	contractOrder := &contract.ContractOrder{
		GVA_MODEL: global.GVA_MODEL{
			ID: orderId,
		},
		OrderNumber:   "",
		OrderTime:     time.Now(),
		UserId:        userId,
		StockId:       contractOrderReq.StockId,
		StockName:     contractOrderReq.StockName,
		OrderType:     contractOrderReq.OrderType,
		OpenPrice:     openPrice.InexactFloat64(),
		OperationType: contractOrderReq.OperationType,
		Quantity:      quantity.InexactFloat64(),
		Fee:           Zero,
		CreatedBy:     userId,
	}
	// 市价单直接生成持仓，限价单生成委托
	switch contractOrderReq.OrderType {
	case contract.MarketOrder:
		{
			// 看看有没有历史持仓
			if oldPosition.ID == 0 {
				// 没有历史持仓
				contractPosition := contract.ContractPosition{
					GVA_MODEL: global.GVA_MODEL{
						ID: uint(global.Snowflake.Generate()),
					},
					UserId:          userId,
					StockId:         contractOrderReq.StockId,
					StockName:       contractOrderReq.StockName,
					PositionTime:    time.Now(),
					OpenPrice:       openPrice.InexactFloat64(),
					LeverageRatio:   contractOrderReq.LeverageRatio,
					Quantity:        quantity.InexactFloat64(),
					Margin:          occMargin,
					PositionAmount:  openPrice.Mul(quantity).InexactFloat64(),
					ForceClosePrice: contractPositionService.GetForceClosePrice(openPrice, quantity, contractAccount.TotalMargin, unrealizedProfitLoss, positionType, priceSize),
					PositionType:    positionType,
					PositionStatus:  contract.Unclosed,
					CreatedBy:       userId,
				}
				err = tx.Create(&contractPosition).Error
				if err != nil {
					tx.Rollback()
					return err
				}
			} else {
				// 看看之前的杠杆有没有更新成功
				if oldPosition.LeverageRatio != contractOrderReq.LeverageRatio {
					err = errors.New("杠杆更新失败")
					return
				}
				// 有历史持仓，要排除掉自身的未实现盈亏
				_, single, err := contractPositionService.GetUnrealizedProfitLossSingle(oldPosition.StockId, NewFromFloat(oldPosition.Quantity), NewFromFloat(oldPosition.OpenPrice), oldPosition.PositionType)
				if err != nil {
					tx.Rollback()
					return err
				}
				if oldPosition.PositionStatus == contract.Closed {
					oldPosition.PositionTime = time.Now()
				}
				oldPosition.OpenPrice, oldPosition.Quantity, occMargin = contractPositionService.GetNewPositionInfo(NewFromFloat(oldPosition.OpenPrice), NewFromFloat(oldPosition.Quantity), oldPosition.Margin, openPrice, quantity, contractOrderReq.LeverageRatio, positionType, priceSize, quantitySize)
				//oldPosition.LeverageRatio = contractOrderReq.LeverageRatio
				oldPosition.Margin = oldPosition.Margin.Add(occMargin)
				oldPosition.PositionAmount = NewFromFloat(oldPosition.OpenPrice).Mul(NewFromFloat(oldPosition.Quantity)).InexactFloat64()
				oldPosition.ForceClosePrice = contractPositionService.GetForceClosePrice(NewFromFloat(oldPosition.OpenPrice), NewFromFloat(oldPosition.Quantity), contractAccount.TotalMargin, unrealizedProfitLoss.Sub(single), positionType, priceSize)
				oldPosition.PositionStatus = contract.Unclosed
				oldPosition.UpdatedBy = userId
				err = tx.Model(&contract.ContractPosition{}).Where("id = ?", oldPosition.ID).Select("PositionTime", "OpenPrice", "Quantity", "Margin", "PositionAmount", "ForceClosePrice", "PositionStatus", "UpdatedBy").Updates(&oldPosition).Error
				if err != nil {
					tx.Rollback()
					return err
				}
			}
			contractOrder.OrderStatus = contract.FullyFilled
			contractAccount.UsedMargin = contractAccount.UsedMargin.Add(occMargin)
		}
	case contract.LimitOrder:
		{
			// 生成委托单
			contractEntrust := contract.ContractEntrust{
				GVA_MODEL: global.GVA_MODEL{
					ID: uint(global.Snowflake.Generate()),
				},
				UserId:        userId,
				OrderId:       orderId,
				StockId:       contractOrderReq.StockId,
				StockName:     contractOrderReq.StockName,
				TriggerType:   contract.Limit,
				TriggerPrice:  openPrice.InexactFloat64(),
				Margin:        occMargin,
				OperationType: contractOrderReq.OperationType,
				Quantity:      quantity.InexactFloat64(),
				EntrustStatus: contract.Untriggered,
				CreatedBy:     userId,
			}
			err = tx.Create(&contractEntrust).Error
			if err != nil {
				tx.Rollback()
				return err
			}
			contractOrder.OrderStatus = contract.Pending
			contractAccount.FrozenMargin = contractAccount.FrozenMargin.Add(occMargin)
		}
	}
	// 生成订单
	err = tx.Create(contractOrder).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// 更新账户金额
	contractAccount.AvailableMargin = contractAccount.AvailableMargin.Sub(occMargin)
	contractAccount.UpdatedBy = userId
	err = tx.Model(&contract.ContractAccount{}).Where("id = ?", contractAccount.ID).Select("AvailableMargin", "UsedMargin", "FrozenMargin", "UpdatedBy").Updates(contractAccount).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// 提交事务
	err = tx.Commit().Error
	return err
}

// DeleteContractOrder 删除contractOrder表记录
// Author [yourname](https://github.com/yourname)
func (contractOrderService *ContractOrderService) DeleteContractOrder(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&contract.ContractOrder{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&contract.ContractOrder{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteContractOrderByIds 批量删除contractOrder表记录
// Author [yourname](https://github.com/yourname)
func (contractOrderService *ContractOrderService) DeleteContractOrderByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&contract.ContractOrder{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&contract.ContractOrder{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateContractOrder 更新contractOrder表记录
// Author [yourname](https://github.com/yourname)
func (contractOrderService *ContractOrderService) UpdateContractOrder(contractOrder contract.ContractOrder) (err error) {
	err = global.GVA_DB.Model(&contract.ContractOrder{}).Where("id = ?", contractOrder.ID).Updates(&contractOrder).Error
	return err
}

// GetContractOrder 根据ID获取contractOrder表记录
// Author [yourname](https://github.com/yourname)
func (contractOrderService *ContractOrderService) GetContractOrder(ID string) (contractOrder contract.ContractOrder, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&contractOrder).Error
	return
}

// GetContractOrderInfoList 分页获取contractOrder表记录
// Author [yourname](https://github.com/yourname)
func (contractOrderService *ContractOrderService) GetContractOrderInfoList(info contractReq.ContractOrderSearch, userId uint) (resList []contractRes.ContractOrderRes, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&contract.ContractOrder{}).Where("user_id = ?", userId)
	var contractOrders []contract.ContractOrder
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

	err = db.Find(&contractOrders).Error
	if err != nil {
		return
	}

	for _, order := range contractOrders {
		resList = append(resList, contractRes.ContractOrderRes{
			ID:                 strconv.Itoa(int(order.ID)),
			OrderNumber:        order.OrderNumber,
			OrderTime:          order.OrderTime.UnixMilli(),
			StockId:            strconv.Itoa(int(order.StockId)),
			StockName:          order.StockName,
			OrderType:          order.OrderType,
			OpenPrice:          NewFromFloat(order.OpenPrice),
			ClosePrice:         NewFromFloat(order.ClosePrice),
			OperationType:      order.OperationType,
			Quantity:           NewFromFloat(order.Quantity),
			OrderStatus:        order.OrderStatus,
			Fee:                order.Fee,
			RealizedProfitLoss: order.RealizedProfitLoss,
		})
	}
	return resList, total, err
}
func (contractOrderService *ContractOrderService) GetContractOrderPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

/*// GetMaxOpenQuantity 获取最大开仓数量
func (contractOrderService *ContractOrderService) GetMaxOpenQuantity(openPrice float64, leverageRatio int, availableMargin float64, unrealizedProfitLoss float64) (quantity float64) {
	// 开户只能用保证金和未实现盈亏的80%
	availableMargin = (availableMargin + unrealizedProfitLoss) * 0.8
	// 持仓金额
	positionAmount := availableMargin * float64(leverageRatio)
	// 持仓数量，向下取整
	quantity = math.Floor(positionAmount / openPrice)
	return quantity
}*/

// CloseContractOrder 平仓方法
// Author [yourname](https://github.com/yourname)
func (contractOrderService *ContractOrderService) CloseContractOrder(contractCloseReq *contractReq.ContractCloseReq, contractAccount *contract.ContractAccount, userId uint) (err error) {
	// 历史持仓
	var oldPosition contract.ContractPosition
	err = global.GVA_DB.Where("user_id = ? AND id = ? AND position_status = ?", userId, contractCloseReq.PositionId, contract.Unclosed).First(&oldPosition).Error
	if err != nil {
		return err
	}
	// 操作类型
	operationType := contract.CloseLong
	if oldPosition.PositionType == contract.Short {
		operationType = contract.CloseShort
	}
	// 判断数量
	if contractCloseReq.Quantity.GreaterThan(NewFromFloat(oldPosition.Quantity)) {
		err = errors.New(fmt.Sprintf("最大可平仓数量%.2f", oldPosition.Quantity))
		return
	}
	if !contractCloseReq.Quantity.IsPositive() {
		err = errors.New("平仓数量必须大于0")
		return
	}
	// 查一下股票对应的精度
	var stock symbolModel.Symbols
	err = global.GVA_DB.Where("id = ?", oldPosition.StockId).First(&stock).Error
	if err != nil {
		return err
	}
	ticketSize := NewFromFloat(*stock.TicketSize)
	//ticketNumSize := NewFromFloat(*stock.TicketNumSize)
	// 获得小数点位数
	priceSize := int32(decimalPlaces(ticketSize))
	//quantitySize := int32(decimalPlaces(ticketNumSize))
	// 平仓价,计算已实现盈亏
	contractPositionService := ContractPositionService{}
	closePrice, realizedProfitLoss, err := contractPositionService.GetUnrealizedProfitLossSingle(oldPosition.StockId, contractCloseReq.Quantity, NewFromFloat(oldPosition.OpenPrice), oldPosition.PositionType)
	if err != nil {
		return err
	}
	// 计算未实现盈亏
	unrealizedProfitLoss, err := contractPositionService.GetUnrealizedProfitLoss(userId)
	if err != nil {
		return err
	}
	// 有历史持仓，要排除掉自身的未实现盈亏
	_, single, err := contractPositionService.GetUnrealizedProfitLossSingle(oldPosition.StockId, NewFromFloat(oldPosition.Quantity), NewFromFloat(oldPosition.OpenPrice), oldPosition.PositionType)
	if err != nil {
		return err
	}
	// 保证金
	newQuantity := NewFromFloat(oldPosition.Quantity).Sub(contractCloseReq.Quantity)
	newMargin := newQuantity.Mul(NewFromFloat(oldPosition.OpenPrice)).Div(NewFromInt(int64(oldPosition.LeverageRatio))).RoundCeil(2)
	diffMargin := newMargin.Sub(oldPosition.Margin)
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
	// 订单数据
	contractOrder := &contract.ContractOrder{
		GVA_MODEL: global.GVA_MODEL{
			ID: uint(global.Snowflake.Generate()),
		},
		OrderNumber:        "",
		OrderTime:          time.Now(),
		UserId:             userId,
		StockId:            oldPosition.StockId,
		StockName:          oldPosition.StockName,
		OrderType:          contract.CloseOrder,
		OpenPrice:          oldPosition.OpenPrice,
		ClosePrice:         closePrice.InexactFloat64(),
		OperationType:      operationType,
		Quantity:           contractCloseReq.Quantity.InexactFloat64(),
		OrderStatus:        contract.FullyFilled,
		Fee:                Zero,
		RealizedProfitLoss: realizedProfitLoss,
		CreatedBy:          userId,
	}
	// 历史持仓更新
	if newQuantity.IsZero() {
		oldPosition.PositionStatus = contract.Closed
		newMargin = Zero
		diffMargin = oldPosition.Margin.Neg()
	} else {
		oldPosition.PositionStatus = contract.Unclosed
	}
	oldPosition.Quantity = newQuantity.InexactFloat64()
	oldPosition.Margin = newMargin
	oldPosition.PositionAmount = NewFromFloat(oldPosition.OpenPrice).Mul(NewFromFloat(oldPosition.Quantity)).InexactFloat64()
	oldPosition.ForceClosePrice = contractPositionService.GetForceClosePrice(NewFromFloat(oldPosition.OpenPrice), NewFromFloat(oldPosition.Quantity), contractAccount.TotalMargin.Add(realizedProfitLoss), unrealizedProfitLoss.Sub(single), oldPosition.PositionType, priceSize)
	oldPosition.UpdatedBy = userId
	err = tx.Model(&contract.ContractPosition{}).Where("id = ?", oldPosition.ID).Select("Quantity", "Margin", "PositionAmount", "ForceClosePrice", "PositionStatus", "UpdatedBy").Updates(&oldPosition).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// 生成订单
	err = tx.Create(contractOrder).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// 更新账户金额
	contractAccount.TotalMargin = contractAccount.TotalMargin.Add(realizedProfitLoss)
	contractAccount.AvailableMargin = contractAccount.AvailableMargin.Sub(diffMargin).Add(realizedProfitLoss)
	contractAccount.UsedMargin = contractAccount.UsedMargin.Add(diffMargin)
	contractAccount.RealizedProfitLoss = contractAccount.RealizedProfitLoss.Add(realizedProfitLoss)
	contractAccount.UpdatedBy = userId
	err = tx.Model(&contract.ContractAccount{}).Where("id = ?", contractAccount.ID).Select("TotalMargin", "AvailableMargin", "UsedMargin", "RealizedProfitLoss", "UpdatedBy").Updates(contractAccount).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// 提交事务
	err = tx.Commit().Error
	return err
}

// CloseAllContractOrder 一键平仓方法
func (contractOrderService *ContractOrderService) CloseAllContractOrder(contractAccount *contract.ContractAccount, userId uint) (err error) {
	var positions []contract.ContractPosition
	err = global.GVA_DB.Model(&contract.ContractPosition{}).Where("user_id = ? AND position_status = ?", userId, contract.Unclosed).Find(&positions).Error
	if err != nil {
		return
	}
	for _, position := range positions {
		contractCloseReq := contractReq.ContractCloseReq{
			PositionId: strconv.Itoa(int(position.ID)),
			Quantity:   NewFromFloat(position.Quantity),
		}
		err = contractOrderService.CloseContractOrder(&contractCloseReq, contractAccount, userId)
		if err != nil {
			return
		}
	}
	return err
}
