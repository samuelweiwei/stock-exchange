package symbol

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/symbol"
	symbolReq "github.com/flipped-aurora/gin-vue-admin/server/model/symbol/request"
	"github.com/gorilla/websocket"

	// polygonws "github.com/polygon-io/client-go/websocket"
	"go.uber.org/zap"
)

type SymbolsService struct{}

// CreateSymbols 创建symbols表记录
// Author [yourname](https://github.com/yourname)
func (symbolsService *SymbolsService) CreateSymbols(symbols *symbol.Symbols) (err error) {
	err = global.GVA_DB.Create(symbols).Error
	return err
}

// DeleteSymbols 删除symbols表记录
// Author [yourname](https://github.com/yourname)
func (symbolsService *SymbolsService) DeleteSymbols(id string) (err error) {
	// 先获取symbol信息，以便后续删除Redis数据
	var symbol symbol.Symbols
	if err = global.GVA_DB.Where("id = ?", id).First(&symbol).Error; err != nil {
		return err
	}

	// 删除数据库记录
	if err = global.GVA_DB.Delete(&symbol, "id = ?", id).Error; err != nil {
		return err
	}

	// 如果有type字段，删除Redis中的相关数据
	if symbol.Type != nil && symbol.Symbol != "" {
		var baseKey string
		switch *symbol.Type {
		case 0: // 股票
			baseKey = fmt.Sprintf("symbol:stock:%s", symbol.Symbol)
		case 1: // 加密货币
			baseKey = fmt.Sprintf("symbol:crypto:%s", symbol.Symbol)
		case 2: // 外汇
			baseKey = fmt.Sprintf("symbol:forex:%s", symbol.Symbol)
		}

		// 使用pipeline批量删除相关的Redis键
		ctx := context.Background()
		pipe := global.GVA_REDIS.Pipeline()

		// 删除价格数据
		pipe.Del(ctx, baseKey)
		// 删除精度设置
		pipe.Del(ctx, baseKey+":precision")
		pipe.Del(ctx, baseKey+":num_precision")
		// 删除前一天收盘价
		pipe.Del(ctx, baseKey+":previous_close")

		// 删除所有时间间隔的K线数据
		intervals := []string{"1m", "5m", "15m", "30m", "1h", "2h", "4h", "1d", "1w", "1M"}
		for _, interval := range intervals {
			klineKey := fmt.Sprintf("%s:kline:%s", baseKey, interval)
			pipe.Del(ctx, klineKey)
		}

		if _, err := pipe.Exec(ctx); err != nil {
			global.GVA_LOG.Error("Failed to delete Redis data",
				zap.String("symbol", symbol.Symbol),
				zap.Error(err))
		}
	}

	return nil
}

// DeleteSymbolsByIds 批量删除symbols表记录
// Author [yourname](https://github.com/yourname)
func (symbolsService *SymbolsService) DeleteSymbolsByIds(ids []string) (err error) {
	// 先获取所有要删除的symbols信息
	var symbols []symbol.Symbols
	if err = global.GVA_DB.Where("id in ?", ids).Find(&symbols).Error; err != nil {
		return err
	}

	// 删除数据库记录
	if err = global.GVA_DB.Delete(&[]symbol.Symbols{}, "id in ?", ids).Error; err != nil {
		return err
	}

	// 删除Redis中的数据
	ctx := context.Background()
	pipe := global.GVA_REDIS.Pipeline()

	for _, symbol := range symbols {
		if symbol.Type != nil && symbol.Symbol != "" {
			var baseKey string
			switch *symbol.Type {
			case 0: // 股票
				baseKey = fmt.Sprintf("symbol:stock:%s", symbol.Symbol)
			case 1: // 加密货币
				baseKey = fmt.Sprintf("symbol:crypto:%s", symbol.Symbol)
			case 2: // 外汇
				baseKey = fmt.Sprintf("symbol:forex:%s", symbol.Symbol)
			}

			// 删除价格数据
			pipe.Del(ctx, baseKey)
			// 删除精度设置
			pipe.Del(ctx, baseKey+":precision")
			pipe.Del(ctx, baseKey+":num_precision")
			// 删除前一天收盘价
			pipe.Del(ctx, baseKey+":previous_close")

			// 删除所有时间间隔的K线数据
			intervals := []string{"1m", "5m", "15m", "30m", "1h", "2h", "4h", "1d", "1w", "1M"}
			for _, interval := range intervals {
				klineKey := fmt.Sprintf("%s:kline:%s", baseKey, interval)
				pipe.Del(ctx, klineKey)
			}
		}
	}

	// 执行Redis删除操作
	if _, err := pipe.Exec(ctx); err != nil {
		global.GVA_LOG.Error("Failed to batch delete Redis data", zap.Error(err))
	}

	return nil
}

// UpdateSymbols 更新symbols表记录
// Author [yourname](https://github.com/yourname)
func (symbolsService *SymbolsService) UpdateSymbols(symbols symbol.Symbols) (err error) {
	// 更新数据库
	err = global.GVA_DB.Model(&symbol.Symbols{}).Where("id = ?", symbols.Id).Updates(&symbols).Error
	if err != nil {
		return err
	}

	// 获取完整的symbol信息
	var updatedSymbol symbol.Symbols
	if err = global.GVA_DB.Where("id = ?", symbols.Id).First(&updatedSymbol).Error; err != nil {
		return err
	}

	// 如果有type字段，更新Redis中的精度信息
	if updatedSymbol.Type != nil {
		if err = setSymbolPrecision(updatedSymbol.Symbol, *updatedSymbol.Type,
			updatedSymbol.TicketSize, updatedSymbol.TicketNumSize); err != nil {
			global.GVA_LOG.Error("Failed to update precision in Redis",
				zap.String("symbol", updatedSymbol.Symbol),
				zap.Error(err))
		}
	}

	return nil
}

// GetSymbols 根据id获取Symbols记录
func (symbolsService *SymbolsService) GetSymbols(id string) (symbols *symbol.Symbols, err error) {
	symbols = &symbol.Symbols{}
	err = global.GVA_DB.Where("id = ?", id).First(symbols).Error
	if err != nil {
		return nil, err
	}

	// 根据类型构建Redis key
	var redisKey string
	if symbols.Type == nil {
		return nil, fmt.Errorf("币种类型未知")
	}

	switch *symbols.Type {
	case 0: // 股票
		redisKey = fmt.Sprintf("symbol:stock:%s", symbols.Symbol)
	case 1: // 加密货币
		redisKey = fmt.Sprintf("symbol:crypto:%s", symbols.Symbol)
	case 2: // 外汇
		redisKey = fmt.Sprintf("symbol:forex:%s", symbols.Symbol)
	default:
		return nil, fmt.Errorf("未知的币种类型: %d", *symbols.Type)
	}

	// 从Redis获取价格
	if price, err := global.GVA_REDIS.Get(context.Background(), redisKey).Result(); err == nil {
		priceFloat, _ := strconv.ParseFloat(price, 64)
		symbols.CurrentPrice = &priceFloat
	} else {
		// 如果Redis中没有价格,保持数据库中的current_price不变
		global.GVA_LOG.Error("从Redis获取价格失败",
			zap.String("symbol", symbols.Symbol),
			zap.Error(err))
	}

	return symbols, nil
}

// GetSymbolsInfoList 分页获取Symbols记录
func (symbolsService *SymbolsService) GetSymbolsInfoList(info symbolReq.SymbolsSearch) (list interface{}, total int64, err error) {
	// 创建db
	db := global.GVA_DB.Model(&symbol.Symbols{})
	var symbolsList []symbol.Symbols

	// 添加查询条件
	if info.Symbol != "" {
		db = db.Where("symbol LIKE ?", "%"+info.Symbol+"%")
	}
	if info.Type != nil {
		db = db.Where("type = ?", *info.Type)
	}

	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 如果总数为0，直接返回空结果
	if total == 0 {
		return []interface{}{}, 0, nil
	}

	// 使用和 advisor_prod 相同的分页方式
	err = info.Paginate()(db).Find(&symbolsList).Error
	if err != nil {
		return nil, 0, err
	}

	// 为每个symbol加价格信息
	enrichedList := make([]map[string]interface{}, len(symbolsList))
	for i, symbol := range symbolsList {
		enrichedSymbol := map[string]interface{}{
			"id":             symbol.Id,
			"createdAt":      symbol.CreatedAt,
			"updatedAt":      symbol.UpdatedAt,
			"symbol":         symbol.Symbol,
			"corporation":    symbol.Corporation,
			"industry":       symbol.Industry,
			"exchange":       symbol.Exchange,
			"marketCap":      symbol.MarketCap,
			"listDate":       symbol.ListDate,
			"description":    symbol.Description,    // 添加公司描述
			"sicDescription": symbol.SicDescription, // 添加行业描述
			"averageVolume":  symbol.AverageVolume,
			"changeRatio":    symbol.ChangeRatio,
			"peRatio":        symbol.PeRatio,
			"icon":           processIconURL(symbol.Icon),
			"type":           symbol.Type,
			"status":         symbol.Status,
			"ticketSize":     symbol.TicketSize,
			"ticketNumSize":  symbol.TicketNumSize,
			"sort":           symbol.Sort,
		}

		// 根据类型构建Redis key
		if symbol.Type != nil {
			var redisKey string
			switch *symbol.Type {
			case 0: // 股票
				redisKey = fmt.Sprintf("symbol:stock:%s", symbol.Symbol)
			case 1: // 加密货币
				redisKey = fmt.Sprintf("symbol:crypto:%s", symbol.Symbol)
			case 2: // 外汇
				redisKey = fmt.Sprintf("symbol:forex:%s", symbol.Symbol)
			}

			// 从Redis获取价格
			if price, err := global.GVA_REDIS.Get(context.Background(), redisKey).Result(); err == nil {
				priceFloat, _ := strconv.ParseFloat(price, 64)
				enrichedSymbol["currentPrice"] = formatPrice(priceFloat, *symbol.Type, symbol.Symbol)
			} else if symbol.CurrentPrice != nil {
				enrichedSymbol["currentPrice"] = formatPrice(*symbol.CurrentPrice, *symbol.Type, symbol.Symbol)
			}
		}

		enrichedList[i] = enrichedSymbol
	}

	return enrichedList, total, err
}

