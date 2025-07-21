package task

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	symbolService "github.com/flipped-aurora/gin-vue-admin/server/service/symbol"
	"go.uber.org/zap"
)

// UpdateSymbolsKlineHistory 更新所有symbols的历史K线数据
func UpdateSymbolsKlineHistory() error {
	defer func() {
		if err := recover(); err != nil {
			global.GVA_LOG.Error(fmt.Sprintf("Recover from UpdateSymbolsKlineHistory panic: %v", err))
		}
	}()
	symbolsService := &symbolService.SymbolsService{}

	symbols, err := symbolsService.GetAllSymbols()
	if err != nil {
		global.GVA_LOG.Error("获取symbols失败", zap.Error(err))
		return err
	}

	// 直接按类型分组所有symbols，不再使用批次
	symbolsByType := make(map[int][]string)
	for _, s := range symbols {
		if s.Type != nil {
			symbolsByType[*s.Type] = append(symbolsByType[*s.Type], s.Symbol)
		}
	}

	// 记录处理的symbols数量

	// 定义需要获取的时间周期
	intervals := []string{"1m", "1h", "4h", "1d", "1w", "1M"}

	client := &http.Client{
		Timeout: time.Second * 30,
		Transport: &http.Transport{
			// 禁用 HTTP/2
			ForceAttemptHTTP2: false,
			// 设置连接池
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
			IdleConnTimeout:     90 * time.Second,
			// 启用 TCP keepalive
			DisableKeepAlives: false,
		},
	}

	ctx := context.Background()

	// 第一层循环：按资产类型处理
	for typeID, symbols := range symbolsByType {

		// 第二层循环：按时间周期处理
		for _, interval := range intervals {
			for _, symbol := range symbols {

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
					startTime = now.AddDate(0, 0, -3) // 获取最近1天的分钟数据
				case "1h":
					startTime = now.AddDate(0, -1, 0) // 获取最近1个月的小时数据
				case "4h":
					startTime = now.AddDate(-2, 0, 0)
				case "1d":
					startTime = now.AddDate(0, 0, -1000)
				case "1w":
					startTime = now.AddDate(-5, 0, 0)
				case "1M":
					startTime = now.AddDate(-5, 0, 0)
				}

				// 修改URL，确保limit参数正确
				url := fmt.Sprintf("%s/v2/aggs/ticker/%s/range/%d/%s/%d/%d?apiKey=%s&limit=5000&sort=desc&adjusted=true", // 增加limit到50000
					global.GVA_CONFIG.Polygon.BaseURL,
					ticker,
					multiplier,
					timespan,
					startTime.UnixMilli(),
					now.UnixMilli(),
					global.GVA_CONFIG.Polygon.APIKey)

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

				if len(result.Results) == 0 {
					global.GVA_LOG.Warn("未获取到数据",
						zap.Int("typeID", typeID),
						zap.String("symbol", symbol),
						zap.String("interval", interval))
					continue
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
				historyKey := fmt.Sprintf("%s:kline:%s:history", baseKey, interval)

				// 将所有K线数据序列化
				var klineDataList []map[string]interface{}
				for _, kline := range result.Results {
					klineData := map[string]interface{}{
						"o": kline.O,
						"h": kline.H,
						"l": kline.L,
						"c": kline.C,
						"v": kline.V,
						"t": kline.T,
					}
					klineDataList = append(klineDataList, klineData)
				}

				// 序列化K线数据列表
				if jsonData, err := json.Marshal(klineDataList); err == nil {
					// 直接使用 Set 更新数据
					if err := global.GVA_REDIS.Set(ctx, historyKey, jsonData, 0).Err(); err != nil {
						global.GVA_LOG.Error("更新Redis历史K线数据失败",
							zap.String("symbol", symbol),
							zap.String("interval", interval),
							zap.Error(err))
					}
				} else {
					global.GVA_LOG.Error("序列化K线数据失败",
						zap.String("symbol", symbol),
						zap.String("interval", interval),
						zap.Error(err))
				}

				// 避免触发API限制
				time.Sleep(time.Second * 1)
			}
		}
	}

	return nil
}
