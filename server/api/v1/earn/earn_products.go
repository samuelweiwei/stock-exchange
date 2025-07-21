package earn

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/i18n"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/earn"
	earnReq "github.com/flipped-aurora/gin-vue-admin/server/model/earn/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EarnProductsApi struct{}

// CreateEarnProducts 后台创建理财产品
// @Tags EarnProducts
// @Summary 后台创建理财产品
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body earn.EarnProducts true "创建理财产品"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /earn/products/create [post]
func (earnProductsApi *EarnProductsApi) CreateEarnProducts(c *gin.Context) {
	var earnProduct earn.EarnProducts
	err := c.ShouldBindJSON(&earnProduct)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if earnProduct.Type == earn.Flexible && !earnProduct.PenaltyRatio.IsZero() {
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.FlexibleEarnStakeProductPenaltyRatioMustZero, response.ERROR), c)
		return
	}
	err = earnProductsService.CreateEarnProducts(&earnProduct)
	if err != nil {
		global.GVA_LOG.Error("create failed!", zap.Error(err), zap.Any("earnProduct", earnProduct))
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.CreationFailed, response.ERROR), c)
		return
	}
	response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}

// DeleteEarnProducts 后台删除理财产品
// @Tags EarnProducts
// @Summary 后台删除理财产品
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id query string true "删除理财产品"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /earn/products/delete [delete]
func (earnProductsApi *EarnProductsApi) DeleteEarnProducts(c *gin.Context) {
	id := c.Query("id")
	err := earnProductsService.DeleteEarnProducts(id)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.DeleteFailed, response.ERROR), c)
		return
	}
	response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}

// DeleteEarnProductsByIds 批量删除earnProducts表
// @Tags EarnProducts
// @Summary 批量删除earnProducts表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /earnProducts/deleteEarnProductsByIds [delete]
func (earnProductsApi *EarnProductsApi) DeleteEarnProductsByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	err := earnProductsService.DeleteEarnProductsByIds(ids)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Failed, response.ERROR), c)
		return
	}
	response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}

// UpdateEarnProducts 后台更新理财产品
// @Tags EarnProducts
// @Summary 后台更新理财产品
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body earn.EarnProducts true "更新理财产品"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /earn/products/edit [put]
func (earnProductsApi *EarnProductsApi) UpdateEarnProducts(c *gin.Context) {
	var earnProducts earn.EarnProducts
	err := c.ShouldBindJSON(&earnProducts)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = earnProductsService.UpdateEarnProducts(earnProducts)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err), zap.Any("earnProducts", earnProducts))
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.UpdateFailed, response.UPDATE_FAILED), c)
		return
	}
	response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}

// FindEarnProducts 用id查询earnProducts表
// @Tags EarnProducts
// @Summary 用id查询earnProducts表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query earn.EarnProducts true "用id查询earnProducts表"
// @Success 200 {object} response.Response{data=earn.EarnProducts,msg=string} "查询成功"
// @Router /earnProducts/findEarnProducts [get]
func (earnProductsApi *EarnProductsApi) FindEarnProducts(c *gin.Context) {
	id := c.Query("id")
	reearnProducts, err := earnProductsService.GetEarnProducts(id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reearnProducts, c)
}

// GetEarnProductsList 后台获取理财产品列表
// @Tags EarnProducts
// @Summary 后台获取理财产品列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query earnReq.EarnProductsSearch true "获取理财产品列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /earn/products/list [get]
func (earnProductsApi *EarnProductsApi) GetEarnProductsList(c *gin.Context) {
	var pageInfo earnReq.EarnProductsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := earnProductsService.GetEarnProductsInfoList(pageInfo)
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

// GetFrontEarnProductsList 前端用户获取自已理财产品
// @Tags EarnProducts
// @Summary  前端用户获取自已理财产品
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query earnReq.EarnProductsSearch true "前端用户获取自已理财产品"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /front/earn/product/list [get]
func (earnProductsApi *EarnProductsApi) GetFrontEarnProductsList(c *gin.Context) {
	var pageInfo earnReq.EarnProductsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pageInfo.Uid = utils.GetUserID(c)
	list, total, err := earnProductsService.GetFrontEarnProductsInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("get earn products list error", zap.Error(err), zap.Any("req", pageInfo))
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ServerError, response.ERROR), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}

// GetEarnProductsPublic 不需要鉴权的earnProducts表接口
// @Tags EarnProducts
// @Summary 不需要鉴权的earnProducts表接口
// @accept application/json
// @Produce application/json
// @Param data query earnReq.EarnProductsSearch true "分页获取earnProducts表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /earnProducts/getEarnProductsPublic [get]
func (earnProductsApi *EarnProductsApi) GetEarnProductsPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	earnProductsService.GetEarnProductsPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的earnProducts表接口信息",
	}, "获取成功", c)
}