// 在 SymbolsService 结构体中 GetAllSymbols 方法
func (symbolsService *SymbolsService) GetAllSymbols() ([]symbol.Symbols, error) {
	var symbolsList []symbol.Symbols
	err := global.GVA_DB.Find(&symbolsList).Error
	if err != nil {
		return nil, err
	}
	return symbolsList, nil
}

// WSResponse WebSocket响应结构
type WSResponse struct {
	Type string      `json:"type"`          // 消息类型：connected/subscribed/price_update
	Data interface{} `json:"data"`          // 数据内容
	Msg  string      `json:"msg,omitempty"` // 可选的消息说明
}

// PriceData 价格数据结构
type PriceData struct {
	Price       float64 `json:"price"`
	ChangeRatio float64 `json:"changeRatio"` // 涨跌幅
}

// SubscribeSymbolPrices 处理WebSocket连接和价格订阅
func (symbolsService *SymbolsService) SubscribeSymbolPrices(conn *websocket.Conn) {
	// 用于存储所有订阅的symbols和它们的价格数据
	prices := make(map[string]PriceData)
	// 用于存储每个类型的订阅symbols
	subscribedSymbols := make(map[int]map[string]bool)

	defer conn.Close()

	// 发送连接成功消息
	response := WSResponse{
		Type: "connected",
		Msg:  "WebSocket connected successfully",
	}
	if err := conn.WriteJSON(response); err != nil {
		global.GVA_LOG.Error("write json error:", zap.Error(err))
		return
	}

	// 创建定时器，定期更新价格
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// 创建一个用于优雅退出的channel
	done := make(chan struct{})
	defer close(done)

	// 启动goroutine处价格更新
	go func() {
		for {
			select {
			case <-ticker.C:
				// 更新所有已订阅的价格
				for typeID, symbols := range subscribedSymbols {
					if len(symbols) > 0 {
						symbolsList := make([]string, 0, len(symbols))
						for symbol := range symbols {
							symbolsList = append(symbolsList, symbol)
						}

						// 获取最新价格和前一天收盘价
						newPrices := make(map[string]PriceData)
						for _, symbol := range symbolsList {
							var (
								currentKey       string
								previousCloseKey string
							)

							switch typeID {
							case 0:
								currentKey = fmt.Sprintf("symbol:stock:%s", symbol)
								previousCloseKey = fmt.Sprintf("symbol:stock:%s:previous_close", symbol)
							case 1:
								currentKey = fmt.Sprintf("symbol:crypto:%s", symbol)
								previousCloseKey = fmt.Sprintf("symbol:crypto:%s:previous_close", symbol)
							case 2:
								currentKey = fmt.Sprintf("symbol:forex:%s", symbol)
								previousCloseKey = fmt.Sprintf("symbol:forex:%s:previous_close", symbol)
							}

							// 获取当前价格
							currentPrice, err := global.GVA_REDIS.Get(context.Background(), currentKey).Float64()
							if err != nil {
								// 尝试从1分钟K线获取最新价格
								klineKey := fmt.Sprintf("%s:kline:1m", currentKey)
								if klineStr, err := global.GVA_REDIS.Get(context.Background(), klineKey).Result(); err == nil {
									// 解析单个K线数据
									var klineData struct {
										Close  float64 `json:"c"`
										High   float64 `json:"h"`
										Low    float64 `json:"l"`
										Open   float64 `json:"o"`
										Time   int64   `json:"t"`
										Volume float64 `json:"v"`
									}
									if err := json.Unmarshal([]byte(klineStr), &klineData); err == nil {
										currentPrice = klineData.Close
									} else {
										continue
									}
								} else {
									continue
								}
							}

							// 格式化价格
							currentPrice = formatPrice(currentPrice, typeID, symbol)

							// 获取前一天收盘价
							previousClose, err := global.GVA_REDIS.Get(context.Background(), previousCloseKey).Float64()

							// 计算涨跌幅
							changeRatio := 0.0
							if previousClose != 0 {
								// 计算涨跌幅并保留两位小数
								changeRatio = math.Round(((currentPrice-previousClose)/previousClose*100)*100) / 100
							}

							newPrices[symbol] = PriceData{
								Price:       currentPrice,
								ChangeRatio: changeRatio,
							}
						}

						// 更新总价map
						for symbol, data := range newPrices {
							prices[symbol] = data
						}
					}
				}

				// 只有在有订阅数据时才发送更新
				if len(prices) > 0 {
					response := WSResponse{
						Type: "price_update",
						Data: prices,
					}
					if err := conn.WriteJSON(response); err != nil {
						global.GVA_LOG.Error("write json error:", zap.Error(err))
						return
					}
				}
			case <-done:
				return
			}
		}
	}()

	// 处理订阅和取消订阅的消息
	for {
		var msg SubscribeMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			global.GVA_LOG.Error("read message error:", zap.Error(err))
			break
		}

		// 确保subscribedSymbols中有对应type的map
		if subscribedSymbols[msg.Type] == nil {
			subscribedSymbols[msg.Type] = make(map[string]bool)
		}

		switch msg.Action {
		case "subscribe":
			// 添加新的订阅
			for _, symbol := range msg.Symbols {
				subscribedSymbols[msg.Type][symbol] = true
			}

			// 送订阅确认消息
			response := WSResponse{
				Type: "subscribed",
				Data: map[string]interface{}{
					"symbols":    msg.Symbols,
					"symbolType": msg.Type,
				},
				Msg: "Subscription updated",
			}
			if err := conn.WriteJSON(response); err != nil {
				global.GVA_LOG.Error("write json error:", zap.Error(err))
				break
			}

		case "unsubscribe":
			// 如果symbols为空，取消该type的所有订阅
			if len(msg.Symbols) == 0 {
				// 先删除这个type下所有symbol价格
				for symbol := range subscribedSymbols[msg.Type] {
					delete(prices, symbol)
				}
				delete(subscribedSymbols, msg.Type)
			} else {
				// 取消指定symbols的订阅
				for _, symbol := range msg.Symbols {
					delete(subscribedSymbols[msg.Type], symbol)
					delete(prices, symbol)
				}
				// 如果该type没有任何订阅了，清除该type
				if len(subscribedSymbols[msg.Type]) == 0 {
					delete(subscribedSymbols, msg.Type)
				}
			}

			// 发送取消订阅确认消息
			response := WSResponse{
				Type: "unsubscribed",
				Data: map[string]interface{}{
					"symbols":    msg.Symbols,
					"symbolType": msg.Type,
				},
				Msg: "Unsubscription confirmed",
			}
			if err := conn.WriteJSON(response); err != nil {
				global.GVA_LOG.Error("write json error:", zap.Error(err))
				break
			}
		}
	}
}

// getPricesForSymbols 根据symbol类型获取价格
func getPricesForSymbols(symbols []string, symbolType int) map[string]float64 {
	prices := make(map[string]float64)
	ctx := context.Background()

	for _, symbol := range symbols {
		var redisKey string
		switch symbolType {
		case 0: // 股票
			redisKey = fmt.Sprintf("symbol:stock:%s", symbol)
		case 1: // 加密货币
			redisKey = fmt.Sprintf("symbol:crypto:%s", symbol)
		case 2: // 外汇
			redisKey = fmt.Sprintf("symbol:forex:%s", symbol)
		}

		// 从Redis获取价格
		if price, err := global.GVA_REDIS.Get(ctx, redisKey).Result(); err == nil {
			if priceFloat, err := strconv.ParseFloat(price, 64); err == nil {
				prices[symbol] = priceFloat
			}
		}
	}

	return prices
}

