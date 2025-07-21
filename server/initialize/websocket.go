package initialize

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/service/symbol"
	polygonws "github.com/polygon-io/client-go/websocket"
	"github.com/polygon-io/client-go/websocket/models"
	"go.uber.org/zap"
)

// wsInitialized 标记 WebSocket 是否已初始化
var wsInitialized bool

// WebsocketInit 初始化 WebSocket 连接
func WebsocketInit(ctx context.Context) error {
	// 验证上下文
	if ctx == nil {
		global.GVA_LOG.Error("Context is nil")
		return errors.New("context is nil")
	}

	// 验证 API 密钥
	apiKey := global.GVA_CONFIG.Polygon.APIKey
	if apiKey == "" {
		global.GVA_LOG.Error("Polygon API key not configured")
		return errors.New("polygon API key not configured")
	}

	// 从 symbol service 获取需要订阅的交易对
	var symbolService symbol.SymbolsService
	symbols, err := symbolService.GetAllSymbols()
	if err != nil {
		global.GVA_LOG.Error("Failed to get symbols", zap.Error(err))
		return err
	}

	// 根据类型分类交易对
	var stockSymbols, forexSymbols, cryptoSymbols []string
	for _, s := range symbols {
		if s.Type == nil {
			global.GVA_LOG.Warn("Symbol type is nil", zap.String("symbol", s.Symbol))
			continue
		}

		switch *s.Type {
		case 1: // 加密货币
			// 将 BTC/USD 转换为 X:BTC-USD
			formattedSymbol := strings.Replace(s.Symbol, "/", "-", -1)
			cryptoSymbols = append(cryptoSymbols, "X:"+formattedSymbol)
		case 2: // 外汇
			// 将 EUR/USD 转换为 C:EUR-USD
			formattedSymbol := strings.Replace(s.Symbol, "/", "-", -1)
			forexSymbols = append(forexSymbols, "C:"+formattedSymbol)
		case 0: // 股票
			stockSymbols = append(stockSymbols, s.Symbol)
		default:
			global.GVA_LOG.Warn("Unknown symbol type",
				zap.String("symbol", s.Symbol),
				zap.Int("type", int(*s.Type)))
		}
	}

	// 启动各类 WebSocket 客户端
	if len(stockSymbols) > 0 {
		go func() {
			if err := startStockWebSocket(ctx, apiKey, stockSymbols); err != nil {
				global.GVA_LOG.Error("Stock WebSocket error", zap.Error(err))
			}
		}()
	}

	if len(cryptoSymbols) > 0 {
		go func() {
			if err := startCryptoWebSocket(ctx, apiKey, cryptoSymbols); err != nil {
				global.GVA_LOG.Error("Crypto WebSocket error", zap.Error(err))
			}
		}()
	}

	if len(forexSymbols) > 0 {
		go func() {
			if err := startForexWebSocket(ctx, apiKey, forexSymbols); err != nil {
				global.GVA_LOG.Error("Forex WebSocket error", zap.Error(err))
			}
		}()
	}

	wsInitialized = true
	global.GVA_LOG.Info("WebSocket initialization completed")
	return nil
}

// startStockWebSocket 启动股票市场的 WebSocket 连接
func startStockWebSocket(ctx context.Context, apiKey string, symbols []string) error {
	client, err := polygonws.New(polygonws.Config{
		APIKey: apiKey,
		Feed:   polygonws.Delayed,
		Market: polygonws.Stocks,
	})
	if err != nil {
		return err
	}

	global.GVA_WS_MUTEX.Lock()
	global.GVA_WS_STOCK = client
	global.GVA_WS_MUTEX.Unlock()

	// 建立连接
	if err := client.Connect(); err != nil {
		global.GVA_LOG.Error("Failed to stock connect", zap.Error(err))
		return err
	}

	// 订阅交易数据
	if err := client.Subscribe(polygonws.StocksSecAggs, symbols...); err != nil {
		global.GVA_LOG.Error("Failed to subscribe to trades", zap.Error(err))
		return err
	}

	// 订阅报价数据
	if err := client.Subscribe(polygonws.StocksQuotes, symbols...); err != nil {
		global.GVA_LOG.Error("Failed to subscribe to quotes", zap.Error(err))
		return err
	}

	// 启动消息处理
	go func() {
		defer func() {
			client.Close()
			if err := recover(); err != nil {
				global.GVA_LOG.Error(fmt.Sprintf("Recover Stock websocket error %v", err))
			}
		}()

		for {
			select {
			case <-ctx.Done():
				return
			case err, ok := <-client.Error():
				if !ok {
					return
				}
				global.GVA_LOG.Error("Stock websocket error", zap.Error(err))
			case out, more := <-client.Output():
				if !more {
					return
				}

				// 处理聚合数据并更新 Redis
				switch data := out.(type) {
				case models.EquityAgg:
					// global.GVA_LOG.Info("Stock Aggregates",
					// 	zap.String("symbol", data.Symbol),
					// 	zap.Float64("close", data.Close))

					redisKey := fmt.Sprintf("symbol:stock:%s", data.Symbol)
					priceStr := fmt.Sprintf("%v", data.Close)
					if err := global.GVA_REDIS.Set(ctx, redisKey, priceStr, 0).Err(); err != nil {
						global.GVA_LOG.Error("Redis error", zap.Error(err))
					}
				}
			}
		}
	}()

	return nil
}

