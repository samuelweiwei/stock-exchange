package userfund

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/userfund"
	userfundReq "github.com/flipped-aurora/gin-vue-admin/server/model/userfund/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type RechargeChannelsService struct{}

// CreateRechargeChannels 创建rechargeChannels表记录
// Author [yourname](https://github.com/yourname)
func (rechargeChannelsService *RechargeChannelsService) CreateRechargeChannels(rechargeChannels *userfund.RechargeChannels) (err error) {
	err = global.GVA_DB.Create(rechargeChannels).Error
	return err
}

// DeleteRechargeChannels 删除rechargeChannels表记录
// Author [yourname](https://github.com/yourname)
func (rechargeChannelsService *RechargeChannelsService) DeleteRechargeChannels(ID string) (err error) {
	err = global.GVA_DB.Delete(&userfund.RechargeChannels{}, "id = ?", ID).Error
	return err
}

// DeleteRechargeChannelsByIds 批量删除rechargeChannels表记录
// Author [yourname](https://github.com/yourname)
func (rechargeChannelsService *RechargeChannelsService) DeleteRechargeChannelsByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]userfund.RechargeChannels{}, "id in ?", IDs).Error
	return err
}

// UpdateRechargeChannels 更新rechargeChannels表记录
// Author [yourname](https://github.com/yourname)
func (rechargeChannelsService *RechargeChannelsService) UpdateRechargeChannels(rechargeChannels userfund.RechargeChannels) (err error) {
	err = global.GVA_DB.Model(&userfund.RechargeChannels{}).Where("id = ?", rechargeChannels.ID).Updates(&rechargeChannels).Error
	return err
}

// GetRechargeChannels 根据ID获取rechargeChannels表记录
// Author [yourname](https://github.com/yourname)
func (rechargeChannelsService *RechargeChannelsService) GetRechargeChannels(ID string) (rechargeChannels userfund.RechargeChannels, err error) {
	// 先查询充值渠道信息
	err = global.GVA_DB.Where("id = ?", ID).First(&rechargeChannels).Error
	if err != nil {
		return
	}

	// 如果有关联的货币ID，查询货币信息
	if rechargeChannels.CoinId != nil {
		var currency userfund.Currencies
		err = global.GVA_DB.Where("id = ?", *rechargeChannels.CoinId).First(&currency).Error
		if err != nil {
			return
		}
		// 组装货币信息
		rechargeChannels.Currency = currency.Currency
		rechargeChannels.CurrencyIcon = currency.Icon
		rechargeChannels.CoinType = currency.CoinType
		rechargeChannels.TicketSize = currency.TicketSize
		rechargeChannels.TicketNumSize = currency.TicketNumSize
		rechargeChannels.MinRechargeNum = currency.MinRechargeNum

		// 获取实时价格并按精度处理
		currencyService := CurrenciesService{}
		realTimePrice, err := currencyService.GetRealTimePrice(&currency)

		if err != nil {
			global.GVA_LOG.Error("获取实时价格失败", zap.Error(err))
			rechargeChannels.PriceUsdt = currency.PriceUsdt
		} else {
			// 根据TicketSize处理价格精度

			if currency.TicketSize.Equal(decimal.Zero) {
				decimalPlaces := utils.GetDecimalPlaces2(currency.TicketSize)
				processedPrice := utils.RoundDecimal(realTimePrice, int32(decimalPlaces))
				rechargeChannels.PriceUsdt = processedPrice
			} else {
				rechargeChannels.PriceUsdt = realTimePrice
			}
		}
	}
	return
}

// GetRechargeChannelsInfoList 分页获取rechargeChannels表记录
func (rechargeChannelsService *RechargeChannelsService) GetRechargeChannelsInfoList(info userfundReq.RechargeChannelsSearch) (list []userfund.RechargeChannels, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&userfund.RechargeChannels{})

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

	err = db.Find(&list).Error
	if err != nil {
		return
	}

	// 收集所有的 CoinId
	var coinIds []int
	for _, channel := range list {
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
		for i := range list {
			if list[i].CoinId != nil {
				if currency, exists := currencyMap[*list[i].CoinId]; exists {
					list[i].Currency = currency.Currency
					list[i].CurrencyIcon = currency.Icon
					list[i].CoinType = currency.CoinType
					list[i].TicketSize = currency.TicketSize
					list[i].TicketNumSize = currency.TicketNumSize
					list[i].MinRechargeNum = currency.MinRechargeNum

					// 获取实时价格并按精度处理
					currencyService := CurrenciesService{}
					realTimePrice, err := currencyService.GetRealTimePrice(&currency)
					if err != nil {
						global.GVA_LOG.Error("获取实时价格失败", zap.Error(err))
						list[i].PriceUsdt = currency.PriceUsdt
					} else {
						// 根据TicketSize处理价格精度
						if currency.TicketSize.Equal(decimal.Zero) {
							decimalPlaces := utils.GetDecimalPlaces2(currency.TicketSize)
							processedPrice := utils.RoundDecimal(realTimePrice, int32(decimalPlaces))
							list[i].PriceUsdt = processedPrice
						} else {
							list[i].PriceUsdt = realTimePrice
						}
					}
				}
			}
		}
	}

	return list, total, nil
}
func (rechargeChannelsService *RechargeChannelsService) GetRechargeChannelsPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// Paynotify 支付回调接口
// Author [yourname](https://github.com/yourname)
func (rechargeChannelsService *RechargeChannelsService) Paynotify() (err error) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&userfund.RechargeChannels{})
	return db.Error
}

// GetRechargeChannelsInfoList 分页获取rechargeChannels表记录
func (rechargeChannelsService *RechargeChannelsService) GetRechargeChannelsInfoList2(info userfundReq.RechargeChannelsSearch) (list []userfund.RechargeChannels, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&userfund.RechargeChannels{})

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

	err = db.Find(&list).Error
	if err != nil {
		return
	}

	// 收集所有的 CoinId
	var coinIds []int
	for _, channel := range list {
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
		for i := range list {
			if list[i].CoinId != nil {
				if currency, exists := currencyMap[*list[i].CoinId]; exists {
					list[i].Currency = currency.Currency
					list[i].CurrencyIcon = currency.Icon
					list[i].CoinType = currency.CoinType
					list[i].TicketSize = currency.TicketSize
					list[i].TicketNumSize = currency.TicketNumSize
					list[i].MinRechargeNum = currency.MinRechargeNum

					// 获取实时价格并按精度处理
					currencyService := CurrenciesService{}
					realTimePrice, err := currencyService.GetRealTimePrice(&currency)
					if err != nil {
						global.GVA_LOG.Error("获取实时价格失败", zap.Error(err))
						list[i].PriceUsdt = currency.PriceUsdt
					} else {
						// 根据TicketSize处理价格精度
						if currency.TicketSize.Equal(decimal.Zero) {
							decimalPlaces := utils.GetDecimalPlaces2(currency.TicketSize)
							processedPrice := utils.RoundDecimal(realTimePrice, int32(decimalPlaces))
							list[i].PriceUsdt = processedPrice
						} else {
							list[i].PriceUsdt = realTimePrice
						}
					}
				}
			}
		}
	}

	return list, total, nil
}
