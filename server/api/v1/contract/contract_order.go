package contract

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/i18n"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/contract"
	contractReq "github.com/flipped-aurora/gin-vue-admin/server/model/contract/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ContractOrderApi struct{}

// CreateContractOrder 创建contractOrder表
// @Tags ContractOrder
// @Summary 创建contractOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body contractReq.ContractOrderReq true "创建contractOrder表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /contractOrder/createContractOrder [post]
func (contractOrderApi *ContractOrderApi) CreateContractOrder(c *gin.Context) {
	var contractOrderReq contractReq.ContractOrderReq
	err := c.ShouldBindJSON(&contractOrderReq)
	if err != nil {
		global.GVA_LOG.Error(i18n.Message(request.GetLanguageTag(c), i18n.ContractOrderOpenError, response.ERROR), zap.Error(err))
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ContractOrderOpenError, response.ERROR)+err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	contractAccount, err := contractAccountService.GetContractAccountByUserId(userID)
	if err != nil {
		global.GVA_LOG.Error(i18n.Message(request.GetLanguageTag(c), i18n.ContractOrderOpenError, response.ERROR), zap.Error(err))
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ContractOrderOpenError, response.ERROR)+err.Error(), c)
		return
	}
	err = contractOrderService.CreateContractOrder(&contractOrderReq, &contractAccount, userID)
	if err != nil {
		global.GVA_LOG.Error(i18n.Message(request.GetLanguageTag(c), i18n.ContractOrderOpenError, response.ERROR), zap.Error(err))
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ContractOrderOpenError, response.ERROR)+err.Error(), c)
		return
	}
	response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ContractOrderOpenSuccess, response.SUCCESS), c)
}

// DeleteContractOrder 删除contractOrder表
// @Tags ContractOrder
// @Summary 删除contractOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body contract.ContractOrder true "删除contractOrder表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /contractOrder/deleteContractOrder [delete]
func (contractOrderApi *ContractOrderApi) DeleteContractOrder(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := contractOrderService.DeleteContractOrder(ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteContractOrderByIds 批量删除contractOrder表
// @Tags ContractOrder
// @Summary 批量删除contractOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /contractOrder/deleteContractOrderByIds [delete]
func (contractOrderApi *ContractOrderApi) DeleteContractOrderByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := contractOrderService.DeleteContractOrderByIds(IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateContractOrder 更新contractOrder表
// @Tags ContractOrder
// @Summary 更新contractOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body contract.ContractOrder true "更新contractOrder表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /contractOrder/updateContractOrder [put]
func (contractOrderApi *ContractOrderApi) UpdateContractOrder(c *gin.Context) {
	var contractOrder contract.ContractOrder
	err := c.ShouldBindJSON(&contractOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	contractOrder.UpdatedBy = utils.GetUserID(c)
	err = contractOrderService.UpdateContractOrder(contractOrder)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindContractOrder 用id查询contractOrder表
// @Tags ContractOrder
// @Summary 用id查询contractOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query contract.ContractOrder true "用id查询contractOrder表"
// @Success 200 {object} response.Response{data=contract.ContractOrder,msg=string} "查询成功"
// @Router /contractOrder/findContractOrder [get]
func (contractOrderApi *ContractOrderApi) FindContractOrder(c *gin.Context) {
	ID := c.Query("ID")
	recontractOrder, err := contractOrderService.GetContractOrder(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(recontractOrder, c)
}

// GetContractOrderList 分页获取contractOrder表列表
// @Tags ContractOrder
// @Summary 分页获取contractOrder表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query contractReq.ContractOrderSearch true "分页获取contractOrder表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /contractOrder/getContractOrderList [get]
func (contractOrderApi *ContractOrderApi) GetContractOrderList(c *gin.Context) {
	var pageInfo contractReq.ContractOrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		global.GVA_LOG.Error(i18n.Message(request.GetLanguageTag(c), i18n.ContractQueryError, response.ERROR), zap.Error(err))
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ContractQueryError, response.ERROR)+err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	list, total, err := contractOrderService.GetContractOrderInfoList(pageInfo, userID)
	if err != nil {
		global.GVA_LOG.Error(i18n.Message(request.GetLanguageTag(c), i18n.ContractQueryError, response.ERROR), zap.Error(err))
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ContractQueryError, response.ERROR)+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, i18n.Message(request.GetLanguageTag(c), i18n.ContractQuerySuccess, response.SUCCESS), c)
}

// GetContractOrderPublic 不需要鉴权的contractOrder表接口
// @Tags ContractOrder
// @Summary 不需要鉴权的contractOrder表接口
// @accept application/json
// @Produce application/json
// @Param data query contractReq.ContractOrderSearch true "分页获取contractOrder表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /contractOrder/getContractOrderPublic [get]
func (contractOrderApi *ContractOrderApi) GetContractOrderPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	contractOrderService.GetContractOrderPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的contractOrder表接口信息",
	}, "获取成功", c)
}

// CloseContractOrder 平仓方法
// @Tags ContractOrder
// @Summary 平仓方法
// @accept application/json
// @Produce application/json
// @Param data body contractReq.ContractOrderReq true "成功"
// @Success 200 {object} response.Response{msg=string} "成功"
// @Router /contractOrder/closeContractOrder [POST]
func (contractOrderApi *ContractOrderApi) CloseContractOrder(c *gin.Context) {
	var contractCLoseReq contractReq.ContractCloseReq
	err := c.ShouldBindJSON(&contractCLoseReq)
	if err != nil {
		global.GVA_LOG.Error(i18n.Message(request.GetLanguageTag(c), i18n.ContractOrderCloseError, response.ERROR), zap.Error(err))
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ContractOrderCloseError, response.ERROR)+err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	contractAccount, err := contractAccountService.GetContractAccountByUserId(userID)
	if err != nil {
		global.GVA_LOG.Error(i18n.Message(request.GetLanguageTag(c), i18n.ContractOrderCloseError, response.ERROR), zap.Error(err))
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ContractOrderCloseError, response.ERROR)+err.Error(), c)
		return
	}
	err = contractOrderService.CloseContractOrder(&contractCLoseReq, &contractAccount, userID)
	if err != nil {
		global.GVA_LOG.Error(i18n.Message(request.GetLanguageTag(c), i18n.ContractOrderCloseError, response.ERROR), zap.Error(err))
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ContractOrderCloseError, response.ERROR)+err.Error(), c)
		return
	}
	response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ContractOrderCloseSuccess, response.SUCCESS), c)
}

// CloseAllContractOrder 一键平仓方法
// @Tags CloseAllContractOrder
// @Summary 一键平仓方法
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "成功"
// @Router /contractOrder/closeAllContractOrder [POST]
func (contractOrderApi *ContractOrderApi) CloseAllContractOrder(c *gin.Context) {
	userID := utils.GetUserID(c)
	contractAccount, err := contractAccountService.GetContractAccountByUserId(userID)
	if err != nil {
		global.GVA_LOG.Error(i18n.Message(request.GetLanguageTag(c), i18n.ContractOrderCloseError, response.ERROR), zap.Error(err))
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ContractOrderCloseError, response.ERROR)+err.Error(), c)
		return
	}
	err = contractOrderService.CloseAllContractOrder(&contractAccount, userID)
	if err != nil {
		global.GVA_LOG.Error(i18n.Message(request.GetLanguageTag(c), i18n.ContractOrderCloseError, response.ERROR), zap.Error(err))
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ContractOrderCloseError, response.ERROR)+err.Error(), c)
		return
	}
	response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ContractOrderCloseSuccess, response.SUCCESS), c)
}
