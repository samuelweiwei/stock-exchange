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

type AdvisorProdApi struct {
}

// CreateProd 创建导师产品
//
//	@Tags		AdvisorProd
//	@Summary	后端-创建导师产品
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		request.AdvisorProdCreateReq	true	"创建参数"
//	@Success	200		{object}	response.Response{msg=string}	"创建成功"
//	@Router		/advisor/prod [post]
func (advisorProdApi *AdvisorProdApi) CreateProd(c *gin.Context) {
	var req request.AdvisorProdCreateReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	if err = advisorProdService.CreateProduct(&req); err != nil {
		_ = c.Error(err)
		return
	}

	response.OkWithMessage(i18n.Success, c)
}

// UpdateProd 更新导师产品
//
//	@Tags		AdvisorProd
//	@Summary	后端-更新导师产品
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		id		path		int								true	"产品ID"
//	@Param		data	body		request.AdvisorProdUpdateReq	true	"更新参数"
//	@Success	200		{object}	response.Response{msg=string}	"更新成功"
//	@Router		/advisor/prod/:id [put]
func (advisorProdApi *AdvisorProdApi) UpdateProd(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.Error(err)
		return
	}

	var req request.AdvisorProdUpdateReq
	err = c.ShouldBindJSON(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	req.Id = uint(id)
	err = advisorProdService.UpdateProduct(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	response.OkWithMessage(i18n.Success, c)
}

// DeleteProd 删除导师产品
//
//	@Tags		AdvisorProd
//	@Summary	后端-删除导师产品
//	@Security	ApiKeyAuth
//	@Produce	application/json
//	@Param		id	path		int								true	"产品ID"
//	@Success	200	{object}	response.Response{msg=string}	"删除成功"
//	@Router		/advisor/prod/:id [delete]
func (advisorProdApi *AdvisorProdApi) DeleteProd(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.Error(err)
		return
	}

	err = advisorProdService.DeleteProduct(uint(id))
	if err != nil {
		_ = c.Error(err)
		return
	}

	response.OkWithMessage(i18n.Success, c)
}

// PageQuery  分页查询导师产品信息
//
//	@Tags		AdvisorProd
//	@Summary	后端-分页查询导师产品信息
//	@Security	ApiKeyAuth
//	@Produce	application/json
//	@Param		data	query		request.AdvisorProdPageQueryReq																true	"查询参数"
//	@Success	200		{object}	response.Response{msg=string, data=response.PageResult{list=[]response.AdvisorPageData}}	"查询成功"
//	@Router		/advisor/prod/page [get]
func (advisorProdApi *AdvisorProdApi) PageQuery(c *gin.Context) {
	var req request.AdvisorProdPageQueryReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	pageDataList, total, err := advisorProdService.PageQuery(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	response.OkWithDetailed(&response.PageResult{
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    total,
		List:     pageDataList,
	}, i18n.Success, c)
}

// GetAdvisorProdById   查询导师产品详情
//
//	@Tags		AdvisorProd
//	@Summary	后端-查询导师产品详情
//	@Security	ApiKeyAuth
//	@Produce	application/json
//	@Param		advisorProdId	path		int																							true	"导师产品ID"
//	@Success	200				{object}	response.Response{msg=string, data=response.PageResult{list=[]response.AdvisorPageData}}	"查询成功"
//	@Router		/advisor/prod/:advisorProdId [get]
func (advisorProdApi *AdvisorProdApi) GetAdvisorProdById(c *gin.Context) {
	advisorProdId, err := strconv.Atoi(c.Param("advisorProdId"))
	if err != nil {
		_ = c.Error(err)
		return
	}

	resp, err := advisorProdService.GetAdvisorProdById(uint(advisorProdId))
	if err != nil {
		_ = c.Error(err)
		return
	}

	response.OkWithDetailed(resp, i18n.Success, c)
}

// ListByAdvisorId 根据导师ID查询导师产品列表
//
//	@Tags		AdvisorProd
//	@Summary	前端-根据导师ID查询导师产品列表
//	@accept		application/json
//	@Produce	application/json
//	@Param		advisorId	path		int																	true	"导师ID"
//	@Success	200			{object}	response.Response{msg=string, data=[]response.AdvisorProdListData}	"查询成功"
//	@Router		/advisor/prods/:advisorId [get]
func (advisorProdApi *AdvisorProdApi) ListByAdvisorId(c *gin.Context) {
	advisorId, err := strconv.Atoi(c.Param("advisorId"))
	if err != nil {
		_ = c.Error(err)
		return
	}

	list, err := advisorProdService.ListByAdvisorId(uint(advisorId), utils.GetUserIDFrontUser(c))
	if err != nil {
		_ = c.Error(err)
		return
	}

	lan := cRequest.GetLanguageTag(c)
	for _, d := range list {
		d.ProductName = i18n.Message(lan, d.ProductName, 0)
		d.AutoRenewText = i18n.Message(lan, i18n.AutoRenewPrefix+strconv.Itoa(int(d.AutoRenew)), 0)
	}

	response.OkWithDetailed(list, i18n.Success, c)
}