// BatchUpdateSymbols symbols表记录
func (symbolsService *SymbolsService) BatchUpdateSymbols(symbols []symbol.Symbols) error {
	if len(symbols) == 0 {
		return nil
	}

	// 构建 CASE WHEN 句
	var (
		cases     []string
		valueArgs []interface{}
	)

	// 收集所有需要更新的ID
	var ids []interface{}
	for _, s := range symbols {
		if s.Id != nil {
			ids = append(ids, *s.Id)
		}
	}

	// 构建批量更新SQL
	sql := `UPDATE symbols SET 
			corporation = CASE id `
	for _, s := range symbols {
		if s.Id != nil {
			cases = append(cases, "WHEN ? THEN ?")
			valueArgs = append(valueArgs, *s.Id, s.Corporation)
		}
	}
	sql += strings.Join(cases, " ") + " END, "

	// 重置cases数
	cases = cases[:0]

	// Exchange
	sql += "exchange = CASE id "
	for _, s := range symbols {
		if s.Id != nil {
			cases = append(cases, "WHEN ? THEN ?")
			valueArgs = append(valueArgs, *s.Id, s.Exchange)
		}
	}
	sql += strings.Join(cases, " ") + " END, "

	// Market Cap
	cases = cases[:0]
	sql += "market_cap = CASE id "
	for _, s := range symbols {
		if s.Id != nil {
			cases = append(cases, "WHEN ? THEN ?")
			valueArgs = append(valueArgs, *s.Id, s.MarketCap)
		}
	}
	sql += strings.Join(cases, " ") + " END, "

	// Industry
	cases = cases[:0]
	sql += "industry = CASE id "
	for _, s := range symbols {
		if s.Id != nil {
			cases = append(cases, "WHEN ? THEN ?")
			valueArgs = append(valueArgs, *s.Id, s.Industry)
		}
	}
	sql += strings.Join(cases, " ") + " END, "

	// List Date
	cases = cases[:0]
	sql += "list_date = CASE id "
	for _, s := range symbols {
		if s.Id != nil {
			cases = append(cases, "WHEN ? THEN ?")
			valueArgs = append(valueArgs, *s.Id, s.ListDate)
		}
	}
	sql += strings.Join(cases, " ") + " END, "

	// Status
	cases = cases[:0]
	sql += "status = CASE id "
	for _, s := range symbols {
		if s.Id != nil {
			cases = append(cases, "WHEN ? THEN ?")
			valueArgs = append(valueArgs, *s.Id, s.Status)
		}
	}
	sql += strings.Join(cases, " ") + " END, "

	// TicketSize
	cases = cases[:0]
	sql += "ticket_size = CASE id "
	for _, s := range symbols {
		if s.Id != nil {
			cases = append(cases, "WHEN ? THEN ?")
			valueArgs = append(valueArgs, *s.Id, s.TicketSize)
		}
	}
	sql += strings.Join(cases, " ") + " END, "

	// TicketNumSize
	cases = cases[:0]
	sql += "ticket_num_size = CASE id "
	for _, s := range symbols {
		if s.Id != nil {
			cases = append(cases, "WHEN ? THEN ?")
			valueArgs = append(valueArgs, *s.Id, s.TicketNumSize)
		}
	}
	sql += strings.Join(cases, " ") + " END, "

	// Sort
	cases = cases[:0]
	sql += "sort = CASE id "
	for _, s := range symbols {
		if s.Id != nil {
			cases = append(cases, "WHEN ? THEN ?")
			valueArgs = append(valueArgs, *s.Id, s.Sort)
		}
	}
	sql += strings.Join(cases, " ") + " END "

	// WHERE 子句
	sql += " WHERE id IN ?"
	valueArgs = append(valueArgs, ids)

	// 执行更新
	return global.GVA_DB.Exec(sql, valueArgs...).Error
}

// GetSymbolsPublic 获取公的symbols信息
func (symbolsService *SymbolsService) GetSymbolsPublic(info symbolReq.SymbolsPublicSearch) (interface{}, int64, error) {
	db := global.GVA_DB.Model(&symbol.Symbols{})
	var (
		total    int64
		orderDir = "ASC"
	)

	// 打印查询参数
	global.GVA_LOG.Info("查询参数",
		zap.String("orderBy", info.OrderBy),
		zap.String("order", info.Order),
		zap.Int("page", info.Page),
		zap.Int("pageSize", info.PageSize))

	// 判断是升序还是降序
	if strings.ToUpper(info.Order) != orderDir {
		orderDir = "DESC"
	}

	// 添加排序
	switch info.OrderBy {
	case "changeRatio":
		db = db.Order("change_ratio " + orderDir)
	case "marketCap":
		db = db.Order("market_cap IS NULL").
			Order("market_cap " + orderDir)
	default:
		// 默认按市值序
		if len(info.OrderBy) > 0 {
			db = db.Order(info.OrderBy + " " + orderDir)

		} else {
			db = db.Order("market_cap IS NULL").
				Order("market_cap " + orderDir)
		}
	}

	// 计算总数
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 如果总数为0，直接返回空结果
	if total == 0 {
		return []interface{}{}, 0, nil
	}

	var symbolsList []symbol.Symbols
	// 使用 Paginate 方法进行分页查询
	err = info.Paginate()(db).Find(&symbolsList).Error
	if err != nil {
		return nil, 0, err
	}

	// 创建enriched列表
	enrichedList := make([]map[string]interface{}, len(symbolsList))
	for i, symbol := range symbolsList {
		enrichedSymbol := map[string]interface{}{
			"id":             symbol.Id,
			"symbol":         symbol.Symbol,
			"corporation":    symbol.Corporation,
			"industry":       symbol.Industry,
			"exchange":       symbol.Exchange,
			"marketCap":      symbol.MarketCap,
			"listDate":       symbol.ListDate,
			"description":    symbol.Description,    // 添加公司描述
			"sicDescription": symbol.SicDescription, // 添加行业描述
			"averageVolume":  symbol.AverageVolume,
			"changeRatio":    symbol.ChangeRatio,
			"peRatio":        symbol.PeRatio,
			"icon":           processIconURL(symbol.Icon),
			"type":           symbol.Type,
			"status":         symbol.Status,
			"ticketSize":     symbol.TicketSize,
			"ticketNumSize":  symbol.TicketNumSize,
			"sort":           symbol.Sort,
		}

		// 获取实时价格
		if symbol.Type != nil {
			var redisKey string
			switch *symbol.Type {
			case 0: // 股票
				redisKey = fmt.Sprintf("symbol:stock:%s", symbol.Symbol)
			case 1: // 加密货币
				redisKey = fmt.Sprintf("symbol:crypto:%s", symbol.Symbol)
			case 2: // 外汇
				redisKey = fmt.Sprintf("symbol:forex:%s", symbol.Symbol)
			}

			// 从Redis获取价格
			if price, err := global.GVA_REDIS.Get(context.Background(), redisKey).Result(); err == nil {
				priceFloat, _ := strconv.ParseFloat(price, 64)
				enrichedSymbol["currentPrice"] = formatPrice(priceFloat, *symbol.Type, symbol.Symbol)
			} else if symbol.CurrentPrice != nil {
				enrichedSymbol["currentPrice"] = formatPrice(*symbol.CurrentPrice, *symbol.Type, symbol.Symbol)
			}
		}

		enrichedList[i] = enrichedSymbol
	}

	return enrichedList, total, nil
}

// 股票响应结构
type StockDetails struct {
	Results struct {
		Ticker         string  `json:"ticker"`
		Name           string  `json:"name"`
		Market         string  `json:"market"`
		PrimaryExch    string  `json:"primary_exchange"`
		Type           string  `json:"type"`
		Active         bool    `json:"active"`
		MarketCap      float64 `json:"market_cap"`
		ListDate       string  `json:"list_date"`
		Description    string  `json:"description"`
		SicDescription string  `json:"sic_description"`
		Branding       struct {
			IconUrl string `json:"icon_url"`
		} `json:"branding"`
	} `json:"results"`
	Status string `json:"status"`
}

// 加密货币响应结构
type CryptoDetails struct {
	Results struct {
		Ticker             string `json:"ticker"`
		Name               string `json:"name"`
		Market             string `json:"market"`
		Active             bool   `json:"active"`
		CurrencySymbol     string `json:"currency_symbol"`
		BaseCurrencySymbol string `json:"base_currency_symbol"`
		BaseCurrencyName   string `json:"base_currency_name"`
		Branding           struct {
			IconUrl string `json:"icon_url"`
		} `json:"branding"`
	} `json:"results"`
	Status string `json:"status"`
}

// 外汇响应结构
type ForexDetails struct {
	Results struct {
		Ticker             string `json:"ticker"`
		Name               string `json:"name"`
		Market             string `json:"market"`
		Locale             string `json:"locale"`
		Active             bool   `json:"active"`
		CurrencySymbol     string `json:"currency_symbol"`
		CurrencyName       string `json:"currency_name"`
		BaseCurrencySymbol string `json:"base_currency_symbol"`
		BaseCurrencyName   string `json:"base_currency_name"`
	} `json:"results"`
	Status string `json:"status"`
}

