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

type ContractEntrustApi struct{}

// CreateContractEntrust 创建contractEntrust表
// @Tags ContractEntrust
// @Summary 创建contractEntrust表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body contract.ContractEntrust true "创建contractEntrust表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /contractEntrust/createContractEntrust [post]
func (contractEntrustApi *ContractEntrustApi) CreateContractEntrust(c *gin.Context) {
	var contractEntrustReq contractReq.ContractEntrustReq
	err := c.ShouldBindJSON(&contractEntrustReq)
	if err != nil {
		global.GVA_LOG.Error(i18n.Message(request.GetLanguageTag(c), i18n.ContractEntrustError, response.ERROR), zap.Error(err))
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ContractEntrustError, response.ERROR)+err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	err = contractEntrustService.CreateContractEntrust(&contractEntrustReq, userID)
	if err != nil {
		global.GVA_LOG.Error(i18n.Message(request.GetLanguageTag(c), i18n.ContractEntrustError, response.ERROR), zap.Error(err))
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ContractEntrustError, response.ERROR)+err.Error(), c)
		return
	}
	response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ContractEntrustSuccess, response.SUCCESS), c)
}

// DeleteContractEntrust 删除contractEntrust表
// @Tags ContractEntrust
// @Summary 删除contractEntrust表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body contract.ContractEntrust true "删除contractEntrust表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /contractEntrust/deleteContractEntrust [post]
func (contractEntrustApi *ContractEntrustApi) DeleteContractEntrust(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := contractEntrustService.DeleteContractEntrust(ID, userID)
	if err != nil {
		global.GVA_LOG.Error(i18n.Message(request.GetLanguageTag(c), i18n.ContractEntrustDeleteError, response.ERROR), zap.Error(err))
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ContractEntrustDeleteError, response.ERROR)+err.Error(), c)
		return
	}
	response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ContractEntrustDeleteSuccess, response.SUCCESS), c)
}

// DeleteAllContractEntrust 一键撤销
// @Tags ContractEntrust
// @Summary 一键撤销
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /contractEntrust/deleteContractEntrust [post]
func (contractEntrustApi *ContractEntrustApi) DeleteAllContractEntrust(c *gin.Context) {
	userID := utils.GetUserID(c)
	err := contractEntrustService.DeleteAllContractEntrust(userID)
	if err != nil {
		global.GVA_LOG.Error(i18n.Message(request.GetLanguageTag(c), i18n.ContractEntrustDeleteError, response.ERROR), zap.Error(err))
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ContractEntrustDeleteError, response.ERROR)+err.Error(), c)
		return
	}
	response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ContractEntrustDeleteSuccess, response.SUCCESS), c)
}

// DeleteContractEntrustByIds 批量删除contractEntrust表
// @Tags ContractEntrust
// @Summary 批量删除contractEntrust表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /contractEntrust/deleteContractEntrustByIds [delete]
func (contractEntrustApi *ContractEntrustApi) DeleteContractEntrustByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := contractEntrustService.DeleteContractEntrustByIds(IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateContractEntrust 更新contractEntrust表
// @Tags ContractEntrust
// @Summary 更新contractEntrust表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body contract.ContractEntrust true "更新contractEntrust表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /contractEntrust/updateContractEntrust [put]
func (contractEntrustApi *ContractEntrustApi) UpdateContractEntrust(c *gin.Context) {
	var contractEntrust contract.ContractEntrust
	err := c.ShouldBindJSON(&contractEntrust)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	contractEntrust.UpdatedBy = utils.GetUserID(c)
	err = contractEntrustService.UpdateContractEntrust(contractEntrust)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindContractEntrust 用id查询contractEntrust表
// @Tags ContractEntrust
// @Summary 用id查询contractEntrust表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query contract.ContractEntrust true "用id查询contractEntrust表"
// @Success 200 {object} response.Response{data=contract.ContractEntrust,msg=string} "查询成功"
// @Router /contractEntrust/findContractEntrust [get]
func (contractEntrustApi *ContractEntrustApi) FindContractEntrust(c *gin.Context) {
	ID := c.Query("ID")
	recontractEntrust, err := contractEntrustService.GetContractEntrust(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(recontractEntrust, c)
}

// GetContractEntrustList 分页获取contractEntrust表列表
// @Tags ContractEntrust
// @Summary 分页获取contractEntrust表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query contractReq.ContractEntrustSearch true "分页获取contractEntrust表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /contractEntrust/getContractEntrustList [get]
func (contractEntrustApi *ContractEntrustApi) GetContractEntrustList(c *gin.Context) {
	var pageInfo contractReq.ContractEntrustSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		global.GVA_LOG.Error(i18n.Message(request.GetLanguageTag(c), i18n.ContractQueryError, response.ERROR), zap.Error(err))
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ContractQueryError, response.ERROR)+err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	list, total, err := contractEntrustService.GetContractEntrustInfoList(pageInfo, userID)
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

// GetContractEntrustPublic 不需要鉴权的contractEntrust表接口
// @Tags ContractEntrust
// @Summary 不需要鉴权的contractEntrust表接口
// @accept application/json
// @Produce application/json
// @Param data query contractReq.ContractEntrustSearch true "分页获取contractEntrust表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /contractEntrust/getContractEntrustPublic [get]
func (contractEntrustApi *ContractEntrustApi) GetContractEntrustPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	contractEntrustService.GetContractEntrustPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的contractEntrust表接口信息",
	}, "获取成功", c)
}
