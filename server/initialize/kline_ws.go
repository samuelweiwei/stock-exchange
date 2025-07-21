package initialize

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/service/symbol"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

// KlineData 定义K线数据结构
type KlineData struct {
	Symbol    string    `json:"symbol"`
	Open      float64   `json:"o"`
	High      float64   `json:"h"`
	Low       float64   `json:"l"`
	Close     float64   `json:"c"`
	Volume    float64   `json:"v"`
	Timestamp time.Time `json:"t"`
	Interval  string    `json:"interval"`
}

// 添加新的消息结构体
type WebSocketMessage struct {
	EventType string `json:"ev"`
	Status    string `json:"status"`
	Message   string `json:"message"`
}

type WebSocketResponse []struct {
	EventType string `json:"ev"`
	Status    string `json:"status"`
	Message   string `json:"message"`
	// K线数据字段
	Symbol    string    `json:"symbol,omitempty"`
	Open      float64   `json:"o,omitempty"`
	High      float64   `json:"h,omitempty"`
	Low       float64   `json:"l,omitempty"`
	Close     float64   `json:"c,omitempty"`
	Volume    float64   `json:"v,omitempty"`
	Timestamp time.Time `json:"t,omitempty"`
}

// formatSymbol 根据不同类型格式化订阅symbol
func formatSymbol(symbol string, symbolType *int) string {
	if symbolType == nil {
		return symbol
	}

	switch *symbolType {
	case 0: // 股票
		return symbol // 例如: AAPL
	case 1: // 加密货币
		return strings.Replace(symbol, "/", "-", -1) // 例如: BTC-USD
	case 2: // 外汇
		return strings.Replace(symbol, "/", "", -1) // 例如: EURUSD
	default:
		return symbol
	}
}

// 添加重连和连接管理逻辑
type WebSocketClient struct {
	conn      *websocket.Conn
	apiKey    string
	host      string
	path      string
	symbols   map[string]string
	intervals []string
}

func NewWebSocketClient(host, path string, symbols map[string]string, intervals []string) *WebSocketClient {
	return &WebSocketClient{
		apiKey:    global.GVA_CONFIG.Polygon.APIKey,
		host:      host,
		path:      path,
		symbols:   symbols,
		intervals: intervals,
	}
}

func (c *WebSocketClient) connect() error {
	u := url.URL{Scheme: "wss", Host: c.host, Path: c.path}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return fmt.Errorf("websocket connection failed: %v", err)
	}
	c.conn = conn

	// 认证
	authMsg := fmt.Sprintf(`{"action":"auth","params":"%s"}`, c.apiKey)
	if err := c.conn.WriteMessage(websocket.TextMessage, []byte(authMsg)); err != nil {
		return fmt.Errorf("authentication failed: %v", err)
	}

	return nil
}

func (c *WebSocketClient) subscribe() error {
	var prefix string
	switch c.path {
	case "/stocks":
		prefix = "A"
	case "/crypto":
		prefix = "XA"
	case "/forex":
		prefix = "CA"
	}

	// 将所有订阅合并到一个字符串
	var allSubscriptions []string
	for _, symbol := range c.symbols {
		for _, interval := range c.intervals {
			allSubscriptions = append(allSubscriptions, fmt.Sprintf("%s.%s.%s", prefix, symbol, interval))
		}
	}

	// 一次性订阅所有
	subscribeMsg := fmt.Sprintf(`{"action":"subscribe","params":["%s"]}`, strings.Join(allSubscriptions, `","`))

	global.GVA_LOG.Info("Subscribing to symbols",
		zap.String("path", c.path),
		zap.String("subscriptions", subscribeMsg))

	return c.conn.WriteMessage(websocket.TextMessage, []byte(subscribeMsg))
}

