package initialize

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

// InitSymbolsPrices 在服务启动时初始化所有symbols的前一天收盘价到Redis
func InitSymbolsPrices() error {
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

	client := &http.Client{
		Timeout: time.Second * 30,
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

			// 为每个symbol单独请求数据
			pipe := global.GVA_REDIS.Pipeline()
			ctx := context.Background()

			for _, symbol := range currentBatch {
				// 构建ticker
				var ticker string
				switch typeID {
				case 0: // 股票
					ticker = symbol.Symbol
				case 1: // 加密货币
					ticker = "X:" + strings.ReplaceAll(symbol.Symbol, "/", "")
				case 2: // 外汇
					ticker = "C:" + strings.ReplaceAll(symbol.Symbol, "/", "")
				}

				// 只获取前一天的数据
				yesterday := time.Now().AddDate(0, 0, -1)
				dailyUrl := fmt.Sprintf("%s/v2/aggs/ticker/%s/range/1/day/%s/%s?apiKey=%s",
					global.GVA_CONFIG.Polygon.BaseURL,
					ticker,
					yesterday.Format("2006-01-02"),
					yesterday.Format("2006-01-02"),
					global.GVA_CONFIG.Polygon.APIKey)

				dailyResp, err := client.Get(dailyUrl)
				if err != nil {
					global.GVA_LOG.Error("请求日交易数据失败",
						zap.String("symbol", symbol.Symbol),
						zap.Error(err))
					continue
				}

				var dailyData struct {
					Results []struct {
						O float64 `json:"o"`
						C float64 `json:"c"` // 收盘价
						T int64   `json:"t"` // 时间戳
					} `json:"results"`
					Status    string `json:"status"`
					RequestID string `json:"request_id"`
				}

				if err := json.NewDecoder(dailyResp.Body).Decode(&dailyData); err != nil {
					dailyResp.Body.Close()
					global.GVA_LOG.Error("解析日交易数据失败",
						zap.String("symbol", symbol.Symbol),
						zap.Error(err))
					continue
				}
				dailyResp.Body.Close()

				// 如果有数据，更新到Redis
				if len(dailyData.Results) > 0 {
					// 构建Redis key
					var redisKey string
					switch typeID {
					case 0:
						redisKey = fmt.Sprintf("symbol:stock:%s:previous_close", symbol.Symbol)
					case 1:
						redisKey = fmt.Sprintf("symbol:crypto:%s:previous_close", symbol.Symbol)
					case 2:
						redisKey = fmt.Sprintf("symbol:forex:%s:previous_close", symbol.Symbol)
					}

					// 设置前一天收盘价（使用最后一个数据点）
					lastIndex := len(dailyData.Results) - 1
					pipe.Set(ctx, redisKey, dailyData.Results[lastIndex].C, 0)
				}

				// 避免触发API限制
				time.Sleep(time.Second * 1)
			}

			// 执行管道命令
			if _, err := pipe.Exec(ctx); err != nil {
				global.GVA_LOG.Error("批量更新Redis失败", zap.Error(err))
			}
		}
	}

	return nil
}

// InitSymbolsPricesData 初始化所有symbols的前一天收盘价到Redis
func InitSymbolsPricesData() {
	defer func() {
		if err := recover(); err != nil {
			global.GVA_LOG.Error(fmt.Sprintf("Recover from InitSymbolsPricesData Panic: %v", err))
		}
	}()
	// 等待Redis初始化完成
	time.Sleep(2 * time.Second)

	if err := InitSymbolsPrices(); err != nil {
		global.GVA_LOG.Error("初始化symbols前一天收盘价失败", zap.Error(err))
	}
}
