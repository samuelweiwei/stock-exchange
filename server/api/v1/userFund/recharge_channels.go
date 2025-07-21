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

type RechargeChannelsApi struct{}

type AAPayNotifyReq struct {
	Code      int             `json:"code"`
	Message   string          `json:"message"`
	Sign      string          `json:"sign"`
	Timestamp int64           `json:"timestamp"`
	Data      AAPayNotifyData `json:"data"`
}

type AAPayNotifyData struct {
	OrderId              int    `json:"order_id"`
	OrderStatus          int    `json:"order_status"`
	CurrencyCode         string `json:"currency_code"`
	CoinCode             string `json:"coin_code"`
	CoinAddress          string `json:"coin_address"`
	UserOrderId          string `json:"user_order_id"`
	CoinReceiptMoney     string `json:"coin_receipt_money"`
	CurrencyReceiptMoney string `json:"currency_receipt_money"`
}

// CreateRechargeChannels 创建rechargeChannels表
// @Tags RechargeChannels
// @Summary 创建rechargeChannels表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userfund.RechargeChannels true "创建rechargeChannels表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /rechargeChannels/createRechargeChannels [post]
func (rechargeChannelsApi *RechargeChannelsApi) CreateRechargeChannels(c *gin.Context) {
	var rechargeChannels userfund.RechargeChannels
	err := c.ShouldBindJSON(&rechargeChannels)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 验证关联的货币是否存在
	if rechargeChannels.CoinId == nil {
		response.FailWithMessage("货币ID不能为空", c)
		return
	}

	var currency userfund.Currencies
	if err := global.GVA_DB.First(&currency, rechargeChannels.CoinId).Error; err != nil {
		response.FailWithMessage("关联的货币不存在", c)
		return
	}

	err = rechargeChannelsService.CreateRechargeChannels(&rechargeChannels)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteRechargeChannels 删除rechargeChannels表
// @Tags RechargeChannels
// @Summary 删除rechargeChannels表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userfund.RechargeChannels true "删除rechargeChannels表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /rechargeChannels/deleteRechargeChannels [delete]
func (rechargeChannelsApi *RechargeChannelsApi) DeleteRechargeChannels(c *gin.Context) {
	ID := c.Query("ID")
	err := rechargeChannelsService.DeleteRechargeChannels(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteRechargeChannelsByIds 批量删除rechargeChannels表
// @Tags RechargeChannels
// @Summary 批量删除rechargeChannels表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /rechargeChannels/deleteRechargeChannelsByIds [delete]
func (rechargeChannelsApi *RechargeChannelsApi) DeleteRechargeChannelsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := rechargeChannelsService.DeleteRechargeChannelsByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateRechargeChannels 更新rechargeChannels表
// @Tags RechargeChannels
// @Summary 更新rechargeChannels表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userfund.RechargeChannels true "更新rechargeChannels表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /rechargeChannels/updateRechargeChannels [put]
func (rechargeChannelsApi *RechargeChannelsApi) UpdateRechargeChannels(c *gin.Context) {
	var rechargeChannels userfund.RechargeChannels
	err := c.ShouldBindJSON(&rechargeChannels)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 验证关联的货币是否存在
	if rechargeChannels.CoinId != nil {
		var currency userfund.Currencies
		if err := global.GVA_DB.First(&currency, rechargeChannels.CoinId).Error; err != nil {
			response.FailWithMessage("关联的货币不存在", c)
			return
		}
	}

	err = rechargeChannelsService.UpdateRechargeChannels(rechargeChannels)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindRechargeChannels 用id查询rechargeChannels表
// @Tags RechargeChannels
// @Summary 用id查询rechargeChannels表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userfund.RechargeChannels true "用id查询rechargeChannels表"
// @Success 200 {object} response.Response{data=userfund.RechargeChannels,msg=string} "查询成功"
// @Router /rechargeChannels/findRechargeChannels [get]
func (rechargeChannelsApi *RechargeChannelsApi) FindRechargeChannels(c *gin.Context) {
	ID := c.Query("ID")
	rerechargeChannels, err := rechargeChannelsService.GetRechargeChannels(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(rerechargeChannels, c)
}

// GetRechargeChannelsList 分页获取rechargeChannels表列表
// @Tags RechargeChannels
// @Summary 分页获取rechargeChannels表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userfundReq.RechargeChannelsSearch true "分页获取rechargeChannels表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /rechargeChannels/getRechargeChannelsList [get]
func (rechargeChannelsApi *RechargeChannelsApi) GetRechargeChannelsList(c *gin.Context) {
	var pageInfo userfundReq.RechargeChannelsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := rechargeChannelsService.GetRechargeChannelsInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	for i := range list {
		record := &list[i] // 使用指针
		record.CreateAtInt = record.CreatedAt.UnixMilli()
		record.UpdateAtInt = record.UpdatedAt.UnixMilli()
	}
	// 处理返回数据，确保货币信息正确展示
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (rechargeChannelsApi *RechargeChannelsApi) GetRechargeChannelsList2(c *gin.Context) {
	var pageInfo userfundReq.RechargeChannelsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := rechargeChannelsService.GetRechargeChannelsInfoList2(pageInfo)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.OperateFail, response.ERROR), c)
		return
	}

	// 处理返回数据，确保货币信息正确展示
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, i18n.Message(request.GetLanguageTag(c), i18n.OperateSuccess, 0), c)
}

// GetRechargeChannelsPublic 不需要鉴权的rechargeChannels表接口
// @Tags RechargeChannels
// @Summary 不需要鉴权的rechargeChannels表接口
// @accept application/json
// @Produce application/json
// @Param data query userfundReq.RechargeChannelsSearch true "分页获取rechargeChannels表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /rechargeChannels/getRechargeChannelsPublic [get]
func (rechargeChannelsApi *RechargeChannelsApi) GetRechargeChannelsPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	rechargeChannelsService.GetRechargeChannelsPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的rechargeChannels表接口信息",
	}, "获取成功", c)
}
