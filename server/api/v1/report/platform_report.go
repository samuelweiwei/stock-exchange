package report

import (
	"github.com/flipped-aurora/gin-vue-admin/server/i18n"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/report/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
)

type PlatformReportApi struct{}

// GetPlatformReportSummary 查询平台固定报表
//
//	@Tags		PlatformReport
//	@Summary	后端-查询平台固定报表
//	@Security	ApiKeyAuth
//	@Produce	application/json
//	@Success	200	{object}	response.Response{msg=string,data=response.PlatformReportSummary}	"请求成功"
//	@Router		/report/summary [get]
func (a *PlatformReportApi) GetPlatformReportSummary(c *gin.Context) {
	rootUserId := utils.GetUserInfo(c).FrontUserId
	resp, err := platformReportService.GetPlatformReportSummary(rootUserId)
	if err != nil {
		_ = c.Error(err)
		return
	}

	response.OkWithDetailed(resp, i18n.Success, c)
}

// PageQueryPlatformReportDaily 分页查询平台每日报表
//
//	@Tags		PlatformReport
//	@Summary	后端-分页查询平台每日报表
//	@Security	ApiKeyAuth
//	@Produce	application/json
//	@Param		data	query		request.PlatformReportDailyPageQueryReq														true	"查询参数"
//	@Success	200		{object}	response.Response{msg=string,data=response.PageResult{list=[]response.PlatformReportDaily}}	"请求成功"
//	@Router		/report/daily/page [get]
func (a *PlatformReportApi) PageQueryPlatformReportDaily(c *gin.Context) {
	var req request.PlatformReportDailyPageQueryReq
	if err := c.ShouldBindQuery(&req); err != nil {
		_ = c.Error(err)
		return
	}

	req.RootUserId = utils.GetUserInfo(c).FrontUserId
	list, total, err := platformReportService.PageQueryPlatformReportDaily(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	response.OkWithDetailed(response.PageResult{
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    total,
		List:     list,
	}, i18n.Success, c)
}

// GetPlatformDynamicReport 查询平台动态报表
//
//	@Tags		PlatformReport
//	@Summary	后端-查询平台动态报表
//	@Security	ApiKeyAuth
//	@Produce	application/json
//	@Param		data	query		request.PlatformReportDailyGetReq									true	"查询参数"
//	@Success	200		{object}	response.Response{msg=string,data=response.PlatformDynamicReport}	"请求成功"
//	@Router		/report/dynamic [get]
func (a *PlatformReportApi) GetPlatformDynamicReport(c *gin.Context) {
	var req request.PlatformReportDailyGetReq
	if err := c.ShouldBindQuery(&req); err != nil {
		_ = c.Error(err)
		return
	}

	req.RootUserId = utils.GetUserInfo(c).FrontUserId
	report, err := platformReportService.GetPlatformDynamicReport(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	response.OkWithDetailed(report, i18n.Success, c)
}
