package userfund

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/userfund"
	userfundReq "github.com/flipped-aurora/gin-vue-admin/server/model/userfund/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type WithdrawChannelsService struct{}

// CreateWithdrawChannels 创建withdrawChannels表记录
// Author [yourname](https://github.com/yourname)
func (withdrawChannelsService *WithdrawChannelsService) CreateWithdrawChannels(withdrawChannels *userfund.WithdrawChannels) (err error) {
	err = global.GVA_DB.Create(withdrawChannels).Error
	return err
}

// DeleteWithdrawChannels 删除withdrawChannels表记录
// Author [yourname](https://github.com/yourname)
func (withdrawChannelsService *WithdrawChannelsService) DeleteWithdrawChannels(ID string) (err error) {
	err = global.GVA_DB.Delete(&userfund.WithdrawChannels{}, "id = ?", ID).Error
	return err
}

// DeleteWithdrawChannelsByIds 批量删除withdrawChannels表记录
// Author [yourname](https://github.com/yourname)
func (withdrawChannelsService *WithdrawChannelsService) DeleteWithdrawChannelsByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]userfund.WithdrawChannels{}, "id in ?", IDs).Error
	return err
}

// UpdateWithdrawChannels 更新withdrawChannels表记录
// Author [yourname](https://github.com/yourname)
func (withdrawChannelsService *WithdrawChannelsService) UpdateWithdrawChannels(withdrawChannels userfund.WithdrawChannels) (err error) {
	err = global.GVA_DB.Model(&userfund.WithdrawChannels{}).Where("id = ?", withdrawChannels.ID).Updates(&withdrawChannels).Error
	return err
}

// GetWithdrawChannels 根据ID获取withdrawChannels表记录
// Author [yourname](https://github.com/yourname)
func (withdrawChannelsService *WithdrawChannelsService) GetWithdrawChannels(ID string) (withdrawChannels userfund.WithdrawChannels, err error) {
	// 先查询提现渠道信息
	err = global.GVA_DB.Where("id = ?", ID).First(&withdrawChannels).Error
	if err != nil {
		return
	}

	// 如果有关联的货币ID，查询货币信息
	if withdrawChannels.CoinId != nil {
		var currency userfund.Currencies
		err = global.GVA_DB.Where("id = ?", *withdrawChannels.CoinId).First(&currency).Error
		if err != nil {
			return
		}
		// 组装货币信息
		withdrawChannels.Currency = currency.Currency
		withdrawChannels.CurrencyIcon = currency.Icon
		withdrawChannels.CoinType = currency.CoinType
		withdrawChannels.TicketSize = currency.TicketSize
		withdrawChannels.TicketNumSize = currency.TicketNumSize
		withdrawChannels.MinWithdrawNum = currency.MinWithdrawNum

		// 获取实时价格并按精度处理
		currencyService := CurrenciesService{}
		realTimePrice, err := currencyService.GetRealTimePrice(&currency)
		if err != nil {
			global.GVA_LOG.Error("获取实时价格失败", zap.Error(err))
			withdrawChannels.PriceUsdt = currency.PriceUsdt
		} else {
			// 根据TicketSize处理价格精度
			if currency.TicketSize != decimal.Zero {
				decimalPlaces := utils.GetDecimalPlaces2(currency.TicketSize)
				processedPrice := utils.RoundDecimal(realTimePrice, int32(decimalPlaces))
				withdrawChannels.PriceUsdt = processedPrice
			} else {
				withdrawChannels.PriceUsdt = realTimePrice
			}
		}
	}
	return
}

