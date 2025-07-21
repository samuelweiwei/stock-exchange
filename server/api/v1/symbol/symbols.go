package symbol

import (
	"fmt"
	"net/http"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/symbol"
	symbolReq "github.com/flipped-aurora/gin-vue-admin/server/model/symbol/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type SymbolsApi struct{}

// CreateSymbols 创建symbols表
// @Tags Symbols
// @Summary 创建symbols表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body symbol.Symbols true "创建symbols表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /symbols/create [post]
func (symbolsApi *SymbolsApi) CreateSymbols(c *gin.Context) {
	var symbols symbol.Symbols
	err := c.ShouldBindJSON(&symbols)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = symbolsService.CreateSymbols(&symbols)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteSymbols 删除symbols表
// @Tags Symbols
// @Summary 删除symbols表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body symbol.Symbols true "删除symbols表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /symbols/delete [delete]
func (symbolsApi *SymbolsApi) DeleteSymbols(c *gin.Context) {
	id := c.Query("id")
	err := symbolsService.DeleteSymbols(id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteSymbolsByIds 批量删除symbols表
// @Tags Symbols
// @Summary 批量删除symbols表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /symbols/deleteByIds [delete]
func (symbolsApi *SymbolsApi) DeleteSymbolsByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	err := symbolsService.DeleteSymbolsByIds(ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateSymbols 更新symbols表
// @Tags Symbols
// @Summary 更新symbols表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body symbol.Symbols true "更新symbols表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /symbols/update [put]
func (symbolsApi *SymbolsApi) UpdateSymbols(c *gin.Context) {
	var symbols symbol.Symbols
	err := c.ShouldBindJSON(&symbols)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = symbolsService.UpdateSymbols(symbols)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSymbols 用id查询symbols表
// @Tags Symbols
// @Summary 用id查询symbols表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query symbol.Symbols true "用id查询symbols表"
// @Success 200 {object} response.Response{data=symbol.Symbols,msg=string} "查询成功"
// @Router /symbols/detail/{id} [get]
func (symbolsApi *SymbolsApi) FindSymbols(c *gin.Context) {
	id := c.Param("id")
	resymbols, err := symbolsService.GetSymbols(id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(resymbols, c)
}

// FindSymbols 前台-用id查询symbols表
// @Tags Symbols
// @Summary 前台-用id查询symbols表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query symbol.Symbols true "用id查询symbols表"
// @Success 200 {object} response.Response{data=symbol.Symbols,msg=string} "查询成功"
// @Router /symbols/front/detail/{id} [get]
func (symbolsApi *SymbolsApi) FrontSymbolsDetail(c *gin.Context) {
	id := c.Param("id")
	resymbols, err := symbolsService.GetSymbols(id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(resymbols, c)
}

// GetSymbolsList 分页获取symbols表列表
// @Tags Symbols
// @Summary 分页获取symbols表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query symbolReq.SymbolsSearch true "页获取symbols表列表"
// @Param symbol query string false "交易对名称"
// @Param type query int false "交易对类型：0-股票 1-加密货币 2-外汇"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /symbols/list [get]
func (symbolsApi *SymbolsApi) GetSymbolsList(c *gin.Context) {
	var pageInfo symbolReq.SymbolsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := symbolsService.GetSymbolsInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetSymbolsPublic 前台-获取公开的symbols信息
// @Tags Symbols
// @Summary 前台-获取公开的symbols信息
// @accept application/json
// @Produce application/json
// @Param data query symbolReq.SymbolsPublicSearch true "分页参数"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /symbols/public/list [get]
func (symbolsApi *SymbolsApi) GetSymbolsPublic(c *gin.Context) {
	var pageInfo symbolReq.SymbolsPublicSearch

	// 打印原始查询参数
	global.GVA_LOG.Info("原始查询参数",
		zap.String("raw_query", c.Request.URL.RawQuery))

	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		global.GVA_LOG.Error("参数绑定失败!",
			zap.Error(err),
			zap.String("raw_query", c.Request.URL.RawQuery))
		response.FailWithMessage("参数错误: "+err.Error(), c)
		return
	}

	// 打印绑定后的参数
	global.GVA_LOG.Info("绑定后的参数",
		zap.Any("pageInfo", pageInfo))

	list, total, err := symbolsService.GetSymbolsPublic(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// SubscribeSymbolPrices 前台-订阅实时价格
// @Tags PublicSymbols
// @Summary 前台-订阅实时价格推送
// @Description 建立WebSocket连接以接收实时价格推送
// @Description 1. 客户端订阅请求格式:
// @Description    {
// @Description      "req": "symbols",
// @Description      "symbols": ["AAPL","GOOGL"],
// @Description      "type": 0
// @Description    }
// @Description    type说明: 0-股票, 1-加密货币, 2-外汇
// @Description 2. 服务端响应格式:
// @Description    连接成功:
// @Description    {
// @Description      "type": "connected",
// @Description      "msg": "WebSocket connected successfully"
// @Description    }
// @Description    订阅确认:
// @Description    {
// @Description      "type": "subscribed",
// @Description      "symbols": ["AAPL","GOOGL"],
// @Description      "symbolType": "stock",
// @Description      "msg": "Subscription updated"
// @Description    }
// @Description    价格更新:
// @Description    {
// @Description      "type": "price_update",
// @Description      "data": {"AAPL": 150.25, "GOOGL": 2750.50},
// @Description      "timestamp": 1683888999
// @Description    }
// @Accept  json
// @Produce json
// @Success 101 {string} string "WebSocket协成功"
// @Failure 400 {object} response.Response{msg=string} "连接建立失败"
// @Router /symbols/ws [get]
func (s *SymbolsApi) SubscribeSymbolPrices(c *gin.Context) {
	// 升级HTTP连接为WebSocket连接
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // 开发环境允许所有来源
		},
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		global.GVA_LOG.Error("WebSocket upgrade failed:", zap.Error(err))
		return
	}

	// 调用service处理WebSocket连接
	symbolsService.SubscribeSymbolPrices(conn)
}

// AddSymbolsByName 根据股票代码从Polygon获信息并添加到数据库
// @Tags Symbols
// @Summary 根据股票代码添加新的股票信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param symbol query string true "股票代码"
// @Success 200 {object} response.Response{msg=string} "添加成功"
// @Router /symbols/addByName [post]
func (symbolsApi *SymbolsApi) AddSymbolsByName(c *gin.Context) {
	var req symbolReq.AddSymbolRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := symbolsService.AddSymbolsByName(req); err != nil {
		global.GVA_LOG.Error("添加失败!", zap.Error(err))
		response.FailWithMessage("添加失败", c)
		return
	}
	response.OkWithMessage("添加成功", c)
}

// GetSymbolPrice 前台-获取特定 symbol 的实时价格
// @Tags Symbols
// @Summary 前台-获取特定 symbol 的实时价格
// @accept application/json
// @Produce application/json
// @Param symbol path string true "股票代码"
// @Success 200 {object} response.Response{data=string,msg=string} "获取成功"
// @Router /symbols/price/{symbol} [get]
func (s *SymbolsApi) GetSymbolPrice(c *gin.Context) {
	symbol := c.Param("symbol")
	if symbol == "" {
		response.FailWithMessage("Symbol parameter is required", c)
		return
	}

	// 直接使用 symbolsService
	price, err := symbolsService.GetSymbolPrice(symbol)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("Failed to get price for symbol %s: %v", symbol, err), c)
		return
	}

	response.OkWithData(price, c)
}

// KlineRequest K线请求参数
type KlineRequest struct {
	Symbol   string `json:"symbol" binding:"required"`   // 交易对
	Interval string `json:"interval" binding:"required"` // 时间间隔
	// StartTime int64  `json:"start_time" binding:"required"` // 开始时间戳(毫秒)
	// EndTime   int64  `json:"end_time" binding:"required"`   // 结束时间戳(毫秒)
}

// GetKlineData 前台-获取K线数据
// @Tags Symbols
// @Summary 前台-获取K线数据
// @accept application/json
// @Produce application/json
// @Param data body KlineRequest true "K线请求参数"
// @Success 200 {object} response.Response{data=[]symbol.KlineData,msg=string} "获取成功"
// @Router /symbols/kline [post]
func (s *SymbolsApi) GetKlineData(c *gin.Context) {
	var req KlineRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("Invalid request parameters: "+err.Error(), c)
		return
	}

	// 验证interval是否有效
	validIntervals := map[string]bool{
		"1h": true,
		"1m": true,
		"4h": true,
		"1d": true,
		"1w": true,
		"1M": true,
	}
	if !validIntervals[req.Interval] {
		response.FailWithMessage("invalid interval", c)
		return
	}

	// 调用service层获取K线数据
	klineData, err := symbolsService.GetKlineData(req.Symbol, req.Interval)
	if err != nil {
		response.FailWithMessage("Failed to get kline data: "+err.Error(), c)
		return
	}

	response.OkWithData(klineData, c)
}

// GetAllSymbolsSimple 获取所有股票的简化信息
// @Tags Symbols
// @Summary 获取所有股票的简化信息（仅symbol和price）
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=[]map[string]interface{},msg=string} "获取成功"
// @Router /symbols/getAllSimple [get]
func (symbolsApi *SymbolsApi) GetAllSymbolsSimple(c *gin.Context) {
	list, err := symbolsService.GetAllSymbolsSimple()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(list, c)
}

// SearchTickers 查询可添加的ticker列表
// @Tags Symbols
// @Summary 查询可添加的ticker列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param type query int true "类型(0:股票 1:加密货币 2:外汇)"
// @Param search query string false "搜索关键字"
// @Success 200 {object} response.Response{data=[]map[string]interface{},msg=string} "查询成功"
// @Router /symbols/search [get]
func (symbolsApi *SymbolsApi) SearchTickers(c *gin.Context) {
	var req symbolReq.TickerSearchReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), c)
		return
	}

	list, err := symbolsService.SearchTickers(req)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败: "+err.Error(), c)
		return
	}

	response.OkWithData(list, c)
}