// startCryptoWebSocket 启动加密货币市场的 WebSocket 连接
func startCryptoWebSocket(ctx context.Context, apiKey string, symbols []string) error {
	client, err := polygonws.New(polygonws.Config{
		APIKey: apiKey,
		Feed:   polygonws.RealTime,
		Market: polygonws.Crypto,
	})
	if err != nil {
		return err
	}

	global.GVA_WS_MUTEX.Lock()
	global.GVA_WS_CRYPTO = client
	global.GVA_WS_MUTEX.Unlock()

	// 建立连接
	if err := client.Connect(); err != nil {
		global.GVA_LOG.Error("Failed to connect crypto client", zap.Error(err))
		return err
	}

	// 订阅聚合数据
	if err := client.Subscribe(polygonws.CryptoSecAggs, symbols...); err != nil {
		global.GVA_LOG.Error("Failed to subscribe to crypto trades", zap.Error(err))
		return err
	}
	// 订阅报价数据
	if err := client.Subscribe(polygonws.CryptoQuotes, symbols...); err != nil {
		global.GVA_LOG.Error("Failed to subscribe to quotes", zap.Error(err))
		return err
	}
	// 启动消息处理
	go func() {
		defer func() {
			client.Close()
			if err := recover(); err != nil {
				global.GVA_LOG.Error(fmt.Sprintf("Recover Crypto websocket error %v", err))
			}
		}()

		for {
			select {
			case <-ctx.Done():
				return
			case err, ok := <-client.Error():
				if !ok {
					return
				}
				global.GVA_LOG.Error("Crypto websocket error", zap.Error(err))
			case out, more := <-client.Output():
				if !more {
					return
				}

				// 处理聚合数据并更新 Redis
				switch data := out.(type) {
				case models.CurrencyAgg:
					// 从 X:BTC-USD 转回 BTC/USD 格式
					cleanPair := strings.TrimPrefix(data.Pair, "X:")
					cleanPair = strings.Replace(cleanPair, "-", "/", -1)
					// global.GVA_LOG.Info("Crypto Aggregates",
					// 	zap.String("pair", cleanPair),
					// 	zap.Float64("close", data.Close))

					redisKey := fmt.Sprintf("symbol:crypto:%s", cleanPair)
					priceStr := fmt.Sprintf("%v", data.Close)
					if err := global.GVA_REDIS.Set(ctx, redisKey, priceStr, 0).Err(); err != nil {
						global.GVA_LOG.Error("Redis error", zap.Error(err))
					}
				}
			}
		}
	}()

	return nil
}

// startForexWebSocket 启动外汇市场的 WebSocket 连接
func startForexWebSocket(ctx context.Context, apiKey string, symbols []string) error {
	client, err := polygonws.New(polygonws.Config{
		APIKey: apiKey,
		Feed:   polygonws.RealTime,
		Market: polygonws.Forex,
	})
	if err != nil {
		return err
	}

	global.GVA_WS_MUTEX.Lock()
	global.GVA_WS_FOREX = client
	global.GVA_WS_MUTEX.Unlock()

	// 建立连接
	if err := client.Connect(); err != nil {
		global.GVA_LOG.Error("Failed to Forex connect", zap.Error(err))
		return err
	}

	// 订阅聚合数据
	if err := client.Subscribe(polygonws.ForexSecAggs, symbols...); err != nil {
		global.GVA_LOG.Error("Failed to subscribe to ForexSecAggs", zap.Error(err))
		return err
	}

	// 订阅报价数据
	if err := client.Subscribe(polygonws.ForexQuotes, symbols...); err != nil {
		global.GVA_LOG.Error("Failed to subscribe to Forex quotes", zap.Error(err))
		return err
	}

	// 启动消息处理
	go func() {
		defer func() {
			client.Close()
			if err := recover(); err != nil {
				global.GVA_LOG.Error(fmt.Sprintf("Recover Forex websocket error %v", err))
			}
		}()

		for {
			select {
			case <-ctx.Done():
				return
			case err, ok := <-client.Error():
				if !ok {
					return
				}
				global.GVA_LOG.Error("Forex websocket error", zap.Error(err))
			case out, more := <-client.Output():
				if !more {
					return
				}

				// 处理聚合数据并更新 Redis
				switch data := out.(type) {
				case models.CurrencyAgg:
					// 从 C:EUR-USD 转回 EUR/USD 格式
					cleanPair := strings.TrimPrefix(data.Pair, "C:")
					cleanPair = strings.Replace(cleanPair, "-", "/", -1)
					// global.GVA_LOG.Info("Forex Aggregates",
					// 	zap.String("pair", cleanPair),
					// 	zap.Float64("close", data.Close))

					redisKey := fmt.Sprintf("symbol:forex:%s", cleanPair)
					priceStr := fmt.Sprintf("%v", data.Close)
					if err := global.GVA_REDIS.Set(ctx, redisKey, priceStr, 0).Err(); err != nil {
						global.GVA_LOG.Error("Redis error", zap.Error(err))
					}
				}
			}
		}
	}()

	return nil
}

// IsWebsocketInitialized 检查 WebSocket 是否已初始化
func IsWebsocketInitialized() bool {
	return wsInitialized
}