// GetWithdrawChannelsInfoList 分页获取withdrawChannels表记录
// Author [yourname](https://github.com/yourname)
func (withdrawChannelsService *WithdrawChannelsService) GetWithdrawChannelsInfoList(info userfundReq.WithdrawChannelsSearch) (list []userfund.WithdrawChannels, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&userfund.WithdrawChannels{})
	var withdrawChannelss []userfund.WithdrawChannels

	// 添加 CoinId 查询条件
	if info.CoinId != nil {
		db = db.Where("coin_id = ?", *info.CoinId)
	}

	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	db = db.Order("sort_order DESC")
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&withdrawChannelss).Error
	if err != nil {
		return
	}

	// 收集所有的 CoinId
	var coinIds []int
	for _, channel := range withdrawChannelss {
		if channel.CoinId != nil {
			coinIds = append(coinIds, *channel.CoinId)
		}
	}

	// 如果有关联的货币，批量查询货币信息
	if len(coinIds) > 0 {
		var currencies []userfund.Currencies
		err = global.GVA_DB.Where("id IN ?", coinIds).Find(&currencies).Error
		if err != nil {
			return
		}

		// 构建货币信息映射
		currencyMap := make(map[int]userfund.Currencies)
		for _, currency := range currencies {
			if currency.Id != nil {
				currencyMap[*currency.Id] = currency
			}
		}

		// 组装数据时添加精度处理
		for i := range withdrawChannelss {
			if withdrawChannelss[i].CoinId != nil {
				if currency, exists := currencyMap[*withdrawChannelss[i].CoinId]; exists {
					withdrawChannelss[i].Currency = currency.Currency
					withdrawChannelss[i].CurrencyIcon = currency.Icon
					withdrawChannelss[i].CoinType = currency.CoinType
					withdrawChannelss[i].TicketSize = currency.TicketSize
					withdrawChannelss[i].TicketNumSize = currency.TicketNumSize
					withdrawChannelss[i].MinWithdrawNum = currency.MinWithdrawNum

					// 获取实时价格并按精度处理
					currencyService := CurrenciesService{}
					realTimePrice, err := currencyService.GetRealTimePrice(&currency)
					if err != nil {
						global.GVA_LOG.Error("获取实时价格失败", zap.Error(err))
						withdrawChannelss[i].PriceUsdt = currency.PriceUsdt
					} else {
						// 根据TicketSize处理价格精度
						if currency.TicketSize != decimal.Zero {
							decimalPlaces := utils.GetDecimalPlaces2(currency.TicketSize)
							processedPrice := utils.RoundDecimal(realTimePrice, int32(decimalPlaces))
							withdrawChannelss[i].PriceUsdt = processedPrice
						} else {
							withdrawChannelss[i].PriceUsdt = realTimePrice
						}
					}
				}
			}
		}
	}

	return withdrawChannelss, total, nil
}
func (withdrawChannelsService *WithdrawChannelsService) GetWithdrawChannelsPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

func (withdrawChannelsService *WithdrawChannelsService) GetWithdrawChannelsInfoList2(info userfundReq.WithdrawChannelsSearch) (list []userfund.WithdrawChannels, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&userfund.WithdrawChannels{})
	var withdrawChannelss []userfund.WithdrawChannels

	// 添加 CoinId 查询条件
	if info.CoinId != nil {
		db = db.Where("coin_id = ?", *info.CoinId)
	}

	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	db = db.Where("status = ? ", "1")
	db = db.Order("sort_order DESC")
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&withdrawChannelss).Error
	if err != nil {
		return
	}

	// 收集所有的 CoinId
	var coinIds []int
	for _, channel := range withdrawChannelss {
		if channel.CoinId != nil {
			coinIds = append(coinIds, *channel.CoinId)
		}
	}

	// 如果有关联的货币，批量查询货币信息
	if len(coinIds) > 0 {
		var currencies []userfund.Currencies
		err = global.GVA_DB.Where("id IN ?", coinIds).Find(&currencies).Error
		if err != nil {
			return
		}

		// 构建货币信息映射
		currencyMap := make(map[int]userfund.Currencies)
		for _, currency := range currencies {
			if currency.Id != nil {
				currencyMap[*currency.Id] = currency
			}
		}

		// 组装数据时添加精度处理
		for i := range withdrawChannelss {
			if withdrawChannelss[i].CoinId != nil {
				if currency, exists := currencyMap[*withdrawChannelss[i].CoinId]; exists {
					withdrawChannelss[i].Currency = currency.Currency
					withdrawChannelss[i].CurrencyIcon = currency.Icon
					withdrawChannelss[i].CoinType = currency.CoinType
					withdrawChannelss[i].TicketSize = currency.TicketSize
					withdrawChannelss[i].TicketNumSize = currency.TicketNumSize
					withdrawChannelss[i].MinWithdrawNum = currency.MinWithdrawNum

					// 获取实时价格并按精度处理
					currencyService := CurrenciesService{}
					realTimePrice, err := currencyService.GetRealTimePrice(&currency)
					if err != nil {
						global.GVA_LOG.Error("获取实时价格失败", zap.Error(err))
						withdrawChannelss[i].PriceUsdt = currency.PriceUsdt
					} else {
						// 根据TicketSize处理价格精度
						if currency.TicketSize != decimal.Zero {
							decimalPlaces := utils.GetDecimalPlaces2(currency.TicketSize)
							processedPrice := utils.RoundDecimal(realTimePrice, int32(decimalPlaces))
							withdrawChannelss[i].PriceUsdt = processedPrice
						} else {
							withdrawChannelss[i].PriceUsdt = realTimePrice
						}
					}
				}
			}
		}
	}

	return withdrawChannelss, total, nil
}
