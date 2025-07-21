package task

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	symbolService "github.com/flipped-aurora/gin-vue-admin/server/service/symbol"
	"go.uber.org/zap"
)

// UpdateSymbolsKlineLast 更新所有symbols的K线数据
func UpdateSymbolsKlineLast() error {
	defer func() {
		if err := recover(); err != nil {
			global.GVA_LOG.Error(fmt.Sprintf("Recover from UpdateSymbolsKlineLast panic: %v", err))
		}
	}()
	symbolsService := &symbolService.SymbolsService{}

	symbols, err := symbolsService.GetAllSymbols()
	if err != nil {
		global.GVA_LOG.Error("获取symbols失败", zap.Error(err))
		return err
	}

	// 直接按类型分组所有symbols
	symbolsByType := make(map[int][]string)
	for _, s := range symbols {
		if s.Type != nil {
			symbolsByType[*s.Type] = append(symbolsByType[*s.Type], s.Symbol)
		}
	}

	// 记录处理的symbols数量
	global.GVA_LOG.Info("开始处理symbols",
		zap.Int("total_symbols", len(symbols)))
	for typeID, typeSymbols := range symbolsByType {
		global.GVA_LOG.Info("类型统计",
			zap.Int("typeID", typeID),
			zap.Int("count", len(typeSymbols)))
	}

	// 定义需要获取的时间周期
	intervals := []string{"1m", "1h", "4h", "1d", "1w", "1M"}

	client := &http.Client{
		Timeout: time.Second * 60,
		Transport: &http.Transport{
			ForceAttemptHTTP2:     false,
			MaxIdleConns:          200,
			MaxIdleConnsPerHost:   200,
			IdleConnTimeout:       90 * time.Second,
			DisableKeepAlives:     false,
			ResponseHeaderTimeout: 30 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}

	ctx := context.Background()

	var wg sync.WaitGroup

	// 第一层循环：按资产类型处理
	for typeID, symbols := range symbolsByType {
		// typeID: 0(股票), 1(加密货币), 2(外汇)
		// symbols: 该类型下的所有交易对

		// 第二层循环：按时间周期处理
		for _, interval := range intervals {

			// 第三层循环：处理每个具体的交易对
			for _, symbol := range symbols {
				// symbol: 具体的交易对，如 "AAPL"(股票)、"BTC/USD"(加密货币)、"EUR/USD"(外汇)
				// 构建单个ticker的查询参数
				var ticker string
				switch typeID {
				case 0: // 股票
					ticker = symbol
				case 1: // 加密货币
					ticker = "X:" + strings.ReplaceAll(symbol, "/", "")
				case 2: // 外汇
					ticker = "C:" + strings.ReplaceAll(symbol, "/", "")
				}

				// 转换时间间隔为Polygon支持的格式
				multiplier, timespan, err := parseInterval(interval)
				if err != nil {
					global.GVA_LOG.Error("解析时间间隔失败",
						zap.String("interval", interval),
						zap.Error(err))
					continue
				}

				// 获取数据
				now := time.Now()
				var startTime time.Time
				switch interval {
				case "1m":
					startTime = now.AddDate(0, 0, -4) // 增加到4小时，确保能获取到更多数据
				case "1h":
					startTime = now.AddDate(0, 0, -4) // 获取最近2天的小时数据
				case "4h":
					startTime = now.AddDate(0, 0, -7) // 获取最近3天的4小时数据
				case "1d":
					startTime = now.AddDate(0, 0, -7) // 获取最近7天的日线数据
				case "1w":
					startTime = now.AddDate(0, -1, 0) // 获取最近1个月的周线数据
				case "1M":
					startTime = now.AddDate(-1, 0, 0) // 获取最近1年的月线数据
				}

				// 在URL中添加adjusted参数，获取调整后的数据
				url := fmt.Sprintf("%s/v2/aggs/ticker/%s/range/%d/%s/%d/%d?apiKey=%s&sort=desc",
					global.GVA_CONFIG.Polygon.BaseURL,
					ticker,
					multiplier,
					timespan,
					startTime.UnixMilli(),
					now.UnixMilli(),
					global.GVA_CONFIG.Polygon.APIKey)

				time.Sleep(100 * time.Millisecond) // 每个请求间隔100ms

				resp, err := makeRequest(client, url, 3) // 最多重试3次
				if err != nil {
					global.GVA_LOG.Error("请求K线数据失败",
						zap.String("symbol", symbol),
						zap.String("interval", interval),
						zap.Error(err))
					continue
				}

				var result struct {
					Results []struct {
						O      float64 `json:"o"`
						H      float64 `json:"h"`
						L      float64 `json:"l"`
						C      float64 `json:"c"`
						V      float64 `json:"v"`
						T      int64   `json:"t"`
						Ticker string  `json:"T"`
					} `json:"results"`
					Status    string `json:"status"`
					RequestID string `json:"request_id"`
				}

				if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
					resp.Body.Close()
					global.GVA_LOG.Error("解析K线数据失败", zap.Error(err))
					continue
				}
				resp.Body.Close()

				// 为每个结果创建K线数据
				if len(result.Results) > 0 {
					lastKline := result.Results[0]

					// 从ticker中提取实际的symbol
					switch typeID {
					case 1: // 加密货币
						symbol = strings.TrimPrefix(symbol, "X:")
					case 2: // 外汇
						symbol = strings.TrimPrefix(symbol, "C:")
					}

					// 构建Redis key
					var baseKey string
					switch typeID {
					case 0:
						baseKey = fmt.Sprintf("symbol:stock:%s", symbol)
					case 1:
						baseKey = fmt.Sprintf("symbol:crypto:%s", symbol)
					case 2:
						baseKey = fmt.Sprintf("symbol:forex:%s", symbol)
					}

					klineData := map[string]interface{}{
						"o": lastKline.O,
						"h": lastKline.H,
						"l": lastKline.L,
						"c": lastKline.C,
						"v": lastKline.V,
						"t": lastKline.T,
					}

					// 序列化K线数据
					if jsonData, err := json.Marshal(klineData); err == nil {
						// 直接更新K线数据
						redisKey := fmt.Sprintf("%s:kline:%s", baseKey, interval)
						if err := global.GVA_REDIS.Set(ctx, redisKey, jsonData, 0).Err(); err != nil {
							global.GVA_LOG.Error("更新Redis K线数据失败",
								zap.String("symbol", symbol),
								zap.String("interval", interval),
								zap.Error(err))
						}

						// 如果是日线数据，更新最新收盘价
						if interval == "1m" {
							if err := global.GVA_REDIS.Set(ctx, baseKey, lastKline.C, 0).Err(); err != nil {
								global.GVA_LOG.Error("更新Redis最新价格失败",
									zap.String("symbol", symbol),
									zap.Error(err))
							}
						}
					} else {
						global.GVA_LOG.Error("序列化K线数据失败",
							zap.String("symbol", symbol),
							zap.String("interval", interval),
							zap.Error(err))
					}
				} else {
					global.GVA_LOG.Warn("未获取到K线数据",
						zap.String("symbol", symbol),
						zap.String("interval", interval),
						zap.Time("start_time", startTime),
						zap.Time("end_time", now))
					continue
				}
			}
		}
	}

	wg.Wait() // 等待所有goroutine完成

	return nil
}
