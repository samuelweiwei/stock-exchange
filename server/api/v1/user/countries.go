package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/i18n"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	userReq "github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CountriesApi struct{}

// CreateCountries 后台-创建countries表
// @Tags Countries
// @Summary 创建countries表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.Countries true "创建countries表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /countries/createCountries [post]
func (countriesApi *CountriesApi) CreateCountries(c *gin.Context) {
	var countries user.Countries
	err := c.ShouldBindJSON(&countries)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = countriesService.CreateCountries(&countries)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteCountries 后台-删除countries表
// @Tags Countries
// @Summary 删除countries表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.Countries true "删除countries表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /countries/deleteCountries [delete]
func (countriesApi *CountriesApi) DeleteCountries(c *gin.Context) {
	id := c.Query("id")
	err := countriesService.DeleteCountries(id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteCountriesByIds 后台-批量删除countries表
// @Tags Countries
// @Summary 批量删除countries表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /countries/deleteCountriesByIds [delete]
func (countriesApi *CountriesApi) DeleteCountriesByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	err := countriesService.DeleteCountriesByIds(ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateCountries 后台-更新countries表
// @Tags Countries
// @Summary 更新countries表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.Countries true "更新countries表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /countries/updateCountries [put]
func (countriesApi *CountriesApi) UpdateCountries(c *gin.Context) {
	var countries user.Countries
	err := c.ShouldBindJSON(&countries)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = countriesService.UpdateCountries(countries)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindCountries 后台-用id查询countries表
// @Tags Countries
// @Summary 用id查询countries表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query user.Countries true "用id查询countries表"
// @Success 200 {object} response.Response{data=user.Countries,msg=string} "查询成功"
// @Router /countries/findCountries [get]
func (countriesApi *CountriesApi) FindCountries(c *gin.Context) {
	id := c.Query("id")
	recountries, err := countriesService.GetCountries(id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(recountries, c)
}

// GetCountriesList 后台-分页获取countries表列表
// @Tags Countries
// @Summary 后台-分页获取countries表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userReq.CountriesSearch true "分页获取countries表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /countries/getCountriesList [get]
func (countriesApi *CountriesApi) GetCountriesList(c *gin.Context) {
	var pageInfo userReq.CountriesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := countriesService.GetCountriesInfoList(pageInfo)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.FailedObtained, response.ERROR), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}

// GetCountriesPublic 不需要鉴权的countries表接口
// @Tags Countries
// @Summary 不需要鉴权的countries表接口
// @accept application/json
// @Produce application/json
// @Param data query userReq.CountriesSearch true "分页获取countries表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /countries/getCountriesPublic [get]
func (countriesApi *CountriesApi) GetCountriesPublic(c *gin.Context) {
	// 获取原始的查询参数
	query := c.Request.URL.Query()

	// 设置 status=1
	query.Set("status", "1")

	// 重新设置 URL 的 RawQuery
	c.Request.URL.RawQuery = query.Encode()
	countriesApi.GetCountriesList(c)
}
