package system

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/constants"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/i18n"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gin-gonic/gin"
)

type SystemConfigApi struct {
}

// GetSystemConfig 				查询系统配置
// @Tags      					SystemConfig
// @Summary   					查询系统配置
// @Security  					ApiKeyAuth
// @Produce   					application/json
// @Success   					200   {object}  response.Response{msg=string,data=response.SystemConfig}  "请求成功"
// @Router    					/system/config [get]
func (api *SystemConfigApi) GetSystemConfig(c *gin.Context) {
	resp, err := systemConfigService.GetPlatformSystemConfig()
	if err != nil {
		_ = c.Error(err)
		return
	}

	response.OkWithDetailed(resp, i18n.Success, c)
}

// GetPlatformCommissionRate 	查询平台佣金比例
// @Tags      					SystemConfig
// @Summary   					查询平台佣金比例
// @Produce   					application/json
// @Success   					200   {object}  response.Response{msg=string,data=float64}  "请求成功"
// @Router    					/system/platform-commission-rate [get]
func (api *SystemConfigApi) GetPlatformCommissionRate(c *gin.Context) {
	rate, err := global.GVA_REDIS.Get(context.Background(), constants.RedisKeyPlatformCommissionRate).Float64()
	if err != nil {
		_ = c.Error(err)
		return
	}
	response.OkWithDetailed(rate, i18n.Success, c)
}

// SaveSystemConfig 			保存系统配置
// @Tags      					SystemConfig
// @Summary   					保存系统配置
// @Security  					ApiKeyAuth
// @accept						application/json
// @Produce   					application/json
// @Param						data  body      request.SystemConfigSaveReq true "请求参数"
// @Success   					200   {object}  response.Response{msg=string,data=response.SystemConfig}  "请求成功"
// @Router    					/system/config [post]
func (api *SystemConfigApi) SaveSystemConfig(c *gin.Context) {
	var req request.SystemConfigSaveReq
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		return
	}

	err := systemConfigService.SavePlatformSystemConfig(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	response.OkWithMessage(i18n.Success, c)
}

// GetDomainInfo 				查询站点信息
// @Tags      					SystemConfig
// @Summary   					前台-查询站点信息
// @Security  					ApiKeyAuth
// @Produce   					application/json
// @Success   					200   {object}  response.Response{msg=string,data=response.SystemConfig}  "请求成功"
// @Router    					/system/config [get]
func (api *SystemConfigApi) GetDomainInfo(c *gin.Context) {
	resp, err := systemConfigService.GetDomainInfo()
	if err != nil {
		_ = c.Error(err)
		return
	}
	response.OkWithDetailed(resp, i18n.Success, c)
}
