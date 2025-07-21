package order

import (
	"github.com/flipped-aurora/gin-vue-admin/server/i18n"
	cRequest "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order/request"
	_ "github.com/flipped-aurora/gin-vue-admin/server/model/order/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

type AdvisorApi struct{}

// CreateAdvisor 创建导师
//
//	@Tags		Advisor
//	@Summary	后端-创建导师
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		request.AdvisorCreateReq		true	"创建参数"
//	@Success	200		{object}	response.Response{msg=string}	"创建成功"
//	@Router		/advisor [post]
func (advisorApi *AdvisorApi) CreateAdvisor(c *gin.Context) {
	var req request.AdvisorCreateReq

	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		return
	}

	if err := advisorService.Create(&req); err != nil {
		_ = c.Error(err)
		return
	}

	response.OkWithMessage(i18n.Success, c)
}

// GetAdvisor 查询导师详情
//
//	@Tags		Advisor
//	@Summary	前后端-查询导师详情
//	@Produce	application/json
//	@Param		id	path		int															true	"导师记录ID"
//	@Success	200	{object}	response.Response{msg=string,data=response.AdvisorDetail}	"查询成功"
//	@Router		/advisor/:id [get]
func (advisorApi *AdvisorApi) GetAdvisor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.Error(err)
		return
	}

	resp, err := advisorService.Get(id)
	if err != nil {
		_ = c.Error(err)
		return
	}

	if utils.GetUserIDFrontUser(c) > 0 {
		lan := cRequest.GetLanguageTag(c)
		resp.NickName = i18n.Message(lan, resp.NickName, 0)
		resp.Duty = i18n.Message(lan, resp.Duty, 0)
		resp.Intro = i18n.Message(lan, resp.Intro, 0)
		resp.ActiveStatusText = i18n.Message(lan, i18n.ActiveStatusPrefix+strconv.Itoa(int(resp.ActiveStatus)), 0)
	}

	response.OkWithDetailed(resp, i18n.Success, c)
}

// UpdateAdvisor 更新导师
//
//	@Tags		Advisor
//	@Summary	后端-更新导师
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		id		path		int								true	"导师记录ID"
//	@Param		data	body		request.AdvisorUpdateReq		true	"更新参数"
//	@Success	200		{object}	response.Response{msg=string}	"更新成功"
//	@Router		/advisor/:id [put]
func (advisorApi *AdvisorApi) UpdateAdvisor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.Error(err)
		return
	}

	var req request.AdvisorUpdateReq
	err = c.ShouldBindJSON(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	req.Id = id
	err = advisorService.Update(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	response.OkWithMessage(i18n.Success, c)
}

// UpdateAdvisorStatus 更新导师启用状态
//
//	@Tags		Advisor
//	@Summary	后端-更新导师启用状态
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		id		path		int								true	"导师记录ID"
//	@Param		data	body		request.AdvisorStatusUpdateReq	true	"更新参数"
//	@Success	200		{object}	response.Response{msg=string}	"更新成功"
//	@Router		/advisor/status/:id [put]
func (advisorApi *AdvisorApi) UpdateAdvisorStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.Error(err)
		return
	}

	var req request.AdvisorStatusUpdateReq
	err = c.ShouldBindJSON(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	req.Id = id
	err = advisorService.UpdateStatus(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}
	response.OkWithMessage(i18n.Success, c)
}

// PageQuery  分页查询导师信息
//
//	@Tags		Advisor
//	@Summary	前后端-分页查询导师信息
//	@Produce	application/json
//	@Param		data	query		request.AdvisorPageQueryReq																	true	"查询参数"
//	@Success	200		{object}	response.Response{msg=string, data=response.PageResult{list=[]response.AdvisorPageData}}	"查询成功"
//	@Router		/advisor/page [get]
func (advisorApi *AdvisorApi) PageQuery(c *gin.Context) {
	var req request.AdvisorPageQueryReq
	if err := c.ShouldBindQuery(&req); err != nil {
		_ = c.Error(err)
		return
	}

	list, total, err := advisorService.PageQuery(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	if utils.GetUserIDFrontUser(c) > 0 {
		lan := cRequest.GetLanguageTag(c)
		for _, a := range list {
			a.NickName = i18n.Message(lan, a.NickName, 0)
			a.Duty = i18n.Message(lan, a.Duty, 0)
			a.Intro = i18n.Message(lan, a.Intro, 0)
			a.ActiveStatusText = i18n.Message(lan, i18n.ActiveStatusPrefix+strconv.Itoa(int(a.ActiveStatus)), 0)
		}
	}

	response.OkWithDetailed(&response.PageResult{
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    total,
		List:     list,
	}, i18n.Success, c)
}

// PageQueryFront 分页查询导师信息
//
//	@Tags		Advisor
//	@Summary	前端-分页查询导师信息
//	@Produce	application/json
//	@Param		data	query		request.AdvisorPageQueryReq																	true	"查询参数"
//	@Success	200		{object}	response.Response{msg=string, data=response.PageResult{list=[]response.AdvisorPageData}}	"查询成功"
//	@Router		/advisor/front/page [get]
func (advisorApi *AdvisorApi) PageQueryFront(c *gin.Context) {
	var req request.AdvisorPageQueryReq
	if err := c.ShouldBindQuery(&req); err != nil {
		_ = c.Error(err)
		return
	}

	list, total, err := advisorService.PageQuery(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	lan := cRequest.GetLanguageTag(c)
	for _, a := range list {
		a.NickName = i18n.Message(lan, a.NickName, 0)
		a.Duty = i18n.Message(lan, a.Duty, 0)
		a.Intro = i18n.Message(lan, a.Intro, 0)
		a.ActiveStatusText = i18n.Message(lan, i18n.ActiveStatusPrefix+strconv.Itoa(int(a.ActiveStatus)), 0)
	}

	response.OkWithDetailed(&response.PageResult{
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    total,
		List:     list,
	}, i18n.Success, c)
}

// ListAdvisorOptions 查询导师选项列表
//
//	@Tags		Advisor
//	@Summary	前后端-查询导师下拉选项
//	@Produce	application/json
//	@Param		data	query		request.AdvisorListQueryReq										true	"查询参数"
//	@Success	200		{object}	response.Response{msg=string, data=[]response.AdvisorOption}	"查询成功"
//	@Router		/advisor/options [get]
func (advisorApi *AdvisorApi) ListAdvisorOptions(c *gin.Context) {
	var req request.AdvisorListQueryReq
	if err := c.ShouldBindQuery(&req); err != nil {
		_ = c.Error(err)
		return
	}

	options, err := advisorService.ListAdvisors(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	response.OkWithDetailed(options, i18n.Success, c)
}
