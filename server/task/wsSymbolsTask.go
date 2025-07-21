package task

import (
	"context"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	polygonws "github.com/polygon-io/client-go/websocket"
	"go.uber.org/zap"
)

// WebSocket订阅的批次大小
const (
	wsSubscribeBatchSize = 100 // 每批订阅的symbol数量
	wsSubscribeDelay     = 100 // 批次间延迟(毫秒)
)

// UpdateWebSocketSubscriptions 更新WebSocket订阅列表
func UpdateWebSocketSubscriptions(ctx context.Context) error {
	// 获取数据库中的所有symbols
	symbolsService := service.ServiceGroupApp.SymbolServiceGroup.SymbolsService
	symbols, err := symbolsService.GetAllSymbols()
	if err != nil {
		global.GVA_LOG.Error("Failed to get symbols from database", zap.Error(err))
		return err
	}

	// 分类symbols
	stockSymbols := make(map[string]bool)
	cryptoSymbols := make(map[string]bool)
	forexSymbols := make(map[string]bool)

	for _, s := range symbols {
		if s.Type == nil || s.Status != nil && *s.Status == 0 {
			continue // 跳过无类型或未激活的symbol
		}

		switch *s.Type {
		case 0: // 股票
			stockSymbols[s.Symbol] = true
		case 1: // 加密货币
			// 正确格式化加密货币 symbol
			formattedSymbol := "X:" + strings.Replace(s.Symbol, "/", "-", -1)
			cryptoSymbols[formattedSymbol] = true
		case 2: // 外汇
			// 正确格式化外汇 symbol
			formattedSymbol := "C:" + strings.Replace(s.Symbol, "/", "-", -1)
			forexSymbols[formattedSymbol] = true
		}
	}

	// 获取当前WebSocket客户端的订阅列表
	global.GVA_WS_MUTEX.Lock()
	defer global.GVA_WS_MUTEX.Unlock()

	// 更新股票订阅
	if global.GVA_WS_STOCK != nil {
		updateStockSubscriptions(global.GVA_WS_STOCK, stockSymbols)
	}

	// 更新加密货币订阅
	if global.GVA_WS_CRYPTO != nil {
		updateCryptoSubscriptions(global.GVA_WS_CRYPTO, cryptoSymbols)
	}

	// 更新外汇订阅
	if global.GVA_WS_FOREX != nil {
		updateForexSubscriptions(global.GVA_WS_FOREX, forexSymbols)
	}

	return nil
}

// updateStockSubscriptions 更新股票订阅
func updateStockSubscriptions(client *polygonws.Client, newSymbols map[string]bool) {
	// 先取消所有订阅
	if err := client.Unsubscribe(polygonws.StocksSecAggs); err != nil {
		global.GVA_LOG.Error("Failed to unsubscribe stock aggregates", zap.Error(err))
	}
	if err := client.Unsubscribe(polygonws.StocksQuotes); err != nil {
		global.GVA_LOG.Error("Failed to unsubscribe stock quotes", zap.Error(err))
	}

	// 收集所有symbols
	symbols := make([]string, 0, len(newSymbols))
	for symbol := range newSymbols {
		symbols = append(symbols, symbol)
	}

	// 分批订阅
	for i := 0; i < len(symbols); i += wsSubscribeBatchSize {
		end := i + wsSubscribeBatchSize
		if end > len(symbols) {
			end = len(symbols)
		}
		batch := symbols[i:end]

		// 订阅聚合数据
		if err := client.Subscribe(polygonws.StocksSecAggs, batch...); err != nil {
			global.GVA_LOG.Error("Failed to subscribe stock aggregates",
				zap.Strings("symbols", batch),
				zap.Error(err))
		}

		// 订阅报价数据
		if err := client.Subscribe(polygonws.StocksQuotes, batch...); err != nil {
			global.GVA_LOG.Error("Failed to subscribe stock quotes",
				zap.Strings("symbols", batch),
				zap.Error(err))
		}

		// 添加小延迟，避免请求过快
		time.Sleep(time.Millisecond * wsSubscribeDelay)
	}

	global.GVA_LOG.Info("Stock subscriptions updated",
		zap.Int("total_symbols", len(symbols)))
}

// updateCryptoSubscriptions 更新加密货币订阅
func updateCryptoSubscriptions(client *polygonws.Client, newSymbols map[string]bool) {
	// 收集所有symbols
	symbols := make([]string, 0, len(newSymbols))
	for symbol := range newSymbols {
		// 不需要在这里添加 "X:" 前缀，因为 symbol 已经在前面格式化过了
		symbols = append(symbols, symbol)
	}

	// 分批订阅
	for i := 0; i < len(symbols); i += wsSubscribeBatchSize {
		end := i + wsSubscribeBatchSize
		if end > len(symbols) {
			end = len(symbols)
		}
		batch := symbols[i:end]

		// 订阅聚合数据
		if err := client.Subscribe(polygonws.CryptoSecAggs, batch...); err != nil {
			global.GVA_LOG.Error("Failed to subscribe crypto aggregates",
				zap.Strings("symbols", batch),
				zap.Error(err))
		}
		// 订阅报价数据
		if err := client.Subscribe(polygonws.CryptoQuotes, batch...); err != nil {
			global.GVA_LOG.Error("Failed to subscribe crypto quotes",
				zap.Strings("symbols", batch),
				zap.Error(err))
		}

		// 添加小延迟，避免请求过快
		time.Sleep(time.Millisecond * wsSubscribeDelay)
	}

	global.GVA_LOG.Info("Crypto subscriptions updated",
		zap.Int("total_symbols", len(symbols)))
}

// updateForexSubscriptions 更新外汇订阅
func updateForexSubscriptions(client *polygonws.Client, newSymbols map[string]bool) {
	// 收集所有symbols
	symbols := make([]string, 0, len(newSymbols))
	for symbol := range newSymbols {
		// 不需要在这里添加 "C:" 前缀，因为 symbol 已经在前面格式化过了
		symbols = append(symbols, symbol)
	}

	// 分批订阅
	for i := 0; i < len(symbols); i += wsSubscribeBatchSize {
		end := i + wsSubscribeBatchSize
		if end > len(symbols) {
			end = len(symbols)
		}
		batch := symbols[i:end]

		// 订阅聚合数据
		if err := client.Subscribe(polygonws.ForexSecAggs, batch...); err != nil {
			global.GVA_LOG.Error("Failed to subscribe forex aggregates",
				zap.Strings("symbols", batch),
				zap.Error(err))
		}
		// 订阅报价数据
		if err := client.Subscribe(polygonws.ForexQuotes, batch...); err != nil {
			global.GVA_LOG.Error("Failed to subscribe forex quotes",
				zap.Strings("symbols", batch),
				zap.Error(err))
		}

		// 添加小延迟，避免请求过快
		time.Sleep(time.Millisecond * wsSubscribeDelay)
	}

	global.GVA_LOG.Info("Forex subscriptions updated",
		zap.Int("total_symbols", len(symbols)))
}

// contains 检查字符串是否在切片中
func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}