func (c *WebSocketClient) handleMessages(ctx context.Context, marketType string) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			_, message, err := c.conn.ReadMessage()
			if err != nil {
				global.GVA_LOG.Error("Error reading message",
					zap.String("market", marketType),
					zap.Error(err))

				// 尝试重连
				time.Sleep(5 * time.Second)
				if err := c.connect(); err == nil {
					if err := c.subscribe(); err == nil {
						continue
					}
				}
				return
			}

			var response WebSocketResponse
			if err := json.Unmarshal(message, &response); err != nil {
				global.GVA_LOG.Error("Failed to parse message",
					zap.String("market", marketType),
					zap.Error(err),
					zap.String("raw_message", string(message)))
				continue
			}

			for _, msg := range response {
				if msg.EventType == "status" {
					global.GVA_LOG.Info("Status message",
						zap.String("market", marketType),
						zap.String("status", msg.Status),
						zap.String("message", msg.Message))

					// 处理最大连接数错误
					if msg.Status == "max_connections" {
						time.Sleep(30 * time.Second) // 等待较长时间后重试
						if err := c.connect(); err == nil {
							if err := c.subscribe(); err == nil {
								continue
							}
						}
						return
					}
					continue
				}

				// 处理K线数据
				if msg.Symbol != "" {
					klineData := KlineData{
						Symbol:    msg.Symbol,
						Open:      msg.Open,
						High:      msg.High,
						Low:       msg.Low,
						Close:     msg.Close,
						Volume:    msg.Volume,
						Timestamp: msg.Timestamp,
					}

					// 存储到Redis
					key := fmt.Sprintf("kline:%s:%s:%s", marketType, klineData.Interval, klineData.Symbol)
					jsonData, _ := json.Marshal(klineData)
					if err := global.GVA_REDIS.Set(ctx, key, jsonData, 24*time.Hour).Err(); err != nil {
						global.GVA_LOG.Error("Failed to store in Redis",
							zap.String("market", marketType),
							zap.Error(err),
							zap.String("key", key))
					}
				}
			}
		}
	}
}

func handleStockWebSocket(ctx context.Context, symbols map[string]string, intervals []string) {
	client := NewWebSocketClient("delayed.polygon.io", "/stocks", symbols, intervals)
	if err := client.connect(); err != nil {
		global.GVA_LOG.Error("Failed to connect to stock websocket", zap.Error(err))
		return
	}

	if err := client.subscribe(); err != nil {
		global.GVA_LOG.Error("Failed to subscribe to stock symbols", zap.Error(err))
		return
	}

	client.handleMessages(ctx, "stock")
}

func handleCryptoWebSocket(ctx context.Context, symbols map[string]string, intervals []string) {
	client := NewWebSocketClient("socket.polygon.io", "/crypto", symbols, intervals)
	if err := client.connect(); err != nil {
		global.GVA_LOG.Error("Failed to connect to crypto websocket", zap.Error(err))
		return
	}

	if err := client.subscribe(); err != nil {
		global.GVA_LOG.Error("Failed to subscribe to crypto symbols", zap.Error(err))
		return
	}

	client.handleMessages(ctx, "crypto")
}

func handleForexWebSocket(ctx context.Context, symbols map[string]string, intervals []string) {
	client := NewWebSocketClient("socket.polygon.io", "/forex", symbols, intervals)
	if err := client.connect(); err != nil {
		global.GVA_LOG.Error("Failed to connect to forex websocket", zap.Error(err))
		return
	}

	if err := client.subscribe(); err != nil {
		global.GVA_LOG.Error("Failed to subscribe to forex symbols", zap.Error(err))
		return
	}

	client.handleMessages(ctx, "forex")
}

func InitKlineWs(ctx context.Context) {
	global.GVA_LOG.Info("Initializing Kline WebSocket")

	// 获取所有交易对
	var symbolService symbol.SymbolsService
	symbols, err := symbolService.GetAllSymbols()
	if err != nil {
		global.GVA_LOG.Error("Failed to get symbols", zap.Error(err))
		return
	}

	// 按类型分类symbols
	stockSymbols := make(map[string]string)
	cryptoSymbols := make(map[string]string)
	forexSymbols := make(map[string]string)

	for _, symbol := range symbols {
		formattedSymbol := formatSymbol(symbol.Symbol, symbol.Type)
		if symbol.Type == nil {
			continue
		}

		switch *symbol.Type {
		case 0:
			stockSymbols[symbol.Symbol] = formattedSymbol
		case 1:
			cryptoSymbols[symbol.Symbol] = formattedSymbol
		case 2:
			forexSymbols[symbol.Symbol] = formattedSymbol
		}
	}

	global.GVA_LOG.Info("Symbols categorized",
		zap.Int("stocks", len(stockSymbols)),
		zap.Int("crypto", len(cryptoSymbols)),
		zap.Int("forex", len(forexSymbols)))

	// 定义需要订阅的时间间隔
	intervals := []string{"240", "1D", "1W", "1M"}

	// 启动不同类型的WebSocket连接
	if len(stockSymbols) > 0 {
		go handleStockWebSocket(ctx, stockSymbols, intervals)
	}
	if len(cryptoSymbols) > 0 {
		go handleCryptoWebSocket(ctx, cryptoSymbols, intervals)
	}
	if len(forexSymbols) > 0 {
		go handleForexWebSocket(ctx, forexSymbols, intervals)
	}

	<-ctx.Done()
}
