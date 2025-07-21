package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	userReq "github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FrontendUserLoginLogApi struct{}

// GetFrontendUserLoginLogList 分页获取frontendUserLoginLog表列表
// @Tags FrontendUserLoginLog
// @Summary 分页获取frontendUserLoginLog表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userReq.FrontendUserLoginLogSearch true "分页获取frontendUserLoginLog表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /frontendUserLoginLog/getFrontendUserLoginLogList [get]
func (frontendUserLoginLogApi *FrontendUserLoginLogApi) GetFrontendUserLoginLogList(c *gin.Context) {
	var pageInfo userReq.FrontendUserLoginLogSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := frontendUserLoginLogService.GetFrontendUserLoginLogInfoList(pageInfo)
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
