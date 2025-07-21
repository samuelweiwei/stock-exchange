package order

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/enums/fund"
	"github.com/flipped-aurora/gin-vue-admin/server/errorx"
	"github.com/flipped-aurora/gin-vue-admin/server/i18n"
	cRequest "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserFollowAppendOrderApi struct {
}

// QueryMyAppendOrderList 查询我的追单列表
//
//	@Tags		UserFollowAppendOrder
//	@Summary	前端-查询我的追单列表
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		followOrderId	path		int																			true	"跟单订单ID"
//	@Success	200				{object}	response.Response{msg=string,data=[]response.MyFollowAppendOrderListData}	"请求成功"
//	@Router		/user/append/orders/my/:followOrderId [get]
func (api *UserFollowAppendOrderApi) QueryMyAppendOrderList(c *gin.Context) {
	followOrderId, err := strconv.Atoi(c.Param("followOrderId"))
	if err != nil {
		_ = c.Error(err)
		return
	}

	userId := utils.GetUserIDFrontUser(c)
	list, err := userFollowAppendOrderService.QueryMyFollowAppendOrderList(uint(followOrderId), userId)
	if err != nil {
		_ = c.Error(err)
		return
	}

	lan := cRequest.GetLanguageTag(c)
	for _, v := range list {
		v.AppendOrderStatusText = i18n.Message(lan, i18n.AppendOrderStatusPrefix+strconv.Itoa(int(v.AppendOrderStatus)), 0)
	}

	response.OkWithDetailed(list, i18n.Success, c)
}

// PageQuery 分页查询用户追单列表
//
//	@Tags		UserFollowAppendOrder
//	@Summary	后端-分页查询用户追单列表
//	@Security	ApiKeyAuth
//	@Produce	application/json
//	@Param		data	query		request.UserFollowAppendOrderPageQueryReq																true	"请求参数"
//	@Success	200		{object}	response.Response{msg=string,data=response.PageResult{list=[]response.UserFollowAppendOrderPageData}}	"请求成功"
//	@Router		/user/append/order/page [get]
func (api *UserFollowAppendOrderApi) PageQuery(c *gin.Context) {
	var req request.UserFollowAppendOrderPageQueryReq
	if err := c.ShouldBindQuery(&req); err != nil {
		_ = c.Error(err)
		return
	}

	req.RootUserId = utils.GetUserInfo(c).FrontUserId
	list, total, err := userFollowAppendOrderService.PageQuery(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	lan := cRequest.GetLanguageTag(c)
	for _, v := range list {
		v.AppendOrderStatusText = i18n.Message(lan, i18n.AppendOrderStatusPrefix+strconv.Itoa(int(v.AppendOrderStatus)), 0)
		v.UserTypeText = i18n.Message(lan, i18n.UserTypePrefix+strconv.Itoa(int(v.UserType)), 0)
	}

	response.OkWithDetailed(response.PageResult{
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    total,
		List:     list,
	}, i18n.Success, c)
}

// Apply 申请追单
//
//	@Tags		UserFollowAppendOrder
//	@Summary	前端-申请追单
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		followOrderId	path		int										true	"跟单订单ID"
//	@Param		data			body		request.UserFollowAppendOrderApplyReq	true	"请求参数"
//	@Success	200				{object}	response.Response{msg=string}			"请求成功"
//	@Router		/user/append/order/:followOrderId [post]
func (api *UserFollowAppendOrderApi) Apply(c *gin.Context) {
	followOrderId, err := strconv.Atoi(c.Param("followOrderId"))
	if err != nil {
		_ = c.Error(err)
		return
	}

	var req request.UserFollowAppendOrderApplyReq
	if err = c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		return
	}

	req.FollowOrderId, req.UserId = uint(followOrderId), utils.GetUserIDFrontUser(c)
	err = userFollowAppendOrderService.Apply(&req)

	if err != nil {
		if errors.Is(err, fund.BalanceDoesNotEnoughError) {
			_ = c.Error(errorx.New(response.ERROR, i18n.BalanceDoseNotEnough))
		} else {
			_ = c.Error(err)
		}
		return
	}

	response.OkWithMessage(i18n.Success, c)
}

// Approve 审核通过追单
//
//	@Tags		UserFollowAppendOrder
//	@Summary	后端-审核通过追单
//	@Security	ApiKeyAuth
//	@Produce	application/json
//	@Param		appendOrderId	path		int								true	"追单订单ID"
//	@Success	200				{object}	response.Response{msg=string}	"请求成功"
//	@Router		/user/append/order/approve/:appendOrderId [post]
func (api *UserFollowAppendOrderApi) Approve(c *gin.Context) {
	appendOrderId, err := strconv.Atoi(c.Param("appendOrderId"))
	if err != nil {
		_ = c.Error(err)
		return
	}

	err = userFollowAppendOrderService.Approve(uint(appendOrderId))
	if err != nil {
		if errors.Is(err, fund.BalanceDoesNotEnoughError) {
			_ = c.Error(errorx.New(response.ERROR, i18n.BalanceDoseNotEnough))
		} else {
			_ = c.Error(err)
		}
		return
	}

	response.OkWithMessage(i18n.Success, c)
}

// Reject 审核驳回追单
//
//	@Tags		UserFollowAppendOrder
//	@Summary	后端-审核驳回追单
//	@Security	ApiKeyAuth
//	@Produce	application/json
//	@Param		appendOrderId	path		int								true	"追单订单ID"
//	@Success	200				{object}	response.Response{msg=string}	"请求成功"
//	@Router		/user/append/order/reject/:appendOrderId [post]
func (api *UserFollowAppendOrderApi) Reject(c *gin.Context) {
	appendOrderId, err := strconv.Atoi(c.Param("appendOrderId"))
	if err != nil {
		_ = c.Error(err)
		return
	}

	err = userFollowAppendOrderService.Reject(uint(appendOrderId))
	if err != nil {
		_ = c.Error(err)
		return
	}
	response.OkWithMessage(i18n.Success, c)
}
