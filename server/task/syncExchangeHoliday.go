package task

import (
	"encoding/json"
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	polygonMarketHolidayApi = "/v1/marketstatus/upcoming"
	closed                  = "closed"
	nasdaq                  = "NASDAQ"
)

// HolidayInfo  节假日信息
//
//	示例数据:
//	[{
//	   "date": "2020-11-26",
//	   "exchange": "NASDAQ",
//	   "name": "Thanksgiving",
//	   "status": "closed"
//	}]
type HolidayInfo struct {
	Name     string `json:"name"` //节日名称
	Exchange string `json:"exchange"`
	Date     string `json:"date"`
	Status   string `json:"status"`
}

func SyncExchangeHoliday(db *gorm.DB) error {
	holidays, err := getHolidays()
	if err != nil {
		return err
	}

	for _, holiday := range holidays {
		var exchangeHoliday system.StockExchangeHoliday
		date, _ := time.Parse(time.DateOnly, holiday.Date)
		qErr := db.Model(&exchangeHoliday).
			Where("name = ? and exchange = ? and date = cast(? as date)", holiday.Name, holiday.Exchange, date).
			First(&exchangeHoliday).Error
		if qErr != nil && !errors.Is(qErr, gorm.ErrRecordNotFound) {
			global.GVA_LOG.Error("get stock exchange holiday fail", zap.Error(qErr))
			continue
		} else if exchangeHoliday.ID > 0 {
			continue
		}

		exchangeHoliday.Name = holiday.Name
		exchangeHoliday.Exchange = holiday.Exchange
		exchangeHoliday.Date = date
		exchangeHoliday.Status = holiday.Status
		if e := db.Create(&exchangeHoliday).Error; e != nil {
			global.GVA_LOG.Error("create stock holiday fail", zap.Error(e))
		}
	}
	return nil
}

// getHolidays 查询节假日信息
func getHolidays() ([]HolidayInfo, error) {
	api, err := url.Parse(global.GVA_CONFIG.Polygon.BaseURL + polygonMarketHolidayApi)
	if err != nil {
		global.GVA_LOG.Error("polygonMarketHolidayApi parse error.", zap.Error(err))
		return nil, err
	}
	params := url.Values{}
	params.Set("apiKey", global.GVA_CONFIG.Polygon.APIKey)
	api.RawQuery = params.Encode()

	resp, err := http.Get(api.String())
	if err != nil {
		global.GVA_LOG.Error("polygonMarketHolidayApi http.Get error", zap.Error(err))
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	body, _ := io.ReadAll(resp.Body)
	var holidays []HolidayInfo
	err = json.Unmarshal(body, &holidays)
	if err != nil {
		global.GVA_LOG.Error("polygonMarketHolidayApi json.Unmarshal error", zap.Error(err))
		return nil, err
	}

	return holidays, nil
}