// AddSymbolsByName 修改为使用结构体参数
func (symbolsService *SymbolsService) AddSymbolsByName(req symbolReq.AddSymbolRequest) (err error) {
	// 去除可能存在的引号
	symbolStr := strings.Trim(req.Symbol, "\"'")

	// 检查 symbol 是否已存在
	var count int64
	if err := global.GVA_DB.Model(&symbol.Symbols{}).Where("symbol = ?", symbolStr).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("symbol %s already exists", symbolStr)
	}

	// 根据类型构建正确的 ticker 格式
	var tickerForAPI string
	switch req.Type {
	case 0: // 股票
		tickerForAPI = symbolStr
	case 1: // 加密货币
		tickerForAPI = "X:" + symbolStr
	case 2: // 外汇
		tickerForAPI = "C:" + symbolStr
	}

	// 构建API URL
	detailsUrl := fmt.Sprintf("%s/v3/reference/tickers/%s?apiKey=%s",
		global.GVA_CONFIG.Polygon.BaseURL,
		url.QueryEscape(tickerForAPI), // 使用 url.QueryEscape 确保 URL 安全
		global.GVA_CONFIG.Polygon.APIKey)

	global.GVA_LOG.Info("Requesting details",
		zap.String("url", detailsUrl),
		zap.String("originalSymbol", symbolStr),
		zap.String("tickerForAPI", tickerForAPI))

	resp, err := http.Get(detailsUrl)
	if err != nil {
		return fmt.Errorf("failed to get details: %v", err)
	}
	defer resp.Body.Close()

	// 在switch前定义newSymbol
	newSymbol := &symbol.Symbols{
		Type: &req.Type,
	}

	// 根据不同类型解析响应
	switch req.Type {
	case 0: // 股票
		var details StockDetails
		if err := json.NewDecoder(resp.Body).Decode(&details); err != nil {
			return fmt.Errorf("failed to decode stock details: %v", err)
		}
		global.GVA_LOG.Info("Details Response",
			zap.Any("details", details))

		marketCap := int(details.Results.MarketCap)

		// 转换上市日期字符串为time.Time
		listDate, err := time.Parse("2006-01-02", details.Results.ListDate)
		if err == nil {
			newSymbol.ListDate = &listDate
		}

		newSymbol.Symbol = symbolStr
		newSymbol.Corporation = details.Results.Name
		newSymbol.Exchange = details.Results.PrimaryExch
		newSymbol.MarketCap = &marketCap
		newSymbol.Industry = details.Results.SicDescription
		newSymbol.SicDescription = details.Results.SicDescription
		newSymbol.Icon = details.Results.Branding.IconUrl
		newSymbol.Description = details.Results.Description

	case 1: // 加密货币
		var details CryptoDetails
		if err := json.NewDecoder(resp.Body).Decode(&details); err != nil {
			return fmt.Errorf("failed to decode crypto details: %v", err)
		}
		global.GVA_LOG.Info("Details Response",
			zap.Any("details", details))
		newSymbol.Symbol = fmt.Sprintf("%s/%s", details.Results.BaseCurrencySymbol, details.Results.CurrencySymbol)
		newSymbol.Corporation = details.Results.BaseCurrencyName
		newSymbol.Exchange = details.Results.Market
		newSymbol.Icon = details.Results.Branding.IconUrl

	case 2: // 外汇
		var details ForexDetails
		if err := json.NewDecoder(resp.Body).Decode(&details); err != nil {
			return fmt.Errorf("failed to decode forex details: %v", err)
		}
		global.GVA_LOG.Info("Details Response",
			zap.Any("details", details))
		newSymbol.Symbol = fmt.Sprintf("%s/%s", details.Results.BaseCurrencySymbol, details.Results.CurrencySymbol)
		newSymbol.Corporation = details.Results.BaseCurrencyName
		newSymbol.Exchange = details.Results.Market
	}

	// 构建价格查询URL
	var priceSymbol string
	switch req.Type {
	case 0: // 股票
		priceSymbol = symbolStr
	case 1: // 加密货币
		priceSymbol = "X:" + symbolStr
	case 2: // 外汇
		priceSymbol = "C:" + symbolStr
	}

	priceUrl := fmt.Sprintf("%s/v2/aggs/ticker/%s/prev?apiKey=%s",
		global.GVA_CONFIG.Polygon.BaseURL,
		priceSymbol,
		global.GVA_CONFIG.Polygon.APIKey)

	// 取价格数据
	if priceResp, err := http.Get(priceUrl); err == nil {
		defer priceResp.Body.Close()
		var priceData struct {
			Results []struct {
				C float64 `json:"c"` // Close price
				O float64 `json:"o"` // Open price
				V float64 `json:"v"` // Volume
			} `json:"results"`
		}
		if err := json.NewDecoder(priceResp.Body).Decode(&priceData); err == nil && len(priceData.Results) > 0 {
			price := priceData.Results[0].C
			newSymbol.CurrentPrice = &price

			// 计算涨跌幅
			if priceData.Results[0].O != 0 {
				changeRatio := ((priceData.Results[0].C - priceData.Results[0].O) / priceData.Results[0].O) * 100
				newSymbol.ChangeRatio = &changeRatio
			}

			// 设置交易量
			volume := int(priceData.Results[0].V)
			newSymbol.AverageVolume = &volume

			// Redis缓存 - 使用 newSymbol.Symbol 作为 key 的一部分
			var redisKey string
			switch req.Type {
			case 0: // 股票
				redisKey = fmt.Sprintf("symbol:stock:%s", newSymbol.Symbol)
			case 1: // 加密货币
				redisKey = fmt.Sprintf("symbol:crypto:%s", newSymbol.Symbol)
			case 2: // 外汇
				redisKey = fmt.Sprintf("symbol:forex:%s", newSymbol.Symbol)
			}

			err = global.GVA_REDIS.Set(context.Background(), redisKey, price, 0).Err()
			if err != nil {
				global.GVA_LOG.Error("Failed to set price in Redis", zap.Error(err))
			}
		}
	}
	global.GVA_LOG.Info("newSymbol", zap.Any("info", newSymbol))
	// 设置默认精度
	defaultTicketSize := 0.01
	newSymbol.TicketSize = &defaultTicketSize
	defaultTicketNumSize := 1.0
	newSymbol.TicketNumSize = &defaultTicketNumSize

	// 创建记录
	if err = global.GVA_DB.Create(newSymbol).Error; err != nil {
		return err
	}

	// 将精度信息存入Redis
	if err = setSymbolPrecision(newSymbol.Symbol, req.Type, newSymbol.TicketSize, newSymbol.TicketNumSize); err != nil {
		global.GVA_LOG.Error("Failed to set precision in Redis",
			zap.String("symbol", newSymbol.Symbol),
			zap.Error(err))
	}

	// 创建记录成功后，订阅新的symbol
	global.GVA_WS_MUTEX.Lock()
	defer global.GVA_WS_MUTEX.Unlock()

	// var client *polygonws.Client
	// var formattedSymbol string

	// switch req.Type {
	// case 0: // 股票
	// 	client = global.GVA_WS_STOCK
	// 	formattedSymbol = newSymbol.Symbol
	// case 1: // 加密货币
	// 	client = global.GVA_WS_CRYPTO
	// 	formattedSymbol = "X:" + strings.Replace(newSymbol.Symbol, "/", "-", -1)
	// case 2: // 外汇
	// 	client = global.GVA_WS_FOREX
	// 	formattedSymbol = "C:" + strings.Replace(newSymbol.Symbol, "/", "-", -1)
	// }

	// if client != nil {
	// 	// 订阅聚合数据
	// 	switch req.Type {
	// 	case 0:
	// 		if err := client.Subscribe(polygonws.StocksSecAggs, formattedSymbol); err != nil {
	// 			global.GVA_LOG.Error("Failed to subscribe to stock aggregates",
	// 				zap.String("symbol", formattedSymbol),
	// 				zap.Error(err))
	// 		}
	// 		if err := client.Subscribe(polygonws.StocksQuotes, formattedSymbol); err != nil {
	// 			global.GVA_LOG.Error("Failed to subscribe to stock quotes",
	// 				zap.String("symbol", formattedSymbol),
	// 				zap.Error(err))
	// 		}
	// 	case 1:
	// 		if err := client.Subscribe(polygonws.CryptoSecAggs, formattedSymbol); err != nil {
	// 			global.GVA_LOG.Error("Failed to subscribe to crypto aggregates",
	// 				zap.String("symbol", formattedSymbol),
	// 				zap.Error(err))
	// 		}
	// 		if err := client.Subscribe(polygonws.CryptoQuotes, formattedSymbol); err != nil {
	// 			global.GVA_LOG.Error("Failed to subscribe to crypto quotes",
	// 				zap.String("symbol", formattedSymbol),
	// 				zap.Error(err))
	// 		}
	// 	case 2:
	// 		if err := client.Subscribe(polygonws.ForexSecAggs, formattedSymbol); err != nil {
	// 			global.GVA_LOG.Error("Failed to subscribe to forex aggregates",
	// 				zap.String("symbol", formattedSymbol),
	// 				zap.Error(err))
	// 		}
	// 		if err := client.Subscribe(polygonws.ForexQuotes, formattedSymbol); err != nil {
	// 			global.GVA_LOG.Error("Failed to subscribe to forex quotes",
	// 				zap.String("symbol", formattedSymbol),
	// 				zap.Error(err))
	// 		}
	// 	}
	// }

	return nil
}

