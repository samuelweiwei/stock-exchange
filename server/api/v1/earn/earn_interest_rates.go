package earn

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/earn"
    earnReq "github.com/flipped-aurora/gin-vue-admin/server/model/earn/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type EarnInterestRatesApi struct {}



// CreateEarnInterestRates 创建earnInterestRates表
// @Tags EarnInterestRates
// @Summary 创建earnInterestRates表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body earn.EarnInterestRates true "创建earnInterestRates表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /earnInterestRates/createEarnInterestRates [post]
func (earnInterestRatesApi *EarnInterestRatesApi) CreateEarnInterestRates(c *gin.Context) {
	var earnInterestRates earn.EarnInterestRates
	err := c.ShouldBindJSON(&earnInterestRates)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = earnInterestRatesService.CreateEarnInterestRates(&earnInterestRates)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteEarnInterestRates 删除earnInterestRates表
// @Tags EarnInterestRates
// @Summary 删除earnInterestRates表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body earn.EarnInterestRates true "删除earnInterestRates表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /earnInterestRates/deleteEarnInterestRates [delete]
func (earnInterestRatesApi *EarnInterestRatesApi) DeleteEarnInterestRates(c *gin.Context) {
	id := c.Query("id")
	err := earnInterestRatesService.DeleteEarnInterestRates(id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteEarnInterestRatesByIds 批量删除earnInterestRates表
// @Tags EarnInterestRates
// @Summary 批量删除earnInterestRates表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /earnInterestRates/deleteEarnInterestRatesByIds [delete]
func (earnInterestRatesApi *EarnInterestRatesApi) DeleteEarnInterestRatesByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	err := earnInterestRatesService.DeleteEarnInterestRatesByIds(ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateEarnInterestRates 更新earnInterestRates表
// @Tags EarnInterestRates
// @Summary 更新earnInterestRates表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body earn.EarnInterestRates true "更新earnInterestRates表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /earnInterestRates/updateEarnInterestRates [put]
func (earnInterestRatesApi *EarnInterestRatesApi) UpdateEarnInterestRates(c *gin.Context) {
	var earnInterestRates earn.EarnInterestRates
	err := c.ShouldBindJSON(&earnInterestRates)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = earnInterestRatesService.UpdateEarnInterestRates(earnInterestRates)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindEarnInterestRates 用id查询earnInterestRates表
// @Tags EarnInterestRates
// @Summary 用id查询earnInterestRates表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query earn.EarnInterestRates true "用id查询earnInterestRates表"
// @Success 200 {object} response.Response{data=earn.EarnInterestRates,msg=string} "查询成功"
// @Router /earnInterestRates/findEarnInterestRates [get]
func (earnInterestRatesApi *EarnInterestRatesApi) FindEarnInterestRates(c *gin.Context) {
	id := c.Query("id")
	reearnInterestRates, err := earnInterestRatesService.GetEarnInterestRates(id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reearnInterestRates, c)
}

// GetEarnInterestRatesList 分页获取earnInterestRates表列表
// @Tags EarnInterestRates
// @Summary 分页获取earnInterestRates表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query earnReq.EarnInterestRatesSearch true "分页获取earnInterestRates表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /earnInterestRates/getEarnInterestRatesList [get]
func (earnInterestRatesApi *EarnInterestRatesApi) GetEarnInterestRatesList(c *gin.Context) {
	var pageInfo earnReq.EarnInterestRatesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := earnInterestRatesService.GetEarnInterestRatesInfoList(pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetEarnInterestRatesPublic 不需要鉴权的earnInterestRates表接口
// @Tags EarnInterestRates
// @Summary 不需要鉴权的earnInterestRates表接口
// @accept application/json
// @Produce application/json
// @Param data query earnReq.EarnInterestRatesSearch true "分页获取earnInterestRates表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /earnInterestRates/getEarnInterestRatesPublic [get]
func (earnInterestRatesApi *EarnInterestRatesApi) GetEarnInterestRatesPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    earnInterestRatesService.GetEarnInterestRatesPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的earnInterestRates表接口信息",
    }, "获取成功", c)
}
