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

type ContractPositionApi struct{}

// CreateContractPosition 创建contractPosition表
// @Tags ContractPosition
// @Summary 创建contractPosition表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body contract.ContractPosition true "创建contractPosition表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /contractPosition/createContractPosition [post]
func (contractPositionApi *ContractPositionApi) CreateContractPosition(c *gin.Context) {
	var contractPosition contract.ContractPosition
	err := c.ShouldBindJSON(&contractPosition)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	contractPosition.CreatedBy = utils.GetUserID(c)
	err = contractPositionService.CreateContractPosition(&contractPosition)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteContractPosition 删除contractPosition表
// @Tags ContractPosition
// @Summary 删除contractPosition表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body contract.ContractPosition true "删除contractPosition表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /contractPosition/deleteContractPosition [delete]
func (contractPositionApi *ContractPositionApi) DeleteContractPosition(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := contractPositionService.DeleteContractPosition(ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteContractPositionByIds 批量删除contractPosition表
// @Tags ContractPosition
// @Summary 批量删除contractPosition表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /contractPosition/deleteContractPositionByIds [delete]
func (contractPositionApi *ContractPositionApi) DeleteContractPositionByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := contractPositionService.DeleteContractPositionByIds(IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateContractPosition 更新contractPosition表
// @Tags ContractPosition
// @Summary 更新contractPosition表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body contract.ContractPosition true "更新contractPosition表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /contractPosition/updateContractPosition [put]
func (contractPositionApi *ContractPositionApi) UpdateContractPosition(c *gin.Context) {
	var contractPosition contract.ContractPosition
	err := c.ShouldBindJSON(&contractPosition)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	contractPosition.UpdatedBy = utils.GetUserID(c)
	err = contractPositionService.UpdateContractPosition(contractPosition)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindContractPosition 用id查询contractPosition表
// @Tags ContractPosition
// @Summary 用id查询contractPosition表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query contract.ContractPosition true "用id查询contractPosition表"
// @Success 200 {object} response.Response{data=contract.ContractPosition,msg=string} "查询成功"
// @Router /contractPosition/findContractPosition [get]
func (contractPositionApi *ContractPositionApi) FindContractPosition(c *gin.Context) {
	ID := c.Query("ID")
	recontractPosition, err := contractPositionService.GetContractPosition(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(recontractPosition, c)
}

// GetContractPositionList 分页获取contractPosition表列表
// @Tags ContractPosition
// @Summary 分页获取contractPosition表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query contractReq.ContractPositionSearch true "分页获取contractPosition表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /contractPosition/getContractPositionList [get]
func (contractPositionApi *ContractPositionApi) GetContractPositionList(c *gin.Context) {
	var pageInfo contractReq.ContractPositionSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		global.GVA_LOG.Error(i18n.Message(request.GetLanguageTag(c), i18n.ContractQueryError, response.ERROR), zap.Error(err))
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ContractQueryError, response.ERROR)+err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	list, total, err := contractPositionService.GetContractPositionInfoList(pageInfo, userID)
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

// GetContractPositionPublic 不需要鉴权的contractPosition表接口
// @Tags ContractPosition
// @Summary 不需要鉴权的contractPosition表接口
// @accept application/json
// @Produce application/json
// @Param data query contractReq.ContractPositionSearch true "分页获取contractPosition表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /contractPosition/getContractPositionPublic [get]
func (contractPositionApi *ContractPositionApi) GetContractPositionPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	contractPositionService.GetContractPositionPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的contractPosition表接口信息",
	}, "获取成功", c)
}
