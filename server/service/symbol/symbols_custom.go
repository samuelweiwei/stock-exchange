package symbol

import (
	"context"
	"fmt"
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/symbol"
	symbolReq "github.com/flipped-aurora/gin-vue-admin/server/model/symbol/request"
)

type SymbolsCustomService struct{}

// CreateSymbolsCustom 创建symbolsCustom表记录
// Author [yourname](https://github.com/yourname)
func (symbolsCustomService *SymbolsCustomService) CreateSymbolsCustom(symbolsCustom *symbol.SymbolsCustom) (err error) {
	err = global.GVA_DB.Create(symbolsCustom).Error
	return err
}

// DeleteSymbolsCustom 删除symbolsCustom表记录
// Author [yourname](https://github.com/yourname)
func (symbolsCustomService *SymbolsCustomService) DeleteSymbolsCustom(id string) (err error) {
	err = global.GVA_DB.Delete(&symbol.SymbolsCustom{}, "id = ?", id).Error
	return err
}

// DeleteSymbolsCustomByIds 批量删除symbolsCustom表记录
// Author [yourname](https://github.com/yourname)
func (symbolsCustomService *SymbolsCustomService) DeleteSymbolsCustomByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]symbol.SymbolsCustom{}, "id in ?", ids).Error
	return err
}

// UpdateSymbolsCustom 更新symbolsCustom表记录
// Author [yourname](https://github.com/yourname)
func (symbolsCustomService *SymbolsCustomService) UpdateSymbolsCustom(symbolsCustom symbol.SymbolsCustom) (err error) {
	err = global.GVA_DB.Model(&symbol.SymbolsCustom{}).Where("id = ?", symbolsCustom.Id).Updates(&symbolsCustom).Error
	return err
}

// GetSymbolsCustom 根据id获取symbolsCustom表记录
// Author [yourname](https://github.com/yourname)
func (symbolsCustomService *SymbolsCustomService) GetSymbolsCustom(id string) (symbolsCustom symbol.SymbolsCustom, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&symbolsCustom).Error
	return
}

// GetSymbolsCustomInfoList 分页获取SymbolsCustom记录
func (symbolsCustomService *SymbolsCustomService) GetSymbolsCustomInfoList(info symbolReq.SymbolsCustomSearch) (list interface{}, total int64, err error) {
	// 创建db
	db := global.GVA_DB.Model(&symbol.SymbolsCustom{})
	var symbolsCustoms []symbol.SymbolsCustom

	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 如果总数为0，直接返回空结果
	if total == 0 {
		return []interface{}{}, 0, nil
	}

	// 使用 Paginate 方法进行分页查询
	err = info.Paginate()(db).Find(&symbolsCustoms).Error
	if err != nil {
		return nil, 0, err
	}

	// 收集所有 symbolId
	var symbolIds []int
	for _, custom := range symbolsCustoms {
		if custom.SymbolId != nil {
			symbolIds = append(symbolIds, *custom.SymbolId)
		}
	}

	// 批量查询 symbols
	var symbolsList []symbol.Symbols
	err = global.GVA_DB.Where("id IN ?", symbolIds).Find(&symbolsList).Error
	if err != nil {
		return nil, 0, err
	}

	// 创建 symbol 映射，方便快速查找
	symbolsMap := make(map[int]symbol.Symbols)
	for _, s := range symbolsList {
		if s.Id != nil {
			symbolsMap[*s.Id] = s
		}
	}

	// 创建enriched列表
	enrichedList := make([]map[string]interface{}, len(symbolsCustoms))
	for i, customSymbol := range symbolsCustoms {
		enrichedSymbol := map[string]interface{}{
			"id":        customSymbol.Id,
			"createdAt": customSymbol.CreatedAt,
			"updatedAt": customSymbol.UpdatedAt,
			"symbolId":  customSymbol.SymbolId,
			"userId":    customSymbol.UserId,
		}

		// 从映射中获取关联的symbol信息
		if customSymbol.SymbolId != nil {
			if symbolInfo, exists := symbolsMap[*customSymbol.SymbolId]; exists {
				enrichedSymbol["symbol"] = symbolInfo.Symbol
				enrichedSymbol["corporation"] = symbolInfo.Corporation
				enrichedSymbol["industry"] = symbolInfo.Industry
				enrichedSymbol["exchange"] = symbolInfo.Exchange
				enrichedSymbol["marketCap"] = symbolInfo.MarketCap
				enrichedSymbol["listDate"] = symbolInfo.ListDate
				enrichedSymbol["description"] = symbolInfo.Description
				enrichedSymbol["sicDescription"] = symbolInfo.SicDescription
				enrichedSymbol["averageVolume"] = symbolInfo.AverageVolume
				enrichedSymbol["changeRatio"] = symbolInfo.ChangeRatio
				enrichedSymbol["peRatio"] = symbolInfo.PeRatio
				enrichedSymbol["icon"] = symbolInfo.Icon
				enrichedSymbol["type"] = symbolInfo.Type

				// 从Redis获取实时价格
				redisKey := fmt.Sprintf("symbol:stock:%s", symbolInfo.Symbol)
				if price, err := global.GVA_REDIS.Get(context.Background(), redisKey).Result(); err == nil {
					priceFloat, _ := strconv.ParseFloat(price, 64)
					enrichedSymbol["currentPrice"] = priceFloat
				}
			}
		}

		enrichedList[i] = enrichedSymbol
	}

	return enrichedList, total, nil
}

