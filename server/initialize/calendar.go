package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/rickar/cal/v2"
	"go.uber.org/zap"
	"time"
)

const (
	closed = "closed"
	nasdaq = "NASDAQ"
)

// BusinessCalendar 初始化日历
//
//  1. 初始化BusinessCalendar
//  2. 查询stock_exchange_holiday表中nasdaq交易所为closed的日期，添加到BusinessCalendar
func BusinessCalendar() {
	global.Calendar = cal.NewBusinessCalendar()

	var exchangeHolidays []system.StockExchangeHoliday
	err := global.GVA_DB.Model(&system.StockExchangeHoliday{}).
		Where("exchange = ? and status = ?", nasdaq, closed).Find(&exchangeHolidays).Error
	if err != nil {
		global.GVA_LOG.Error("query nasdaq closed holidays error.", zap.Error(err))
		return
	}

	for _, holiday := range exchangeHolidays {
		global.Calendar.AddHoliday(&cal.Holiday{
			Name:      holiday.Name,
			Type:      cal.ObservanceOther,
			StartYear: holiday.Date.Year(),
			EndYear:   holiday.Date.Year(),
			Month:     holiday.Date.Month(),
			Day:       holiday.Date.Day(),
			Func:      cal.CalcDayOfMonth,
		})
	}

	refresh()
}

func refresh() {
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				var exchangeHolidays []system.StockExchangeHoliday
				err := global.GVA_DB.Model(&system.StockExchangeHoliday{}).
					Where("exchange = ? and status = ?", nasdaq, closed).Find(&exchangeHolidays).Error
				if err != nil {
					global.GVA_LOG.Error("query nasdaq closed holidays error.", zap.Error(err))
					continue
				}

				for _, holiday := range exchangeHolidays {
					if !IsInCalendar(holiday) {
						global.Calendar.AddHoliday(&cal.Holiday{
							Name:      holiday.Name,
							Type:      cal.ObservanceOther,
							StartYear: holiday.Date.Year(),
							EndYear:   holiday.Date.Year(),
							Month:     holiday.Date.Month(),
							Day:       holiday.Date.Day(),
							Func:      cal.CalcDayOfMonth,
						})
					}
				}
			}
		}
	}()
}

func IsInCalendar(holiday system.StockExchangeHoliday) bool {
	for _, c := range global.Calendar.Holidays {
		if c.Type == cal.ObservanceOther &&
			c.Name == holiday.Name &&
			c.StartYear == holiday.Date.Year() &&
			c.EndYear == holiday.Date.Year() &&
			c.Month == holiday.Date.Month() &&
			c.Day == holiday.Date.Day() {
			return true
		}
	}
	return false
}
