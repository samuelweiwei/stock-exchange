package userfund

import (
	"context"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/userfund"
	userfundReq "github.com/flipped-aurora/gin-vue-admin/server/model/userfund/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	. "github.com/shopspring/decimal"
)

type CurrenciesService struct{}

// CreateCurrencies 创建currencies表记录
// Author [yourname](https://github.com/yourname)
func (currenciesService *CurrenciesService) CreateCurrencies(currencies *userfund.Currencies) (err error) {
	err = global.GVA_DB.Create(currencies).Error
	return err
}

// DeleteCurrencies 删除currencies表记录
// Author [yourname](https://github.com/yourname)
func (currenciesService *CurrenciesService) DeleteCurrencies(id string) (err error) {
	err = global.GVA_DB.Delete(&userfund.Currencies{}, "id = ?", id).Error
	return err
}

// DeleteCurrenciesByIds 批量删除currencies表记录
// Author [yourname](https://github.com/yourname)
func (currenciesService *CurrenciesService) DeleteCurrenciesByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]userfund.Currencies{}, "id in ?", ids).Error
	return err
}

// UpdateCurrencies 更新currencies表记录
// Author [yourname](https://github.com/yourname)
func (currenciesService *CurrenciesService) UpdateCurrencies(currencies userfund.Currencies) (err error) {
	err = global.GVA_DB.Model(&userfund.Currencies{}).Where("id = ?", currencies.Id).Updates(&currencies).Error
	return err
}

// ProcessPriceByCurrency 根据货币处理价格精度
func (currenciesService *CurrenciesService) ProcessPriceByCurrency(price Decimal, currencyName string) (processedPrice Decimal, err error) {
	var currency userfund.Currencies
	// 查询数据库，根据货币名称获取对应记录
	err = global.GVA_DB.Model(&userfund.Currencies{}).Where("currency = ?", currencyName).First(&currency).Error
	if err != nil {
		return Zero, err
	}

	decimalPlaces := utils.GetDecimalPlaces2(currency.TicketSize)
	processedPrice = utils.TruncateDecimal(price, int32(decimalPlaces))

	return processedPrice, nil
}

// GetCurrencies 根据id获取currencies表记录
// Author [yourname](https://github.com/yourname)
func (currenciesService *CurrenciesService) GetCurrencies(id string) (currencies userfund.Currencies, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&currencies).Error
	return
}

// GetCurrenciesInfoList 分页获取currencies表记录
// Author [yourname](https://github.com/yourname)
func (currenciesService *CurrenciesService) GetCurrenciesInfoList(info userfundReq.CurrenciesSearch) (list []userfund.Currencies, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	// 创建db
	db := global.GVA_DB.Model(&userfund.Currencies{})
	// 添加 CoinId 查询条件
	if info.Currency != "" {
		db = db.Where("currency LIKE ?", "%"+info.Currency+"%")
	}

	if info.CoinType != nil {
		db = db.Where("coin_type  =  ? ", info.CoinType)
	}
	db.Order("created_at desc")
	var currenciess []userfund.Currencies
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&currenciess).Error

	return currenciess, total, err
}

// GetCurrenciesPublic 获取货币列表
func (currenciesService *CurrenciesService) GetCurrenciesPublic(info userfundReq.CurrenciesSearch) (list []userfund.Currencies, err error) {
	db := global.GVA_DB.Model(&userfund.Currencies{})

	// 添加类型筛选条件
	if info.CoinType != nil {
		db = db.Where("coin_type = ?", *info.CoinType)
	}
	if info.Source == "recharge" {
		// 查询已经被渠道关联的币种（假设你要关联的是 channels 表）
		db = db.Joins("INNER JOIN recharge_channels ch ON ch.coin_id = currencies.id and ch.status ='1' and  ch.deleted_at is null ")
	}
	if info.Source == "withdraw" {
		// 查询已经被渠道关联的币种（假设你要关联的是 channels 表）
		db = db.Joins("INNER JOIN withdraw_channels ch ON ch.coin_id = currencies.id and ch.status ='1' and  ch.deleted_at is null  ")
	}

	// 按照创建时间倒序排序
	//db = db.Order("created_at desc")
	db = db.Select("DISTINCT currencies.id, currencies.currency, currencies.icon, currencies.price_usdt, currencies.coin_type, currencies.min_recharge_num,currencies.min_withdraw_num, currencies.created_at").Order("currencies.id asc")
	err = db.Find(&list).Error
	return list, err
}