// GetSymbolPriceById 根据ID获取币种价格
// @Summary 根据ID获取币种价格
// @Description 通过币种ID获取当前价格
// @Tags Symbols
// @Accept json
// @Produce json
// @Param id path string true "币种ID"
// @Success 200 {object} response.Response{data=string,msg=string} "获取成功"
// @Router /symbols/public/price/{id} [get]
func (s *SymbolsApi) GetSymbolPriceById(c *gin.Context) {
	symbolId := c.Param("id")
	if symbolId == "" {
		response.FailWithMessage("ID不能为空", c)
		return
	}

	price, err := symbolsService.GetSymbolPriceById(symbolId)
	if err != nil {
		global.GVA_LOG.Error("获取价格失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(price, c)
}

// AggDataRequest K线聚合数据请求参数
type AggDataRequest struct {
	Symbol   string `json:"symbol" binding:"required"`   // 交易对
	Interval string `json:"interval" binding:"required"` // 时间间隔
}

// GetSymbolAggData 获取K线聚合数据
// @Summary 获取指定symbol的聚合数据
// @Description 获取指定symbol在指定时间间隔的K线数据(时间间隔4h/1d/1w/1M)
// @Tags Symbols
// @Accept json
// @Produce json
// @Param data body AggDataRequest true "K线请求参数"
// @Success 200 {object} response.Response{data=[]KlineData} "成功"
// @Router /symbols/agg [post]
func (symbolsApi *SymbolsApi) GetSymbolAggData(c *gin.Context) {
	var req AggDataRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), c)
		return
	}

	// 验证interval是否有效
	validIntervals := map[string]bool{
		"1h": true,
		"1m": true,
		"4h": true,
		"1d": true,
		"1w": true,
		"1M": true,
	}
	if !validIntervals[req.Interval] {
		response.FailWithMessage("无效的时间间隔", c)
		return
	}

	symbolsService := service.ServiceGroupApp.SymbolServiceGroup.SymbolsService
	klineData, err := symbolsService.GetSymbolAggData(req.Symbol, req.Interval)
	if err != nil {
		global.GVA_LOG.Error("获取K线数据失败",
			zap.String("symbol", req.Symbol),
			zap.String("interval", req.Interval),
			zap.Error(err))
		response.FailWithMessage(fmt.Sprintf("获取K线数据失败: %v", err), c)
		return
	}

	response.OkWithData(klineData, c)
}

