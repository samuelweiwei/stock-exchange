package task

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	symbolModel "github.com/flipped-aurora/gin-vue-admin/server/model/symbol"
	symbolService "github.com/flipped-aurora/gin-vue-admin/server/service/symbol"
	"go.uber.org/zap"
)

// UpdateSymbolsInfo 更新所有类型symbols的基本信息的定时任务
func UpdateSymbolsInfo() error {
	symbolsService := &symbolService.SymbolsService{}

	symbols, err := symbolsService.GetAllSymbols()
	if err != nil {
		global.GVA_LOG.Error("获取symbols失败", zap.Error(err))
		return err
	}

	// 按类型分组symbols
	symbolsByType := make(map[int][]symbolModel.Symbols)
	for _, s := range symbols {
		if s.Type != nil {
			symbolsByType[*s.Type] = append(symbolsByType[*s.Type], s)
		}
	}

	// 为每种类型的symbols分别处理
	for typeID, typeSymbols := range symbolsByType {
		// 将symbols分批处理，每批100个
		batchSize := 100
		for i := 0; i < len(typeSymbols); i += batchSize {
			end := i + batchSize
			if end > len(typeSymbols) {
				end = len(typeSymbols)
			}

			currentBatch := typeSymbols[i:end]

			// 构建查询参数
			tickers := make([]string, len(currentBatch))
			for j, s := range currentBatch {
				switch typeID {
				case 0: // 股票
					tickers[j] = s.Symbol
				case 1: // 加密货币
					// 移除 "/" 并添加 "X:" 前缀
					symbol := strings.ReplaceAll(s.Symbol, "/", "")
					tickers[j] = "X:" + symbol
				case 2: // 外汇
					// 移除 "/" 并添加 "C:" 前缀
					symbol := strings.ReplaceAll(s.Symbol, "/", "")
					tickers[j] = "C:" + symbol
				}
			}

			// 1. 获取基本信息
			url := fmt.Sprintf("%s/v3/reference/tickers?apiKey=%s&tickers=%s&active=true&limit=%d",
				global.GVA_CONFIG.Polygon.BaseURL,
				global.GVA_CONFIG.Polygon.APIKey,
				strings.Join(tickers, ","),
				batchSize)

			client := &http.Client{
				Timeout: time.Second * 30,
			}

			resp, err := client.Get(url)
			if err != nil {
				global.GVA_LOG.Error("请求Polygon API失败", zap.Error(err))
				continue
			}
			defer resp.Body.Close()

			var result struct {
				Results []struct {
					Ticker           string  `json:"ticker"`
					Name             string  `json:"name"`
					Market           string  `json:"market"`
					Locale           string  `json:"locale"`
					PrimaryExch      string  `json:"primary_exchange"`
					Type             string  `json:"type"`
					Active           bool    `json:"active"`
					MarketCap        float64 `json:"market_cap"`
					Industry         string  `json:"sic_description"`
					ListDate         string  `json:"list_date"`
					Description      string  `json:"description"`
					CurrencyName     string  `json:"currency_name"`
					BaseCurrencyName string  `json:"base_currency_name"`
					Branding         struct {
						IconUrl string `json:"icon_url"`
					} `json:"branding"`
				} `json:"results"`
			}

			if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
				global.GVA_LOG.Error("解析响应失败", zap.Error(err))
				continue
			}

			// 2. 获取日交易数据
			now := time.Now()
			dailyUrl := fmt.Sprintf("%s/v2/aggs/ticker/%s/range/1/day/%s/%s?apiKey=%s",
				global.GVA_CONFIG.Polygon.BaseURL,
				strings.Join(tickers, ","),
				now.AddDate(0, 0, -5).Format("2006-01-02"),
				now.Format("2006-01-02"),
				global.GVA_CONFIG.Polygon.APIKey)

			dailyResp, err := client.Get(dailyUrl)
			if err != nil {
				global.GVA_LOG.Error("请求日交易数据失败", zap.Error(err))
				continue
			}
			defer dailyResp.Body.Close()

			var dailyData struct {
				Results map[string][]struct {
					V float64 `json:"v"` // 交易量
					C float64 `json:"c"` // 收盘价
					O float64 `json:"o"` // 开盘价
				} `json:"results"`
			}

			if err := json.NewDecoder(dailyResp.Body).Decode(&dailyData); err != nil {
				global.GVA_LOG.Error("解析日交易数据失败", zap.Error(err))
				continue
			}

			// 创建symbol映射以便快速查找和更新
			symbolMap := make(map[string]*symbolModel.Symbols)
			for i := range currentBatch {
				// 根据类型构建正确的key
				var key string
				switch typeID {
				case 0: // 股票
					key = currentBatch[i].Symbol
				case 1: // 加密货币
					key = "X:" + currentBatch[i].Symbol
				case 2: // 外汇
					key = "C:" + currentBatch[i].Symbol
				}
				symbolMap[key] = &currentBatch[i]
			}

			// 更新symbols数据
			var updatedSymbols []symbolModel.Symbols
			for _, res := range result.Results {
				if sym, ok := symbolMap[res.Ticker]; ok {
					// 根据类型设置不同的字段
					switch typeID {
					case 0: // 股票
						sym.Corporation = res.Name
						sym.Exchange = res.PrimaryExch
						sym.Industry = res.Industry
						sym.Description = res.Description

						// 更新市值
						marketCap := int(res.MarketCap)
						sym.MarketCap = &marketCap

						// 更新上市日期
						if res.ListDate != "" {
							if listDate, err := time.Parse("2006-01-02", res.ListDate); err == nil {
								sym.ListDate = &listDate
							}
						}

					case 1: // 加密货币
						sym.Corporation = res.BaseCurrencyName
						sym.Exchange = res.Market

					case 2: // 外汇
						sym.Corporation = res.BaseCurrencyName
						sym.Exchange = res.Market
					}

					// 更新图标URL
					sym.Icon = res.Branding.IconUrl

					// 更新日交易数据
					if dailyResults, ok := dailyData.Results[res.Ticker]; ok && len(dailyResults) > 0 {
						lastDay := dailyResults[len(dailyResults)-1]

						// 更新交易量
						volume := int(lastDay.V)
						sym.AverageVolume = &volume

						// 更新涨跌幅
						ratio := (lastDay.C - lastDay.O) / lastDay.O * 100
						sym.ChangeRatio = &ratio

						// 更新当前价格
						currentPrice := lastDay.C
						sym.CurrentPrice = &currentPrice

						// 更新Redis缓存
						var redisKey string
						switch typeID {
						case 0:
							redisKey = fmt.Sprintf("symbol:stock:%s", sym.Symbol)
						case 1:
							redisKey = fmt.Sprintf("symbol:crypto:%s", sym.Symbol)
						case 2:
							redisKey = fmt.Sprintf("symbol:forex:%s", sym.Symbol)
						}

						// 使用管道批量更新Redis
						pipe := global.GVA_REDIS.Pipeline()
						ctx := context.Background()

						// 设置当前价格
						pipe.Set(ctx, redisKey, currentPrice, 0)

						// 设置前一天收盘价（如果有多天数据）
						if len(dailyResults) > 1 {
							previousDay := dailyResults[len(dailyResults)-2]
							previousCloseKey := redisKey + ":previous_close"
							pipe.Set(ctx, previousCloseKey, previousDay.C, 0)
						}

						if _, err := pipe.Exec(ctx); err != nil {
							global.GVA_LOG.Error("Failed to set price in Redis",
								zap.String("symbol", sym.Symbol),
								zap.Error(err))
						}
					}

					updatedSymbols = append(updatedSymbols, *sym)
				}
			}

			// 批量更新数据库
			if len(updatedSymbols) > 0 {
				if err := symbolsService.BatchUpdateSymbols(updatedSymbols); err != nil {
					global.GVA_LOG.Error("批量更新失败", zap.Error(err))
				}
			}

			// 避免触发API限制
			time.Sleep(time.Second * 1)
		}
	}

	return nil
}
