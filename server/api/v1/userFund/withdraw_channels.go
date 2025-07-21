package userFund

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/i18n"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/userfund"
	userfundReq "github.com/flipped-aurora/gin-vue-admin/server/model/userfund/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WithdrawChannelsApi struct{}

// CreateWithdrawChannels 创建withdrawChannels表
// @Tags WithdrawChannels
// @Summary 创建withdrawChannels表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userfund.WithdrawChannels true "创建withdrawChannels表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /withdrawChannels/createWithdrawChannels [post]
func (withdrawChannelsApi *WithdrawChannelsApi) CreateWithdrawChannels(c *gin.Context) {
	var withdrawChannels userfund.WithdrawChannels
	err := c.ShouldBindJSON(&withdrawChannels)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 验证关联的货币是否存在
	if withdrawChannels.CoinId == nil {
		response.FailWithMessage("货币ID不能为空", c)
		return
	}

	var currency userfund.Currencies
	if err := global.GVA_DB.First(&currency, withdrawChannels.CoinId).Error; err != nil {
		response.FailWithMessage("关联的货币不存在", c)
		return
	}

	err = withdrawChannelsService.CreateWithdrawChannels(&withdrawChannels)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteWithdrawChannels 删除withdrawChannels表
// @Tags WithdrawChannels
// @Summary 删除withdrawChannels表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userfund.WithdrawChannels true "删除withdrawChannels表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /withdrawChannels/deleteWithdrawChannels [delete]
func (withdrawChannelsApi *WithdrawChannelsApi) DeleteWithdrawChannels(c *gin.Context) {
	ID := c.Query("ID")
	err := withdrawChannelsService.DeleteWithdrawChannels(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteWithdrawChannelsByIds 批量删除withdrawChannels表
// @Tags WithdrawChannels
// @Summary 批量删除withdrawChannels表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /withdrawChannels/deleteWithdrawChannelsByIds [delete]
func (withdrawChannelsApi *WithdrawChannelsApi) DeleteWithdrawChannelsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := withdrawChannelsService.DeleteWithdrawChannelsByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateWithdrawChannels 更新withdrawChannels表
// @Tags WithdrawChannels
// @Summary 更新withdrawChannels表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userfund.WithdrawChannels true "更新withdrawChannels表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /withdrawChannels/updateWithdrawChannels [put]
func (withdrawChannelsApi *WithdrawChannelsApi) UpdateWithdrawChannels(c *gin.Context) {
	var withdrawChannels userfund.WithdrawChannels
	err := c.ShouldBindJSON(&withdrawChannels)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 验证关联的货币是否存在
	if withdrawChannels.CoinId != nil {
		var currency userfund.Currencies
		if err := global.GVA_DB.First(&currency, withdrawChannels.CoinId).Error; err != nil {
			response.FailWithMessage("关联的货币不存在", c)
			return
		}
	}

	err = withdrawChannelsService.UpdateWithdrawChannels(withdrawChannels)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindWithdrawChannels 用id查询withdrawChannels表
// @Tags WithdrawChannels
// @Summary 用id查询withdrawChannels表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userfund.WithdrawChannels true "用id查询withdrawChannels表"
// @Success 200 {object} response.Response{data=userfund.WithdrawChannels,msg=string} "查询成功"
// @Router /withdrawChannels/findWithdrawChannels [get]
func (withdrawChannelsApi *WithdrawChannelsApi) FindWithdrawChannels(c *gin.Context) {
	ID := c.Query("ID")
	rewithdrawChannels, err := withdrawChannelsService.GetWithdrawChannels(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(rewithdrawChannels, c)
}

// GetWithdrawChannelsList 分页获取withdrawChannels表列表
// @Tags WithdrawChannels
// @Summary 分页获取withdrawChannels表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userfundReq.WithdrawChannelsSearch true "分页获取withdrawChannels表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /withdrawChannels/getWithdrawChannelsList [get]
func (withdrawChannelsApi *WithdrawChannelsApi) GetWithdrawChannelsList(c *gin.Context) {
	var pageInfo userfundReq.WithdrawChannelsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := withdrawChannelsService.GetWithdrawChannelsInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	for i := range list {
		record := &list[i] // 使用指针
		record.CreatedAtInt = record.CreatedAt.UnixMilli()
		record.UpdatedAtInt = record.UpdatedAt.UnixMilli()
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (withdrawChannelsApi *WithdrawChannelsApi) GetWithdrawChannelsList2(c *gin.Context) {
	var pageInfo userfundReq.WithdrawChannelsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := withdrawChannelsService.GetWithdrawChannelsInfoList2(pageInfo)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.OperateFail, response.ERROR), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, i18n.Message(request.GetLanguageTag(c), i18n.OperateSuccess, 0), c)
}

// GetWithdrawChannelsPublic 不需要鉴权的withdrawChannels表接口
// @Tags WithdrawChannels
// @Summary 不需要鉴权的withdrawChannels表接口
// @accept application/json
// @Produce application/json
// @Param data query userfundReq.WithdrawChannelsSearch true "分页获取withdrawChannels表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /withdrawChannels/getWithdrawChannelsPublic [get]
func (withdrawChannelsApi *WithdrawChannelsApi) GetWithdrawChannelsPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	withdrawChannelsService.GetWithdrawChannelsPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的withdrawChannels表接口信息",
	}, "获取成功", c)
}