// ValidateSymbol 验证symbol是否可用
// @Tags Symbols
// @Summary 验证symbol是否可用
// @Description 验证指定的symbol在Polygon API中是否存在且处于活跃状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body symbolReq.AddSymbolRequest true "symbol验证请求参数"
// @Success 200 {object} response.Response{msg=string} "验证成功"
// @Failure 400 {object} response.Response{msg=string} "验证失败"
// @Router /symbols/validate [post]
func (s *SymbolsApi) ValidateSymbol(c *gin.Context) {
	var req symbolReq.AddSymbolRequest

	// 绑定请求数据
	if err := c.ShouldBindJSON(&req); err != nil {
		if err2 := c.ShouldBind(&req); err2 != nil {
			response.FailWithMessage("参数错误", c)
			return
		}
	}

	// 手动验证必填字段
	if req.Symbol == "" {
		response.FailWithMessage("Symbol不能为空", c)
		return
	}

	// 验证 Type 字段
	if req.Type < 0 || req.Type > 2 {
		response.FailWithMessage("Type必须是0-2之间的值", c)
		return
	}

	if err := symbolsService.ValidateSymbol(req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("Symbol is valid", c)
}

// SubscribeSymbolAggData WebSocket订阅K线聚合数据
// @Tags Symbols
// @Summary WebSocket订阅K线聚合数据
// @Description 通过WebSocket订阅指定交易对的K线数据
// @Description 1. 客户端订阅请求格式:
// @Description    {
// @Description      "action": "subscribe",
// @Description      "symbol": "AAPL",
// @Description      "interval": "4h",
// @Description      "type": 0
// @Description    }
// @Description    type说明: 0-股票, 1-加密货币, 2-外汇
// @Description 2. 服务端响应格式:
// @Description    连接成功:
// @Description    {
// @Description      "type": "connected",
// @Description      "msg": "WebSocket connected successfully"
// @Description    }
// @Description    订阅确认:
// @Description    {
// @Description      "type": "subscribed",
// @Description      "symbol": "AAPL",
// @Description      "interval": "4h",
// @Description      "msg": "Subscription updated"
// @Description    }
// @Description    K线更新:
// @Description    {
// @Description      "type": "kline_update",
// @Description      "symbol": "AAPL",
// @Description      "interval": "4h",
// @Description      "data": {
// @Description        "o": 150.25,
// @Description        "h": 151.20,
// @Description        "l": 149.80,
// @Description        "c": 150.50,
// @Description        "v": 1000000,
// @Description        "t": 1683888999000
// @Description      }
// @Description    }
// @Accept  json
// @Produce json
// @Success 101 {string} string "WebSocket连接成功"
// @Router /symbols/agg/ws [get]
func (s *SymbolsApi) SubscribeSymbolAggData(c *gin.Context) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		global.GVA_LOG.Error("WebSocket upgrade failed:", zap.Error(err))
		return
	}

	// 调用service处理WebSocket连接
	symbolsService.SubscribeSymbolAggData(conn)
}