// GetExchangeRate 获取货币兑换USDT的汇率
func (currenciesService *CurrenciesService) GetExchangeRate(currency userfund.Currencies) (rate Decimal, err error) {
	// 如果是USDT，直接返回1
	if currency.Currency == "USDT" {
		return NewFromFloat(1.0), nil
	}

	// 获取实时价格
	realTimePrice, err := currenciesService.GetRealTimePrice(&currency)
	if err != nil {
		return Zero, err
	}

	return realTimePrice, nil
}

// CalculateUsdtAmount 计算USDT金额
func (currenciesService *CurrenciesService) CalculateUsdtAmount(amount Decimal, currency userfund.Currencies) (usdtAmount Decimal, err error) {
	rate, err := currenciesService.GetExchangeRate(currency)
	if err != nil {
		return Zero, err
	}

	return amount.Mul(rate), nil
}

// CalculateAmountFromUsdt 根据USDT金额计算指定货币的金额
func (currenciesService *CurrenciesService) CalculateAmountFromUsdt(usdtAmount Decimal, currency userfund.Currencies) (amount Decimal, err error) {
	rate, err := currenciesService.GetExchangeRate(currency)
	if err != nil {
		return Zero, err
	}
	if rate.Equal(Zero) {
		return Zero, fmt.Errorf("汇率不能为0")
	}
	return usdtAmount.Div(rate), nil
}

// GetRealTimePrice 获取货币的实时价格
func (currenciesService *CurrenciesService) GetRealTimePrice(currency *userfund.Currencies) (Decimal, error) {
	// 如果是USDT，直接返回1
	if currency.Currency == "USDT" {
		price := NewFromFloat(1.0)
		return price, nil
	}

	var price Decimal

	// 如果是加密货币(type=1)且不是USDT，从redis获取实时价格
	if currency.CoinType != nil && *currency.CoinType == 1 && currency.Currency != "USDT" {
		// 构建redis key，格式为：symbol:crypto:ETH/USD
		symbol := fmt.Sprintf("%s/USD", currency.Currency)
		redisKey := fmt.Sprintf("symbol:crypto:%s", symbol)

		// 从redis获取价格
		result, err := global.GVA_REDIS.Get(context.Background(), redisKey).Result()
		if err != nil {
			// 如果Redis查询失败，使用数据库中的价格
			if currency.PriceUsdt == Zero {
				return Zero, fmt.Errorf("无法获取%s的价格", currency.Currency)
			}
			price = currency.PriceUsdt
		} else {
			// 转换价格为float64
			//price, err = strconv.ParseFloat(result, 64)
			price, err = NewFromString(result)
			if err != nil {
				// 如果价格解析失败，使用数据库中的价格
				if currency.PriceUsdt == Zero {
					return Zero, fmt.Errorf("无法解析%s的价格", currency.Currency)
				}
				price = currency.PriceUsdt
			}
		}
	} else {
		// 如果是法币(type=2)或其他情况，使用数据库中的价格
		if currency.PriceUsdt == Zero {
			return Zero, fmt.Errorf("未设置%s的价格", currency.Currency)
		}
		price = currency.PriceUsdt
	}

	// 根据TicketSize处理价格精度
	if currency.TicketSize != Zero {
		decimalPlaces := utils.GetDecimalPlaces2(currency.TicketSize)
		processedPrice := utils.RoundDecimal(price, int32(decimalPlaces))
		return processedPrice, nil
	}

	return price, nil
}
