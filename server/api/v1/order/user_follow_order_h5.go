package order

import (
	"github.com/flipped-aurora/gin-vue-admin/server/i18n"
	cRequest "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserFollowOrderH5Api struct{}

// Apply 	  用户申请跟单
//
//	@Tags		UserFollowOrderH5
//	@Summary	前端-用户申请跟单
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		request.UserFollowOrderApplyReq	true	"请求参数"
//	@Success	200		{object}	response.Response{msg=string}	"请求成功"
//	@Router		/user/follow/order/apply [post]
func (api *UserFollowOrderH5Api) Apply(c *gin.Context) {
	var req request.UserFollowOrderApplyReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	req.UserId = utils.GetUserIDFrontUser(c)

	err = userFollowOrderService.Apply(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}
	response.OkWithMessage(i18n.Success, c)
}

// Cancel 	  用户撤回跟单
//
//	@Tags		UserFollowOrderH5
//	@Summary	前端-用户撤回跟单
//	@Security	ApiKeyAuth
//	@Produce	application/json
//	@Param		followOrderId	path		int								true	"跟单ID"
//	@Success	200				{object}	response.Response{msg=string}	"请求成功"
//	@Router		/user/follow/order/cancel/:followOrderId [post]
func (api *UserFollowOrderH5Api) Cancel(c *gin.Context) {
	followOrderId, err := strconv.Atoi(c.Param("followOrderId"))
	if err != nil {
		_ = c.Error(err)
		return
	}

	err = userFollowOrderService.Cancel(uint(followOrderId), utils.GetUserIDFrontUser(c))
	if err != nil {
		_ = c.Error(err)
		return
	}

	response.OkWithMessage(i18n.Success, c)
}

// PageQueryMyFollowOrder  分页查询我的跟单列表
//
//	@Tags		UserFollowOrderH5
//	@Summary	前端-分页查询我的跟单列表
//	@Security	ApiKeyAuth
//	@Produce	application/json
//	@Param		data	query		request.MyFollowOrderPageReq																	true	"查询参数"
//	@Success	200		{object}	response.Response{msg=string, data=response.PageResult{list=[]response.MyFollowOrderPageData}}	"查询成功"
//	@Router		/user/follow/order/page/my [get]
func (api *UserFollowOrderH5Api) PageQueryMyFollowOrder(c *gin.Context) {
	var req request.MyFollowOrderPageReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	req.UserId = utils.GetUserIDFrontUser(c)
	list, total, err := userFollowOrderService.PageQueryMyFollowOrder(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	lan := cRequest.GetLanguageTag(c)
	for _, v := range list {
		v.AdvisorName = i18n.Message(lan, v.AdvisorName, 0)
		v.ProductName = i18n.Message(lan, v.ProductName, 0)
	}

	response.OkWithDetailed(response.PageResult{
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    total,
		List:     list,
	}, i18n.Success, c)
}

// GetMyFollowOrderDetail 	  查询我的跟单详情
//
//	@Tags		UserFollowOrderH5
//	@Summary	前端-查询我的跟单详情
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		followOrderId	path		int																true	"跟单订单ID"
//	@Success	200				{object}	response.Response{msg=string,data=response.MyFollowOrderDetail}	"请求成功"
//	@Router		/user/follow/order/my/:followOrderId [get]
func (api *UserFollowOrderH5Api) GetMyFollowOrderDetail(c *gin.Context) {
	followOrderId, err := strconv.Atoi(c.Param("followOrderId"))
	if err != nil {
		_ = c.Error(err)
		return
	}

	userId := utils.GetUserIDFrontUser(c)
	followOrderDetail, err := userFollowOrderService.GetMyFollowOrderDetail(uint(followOrderId), userId)

	if err != nil {
		_ = c.Error(err)
		return
	}

	lan := cRequest.GetLanguageTag(c)
	followOrderDetail.AdvisorName = i18n.Message(lan, followOrderDetail.AdvisorName, 0)
	followOrderDetail.ProductName = i18n.Message(lan, followOrderDetail.ProductName, 0)

	response.OkWithDetailed(&followOrderDetail, i18n.Success, c)
}

// RetrieveFollowOrder 用户跟单提盈
//
//	@Tags		UserFollowOrderH5
//	@Summary	前端-用户跟单提盈
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		followOrderId	path		int																true	"跟单订单ID"
//	@Param		data			body		request.RetrieveFollowOrderReq									true	"请求参数"
//	@Success	200				{object}	response.Response{msg=string,data=response.MyFollowOrderDetail}	"请求成功"
//	@Router		/user/follow/order/retrieve/:followOrderId [post]
func (api *UserFollowOrderH5Api) RetrieveFollowOrder(c *gin.Context) {
	followOrderId, err := strconv.Atoi(c.Param("followOrderId"))
	if err != nil {
		_ = c.Error(err)
		return
	}

	var req request.RetrieveFollowOrderReq
	if err = c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		return
	}

	req.UserId, req.FollowOrderId = utils.GetUserIDFrontUser(c), uint(followOrderId)
	err = userFollowOrderService.RetrieveFollowOrder(&req)

	if err != nil {
		_ = c.Error(err)
		return
	}

	response.OkWithMessage(i18n.Success, c)
}

// CancelAutoRenew 用户取消自动续期
//
//	@Tags		UserFollowOrderH5
//	@Summary	前端-用户取消自动续期
//	@Security	ApiKeyAuth
//	@Produce	application/json
//	@Param		followOrderId	path		int								true	"跟单订单ID"
//	@Success	200				{object}	response.Response{msg=string}	"请求成功"
//	@Router		/user/follow/order/cancel-renew/:followOrderId [post]
func (api *UserFollowOrderH5Api) CancelAutoRenew(c *gin.Context) {
	followOrderId, err := strconv.Atoi(c.Param("followOrderId"))
	if err != nil {
		_ = c.Error(err)
		return
	}

	err = userFollowOrderService.CancelAutoRenew(uint(followOrderId), utils.GetUserIDFrontUser(c))
	if err != nil {
		_ = c.Error(err)
		return
	}

	response.OkWithMessage(i18n.Success, c)
}

// GetUserProfitSummary 查询用户收益报表
//
//	@Tags		UserFollowOrderH5
//	@Summary	前端-查询用户收益报表
//	@Security	ApiKeyAuth
//	@Produce	application/json
//	@Success	200	{object}	response.Response{msg=string,data=response.UserFollowOrderProfitSummary}	"请求成功"
//	@Router		/user/follow/order/profit/my [get]
func (api *UserFollowOrderH5Api) GetUserProfitSummary(c *gin.Context) {
	userId := utils.GetUserIDFrontUser(c)
	resp := userFollowOrderService.GetUserProfitSummary(userId)
	response.OkWithDetailed(resp, i18n.Success, c)
}

// GetTotalAmount 查询用户跟单总金额
//
//	@Tags		UserFollowOrderH5
//	@Summary	前端-查询用户跟单总金额
//	@Security	ApiKeyAuth
//	@Produce	application/json
//	@Success	200	{object}	response.Response{msg=string,data=float64}	"请求成功"
//	@Router		/user/follow/order/total-amount/my [get]
func (api *UserFollowOrderH5Api) GetTotalAmount(c *gin.Context) {
	userId := utils.GetUserIDFrontUser(c)
	totalFollowAmount := userFollowOrderService.GetTotalAmount(userId)
	response.OkWithDetailed(totalFollowAmount, i18n.Success, c)
}

// PageQueryRetrieveRecord 分页查询用户提盈记录
//
//	@Tags		UserFollowOrderH5
//	@Summary	前端-分页查询用户提盈记录
//	@Security	ApiKeyAuth
//	@Produce	application/json
//	@Param		followOrderId	path		int																										true	"跟单订单ID"
//	@Param		data			query		request.UserFollowOrderRetrieveRecordPageReq															true	"查询参数"
//	@Success	200				{object}	response.Response{msg=string,data=response.PageResult{list=[]response.UserFollowOrderRetrieveRecord}}	"请求成功"
//	@Router		/user/follow/order/retrieve/records/:followOrderId [get]
func (api *UserFollowOrderH5Api) PageQueryRetrieveRecord(c *gin.Context) {
	var req request.UserFollowOrderRetrieveRecordPageReq
	if err := c.ShouldBindQuery(&req); err != nil {
		_ = c.Error(err)
		return
	}

	followOrderId, err := strconv.Atoi(c.Param("followOrderId"))
	if err != nil {
		_ = c.Error(err)
		return
	}

	req.UserId = utils.GetUserIDFrontUser(c)
	req.FollowOrderId = uint(followOrderId)

	records, total, err := userFollowOrderService.PageQueryRetrieveRecord(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	response.OkWithDetailed(response.PageResult{
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    total,
		List:     records,
	}, i18n.Success, c)
}
