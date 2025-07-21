package order

import (
	"github.com/flipped-aurora/gin-vue-admin/server/i18n"
	cRequest "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order/request"
	"github.com/gin-gonic/gin"
	"strconv"
)

type AdvisorStockOrderApi struct {
}

// CreateStockOrder 创建导师开单记录
//
//	@Tags		AdvisorStockOrder
//	@Summary	后端-创建导师开单记录
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		request.AdvisorStockOrderCreateReq	true	"创建参数"
//	@Success	200		{object}	response.Response{msg=string}		"创建成功"
//	@Router		/advisor/stock/order [post]
func (api *AdvisorStockOrderApi) CreateStockOrder(c *gin.Context) {
	var req request.AdvisorStockOrderCreateReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	err = advisorStockOrderService.CreateStockOrder(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	response.OkWithMessage(i18n.Success, c)
}

// PageQuery  分页查询导师开单记录
//
//	@Tags		AdvisorStockOrder
//	@Summary	后端-分页查询导师开单记录
//	@Security	ApiKeyAuth
//	@Produce	application/json
//	@Param		data	query		request.AdvisorStockOrderPageQueryReq																true	"查询参数"
//	@Success	200		{object}	response.Response{msg=string, data=response.PageResult{list=[]response.AdvisorStockOrderPageData}}	"查询成功"
//	@Router		/advisor/stock/order/page [get]
func (api *AdvisorStockOrderApi) PageQuery(c *gin.Context) {
	var req request.AdvisorStockOrderPageQueryReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	list, total, err := advisorStockOrderService.PageQuery(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	lan := cRequest.GetLanguageTag(c)
	for _, v := range list {
		v.StatusText = i18n.Message(lan, i18n.AdvisorStockOrderStatusPrefix+strconv.Itoa(int(v.Status)), 0)
	}

	response.OkWithDetailed(response.PageResult{
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    total,
		List:     list,
	}, i18n.Success, c)
}

// GetSellConfirmDetail  查询导师开单卖出提交确认详情
//
//	@Tags		AdvisorStockOrder
//	@Summary	后端-查询导师开单卖出提交确认详情
//	@Security	ApiKeyAuth
//	@Produce	application/json
//	@Param		stockOrderId	path		int																				true	"导师开单记录ID"
//	@Param		data			query		request.AdvisorStockOrderSellConfirmReq											true	"请求参数"
//	@Success	200				{object}	response.Response{msg=string, data=response.AdvisorStockOrderConfirmSummary}	"请求成功"
//	@Router		/advisor/stock/order/confirm/:stockOrderId [get]
func (api *AdvisorStockOrderApi) GetSellConfirmDetail(c *gin.Context) {
	stockOrderId, err := strconv.Atoi(c.Param("stockOrderId"))
	if err != nil {
		_ = c.Error(err)
		return
	}

	var req request.AdvisorStockOrderSellConfirmReq
	if err = c.ShouldBindQuery(&req); err != nil {
		_ = c.Error(err)
		return
	}

	req.StockOrderId = uint(stockOrderId)

	summary, err := advisorStockOrderService.GetSellConfirmSummary(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	response.OkWithDetailed(summary, i18n.Success, c)
}

// SubmitSellConfirm  卖出提交确认详情
//
//	@Tags		AdvisorStockOrder
//	@Summary	后端-卖出提交确认详情
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		stockOrderId	path		int												true	"导师开单记录ID"
//	@Param		data			body		request.AdvisorStockOrderSellConfirmSubmitReq	true	"请求参数"
//	@Success	200				{object}	response.Response{msg=string}					"请求成功"
//	@Router		/advisor/stock/order/confirm/:stockOrderId [post]
func (api *AdvisorStockOrderApi) SubmitSellConfirm(c *gin.Context) {
	stockOrderId, err := strconv.Atoi(c.Param("stockOrderId"))
	if err != nil {
		_ = c.Error(err)
		return
	}

	var req request.AdvisorStockOrderSellConfirmSubmitReq
	if err = c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		return
	}

	req.StockOrderId = uint(stockOrderId)
	err = advisorStockOrderService.SubmitSellConfirm(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	response.OkWithMessage(i18n.Success, c)
}

// AutoFollow 一键跟单
//
//	@Tags		AdvisorStockOrder
//	@Summary	后端-一键跟单
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		stockOrderId	path		int										true	"导师开单记录ID"
//	@Param		data			body		request.AdvisorStockOrderAutoFollowReq	true	"请求参数"
//	@Success	200				{object}	response.Response{msg=string}			"请求成功"
//	@Router		/advisor/stock/order/auto-follow/:stockOrderId [post]
func (api *AdvisorStockOrderApi) AutoFollow(c *gin.Context) {
	stockOrderId, err := strconv.Atoi(c.Param("stockOrderId"))
	if err != nil {
		_ = c.Error(err)
		return
	}

	var req request.AdvisorStockOrderAutoFollowReq
	if err = c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		return
	}

	req.StockOrderId = uint(stockOrderId)
	err = advisorStockOrderService.AutoFollow(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	response.OkWithMessage(i18n.Success, c)
}
