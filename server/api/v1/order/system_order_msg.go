package order

import (
	"github.com/flipped-aurora/gin-vue-admin/server/i18n"
	cRequest "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order/request"
	"github.com/gin-gonic/gin"
	"strconv"
)

type SystemOrderMsgApi struct {
}

// GetUnReadCount 				查询系统订单站内信未读数量
//
//	@Tags		SystemOrderMsg
//	@Summary	后端-查询系统订单站内信未读数量
//	@Security	ApiKeyAuth
//	@Produce	application/json
//	@Success	200	{object}	response.Response{msg=string,data=int64}	"请求成功"
//	@Router		/system/order/msg/unread/count [get]
func (api *SystemOrderMsgApi) GetUnReadCount(c *gin.Context) {

	count, err := systemOrderMsgService.GetUnReadCount()
	if err != nil {
		_ = c.Error(err)
		return
	}

	response.OkWithDetailed(count, i18n.Success, c)
}

// PageQuery 				    分页查询系统订单站内信列表
//
//	@Tags		SystemOrderMsg
//	@Summary	后端-分页查询系统订单站内信列表
//	@Security	ApiKeyAuth
//	@Produce	application/json
//	@Param		data	query		request.SystemOrderMsgPageQueryReq																true	"请求参数"
//	@Success	200		{object}	response.Response{msg=string,data=response.PageResult{list=[]response.SystemOrderMsgPageData}}	"请求成功"
//	@Router		/system/order/msg/page [get]
func (api *SystemOrderMsgApi) PageQuery(c *gin.Context) {
	var req request.SystemOrderMsgPageQueryReq
	if err := c.ShouldBindQuery(&req); err != nil {
		_ = c.Error(err)
		return
	}

	list, total, err := systemOrderMsgService.PageQuery(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	lan := cRequest.GetLanguageTag(c)
	for _, v := range list {
		v.TypeText = i18n.Message(lan, i18n.MsgTypePrefix+strconv.Itoa(int(v.Type)), 0)
		v.ReadStatusText = i18n.Message(lan, i18n.ReadStatusPrefix+strconv.Itoa(int(v.ReadStatus)), 0)
	}

	response.OkWithDetailed(&response.PageResult{
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    total,
		List:     list,
	}, i18n.Success, c)
}

// SetReadStatus 			设置站内信已读状态
//
//	@Tags		SystemOrderMsg
//	@Summary	后端-设置站内信已读状态
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		request.SystemOrderMsgSetReadStatusReq	true	"请求参数"
//	@Success	200		{object}	response.Response{msg=string}			"请求成功"
//	@Router		/system/order/msg/read [post]
func (api *SystemOrderMsgApi) SetReadStatus(c *gin.Context) {
	var req request.SystemOrderMsgSetReadStatusReq
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		return
	}

	err := systemOrderMsgService.SetMessagesRead(req.MsgIds)
	if err != nil {
		_ = c.Error(err)
		return
	}
	response.OkWithMessage(i18n.Success, c)
}