// UnifiedWebSocket 统一的WebSocket接口
// @Tags Symbols
// @Summary 统一的WebSocket数据订阅接口
// @Description 通过单一WebSocket连接订阅价格和K线数据
// @Description 1. 订阅价格数据请求格式:
// @Description    {
// @Description      "sub_type": "price",
// @Description      "action": "subscribe",
// @Description      "symbols": ["AAPL","GOOGL"],
// @Description      "type": 0
// @Description    }
// @Description 2. 订阅K线数据请求格式:
// @Description    {
// @Description      "sub_type": "kline",
// @Description      "action": "subscribe",
// @Description      "symbol": "AAPL",
// @Description      "interval": "4h",
// @Description      "type": 0
// @Description    }
// @Accept  json
// @Produce json
// @Success 101 {string} string "WebSocket连接成功"
// @Router /symbols/unified/ws [get]
func (s *SymbolsApi) UnifiedWebSocket(c *gin.Context) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		global.GVA_LOG.Error("WebSocket upgrade failed:", zap.Error(err))
		return
	}

	symbolsService.HandleUnifiedWebSocket(conn)
}

// GetSymbolBySymbolName 根据symbol查询详细信息
// @Tags Symbols
// @Summary 根据symbol查询详细信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body symbolReq.SymbolNameRequest true "交易对代码"
// @Success 200 {object} response.Response{data=symbol.Symbols,msg=string} "获取成功"
// @Router /symbols/front/symbol [post]
func (symbolsApi *SymbolsApi) GetSymbolBySymbolName(c *gin.Context) {
	var req symbolReq.SymbolNameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), c)
		return
	}

	if req.Symbol == "" {
		response.FailWithMessage("symbol参数不能为空", c)
		return
	}

	symbolInfo, err := symbolsService.GetSymbolBySymbolName(req.Symbol)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithData(symbolInfo, c)
}