// GetSymbolsCustomPublic 分页获取用户自定义交易对
func (s *SymbolsCustomService) GetSymbolsCustomPublic(info request.PageInfo, userId uint) (list interface{}, total int64, err error) {
	// 创建db
	db := global.GVA_DB.Model(&symbol.SymbolsCustom{})
	var symbolsCustoms []symbol.SymbolsCustom

	// 添加用户ID筛选条件
	db = db.Where("user_id = ?", userId)

	// 获取总数
	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 如果总数为0，直接返回空结果
	if total == 0 {
		return []interface{}{}, 0, nil
	}

	// 分页查询
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	err = db.Limit(limit).Offset(offset).Find(&symbolsCustoms).Error
	if err != nil {
		return nil, 0, err
	}

	// 收集所有 symbolId
	var symbolIds []int
	for _, custom := range symbolsCustoms {
		if custom.SymbolId != nil {
			symbolIds = append(symbolIds, *custom.SymbolId)
		}
	}

	// 批量查询 symbols
	var symbolsList []symbol.Symbols
	err = global.GVA_DB.Where("id IN ?", symbolIds).Find(&symbolsList).Error
	if err != nil {
		return nil, 0, err
	}

	// 创建 symbol 映射，方便快速查找
	symbolsMap := make(map[int]symbol.Symbols)
	for _, s := range symbolsList {
		if s.Id != nil {
			symbolsMap[*s.Id] = s
		}
	}

	// 创建enriched列表
	enrichedList := make([]map[string]interface{}, len(symbolsCustoms))
	for i, customSymbol := range symbolsCustoms {
		enrichedSymbol := map[string]interface{}{
			"id":        customSymbol.Id,
			"createdAt": customSymbol.CreatedAt,
			"updatedAt": customSymbol.UpdatedAt,
			"symbolId":  customSymbol.SymbolId,
			"userId":    customSymbol.UserId,
		}

		// 从映射中获取关联的symbol信息
		if customSymbol.SymbolId != nil {
			if symbolInfo, exists := symbolsMap[*customSymbol.SymbolId]; exists {
				enrichedSymbol["symbol"] = symbolInfo.Symbol
				enrichedSymbol["corporation"] = symbolInfo.Corporation
				enrichedSymbol["industry"] = symbolInfo.Industry
				enrichedSymbol["exchange"] = symbolInfo.Exchange
				enrichedSymbol["marketCap"] = symbolInfo.MarketCap
				enrichedSymbol["listDate"] = symbolInfo.ListDate
				enrichedSymbol["description"] = symbolInfo.Description
				enrichedSymbol["sicDescription"] = symbolInfo.SicDescription
				enrichedSymbol["averageVolume"] = symbolInfo.AverageVolume
				enrichedSymbol["changeRatio"] = symbolInfo.ChangeRatio
				enrichedSymbol["peRatio"] = symbolInfo.PeRatio
				enrichedSymbol["icon"] = symbolInfo.Icon
				enrichedSymbol["type"] = symbolInfo.Type

				// 从Redis获取实时价格
				redisKey := fmt.Sprintf("symbol:stock:%s", symbolInfo.Symbol)
				if price, err := global.GVA_REDIS.Get(context.Background(), redisKey).Result(); err == nil {
					priceFloat, _ := strconv.ParseFloat(price, 64)
					enrichedSymbol["currentPrice"] = priceFloat
				}
			}
		}

		enrichedList[i] = enrichedSymbol
	}

	return enrichedList, total, nil
}

