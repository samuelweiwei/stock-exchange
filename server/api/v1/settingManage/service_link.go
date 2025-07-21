package settingManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/constants"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/i18n"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	settingManageReq "github.com/flipped-aurora/gin-vue-admin/server/model/settingManage/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type ServiceLinkApi struct{}

// CreateServiceLink 创建serviceLink表
// @Tags ServiceLink
// @Summary 后台-创建serviceLink表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body settingManage.ServiceLink true "创建serviceLink表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /serviceLink/createServiceLink [post]
func (serviceLinkApi *ServiceLinkApi) CreateServiceLink(c *gin.Context) {
	var serviceLink settingManageReq.ServiceLinkCreate
	err := c.ShouldBindJSON(&serviceLink)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	serviceLink.CreatedUid = utils.GetUserID(c)
	serviceLink.CreatedTime = time.Now().UnixMilli()
	err = serviceLinkService.CreateServiceLink(&serviceLink)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteServiceLink 删除serviceLink表
// @Tags ServiceLink
// @Summary 后台-删除serviceLink表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body settingManage.ServiceLink true "删除serviceLink表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /serviceLink/deleteServiceLink [delete]
func (serviceLinkApi *ServiceLinkApi) DeleteServiceLink(c *gin.Context) {
	id := c.Query("id")
	err := serviceLinkService.DeleteServiceLink(id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

/*
// DeleteServiceLinkByIds 批量删除serviceLink表
// @Tags
// @Summary 批量删除serviceLink表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /serviceLink/deleteServiceLinkByIds [delete]
*/

func (serviceLinkApi *ServiceLinkApi) DeleteServiceLinkByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	err := serviceLinkService.DeleteServiceLinkByIds(ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateServiceLink 更新serviceLink表
// @Tags ServiceLink
// @Summary 后台-更新serviceLink表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body settingManage.ServiceLink true "更新serviceLink表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /serviceLink/updateServiceLink [put]
func (serviceLinkApi *ServiceLinkApi) UpdateServiceLink(c *gin.Context) {
	var serviceLink settingManageReq.ServiceLinkUpdate
	err := c.ShouldBindJSON(&serviceLink)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	serviceLink.UpdatedUid = utils.GetUserID(c)
	serviceLink.UpdatedTime = time.Now().UnixMilli()
	err = serviceLinkService.UpdateServiceLink(serviceLink)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

/*
// FindServiceLink 用id查询serviceLink表
// @Tags
// @Summary 用id查询serviceLink表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query settingManage.ServiceLink true "用id查询serviceLink表"
// @Success 200 {object} response.Response{data=settingManage.ServiceLink,msg=string} "查询成功"
// @Router /serviceLink/findServiceLink [get]
*/

func (serviceLinkApi *ServiceLinkApi) FindServiceLink(c *gin.Context) {
	id := c.Query("id")
	reServiceLink, err := serviceLinkService.GetServiceLink(id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reServiceLink, c)
}

// GetServiceLinkList 分页获取serviceLink表列表
// @Tags ServiceLink
// @Summary 后台-分页获取serviceLink表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query settingManageReq.ServiceLinkSearch true "分页获取serviceLink表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /serviceLink/getServiceLinkList [get]
func (serviceLinkApi *ServiceLinkApi) GetServiceLinkList(c *gin.Context) {
	var searchInfo settingManageReq.ServiceLinkSearch
	err := c.ShouldBindQuery(&searchInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceLinkService.GetServiceLinkInfoList(searchInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     searchInfo.Page,
		PageSize: searchInfo.PageSize,
	}, "获取成功", c)
}

// GetFrontServiceLinkList 分页获取serviceLink表列表
// @Tags ServiceLink
// @Summary 前台-分页获取serviceLink表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query settingManageReq.ServiceLinkSearch true "分页获取serviceLink表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /serviceLink/getFrontServiceLinkList [get]
func (serviceLinkApi *ServiceLinkApi) GetFrontServiceLinkList(c *gin.Context) {
	var searchInfo settingManageReq.ServiceLinkSearch
	err := c.ShouldBindQuery(&searchInfo)
	searchInfo.Status = new(uint)
	*searchInfo.Status = constants.ServiceLinkStatusOpen
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceLinkService.GetServiceLinkInfoList(searchInfo)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ServerError, response.ERROR), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     searchInfo.Page,
		PageSize: searchInfo.PageSize,
	}, i18n.Message(request.GetLanguageTag(c), i18n.SuccessfullyObtained, response.SUCCESS), c)
}

func (serviceLinkApi *ServiceLinkApi) GetServiceLinkPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	serviceLinkService.GetServiceLinkPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的serviceLink表接口信息",
	}, "获取成功", c)
}
