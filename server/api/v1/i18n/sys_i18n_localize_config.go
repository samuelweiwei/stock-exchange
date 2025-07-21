package i18n

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/i18n"
	i18nReq "github.com/flipped-aurora/gin-vue-admin/server/model/i18n/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysI18nLocalizeConfigApi struct{}

// CreateSysI18nLocalizeConfig 创建sysI18nLocalizeConfig表
// @Tags I18n
// @Summary 创建sysI18nLocalizeConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body i18n.SysI18nLocalizeConfig true "创建sysI18nLocalizeConfig表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /i18n/localize/create [post]
func (sysI18nLocalizeConfigApi *SysI18nLocalizeConfigApi) CreateSysI18nLocalizeConfig(c *gin.Context) {
	var sysI18nLocalizeConfig i18n.SysI18nLocalizeConfig
	err := c.ShouldBindJSON(&sysI18nLocalizeConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = sysI18nLocalizeConfigService.CreateSysI18nLocalizeConfig(&sysI18nLocalizeConfig)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteSysI18nLocalizeConfig 删除sysI18nLocalizeConfig表
// @Tags I18n
// @Summary 删除sysI18nLocalizeConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body i18n.SysI18nLocalizeConfig true "删除sysI18nLocalizeConfig表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /i18n/localize/delete [delete]
func (sysI18nLocalizeConfigApi *SysI18nLocalizeConfigApi) DeleteSysI18nLocalizeConfig(c *gin.Context) {
	id := c.Query("id")
	err := sysI18nLocalizeConfigService.DeleteSysI18nLocalizeConfig(id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteSysI18nLocalizeConfigByIds 批量删除sysI18nLocalizeConfig表
// @Tags I18n
// @Summary 批量删除sysI18nLocalizeConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /i18n/localize/batch/delete [delete]
func (sysI18nLocalizeConfigApi *SysI18nLocalizeConfigApi) DeleteSysI18nLocalizeConfigByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	err := sysI18nLocalizeConfigService.DeleteSysI18nLocalizeConfigByIds(ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateSysI18nLocalizeConfig 更新sysI18nLocalizeConfig表
// @Tags I18n
// @Summary 更新sysI18nLocalizeConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body i18n.SysI18nLocalizeConfig true "更新sysI18nLocalizeConfig表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /i18n/localize/update [put]
func (sysI18nLocalizeConfigApi *SysI18nLocalizeConfigApi) UpdateSysI18nLocalizeConfig(c *gin.Context) {
	var sysI18nLocalizeConfig i18n.SysI18nLocalizeConfig
	err := c.ShouldBindJSON(&sysI18nLocalizeConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = sysI18nLocalizeConfigService.UpdateSysI18nLocalizeConfig(sysI18nLocalizeConfig)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSysI18nLocalizeConfig 用id查询sysI18nLocalizeConfig表
// @Tags I18n
// @Summary 用id查询sysI18nLocalizeConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query i18n.SysI18nLocalizeConfig true "用id查询sysI18nLocalizeConfig表"
// @Success 200 {object} response.Response{data=i18n.SysI18nLocalizeConfig,msg=string} "查询成功"
// @Router /i18n/localize/find [get]
func (sysI18nLocalizeConfigApi *SysI18nLocalizeConfigApi) FindSysI18nLocalizeConfig(c *gin.Context) {
	id := c.Query("id")
	resysI18nLocalizeConfig, err := sysI18nLocalizeConfigService.GetSysI18nLocalizeConfig(id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(resysI18nLocalizeConfig, c)
}

// GetSysI18nLocalizeConfigList 分页获取sysI18nLocalizeConfig表列表
// @Tags I18n
// @Summary 分页获取sysI18nLocalizeConfig表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query i18nReq.SysI18nLocalizeConfigSearch true "分页获取sysI18nLocalizeConfig表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /i18n/localize/list [get]
func (sysI18nLocalizeConfigApi *SysI18nLocalizeConfigApi) GetSysI18nLocalizeConfigList(c *gin.Context) {
	var pageInfo i18nReq.SysI18nLocalizeConfigSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := sysI18nLocalizeConfigService.GetSysI18nLocalizeConfigInfoList(pageInfo)
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

// GetSysI18nLocalizeConfigPublic 不需要鉴权的sysI18nLocalizeConfig表接口
// @Tags I18n
// @Summary 不需要鉴权的sysI18nLocalizeConfig表接口
// @accept application/json
// @Produce application/json
// @Param data query i18nReq.SysI18nLocalizeConfigSearch true "分页获取sysI18nLocalizeConfig表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /i18n/localize/public/get [get]
func (sysI18nLocalizeConfigApi *SysI18nLocalizeConfigApi) GetSysI18nLocalizeConfigPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	sysI18nLocalizeConfigService.GetSysI18nLocalizeConfigPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的sysI18nLocalizeConfig表接口信息",
	}, "获取成功", c)
}