// CreateSymbolsCustomPublic 创建用户自定义交易对
func (s *SymbolsCustomService) CreateSymbolsCustomPublic(symbolsCustom *symbol.SymbolsCustom) (err error) {
	// 验证 symbolId 是否存在
	var count int64
	err = global.GVA_DB.Model(&symbol.Symbols{}).Where("id = ?", symbolsCustom.SymbolId).Count(&count).Error
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("symbol id %v not found", symbolsCustom.SymbolId)
	}

	// 检查是否已经添加过
	err = global.GVA_DB.Model(&symbol.SymbolsCustom{}).
		Where("user_id = ? AND symbol_id = ?", symbolsCustom.UserId, symbolsCustom.SymbolId).
		Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("symbol already added to custom list")
	}

	// 创建记录
	return global.GVA_DB.Create(symbolsCustom).Error
}

// DeleteSymbolsCustomPublic 删除用户自定义交易对
func (s *SymbolsCustomService) DeleteSymbolsCustomPublic(id string, userId uint) (err error) {
	// 验证记录是否存在且属于该用户
	var count int64
	err = global.GVA_DB.Model(&symbol.SymbolsCustom{}).
		Where("id = ? AND user_id = ?", id, userId).
		Count(&count).Error
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("record not found or not owned by user")
	}

	// 删除记录
	return global.GVA_DB.Where("id = ? AND user_id = ?", id, userId).Delete(&symbol.SymbolsCustom{}).Error
}

// InitUserDefaultSymbols 初始化用户默认的自定义交易对（市值前十）
func (s *SymbolsCustomService) InitUserDefaultSymbols(userId int) error {
	// 开启事务
	tx := global.GVA_DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 1. 从 symbols 表中获取市值前十的股票
	var topSymbols []symbol.Symbols
	err := tx.Model(&symbol.Symbols{}).
		Where("market_cap IS NOT NULL"). // 确保市值不为空
		Order("market_cap DESC").        // 按市值降序排序
		Limit(10).                       // 获取前10条
		Find(&topSymbols).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to get top symbols: %v", err)
	}

	// 2. 批量创建用户的自定义交易对记录
	var symbolsCustoms []symbol.SymbolsCustom
	for _, sym := range topSymbols {
		if sym.Id != nil { // 确保 ID 不为空
			symbolsCustom := symbol.SymbolsCustom{
				UserId:   &userId,
				SymbolId: sym.Id,
			}
			symbolsCustoms = append(symbolsCustoms, symbolsCustom)
		}
	}

	// 3. 批量插入数据
	if len(symbolsCustoms) > 0 {
		// 使用事务进行批量插入
		err = tx.Create(&symbolsCustoms).Error
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to create default custom symbols: %v", err)
		}
	}

	// 提交事务
	return tx.Commit().Error
}
