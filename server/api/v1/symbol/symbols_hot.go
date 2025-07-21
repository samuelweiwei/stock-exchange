package symbol

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/symbol"
	symbolReq "github.com/flipped-aurora/gin-vue-admin/server/model/symbol/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SymbolsHotApi struct{}

// CreateSymbolsHot 创建symbolsHot表
// @Tags SymbolsHot
// @Summary 创建symbolsHot表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body symbol.SymbolsHot true "创建symbolsHot表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /symbolsHot/create [post]
func (symbolsHotApi *SymbolsHotApi) CreateSymbolsHot(c *gin.Context) {
	var symbolsHot symbol.SymbolsHot
	err := c.ShouldBindJSON(&symbolsHot)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = symbolsHotService.CreateSymbolsHot(&symbolsHot)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteSymbolsHot 删除symbolsHot表
// @Tags SymbolsHot
// @Summary 删除symbolsHot表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body symbol.SymbolsHot true "删除symbolsHot表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /symbolsHot/delete [delete]
func (symbolsHotApi *SymbolsHotApi) DeleteSymbolsHot(c *gin.Context) {
	id := c.Query("id")
	err := symbolsHotService.DeleteSymbolsHot(id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteSymbolsHotByIds 批量删除symbolsHot表
// @Tags SymbolsHot
// @Summary 批量删除symbolsHot表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /symbolsHot/deleteByIds [delete]
func (symbolsHotApi *SymbolsHotApi) DeleteSymbolsHotByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	err := symbolsHotService.DeleteSymbolsHotByIds(ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateSymbolsHot 更新symbolsHot表
// @Tags SymbolsHot
// @Summary 更新symbolsHot表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body symbol.SymbolsHot true "更新symbolsHot表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /symbolsHot/update [put]
func (symbolsHotApi *SymbolsHotApi) UpdateSymbolsHot(c *gin.Context) {
	var symbolsHot symbol.SymbolsHot
	err := c.ShouldBindJSON(&symbolsHot)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = symbolsHotService.UpdateSymbolsHot(symbolsHot)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSymbolsHot 用id查询symbolsHot表
// @Tags SymbolsHot
// @Summary 用id查询symbolsHot表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query symbol.SymbolsHot true "用id查询symbolsHot表"
// @Success 200 {object} response.Response{data=symbol.SymbolsHot,msg=string} "查询成功"
// @Router /symbolsHot/detail/{id} [get]
func (symbolsHotApi *SymbolsHotApi) FindSymbolsHot(c *gin.Context) {
	id := c.Param("id")
	resymbolsHot, err := symbolsHotService.GetSymbolsHot(id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(resymbolsHot, c)
}

// GetSymbolsHotList 分页获取symbolsHot表列表
// @Tags SymbolsHot
// @Summary 分页获取symbolsHot表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query symbolReq.SymbolsHotSearch true "分页获取symbolsHot表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /symbolsHot/list [get]
func (symbolsHotApi *SymbolsHotApi) GetSymbolsHotList(c *gin.Context) {
	var pageInfo symbolReq.SymbolsHotSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := symbolsHotService.GetSymbolsHotInfoList(pageInfo)
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

// GetSymbolsHotPublic 前台-不需要鉴权的symbolsHot表接口
// @Tags SymbolsHot
// @Summary 前台-不需要鉴权的symbolsHot表接口
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页参数"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /symbolsHot/public/list [get]
func (symbolsHotApi *SymbolsHotApi) GetSymbolsHotPublic(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := symbolsHotService.GetSymbolsHotPublic(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取热门股票失败!", zap.Error(err))
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
