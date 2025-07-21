package initialize

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/gorilla/websocket"
)

// WSClient Binance WebSocket 客户端
type WSClient struct {
	conn      *websocket.Conn
	symbols   []string
	prices    sync.Map
	isRunning bool
	mu        sync.RWMutex
}

// PriceData 价格数据结构
type PriceData struct {
	Symbol    string    `json:"symbol"`
	Price     float64   `json:"price"`
	Timestamp time.Time `json:"timestamp"`
}

// WSResponse Binance WebSocket 响应结构
type WSResponse struct {
	Stream string    `json:"stream"`
	Data   TradeData `json:"data"`
}

// TradeData 交易数据结构
type TradeData struct {
	EventType string `json:"e"` // Event type
	EventTime int64  `json:"E"` // Event time
	Symbol    string `json:"s"` // Symbol
	Price     string `json:"p"` // Price
	Quantity  string `json:"q"` // Quantity
	IsBuyer   bool   `json:"m"` // Is buyer maker?
	TradeID   int64  `json:"t"` // Trade ID
}

var (
	GlobalWSClient *WSClient
	once           sync.Once
)

// InitBinanceWS 初始化 Binance WebSocket 连接
func InitBinanceWS() error {
	var initErr error
	once.Do(func() {
		// TODO: 这里后续会从数据库获取 symbols
		symbols := []string{"btcusdt", "ethusdt", "solusdt"}

		GlobalWSClient = &WSClient{
			symbols:   symbols,
			isRunning: false,
		}
		initErr = GlobalWSClient.connect()
	})
	return initErr
}

// storePriceToRedis 将价格数据存储到Redis
func (w *WSClient) storePriceToRedis(priceData PriceData) error {
	// 构造 Redis key
	key := fmt.Sprintf("symbol:crypto:%s", strings.ToLower(priceData.Symbol))

	// 将价格数据转换为JSON
	jsonData, err := json.Marshal(priceData)
	if err != nil {
		return fmt.Errorf("failed to marshal price data: %v", err)
	}

	// 使用 context.Background()
	ctx := context.Background()
	err = global.GVA_REDIS.Set(ctx, key, string(jsonData), time.Hour).Err()
	if err != nil {
		return fmt.Errorf("failed to store price in redis: %v", err)
	}

	return nil
}

// handleMessages 处理 WebSocket 消息
func (w *WSClient) handleMessages() {
	defer func() {
		w.mu.Lock()
		w.isRunning = false
		w.mu.Unlock()
		if w.conn != nil {
			w.conn.Close()
		}
	}()

	for {
		_, message, err := w.conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			return
		}

		var response WSResponse
		if err := json.Unmarshal(message, &response); err != nil {
			log.Printf("Error unmarshaling message: %v", err)
			continue
		}

		// 解析价格数据
		price, err := strconv.ParseFloat(response.Data.Price, 64)
		if err != nil {
			log.Printf("Error parsing price: %v", err)
			continue
		}

		priceData := PriceData{
			Symbol:    response.Data.Symbol,
			Price:     price,
			Timestamp: time.Unix(0, response.Data.EventTime*int64(time.Millisecond)),
		}

		// 存储到内存
		w.prices.Store(priceData.Symbol, priceData)

		// 存储到Redis
		if err := w.storePriceToRedis(priceData); err != nil {
			log.Printf("Error storing price to Redis: %v", err)
		}
	}
}

// connect 连接到 Binance WebSocket
func (w *WSClient) connect() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.isRunning {
		return fmt.Errorf("websocket client is already running")
	}

	// 构建订阅参数
	streams := make([]string, len(w.symbols))
	for i, symbol := range w.symbols {
		// 使用小写并确保格式正确
		symbol = strings.ToLower(strings.TrimSuffix(symbol, "usdt")) + "usdt"
		streams[i] = fmt.Sprintf("%s@trade", symbol)
		w.symbols[i] = strings.ToUpper(symbol) // 存储大写格式用于后续匹配
	}

	// 构建 WebSocket URL
	url := fmt.Sprintf("wss://stream.binance.com:9443/stream?streams=%s",
		strings.Join(streams, "/"))

	// 连接 WebSocket
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return fmt.Errorf("failed to connect to Binance WebSocket: %v", err)
	}

	w.conn = conn
	w.isRunning = true

	// 启动消息处理
	go w.handleMessages()

	log.Printf("Successfully connected to Binance WebSocket for symbols: %v", w.symbols)
	return nil
}

// GetPrice 获取指定交易对的最新价格
func (w *WSClient) GetPrice(symbol string) (PriceData, error) {
	symbol = strings.ToUpper(symbol)
	if value, ok := w.prices.Load(symbol); ok {
		return value.(PriceData), nil
	}
	return PriceData{}, fmt.Errorf("no price data available for symbol: %s", symbol)
}

// GetAllPrices 获取所有交易对的最新价格
func (w *WSClient) GetAllPrices() map[string]PriceData {
	prices := make(map[string]PriceData)
	w.prices.Range(func(key, value interface{}) bool {
		prices[key.(string)] = value.(PriceData)
		return true
	})
	return prices
}

// IsConnected 检查是否已连接
func (w *WSClient) IsConnected() bool {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.isRunning
}

// Close 关闭 WebSocket 连接
func (w *WSClient) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if !w.isRunning {
		return nil
	}

	if w.conn != nil {
		err := w.conn.Close()
		w.isRunning = false
		return err
	}
	return nil
}
