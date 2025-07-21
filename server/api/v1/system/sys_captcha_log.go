package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/i18n"
	captchaReq "github.com/flipped-aurora/gin-vue-admin/server/model/captcha/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CaptchaLogApi struct{}

// List
// @Tags Captcha
// @Summary 查询验证码发送记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query captchaReq.CaptchaListSearch true "查询验证码发送记录"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /admin/captcha/log/list [get]
func (a *CaptchaLogApi) List(c *gin.Context) {
	var (
		req captchaReq.CaptchaListSearch
	)
	err := c.ShouldBindQuery(&req)
	if err != nil {
		global.GVA_LOG.Error("captcha param error", zap.Error(err), zap.Any("req", req))
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ParamError, response.ERROR), c)
		return
	}
	list, total, err := captchaLogService.AdminList(req)
	if err != nil {
		global.GVA_LOG.Error("get captcha error", zap.Error(err), zap.Any("req", req))
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ServerError, response.ERROR), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}