// GetSymbolPrice 根据symbol类型获取价格
func (s *SymbolsService) GetSymbolPrice(symbolStr string) (string, error) {
	// 1. 先查询symbol的类型
	var symbolRecord struct {
		Type *int
	}
	if err := global.GVA_DB.Table("symbols").Select("type").Where("symbol = ?", symbolStr).First(&symbolRecord).Error; err != nil {
		return "", fmt.Errorf("symbol %s not found in database", symbolStr)
	}

	// 2. 根据类型构建不同的Redis key
	var redisKey string
	if symbolRecord.Type == nil {
		return "", fmt.Errorf("symbol type is nil for %s", symbolStr)
	}

	switch *symbolRecord.Type {
	case 0: // 股票
		redisKey = fmt.Sprintf("symbol:stock:%s", symbolStr)
	case 1: // 加密货币
		redisKey = fmt.Sprintf("symbol:crypto:%s", symbolStr)
	case 2: // 外汇
		redisKey = fmt.Sprintf("symbol:forex:%s", symbolStr)
	default:
		return "", fmt.Errorf("unknown symbol type: %d", *symbolRecord.Type)
	}

	// 3. 从Redis获取价格
	price, err := global.GVA_REDIS.Get(context.Background(), redisKey).Result()
	if err != nil {
		if err == redis.Nil {
			return "", fmt.Errorf("no price data available for symbol %s", symbolStr)
		}
		return "", fmt.Errorf("error getting price data: %v", err)
	}

	return price, nil
}

// KlineData 定义K线数据结构
type KlineData struct {
	Open   float64 `json:"o"`
	High   float64 `json:"h"`
	Low    float64 `json:"l"`
	Close  float64 `json:"c"`
	Volume float64 `json:"v"`
	Time   int64   `json:"t"`
}

// GetKlineData 从Redis获取K线历史数据
func (s *SymbolsService) GetKlineData(symbol, interval string) ([]KlineData, error) {
	// 1. 先查询symbol的类型
	var symbolRecord struct {
		Symbol string
		Type   *int
	}
	if err := global.GVA_DB.Table("symbols").Select("symbol", "type").
		Where("symbol = ?", symbol).First(&symbolRecord).Error; err != nil {
		return nil, fmt.Errorf("symbol %s not found in database", symbol)
	}

	if symbolRecord.Type == nil {
		return nil, fmt.Errorf("symbol type is nil for %s", symbol)
	}

	// 2. 构建Redis key
	var baseKey string
	switch *symbolRecord.Type {
	case 0: // 股票
		baseKey = fmt.Sprintf("symbol:stock:%s", symbol)
	case 1: // 加密货币
		baseKey = fmt.Sprintf("symbol:crypto:%s", symbol)
	case 2: // 外汇
		baseKey = fmt.Sprintf("symbol:forex:%s", symbol)
	default:
		return nil, fmt.Errorf("unknown symbol type: %d", *symbolRecord.Type)
	}

	// 3. 获取历史数据
	historyKey := fmt.Sprintf("%s:kline:%s:history", baseKey, interval)
	data, err := global.GVA_REDIS.Get(context.Background(), historyKey).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get kline history data from redis: %v", err)
	}

	// 4. 解析数据
	var klineDataList []KlineData
	if err := json.Unmarshal([]byte(data), &klineDataList); err != nil {
		return nil, fmt.Errorf("failed to unmarshal kline data: %v", err)
	}

	// 5. 获取实时价格
	realTimePrice, err := global.GVA_REDIS.Get(context.Background(), baseKey).Result()
	if err == nil && len(klineDataList) > 0 {
		// 如果有实时价格且K线列表不为空
		if price, err := strconv.ParseFloat(realTimePrice, 64); err == nil {
			// 更新第一个K线的收盘价（因为是最新的）
			klineDataList[0].Close = price

			// 如果实时价格高于最高价,更新最高价
			if price > klineDataList[0].High {
				klineDataList[0].High = price
			}

			// 如果实时价格低于最低价,更新最低价
			if price < klineDataList[0].Low {
				klineDataList[0].Low = price
			}
		}
	}

	return klineDataList, nil
}

// parseInterval 将前端传来的时间间隔转换为 Polygon 支持的格式
func parseInterval(interval string) (multiplier int, timespan string, err error) {
	switch interval {
	case "1m":
		return 1, "minute", nil
	case "5m":
		return 5, "minute", nil
	case "15m":
		return 15, "minute", nil
	case "30m":
		return 30, "minute", nil
	case "1h":
		return 1, "hour", nil
	case "2h":
		return 2, "hour", nil
	case "4h":
		return 4, "hour", nil
	case "1d":
		return 1, "day", nil
	case "1w":
		return 1, "week", nil
	case "1M":
		return 1, "month", nil
	default:
		return 0, "", fmt.Errorf("unsupported interval: %s", interval)
	}
}

// GetAllSymbolsSimple 取所有股票的简化信息
func (symbolsService *SymbolsService) GetAllSymbolsSimple() ([]map[string]interface{}, error) {
	var symbolsList []symbol.Symbols
	// 查询必要的字段
	err := global.GVA_DB.Select("id, symbol, type, current_price").Find(&symbolsList).Error
	if err != nil {
		return nil, err
	}

	result := make([]map[string]interface{}, len(symbolsList))
	for i, sym := range symbolsList {
		// 创建基本信息
		symbolInfo := map[string]interface{}{
			"id":     sym.Id,
			"symbol": sym.Symbol,
		}

		// 根据类型构建Redis key并获取价格
		if sym.Type != nil {
			var redisKey string
			switch *sym.Type {
			case 0: // 股票
				redisKey = fmt.Sprintf("symbol:stock:%s", sym.Symbol)
			case 1: // 加密货币
				redisKey = fmt.Sprintf("symbol:crypto:%s", sym.Symbol)
			case 2: // 外汇
				redisKey = fmt.Sprintf("symbol:forex:%s", sym.Symbol)
			}

			// 从Redis获取价格
			if price, err := global.GVA_REDIS.Get(context.Background(), redisKey).Result(); err == nil {
				priceFloat, _ := strconv.ParseFloat(price, 64)
				symbolInfo["price"] = priceFloat
			} else {
				// 如果Redis中没有价格,使用数据库中的current_price
				if sym.CurrentPrice != nil {
					symbolInfo["price"] = *sym.CurrentPrice
				}
			}
		}

		result[i] = symbolInfo
	}

	return result, nil
}

// TickerSearchResponse Polygon API响应结构
type TickerSearchResponse struct {
	Results []struct {
		Ticker           string `json:"ticker"`
		Name             string `json:"name"`
		Market           string `json:"market"`
		Active           bool   `json:"active"`
		PrimaryExch      string `json:"primary_exchange,omitempty"`
		Type             string `json:"type,omitempty"`
		CurrencyName     string `json:"currency_name,omitempty"`
		BaseCurrencyName string `json:"base_currency_name,omitempty"`
		Branding         struct {
			IconUrl string `json:"icon_url"`
		} `json:"branding"`
	} `json:"results"`
	Status  string `json:"status"`
	Count   int    `json:"count"`
	NextUrl string `json:"next_url"`
}

// SearchTickers 查询可添加的ticker列表
func (symbolsService *SymbolsService) SearchTickers(req symbolReq.TickerSearchReq) (interface{}, error) {
	// 根据type确定market参数
	var market string
	switch req.Type {
	case 0:
		market = "stocks"
	case 1:
		market = "crypto"
	case 2:
		market = "fx"
	default:
		return nil, fmt.Errorf("invalid type: %d", req.Type)
	}

	// 构建URL
	baseUrl := fmt.Sprintf("%s/v3/reference/tickers", global.GVA_CONFIG.Polygon.BaseURL)

	// 创建URL参数
	params := url.Values{}
	params.Add("market", market)

	// 根类型处理搜索关键字
	if req.Search != "" {
		var searchTerm string
		switch req.Type {
		case 0: // 股票
			searchTerm = req.Search
		case 1: // 加货币
			searchTerm = "X:" + req.Search
		case 2: // 外汇
			searchTerm = "C:" + req.Search
		}
		params.Add("search", searchTerm)
	}

	params.Add("active", "true") // 只查询活跃的
	params.Add("limit", "50")    // 限制返回数量
	params.Add("apiKey", global.GVA_CONFIG.Polygon.APIKey)

	// 发送请求
	resp, err := http.Get(baseUrl + "?" + params.Encode())
	if err != nil {
		return nil, fmt.Errorf("failed to search tickers: %v", err)
	}
	defer resp.Body.Close()

	var response TickerSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	// 转换为前端所需格式
	result := make([]map[string]interface{}, 0)
	for _, ticker := range response.Results {
		// 检查symbol是否已存在
		var count int64
		if err := global.GVA_DB.Model(&symbol.Symbols{}).Where("symbol = ?", ticker.Ticker).Count(&count).Error; err != nil {
			global.GVA_LOG.Error("检查symbol是否存在失败", zap.Error(err))
			continue
		}

		item := map[string]interface{}{
			"ticker":   ticker.Ticker,
			"name":     ticker.Name,
			"market":   market,
			"exchange": ticker.PrimaryExch,
			"icon":     processIconURL(ticker.Branding.IconUrl),
			"exists":   count > 0, // 标记是否已存在
			"type":     req.Type,  // 添加type字段
		}

		// 根据类型添加特定字段
		switch req.Type {
		case 0: // 股票
			item["type"] = ticker.Type
		case 1, 2: // 加密货币或外汇
			if ticker.BaseCurrencyName != "" {
				item["name"] = ticker.BaseCurrencyName
			}
		}

		result = append(result, item)
	}

	return result, nil
}

