package order

import (
	"github.com/flipped-aurora/gin-vue-admin/server/i18n"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
)

type UserProfitShareApi struct {
}

// QueryMyProfitShareRecord  查询我的分成盈利情况
//
//	@Tags		UserProfitShare
//	@Summary	前端-查询我的分成盈利情况
//	@Security	ApiKeyAuth
//	@Produce	application/json
//	@param		data	query		request.MyUserProfitShareRecordPageQueryReq														true	"请求参数"
//	@Success	200		{object}	response.Response{msg=string,data=response.PageResult{list=[]response.MyUserProfitShareRecord}}	"请求成功"
//	@Router		/user/profit/share/my [get]
func (api *UserProfitShareApi) QueryMyProfitShareRecord(c *gin.Context) {
	var req request.MyUserProfitShareRecordPageQueryReq
	if err := c.ShouldBindQuery(&req); err != nil {
		_ = c.Error(err)
		return
	}

	req.UserId = utils.GetUserIDFrontUser(c)
	records, total, err := userProfitShareService.QueryMyProfitShareRecord(&req)

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
