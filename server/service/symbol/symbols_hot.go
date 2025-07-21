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

type SymbolsHotService struct{}

// CreateSymbolsHot 创建symbolsHot表记录
// Author [yourname](https://github.com/yourname)
func (symbolsHotService *SymbolsHotService) CreateSymbolsHot(symbolsHot *symbol.SymbolsHot) (err error) {
	err = global.GVA_DB.Create(symbolsHot).Error
	return err
}

// DeleteSymbolsHot 删除symbolsHot表记录
// Author [yourname](https://github.com/yourname)
func (symbolsHotService *SymbolsHotService) DeleteSymbolsHot(id string) (err error) {
	err = global.GVA_DB.Delete(&symbol.SymbolsHot{}, "id = ?", id).Error
	return err
}

// DeleteSymbolsHotByIds 批量删除symbolsHot表记录
// Author [yourname](https://github.com/yourname)
func (symbolsHotService *SymbolsHotService) DeleteSymbolsHotByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]symbol.SymbolsHot{}, "id in ?", ids).Error
	return err
}

// UpdateSymbolsHot 更新symbolsHot表记录
// Author [yourname](https://github.com/yourname)
func (symbolsHotService *SymbolsHotService) UpdateSymbolsHot(symbolsHot symbol.SymbolsHot) (err error) {
	err = global.GVA_DB.Model(&symbol.SymbolsHot{}).Where("id = ?", symbolsHot.Id).Updates(&symbolsHot).Error
	return err
}

// GetSymbolsHot 根据id获取symbolsHot表记录
// Author [yourname](https://github.com/yourname)
func (symbolsHotService *SymbolsHotService) GetSymbolsHot(id string) (symbolsHot symbol.SymbolsHot, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&symbolsHot).Error
	return
}

// GetSymbolsHotInfoList 分页获取symbolsHot表记录
func (symbolsHotService *SymbolsHotService) GetSymbolsHotInfoList(info symbolReq.SymbolsHotSearch) (list interface{}, total int64, err error) {
	db := global.GVA_DB.Model(&symbol.SymbolsHot{})
	var symbolsHots []symbol.SymbolsHot

	if info.Symbol != "" || info.Type != "" {
		db = db.Joins("INNER JOIN symbols ON symbols_hot.symbol_id = symbols.id")

		if info.Symbol != "" {
			db = db.Where("symbols.symbol LIKE ?", "%"+info.Symbol+"%")
		}
		if info.Type != "" {
			db = db.Where("symbols.type = ?", info.Type)
		}
	}

	db = db.Order("sort DESC")

	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 如果总数为0，直接返回空结果
	if total == 0 {
		return []interface{}{}, 0, nil
	}

	// 使用 Paginate 方法进行分页查询
	err = info.Paginate()(db).Find(&symbolsHots).Error
	if err != nil {
		return nil, 0, err
	}

	// 收集所有 symbolId
	var symbolIds []int
	for _, hot := range symbolsHots {
		if hot.SymbolId != nil {
			symbolIds = append(symbolIds, *hot.SymbolId)
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
	enrichedList := make([]map[string]interface{}, len(symbolsHots))
	for i, hotSymbol := range symbolsHots {
		enrichedSymbol := map[string]interface{}{
			"id":        hotSymbol.Id,
			"createdAt": hotSymbol.CreatedAt,
			"updatedAt": hotSymbol.UpdatedAt,
			"symbolId":  hotSymbol.SymbolId,
			"sort":      hotSymbol.Sort,
		}

		// 从映射中获取关联的symbol信息
		if hotSymbol.SymbolId != nil {
			if symbolInfo, exists := symbolsMap[*hotSymbol.SymbolId]; exists {
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
				enrichedSymbol["icon"] = processIconURL(symbolInfo.Icon)
				enrichedSymbol["type"] = symbolInfo.Type
				enrichedSymbol["status"] = symbolInfo.Status
				enrichedSymbol["ticketSize"] = symbolInfo.TicketSize
				enrichedSymbol["ticketNumSize"] = symbolInfo.TicketNumSize

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

// GetSymbolsHotPublic 获取热门股票列表（添加分页）
func (symbolsHotService *SymbolsHotService) GetSymbolsHotPublic(info request.PageInfo) (list interface{}, total int64, err error) {
	// 创建db
	db := global.GVA_DB.Model(&symbol.SymbolsHot{})
	var symbolsHots []symbol.SymbolsHot

	// 添加排序
	db = db.Order("sort DESC") // 按照sort字段降序排序

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
	err = db.Limit(limit).Offset(offset).Find(&symbolsHots).Error
	if err != nil {
		return nil, 0, err
	}

	// 收集所有 symbolId
	var symbolIds []int
	for _, hot := range symbolsHots {
		if hot.SymbolId != nil {
			symbolIds = append(symbolIds, *hot.SymbolId)
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
	enrichedList := make([]map[string]interface{}, len(symbolsHots))
	for i, hotSymbol := range symbolsHots {
		enrichedSymbol := map[string]interface{}{
			"id":        hotSymbol.Id,
			"createdAt": hotSymbol.CreatedAt,
			"updatedAt": hotSymbol.UpdatedAt,
			"symbolId":  hotSymbol.SymbolId,
			"sort":      hotSymbol.Sort,
		}

		// 从映射中获取关联的symbol信息
		if hotSymbol.SymbolId != nil {
			if symbolInfo, exists := symbolsMap[*hotSymbol.SymbolId]; exists {
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
				enrichedSymbol["icon"] = processIconURL(symbolInfo.Icon)
				enrichedSymbol["type"] = symbolInfo.Type
				enrichedSymbol["status"] = symbolInfo.Status
				enrichedSymbol["ticketSize"] = symbolInfo.TicketSize
				enrichedSymbol["ticketNumSize"] = symbolInfo.TicketNumSize

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