// GetSymbolPriceById 根据ID获取币种价格
func (s *SymbolsService) GetSymbolPriceById(symbolId string) (interface{}, error) {
	var symbolRecord struct {
		Symbol       string
		Type         *int
		CurrentPrice *float64
	}

	if err := global.GVA_DB.Table("symbols").Select("symbol", "type", "current_price").
		Where("id = ?", symbolId).First(&symbolRecord).Error; err != nil {
		return nil, fmt.Errorf("币种不存在")
	}

	if symbolRecord.Type == nil {
		return nil, fmt.Errorf("币种类型未知")
	}

	// 根据类型构建Redis key
	var redisKey string
	switch *symbolRecord.Type {
	case 0: // 股票
		redisKey = fmt.Sprintf("symbol:stock:%s", symbolRecord.Symbol)
	case 1: // 加密货币
		redisKey = fmt.Sprintf("symbol:crypto:%s", symbolRecord.Symbol)
	case 2: // 外汇
		redisKey = fmt.Sprintf("symbol:forex:%s", symbolRecord.Symbol)
	default:
		return nil, fmt.Errorf("未知的币种类型: %d", *symbolRecord.Type)
	}

	// 从Redis获取价格
	if price, err := global.GVA_REDIS.Get(context.Background(), redisKey).Result(); err == nil {
		priceFloat, _ := strconv.ParseFloat(price, 64)
		return priceFloat, nil
	}

	if symbolRecord.CurrentPrice != nil {
		return *symbolRecord.CurrentPrice, nil
	}

	return nil, fmt.Errorf("暂无价格数据")
}

// 添加新的响应结构
type MultiTypeSymbolPrices struct {
	Type0 map[string]float64 `json:"type0,omitempty"`
	Type1 map[string]float64 `json:"type1,omitempty"`
	Type2 map[string]float64 `json:"type2,omitempty"`
}

// 添加新消息结构体
type SubscribeMessage struct {
	Action  string   `json:"action"` // "subscribe" 或 "unsubscribe"
	Symbols []string `json:"symbols"`
	Type    int      `json:"type"`
}

