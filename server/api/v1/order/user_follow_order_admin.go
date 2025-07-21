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

type UserFollowOrderAdminApi struct {
}

// PageQuery  分页查询用户跟单列表
//
//	@Tags		UserFollowOrderAdmin
//	@Summary	后端-分页查询用户跟单列表
//	@Security	ApiKeyAuth
//	@Produce	application/json
//	@Param		data	query		request.UserFollowOrderPageQueryReq																	true	"查询参数"
//	@Success	200		{object}	response.Response{msg=string, data=response.PageResult{list=[]response.UserFollowOrderPageData}}	"查询成功"
//	@Router		/user/follow/order/page [get]
func (api *UserFollowOrderAdminApi) PageQuery(c *gin.Context) {
	var req request.UserFollowOrderPageQueryReq
	if err := c.ShouldBindQuery(&req); err != nil {
		_ = c.Error(err)
		return
	}

	req.RootUserId = utils.GetUserInfo(c).FrontUserId
	list, total, err := userFollowOrderService.PageQuery(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	lan := cRequest.GetLanguageTag(c)
	for _, v := range list {
		v.FollowOrderStatusText = i18n.Message(lan, i18n.FollowOrderStatusPrefix+strconv.Itoa(int(v.FollowOrderStatus)), 0)
		v.StockStatusText = i18n.Message(lan, i18n.FollowOrderStockStatusPrefix+strconv.Itoa(int(v.StockStatus)), 0)
		v.UserTypeText = i18n.Message(lan, i18n.UserTypePrefix+strconv.Itoa(int(v.UserType)), 0)
	}

	response.OkWithDetailed(response.PageResult{
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    total,
		List:     list,
	}, i18n.Success, c)
}

// Approve    用户跟单审核通过
//
//	@Tags		UserFollowOrderAdmin
//	@Summary	后端-用户跟单审核通过
//	@Security	ApiKeyAuth
//	@Produce	application/json
//	@Param		followOrderId	path		int								true	"用户跟单ID"
//	@Success	200				{object}	response.Response{msg=string}	"请求成功"
//	@Router		/user/follow/order/approve/:followOrderId [post]
func (api *UserFollowOrderAdminApi) Approve(c *gin.Context) {
	followOrderId, err := strconv.Atoi(c.Param("followOrderId"))
	if err != nil {
		_ = c.Error(err)
		return
	}

	err = userFollowOrderService.Approve(uint(followOrderId))
	if err != nil {
		_ = c.Error(err)
		return
	}

	response.OkWithMessage(i18n.Success, c)
}

// ApproveAll   一键审核用户跟单
//
//	@Tags		UserFollowOrderAdmin
//	@Summary	后端-一键审核用户跟单
//	@Security	ApiKeyAuth
//	@Produce	application/json
//	@Success	200	{object}	response.Response{msg=string}	"请求成功"
//	@Router		/user/follow/order/approve-all [post]
func (api *UserFollowOrderAdminApi) ApproveAll(c *gin.Context) {
	err := userFollowOrderService.ApproveAll()
	if err != nil {
		_ = c.Error(err)
		return
	}
	response.OkWithMessage(i18n.Success, c)
}

// Reject     用户跟单审核驳回
//
//	@Tags		UserFollowOrderAdmin
//	@Summary	后端-用户跟单审核驳回
//	@Security	ApiKeyAuth
//	@Produce	application/json
//	@Param		followOrderId	path		int								true	"用户跟单ID"
//	@Success	200				{object}	response.Response{msg=string}	"请求成功"
//	@Router		/user/follow/order/reject/:followOrderId [post]
func (api *UserFollowOrderAdminApi) Reject(c *gin.Context) {
	followOrderId, err := strconv.Atoi(c.Param("followOrderId"))
	if err != nil {
		_ = c.Error(err)
		return
	}

	err = userFollowOrderService.Reject(uint(followOrderId))
	if err != nil {
		_ = c.Error(err)
		return
	}

	response.OkWithMessage(i18n.Success, c)
}

// GetFollowConfirmDetail    用户跟单确认详情
//
//	@Tags		UserFollowOrderAdmin
//	@Summary	后端-用户跟单确认详情
//	@Security	ApiKeyAuth
//	@Produce	application/json
//	@Param		followOrderId	path		int																	true	"用户跟单ID"
//	@Success	200				{object}	response.Response{msg=string,data=response.UserFollowConfirmDetail}	"请求成功"
//	@Router		/user/follow/order/confirm/:followOrderId [get]
func (api *UserFollowOrderAdminApi) GetFollowConfirmDetail(c *gin.Context) {
	followOrderId, err := strconv.Atoi(c.Param("followOrderId"))
	if err != nil {
		_ = c.Error(err)
		return
	}

	detail, err := userFollowOrderService.GetFollowConfirmDetail(uint(followOrderId))
	if err != nil {
		_ = c.Error(err)
		return
	}

	response.OkWithDetailed(&detail, i18n.Success, c)
}

// SubmitFollowConfirm    	 用户跟单确认提交
//
//	@Tags		UserFollowOrderAdmin
//	@Summary	后端-用户跟单确认提交
//	@Security	ApiKeyAuth
//	@Produce	application/json
//	@Param		followOrderId	path		int										true	"用户跟单ID"
//	@Param		data			body		request.UserFollowOrderConfirmSubmitReq	true	"请求参数"
//	@Success	200				{object}	response.Response{msg=string}			"请求成功"
//	@Router		/user/follow/order/confirm/:followOrderId [post]
func (api *UserFollowOrderAdminApi) SubmitFollowConfirm(c *gin.Context) {
	followOrderId, err := strconv.Atoi(c.Param("followOrderId"))
	if err != nil {
		_ = c.Error(err)
		return
	}

	var req request.UserFollowOrderConfirmSubmitReq
	if err = c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		return
	}

	req.FollowOrderId = uint(followOrderId)
	err = userFollowOrderService.SubmitFollowConfirm(&req)

	if err != nil {
		_ = c.Error(err)
		return
	}

	response.OkWithMessage(i18n.Success, c)
}
