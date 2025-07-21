package symbol

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/symbol"
	symbolReq "github.com/flipped-aurora/gin-vue-admin/server/model/symbol/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SymbolsCustomApi struct{}

// CreateSymbolsCustom 创建symbolsCustom表
// @Tags SymbolsCustom
// @Summary 创建symbolsCustom表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body symbol.SymbolsCustom true "创建symbolsCustom表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /symbolsCustom/create [post]
func (symbolsCustomApi *SymbolsCustomApi) CreateSymbolsCustom(c *gin.Context) {
	var symbolsCustom symbol.SymbolsCustom
	err := c.ShouldBindJSON(&symbolsCustom)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = symbolsCustomService.CreateSymbolsCustom(&symbolsCustom)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteSymbolsCustom 删除symbolsCustom表
// @Tags SymbolsCustom
// @Summary 删除symbolsCustom表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body symbol.SymbolsCustom true "删除symbolsCustom表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /symbolsCustom/delete [delete]
func (symbolsCustomApi *SymbolsCustomApi) DeleteSymbolsCustom(c *gin.Context) {
	id := c.Query("id")
	err := symbolsCustomService.DeleteSymbolsCustom(id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteSymbolsCustomByIds 批量删除symbolsCustom表
// @Tags SymbolsCustom
// @Summary 批量删除symbolsCustom表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /symbolsCustom/deleteByIds [delete]
func (symbolsCustomApi *SymbolsCustomApi) DeleteSymbolsCustomByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	err := symbolsCustomService.DeleteSymbolsCustomByIds(ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateSymbolsCustom 更新symbolsCustom表
// @Tags SymbolsCustom
// @Summary 更新symbolsCustom表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body symbol.SymbolsCustom true "更新symbolsCustom表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /symbolsCustom/update [put]
func (symbolsCustomApi *SymbolsCustomApi) UpdateSymbolsCustom(c *gin.Context) {
	var symbolsCustom symbol.SymbolsCustom
	err := c.ShouldBindJSON(&symbolsCustom)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = symbolsCustomService.UpdateSymbolsCustom(symbolsCustom)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSymbolsCustom 用id查询symbolsCustom表
// @Tags SymbolsCustom
// @Summary 用id查询symbolsCustom表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query symbol.SymbolsCustom true "用id查询symbolsCustom表"
// @Success 200 {object} response.Response{data=symbol.SymbolsCustom,msg=string} "查询成功"
// @Router /symbolsCustom/detail/{id} [get]
func (symbolsCustomApi *SymbolsCustomApi) FindSymbolsCustom(c *gin.Context) {
	id := c.Param("id")
	resymbolsCustom, err := symbolsCustomService.GetSymbolsCustom(id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(resymbolsCustom, c)
}

// GetSymbolsCustomList 分页获取symbolsCustom表列表
// @Tags SymbolsCustom
// @Summary 分页获取symbolsCustom表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query symbolReq.SymbolsCustomSearch true "分页获取symbolsCustom表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /symbolsCustom/list [get]
func (symbolsCustomApi *SymbolsCustomApi) GetSymbolsCustomList(c *gin.Context) {
	var pageInfo symbolReq.SymbolsCustomSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := symbolsCustomService.GetSymbolsCustomInfoList(pageInfo)
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

// GetSymbolsCustomPublic 前台-获取用户自定义交易对
// @Tags SymbolsCustom
// @Summary 前台-获取用户自定义交易对
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页参数"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /symbolsCustom/frontend/list [get]
func (symbolsCustomApi *SymbolsCustomApi) GetSymbolsCustomPublic(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserIDFrontUser(c)
	if userId == 0 {
		response.FailWithMessage("用户未登录", c)
		return
	}

	list, total, err := symbolsCustomService.GetSymbolsCustomPublic(pageInfo, userId)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// CreateSymbolsCustomPublic 前台-创建用户自定义交易对
// @Tags SymbolsCustom
// @Summary 前台-创建用户自定义交易对
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body symbolReq.CreateSymbolsCustomRequest true "symbolId: 交易对ID"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /symbolsCustom/frontend/create [post]
func (symbolsCustomApi *SymbolsCustomApi) CreateSymbolsCustomPublic(c *gin.Context) {
	var req symbolReq.CreateSymbolsCustomRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 从Token中获取用户ID
	userId := utils.GetUserIDFrontUser(c)
	if userId == 0 {
		response.FailWithMessage("用户未登录", c)
		return
	}

	// 构造 SymbolsCustom 对象
	userIdInt := int(userId)
	symbolsCustom := &symbol.SymbolsCustom{
		SymbolId: &req.SymbolId,
		UserId:   &userIdInt,
	}

	err = symbolsCustomService.CreateSymbolsCustomPublic(symbolsCustom)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteSymbolsCustomPublic 前台-删除用户自定义交易对
// @Tags SymbolsCustom
// @Summary 前台-删除用户自定义交易对
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id query string true "自定义交易对ID"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /symbolsCustom/frontend/delete [delete]
func (symbolsCustomApi *SymbolsCustomApi) DeleteSymbolsCustomPublic(c *gin.Context) {
	id := c.Query("id")

	// 从Token中获取用户ID
	userId := utils.GetUserIDFrontUser(c)
	if userId == 0 {
		response.FailWithMessage("用户未登录", c)
		return
	}

	err := symbolsCustomService.DeleteSymbolsCustomPublic(id, userId)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// InitUserDefaultSymbols 初始化用户默认的自定义交易对
// @Tags SymbolsCustom
// @Summary 初始化用户默认的自定义交易对（市值前十）
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "初始化成功"
// @Router /symbolsCustom/frontend/initUserSymbols [post]
func (symbolsCustomApi *SymbolsCustomApi) InitUserDefaultSymbols(c *gin.Context) {
	// 从Token中获取用户ID
	userId := utils.GetUserIDFrontUser(c)
	if userId == 0 {
		response.FailWithMessage("用户未登录", c)
		return
	}
	err := symbolsCustomService.InitUserDefaultSymbols(int(userId))
	if err != nil {
		global.GVA_LOG.Error("初始化默认自选股失败!", zap.Error(err))
		response.FailWithMessage("初始化失败: "+err.Error(), c)
		return
	}
	response.OkWithMessage("初始化成功", c)
}

// InitUserDefaultSymbolsAdmin 初始化用户默认的自定义交易对
// @Tags SymbolsCustom
// @Summary 初始化用户默认的自定义交易对（市值前十）
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "初始化成功"
// @Router /symbolsCustom/initUserSymbols/{userId} [post]
func (symbolsCustomApi *SymbolsCustomApi) InitUserDefaultSymbolsAdmin(c *gin.Context) {
	// 从URL参数获取用户ID
	userId := c.Param("userId")
	if userId == "" {
		response.FailWithMessage("用户ID不能为空", c)
		return
	}

	// 将字符串ID转换为整数
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		response.FailWithMessage("无效的用户ID", c)
		return
	}

	err = symbolsCustomService.InitUserDefaultSymbols(userIdInt)
	if err != nil {
		global.GVA_LOG.Error("初始化默认自选股失败!", zap.Error(err))
		response.FailWithMessage("初始化失败: "+err.Error(), c)
		return
	}
	response.OkWithMessage("初始化成功", c)
}