// GetSymbolKlineData 获取指定symbol的聚合数据
func (s *SymbolsService) GetSymbolAggData(symbol string, interval string) ([]KlineData, error) {
	// 1. 先查询symbol类型
	var symbolRecord struct {
		Type *int
	}
	if err := global.GVA_DB.Table("symbols").Select("type").Where("symbol = ?", symbol).First(&symbolRecord).Error; err != nil {
		return nil, fmt.Errorf("symbol %s not found in database", symbol)
	}

	// 2. 构建Redis key
	var redisKey string
	switch *symbolRecord.Type {
	case 0:
		redisKey = fmt.Sprintf("symbol:stock:%s:kline:%s", symbol, interval)
	case 1:
		redisKey = fmt.Sprintf("symbol:crypto:%s:kline:%s", symbol, interval)
	case 2:
		redisKey = fmt.Sprintf("symbol:forex:%s:kline:%s", symbol, interval)
	default:
		return nil, fmt.Errorf("unknown symbol type: %d", *symbolRecord.Type)
	}

	// 3. 从Redis获取K线数据
	jsonData, err := global.GVA_REDIS.Get(context.Background(), redisKey).Result()
	if err == nil {
		// Redis有数据，直接解析返回
		var klineData []KlineData
		if err := json.Unmarshal([]byte(jsonData), &klineData); err != nil {
			return nil, fmt.Errorf("error unmarshaling kline data: %v", err)
		}
		return klineData, nil
	}

	// Redis中没有数据，从Polygon获取
	var ticker string
	switch *symbolRecord.Type {
	case 0: // 股票
		ticker = symbol
	case 1: // 加密货币
		ticker = "X:" + strings.ReplaceAll(symbol, "/", "")
	case 2: // 外汇
		ticker = "C:" + strings.ReplaceAll(symbol, "/", "")
	}

	// 转换时间间隔
	multiplier, timespan, err := parseInterval(interval)
	if err != nil {
		return nil, err
	}

	// 计算时间范围
	now := time.Now()
	var startTime time.Time
	switch interval {
	case "1m":
		startTime = now.Add(-time.Hour * 2) // 获取最近2小时的分钟数据
	case "1h":
		startTime = now.AddDate(0, 0, -2) // 获取最近2天的小时数据
	case "4h":
		startTime = now.AddDate(0, 0, -3) // 获取最近3天的4小时数据
	case "1d":
		startTime = now.AddDate(0, 0, -7) // 获取最近7天的日线数据
	case "1w":
		startTime = now.AddDate(0, -1, 0) // 获取最近1个月的周线数据
	case "1M":
		startTime = now.AddDate(-1, 0, 0) // 获取最近1年的月线数据
	default:
		startTime = now.AddDate(0, 0, -1) // 默认获取最近1天的数据
	}

	// 构建API请求URL - 移除 limit=1，让我们获取更多数据
	url := fmt.Sprintf("%s/v2/aggs/ticker/%s/range/%d/%s/%d/%d?apiKey=%s&sort=desc",
		global.GVA_CONFIG.Polygon.BaseURL,
		ticker,
		multiplier,
		timespan,
		startTime.UnixMilli(),
		now.UnixMilli(),
		global.GVA_CONFIG.Polygon.APIKey)

	// 添加日志
	global.GVA_LOG.Info("Requesting Polygon data",
		zap.String("url", url),
		zap.String("symbol", symbol),
		zap.String("interval", interval))

	// 发送请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get kline data from Polygon: %v", err)
	}
	defer resp.Body.Close()

	// 解析响应
	var result struct {
		Results []struct {
			O float64 `json:"o"`
			H float64 `json:"h"`
			L float64 `json:"l"`
			C float64 `json:"c"`
			V float64 `json:"v"`
			T int64   `json:"t"`
		} `json:"results"`
		Status    string `json:"status"`
		ErrorMsg  string `json:"error"`
		RequestID string `json:"request_id"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode Polygon response: %v", err)
	}

	// 添加响应日志
	global.GVA_LOG.Info("Polygon response",
		zap.Any("results_count", len(result.Results)),
		zap.String("status", result.Status),
		zap.String("error", result.ErrorMsg))

	// 如果没有数据，返回空切片
	if len(result.Results) == 0 {
		return []KlineData{}, nil
	}

	// 构造返回数据 - 使用最后一个（最新的）K线数据
	lastKline := result.Results[0] // 因为已经是 sort=desc，第一个就是最新的
	klineData := []KlineData{
		{
			Open:   lastKline.O,
			High:   lastKline.H,
			Low:    lastKline.L,
			Close:  lastKline.C,
			Volume: lastKline.V,
			Time:   lastKline.T,
		},
	}

	// 将数据存入 Redis
	if jsonData, err := json.Marshal(klineData); err == nil {
		if err := global.GVA_REDIS.Set(context.Background(), redisKey, jsonData, 0).Err(); err != nil {
			global.GVA_LOG.Error("Failed to cache kline data in Redis",
				zap.String("symbol", symbol),
				zap.String("interval", interval),
				zap.Error(err))
		}
	}

	return klineData, nil
}

// ValidateSymbol 验证symbol是否可用
func (symbolsService *SymbolsService) ValidateSymbol(req symbolReq.AddSymbolRequest) error {
	// 去除可能存在的引号
	symbolStr := strings.Trim(req.Symbol, "\"'")

	// 根据类型构建正确的 ticker 格式
	var tickerForAPI string
	switch req.Type {
	case 0: // 股票
		tickerForAPI = symbolStr
	case 1: // 加密货币
		tickerForAPI = "X:" + symbolStr
	case 2: // 外汇
		tickerForAPI = "C:" + symbolStr
	default:
		return fmt.Errorf("invalid symbol type: %d", req.Type)
	}

	// 构建API URL
	detailsUrl := fmt.Sprintf("%s/v3/reference/tickers/%s?apiKey=%s",
		global.GVA_CONFIG.Polygon.BaseURL,
		url.QueryEscape(tickerForAPI),
		global.GVA_CONFIG.Polygon.APIKey)

	global.GVA_LOG.Info("Validating symbol",
		zap.String("url", detailsUrl),
		zap.String("originalSymbol", symbolStr),
		zap.String("tickerForAPI", tickerForAPI))

	// 发送请求
	resp, err := http.Get(detailsUrl)
	if err != nil {
		return fmt.Errorf("failed to validate symbol: %v", err)
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("symbol validation failed: invalid symbol or API error (status: %d) %s", resp.StatusCode, detailsUrl)
	}

	// 根据不同类型解析响应以验证数据的有效性
	switch req.Type {
	case 0: // 股票
		var details StockDetails
		if err := json.NewDecoder(resp.Body).Decode(&details); err != nil {
			return fmt.Errorf("failed to decode stock details: %v", err)
		}
		if !details.Results.Active {
			return fmt.Errorf("stock symbol is inactive")
		}

	case 1: // 加货币
		var details CryptoDetails
		if err := json.NewDecoder(resp.Body).Decode(&details); err != nil {
			return fmt.Errorf("failed to decode crypto details: %v", err)
		}
		if !details.Results.Active {
			return fmt.Errorf("crypto symbol is inactive")
		}

	case 2: // 外汇
		var details ForexDetails
		if err := json.NewDecoder(resp.Body).Decode(&details); err != nil {
			return fmt.Errorf("failed to decode forex details: %v", err)
		}
		if !details.Results.Active {
			return fmt.Errorf("forex symbol is inactive")
		}
	}

	return nil // 验证通过
}

// 添加新的辅助函数用于处理icon URL
func processIconURL(iconURL string) string {
	if iconURL == "" {
		return ""
	}

	if strings.HasPrefix(iconURL, global.GVA_CONFIG.Polygon.BaseURL) {
		// 如果URL已经包含apiKey参数，不要重复添加
		if !strings.Contains(iconURL, "apiKey=") {
			separator := "?"
			if strings.Contains(iconURL, "?") {
				separator = "&"
			}
			return iconURL + separator + "apiKey=" + global.GVA_CONFIG.Polygon.APIKey
		}
	}
	return iconURL
}

// 添加辅助函数用于设置Redis中的精度信息
func setSymbolPrecision(symbol string, symbolType int, ticketSize, ticketNumSize *float64) error {
	var baseKey string
	switch symbolType {
	case 0:
		baseKey = fmt.Sprintf("symbol:stock:%s", symbol)
	case 1:
		baseKey = fmt.Sprintf("symbol:crypto:%s", symbol)
	case 2:
		baseKey = fmt.Sprintf("symbol:forex:%s", symbol)
	}

	ctx := context.Background()
	pipe := global.GVA_REDIS.Pipeline()

	// 设置价格精度
	size := 0.01 // 默认价格精度
	if ticketSize != nil {
		size = *ticketSize
	}
	pipe.Set(ctx, baseKey+":precision", size, 0)

	// 设置数量精度
	numSize := 1.0 // 默认数量精度
	if ticketNumSize != nil {
		numSize = *ticketNumSize
	}
	pipe.Set(ctx, baseKey+":num_precision", numSize, 0)

	_, err := pipe.Exec(ctx)
	return err
}

// 添加格式化价格的辅助函数
func formatPrice(price float64, symbolType int, symbol string) float64 {
	// 从Redis获取精度
	var redisKey string
	switch symbolType {
	case 0:
		redisKey = fmt.Sprintf("symbol:stock:%s:precision", symbol)
	case 1:
		redisKey = fmt.Sprintf("symbol:crypto:%s:precision", symbol)
	case 2:
		redisKey = fmt.Sprintf("symbol:forex:%s:precision", symbol)
	}

	// 获取精度
	precision, err := global.GVA_REDIS.Get(context.Background(), redisKey).Float64()
	if err != nil {
		precision = 0.01 // 默认精度
	}

	// 计算需要保留的小数位数
	decimalPlaces := 0
	tempSize := precision
	for tempSize < 1 {
		decimalPlaces++
		tempSize *= 10
	}

	// 使用math.Round来保留指定位数的小数
	scale := math.Pow(10, float64(decimalPlaces))
	return math.Round(price*scale) / scale
}

// AggDataSubscription K线订阅消息
type AggDataSubscription struct {
	Action   string `json:"action"`   // subscribe/unsubscribe
	Symbol   string `json:"symbol"`   // 交易对
	Interval string `json:"interval"` // 时间间隔
	Type     int    `json:"type"`     // 交易对类型：0-股票 1-加密货币 2-外汇
}

// AggDataResponse K线数据响应
type AggDataResponse struct {
	Type     string     `json:"type"`     // connected/subscribed/kline_update
	Symbol   string     `json:"symbol"`   // 交易对
	Interval string     `json:"interval"` // 时间间隔
	Data     *KlineData `json:"data"`     // K线数据
	Msg      string     `json:"msg,omitempty"`
}

// SubscribeSymbolAggData 处理K线数据WebSocket订阅
func (s *SymbolsService) SubscribeSymbolAggData(conn *websocket.Conn) {
	defer conn.Close()

	// 发送连接成功消息
	resp := AggDataResponse{
		Type: "connected",
		Msg:  "WebSocket connected successfully",
	}
	if err := conn.WriteJSON(resp); err != nil {
		global.GVA_LOG.Error("Failed to send connected message", zap.Error(err))
		return
	}

	// 创建一个定时器，用于定期检查最新数据
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	// 保存当前订阅的信息
	var currentSymbol string
	var currentInterval string
	var currentType int

	// 创建一个channel用于接收客户端消息
	messageChan := make(chan AggDataSubscription)

	// 启动一个goroutine来读取客户端消息
	go func() {
		for {
			var sub AggDataSubscription
			if err := conn.ReadJSON(&sub); err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					global.GVA_LOG.Error("WebSocket read error", zap.Error(err))
				}
				close(messageChan)
				return
			}
			messageChan <- sub
		}
	}()

	// 主循环
	for {
		select {
		case <-ticker.C:
			// 如果有订阅的symbol，则获取最新数据
			if currentSymbol != "" && currentInterval != "" {
				klineData, err := s.getLatestKlineDataWithType(currentSymbol, currentInterval, currentType)
				if err != nil {
					global.GVA_LOG.Error("Failed to get latest kline data",
						zap.String("symbol", currentSymbol),
						zap.String("interval", currentInterval),
						zap.Int("type", currentType),
						zap.Error(err))
					continue
				}

				// 发送最新数据
				resp := AggDataResponse{
					Type:     "kline_update",
					Symbol:   currentSymbol,
					Interval: currentInterval,
					Data:     klineData,
				}
				if err := conn.WriteJSON(resp); err != nil {
					global.GVA_LOG.Error("Failed to send kline update", zap.Error(err))
					return
				}
			}

		case sub, ok := <-messageChan:
			if !ok {
				// channel已关闭，退出
				return
			}

			// 处理订阅/取消订阅
			switch sub.Action {
			case "subscribe":
				currentSymbol = sub.Symbol
				currentInterval = sub.Interval
				currentType = sub.Type
				resp := AggDataResponse{
					Type:     "subscribed",
					Symbol:   sub.Symbol,
					Interval: sub.Interval,
					Msg:      "Subscription updated",
				}
				if err := conn.WriteJSON(resp); err != nil {
					global.GVA_LOG.Error("Failed to send subscription confirmation", zap.Error(err))
					return
				}

			case "unsubscribe":
				currentSymbol = ""
				currentInterval = ""
				currentType = 0
				resp := AggDataResponse{
					Type: "unsubscribed",
					Msg:  "Unsubscribed successfully",
				}
				if err := conn.WriteJSON(resp); err != nil {
					global.GVA_LOG.Error("Failed to send unsubscription confirmation", zap.Error(err))
					return
				}
			}
		}
	}
}

// getLatestKlineDataWithType 使用传入的type获取最新的K线数据
func (s *SymbolsService) getLatestKlineDataWithType(symbol, interval string, symbolType int) (*KlineData, error) {
	// 1. 直接构建Redis key
	var baseKey string
	switch symbolType {
	case 0:
		baseKey = fmt.Sprintf("symbol:stock:%s", symbol)
	case 1:
		baseKey = fmt.Sprintf("symbol:crypto:%s", symbol)
	case 2:
		baseKey = fmt.Sprintf("symbol:forex:%s", symbol)
	default:
		return nil, fmt.Errorf("unknown symbol type: %d", symbolType)
	}

	// 2. 获取最新K线数据
	redisKey := fmt.Sprintf("%s:kline:%s", baseKey, interval)
	data, err := global.GVA_REDIS.Get(context.Background(), redisKey).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get kline data: %v", err)
	}

	// 3. 获取实时价格
	realTimePrice, err := global.GVA_REDIS.Get(context.Background(), baseKey).Result()
	if err != nil {
		global.GVA_LOG.Error("Failed to get real-time price",
			zap.String("symbol", symbol),
			zap.Error(err))
	}

	// 4. 解析K线数据
	var klineData KlineData
	if err := json.Unmarshal([]byte(data), &klineData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal kline data: %v", err)
	}

	// 5. 如果有实时价格,更新K线收盘价
	if realTimePrice != "" {
		if price, err := strconv.ParseFloat(realTimePrice, 64); err == nil {
			// 更新收盘价为最新实时价格
			klineData.Close = price

			// 如果实时价格高于最高价,更新最高价
			if price > klineData.High {
				klineData.High = price
			}

			// 如果实时价格低于最低价,更新最低价
			if price < klineData.Low {
				klineData.Low = price
			}
		}
	}

	return &klineData, nil
}

// UnifiedSubscription 统一的订阅消息结构
type UnifiedSubscription struct {
	SubType  string   `json:"sub_type"`           // 订阅类型：price/kline
	Action   string   `json:"action"`             // subscribe/unsubscribe
	Symbols  []string `json:"symbols,omitempty"`  // 价格订阅用的symbols列表
	Symbol   string   `json:"symbol,omitempty"`   // K线订阅用的单个symbol
	Interval string   `json:"interval,omitempty"` // K线订阅的时间间隔
	Type     int      `json:"type"`               // 交易对类型：0-股票 1-加密货币 2-外汇
}

// UnifiedResponse 统一的响应消息结构
type UnifiedResponse struct {
	SubType  string      `json:"sub_type"`           // 订阅类型：price/kline
	Type     string      `json:"type"`               // 消息类型：connected/subscribed/price_update/kline_update
	Symbol   string      `json:"symbol,omitempty"`   // K线数据相关的symbol
	Symbols  []string    `json:"symbols,omitempty"`  // 价格订阅相关的symbols
	Interval string      `json:"interval,omitempty"` // K线数据的时间间隔
	Data     interface{} `json:"data,omitempty"`     // 数据内容
	Msg      string      `json:"msg,omitempty"`      // 消息说明
}

// 修改 klineSubscription 结构，支持多个K线订阅
type KlineSubscription struct {
	Symbol   string
	Interval string
	Type     int
}

// HandleUnifiedWebSocket 处理统一的WebSocket连接
func (s *SymbolsService) HandleUnifiedWebSocket(conn *websocket.Conn) {
	defer conn.Close()

	// 设置连接参数
	conn.SetReadLimit(512)
	conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	// 发送连接成功消息
	resp := UnifiedResponse{
		SubType: "system",
		Type:    "connected",
		Msg:     "WebSocket connected successfully",
	}
	if err := conn.WriteJSON(resp); err != nil {
		global.GVA_LOG.Error("Failed to send connected message", zap.Error(err))
		return
	}

	// 创建定时器
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	// 创建ping定时器
	pingTicker := time.NewTicker(30 * time.Second)
	defer pingTicker.Stop()

	// 管理订阅状态
	priceSubscriptions := make(map[int]map[string]bool)      // 价格订阅
	klineSubscriptions := make(map[string]KlineSubscription) // K线订阅，key为"symbol:interval"

	// 创建消息通道和错误通道
	messageChan := make(chan UnifiedSubscription)
	errChan := make(chan error)
	done := make(chan struct{})
	defer close(done)

	// 启动goroutine处理消息读取
	go func() {
		defer close(messageChan)
		defer close(errChan)
		for {
			select {
			case <-done:
				return
			default:
				var sub UnifiedSubscription
				err := conn.ReadJSON(&sub)
				if err != nil {
					if websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
						global.GVA_LOG.Info("WebSocket closed normally")
					} else {
						global.GVA_LOG.Error("WebSocket read error", zap.Error(err))
					}
					errChan <- err
					return
				}
				messageChan <- sub
			}
		}
	}()

	// 主循环
	for {
		select {
		case <-pingTicker.C:
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				global.GVA_LOG.Error("Failed to send ping", zap.Error(err))
				return
			}

		case err := <-errChan:
			global.GVA_LOG.Error("WebSocket error received", zap.Error(err))
			return

		case <-done:
			return

		case sub, ok := <-messageChan:
			if !ok {
				return
			}

			switch sub.SubType {
			case "price":
				handlePriceSubscription(conn, sub, priceSubscriptions)
			case "kline":
				handleKlineSubscription(conn, sub, klineSubscriptions)
			}

		case <-ticker.C:
			// 更新价格数据
			if err := updatePriceData(conn, priceSubscriptions); err != nil {
				global.GVA_LOG.Error("Failed to update price data", zap.Error(err))
				return
			}

			// 更新K线数据
			if err := updateKlineData(conn, s, klineSubscriptions); err != nil {
				global.GVA_LOG.Error("Failed to update kline data", zap.Error(err))
				return
			}
		}
	}
}

// 处理价格订阅
func handlePriceSubscription(conn *websocket.Conn, sub UnifiedSubscription, priceSubscriptions map[int]map[string]bool) error {
	if len(sub.Symbols) == 0 {
		return sendResponse(conn, UnifiedResponse{
			SubType: "price",
			Type:    "error",
			Msg:     "Symbols list cannot be empty",
		})
	}

	if priceSubscriptions[sub.Type] == nil {
		priceSubscriptions[sub.Type] = make(map[string]bool)
	}

	switch sub.Action {
	case "subscribe":
		for _, symbol := range sub.Symbols {
			priceSubscriptions[sub.Type][symbol] = true
		}
		return sendResponse(conn, UnifiedResponse{
			SubType: "price",
			Type:    "subscribed",
			Symbols: sub.Symbols,
			Data:    map[string]interface{}{"type": sub.Type},
			Msg:     "Price subscription updated",
		})

	case "unsubscribe":
		if len(sub.Symbols) == 0 {
			delete(priceSubscriptions, sub.Type)
		} else {
			for _, symbol := range sub.Symbols {
				delete(priceSubscriptions[sub.Type], symbol)
			}
			if len(priceSubscriptions[sub.Type]) == 0 {
				delete(priceSubscriptions, sub.Type)
			}
		}
		return sendResponse(conn, UnifiedResponse{
			SubType: "price",
			Type:    "unsubscribed",
			Symbols: sub.Symbols,
			Msg:     "Price unsubscription confirmed",
		})
	}
	return nil
}

// 处理K线订阅
func handleKlineSubscription(conn *websocket.Conn, sub UnifiedSubscription, klineSubscriptions map[string]KlineSubscription) error {
	subscriptionKey := fmt.Sprintf("%s:%s", sub.Symbol, sub.Interval)

	switch sub.Action {
	case "subscribe":
		klineSubscriptions[subscriptionKey] = KlineSubscription{
			Symbol:   sub.Symbol,
			Interval: sub.Interval,
			Type:     sub.Type,
		}
		return sendResponse(conn, UnifiedResponse{
			SubType:  "kline",
			Type:     "subscribed",
			Symbol:   sub.Symbol,
			Interval: sub.Interval,
			Msg:      "Kline subscription updated",
		})

	case "unsubscribe":
		delete(klineSubscriptions, subscriptionKey)
		return sendResponse(conn, UnifiedResponse{
			SubType: "kline",
			Type:    "unsubscribed",
			Symbol:  sub.Symbol,
			Msg:     "Kline unsubscription confirmed",
		})
	}
	return nil
}

// 更新价格数据
func updatePriceData(conn *websocket.Conn, priceSubscriptions map[int]map[string]bool) error {
	if len(priceSubscriptions) == 0 {
		return nil
	}

	prices := make(map[string]PriceData)
	for typeID, symbols := range priceSubscriptions {
		if len(symbols) == 0 {
			continue
		}

		for symbol := range symbols {
			var currentKey, previousCloseKey string
			switch typeID {
			case 0:
				currentKey = fmt.Sprintf("symbol:stock:%s", symbol)
				previousCloseKey = fmt.Sprintf("symbol:stock:%s:previous_close", symbol)
			case 1:
				currentKey = fmt.Sprintf("symbol:crypto:%s", symbol)
				previousCloseKey = fmt.Sprintf("symbol:crypto:%s:previous_close", symbol)
			case 2:
				currentKey = fmt.Sprintf("symbol:forex:%s", symbol)
				previousCloseKey = fmt.Sprintf("symbol:forex:%s:previous_close", symbol)
			}

			if currentPrice, err := global.GVA_REDIS.Get(context.Background(), currentKey).Float64(); err == nil {
				previousClose, _ := global.GVA_REDIS.Get(context.Background(), previousCloseKey).Float64()
				changeRatio := 0.0
				if previousClose != 0 {
					changeRatio = math.Round(((currentPrice-previousClose)/previousClose*100)*100) / 100
				}
				prices[symbol] = PriceData{
					Price:       formatPrice(currentPrice, typeID, symbol),
					ChangeRatio: changeRatio,
				}
			}
		}
	}

	if len(prices) > 0 {
		return sendResponse(conn, UnifiedResponse{
			SubType: "price",
			Type:    "price_update",
			Data:    prices,
		})
	}
	return nil
}

// 更新K线数据
func updateKlineData(conn *websocket.Conn, s *SymbolsService, klineSubscriptions map[string]KlineSubscription) error {
	for _, sub := range klineSubscriptions {
		klineData, err := s.getLatestKlineDataWithType(
			sub.Symbol,
			sub.Interval,
			sub.Type,
		)
		if err != nil {
			global.GVA_LOG.Error("Failed to get kline data",
				zap.String("symbol", sub.Symbol),
				zap.Error(err))
			continue
		}

		if err := sendResponse(conn, UnifiedResponse{
			SubType:  "kline",
			Type:     "kline_update",
			Symbol:   sub.Symbol,
			Interval: sub.Interval,
			Data:     klineData,
		}); err != nil {
			return err
		}
	}
	return nil
}

// 发送WebSocket响应
func sendResponse(conn *websocket.Conn, response UnifiedResponse) error {
	return conn.WriteJSON(response)
}

// GetSymbolBySymbolName 根据symbol字符串查询详细信息
func (symbolsService *SymbolsService) GetSymbolBySymbolName(symbolName string) (*symbol.Symbols, error) {
	var symbolInfo symbol.Symbols

	// 查询数据库
	err := global.GVA_DB.Where("symbol = ?", symbolName).First(&symbolInfo).Error
	if err != nil {
		return nil, fmt.Errorf("symbol not found: %v", err)
	}

	// 如果有type字段,获取实时价格
	if symbolInfo.Type != nil {
		var redisKey string
		switch *symbolInfo.Type {
		case 0: // 股票
			redisKey = fmt.Sprintf("symbol:stock:%s", symbolName)
		case 1: // 加密货币
			redisKey = fmt.Sprintf("symbol:crypto:%s", symbolName)
		case 2: // 外汇
			redisKey = fmt.Sprintf("symbol:forex:%s", symbolName)
		}

		// 从Redis获取价格
		if price, err := global.GVA_REDIS.Get(context.Background(), redisKey).Result(); err == nil {
			priceFloat, _ := strconv.ParseFloat(price, 64)
			symbolInfo.CurrentPrice = &priceFloat
		}
	}

	return &symbolInfo, nil
}
