package contract

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/i18n"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/contract"
	contractReq "github.com/flipped-aurora/gin-vue-admin/server/model/contract/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ContractLeverageApi struct{}

// CreateContractLeverage 创建contractLeverage表
// @Tags ContractLeverage
// @Summary 创建contractLeverage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body contract.ContractLeverage true "创建contractLeverage表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /contractLeverage/createContractLeverage [post]
func (contractLeverageApi *ContractLeverageApi) CreateContractLeverage(c *gin.Context) {
	var contractLeverage contract.ContractLeverage
	err := c.ShouldBindJSON(&contractLeverage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	contractLeverage.CreatedBy = utils.GetUserID(c)
	err = contractLeverageService.CreateContractLeverage(&contractLeverage)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteContractLeverage 删除contractLeverage表
// @Tags ContractLeverage
// @Summary 删除contractLeverage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body contract.ContractLeverage true "删除contractLeverage表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /contractLeverage/deleteContractLeverage [delete]
func (contractLeverageApi *ContractLeverageApi) DeleteContractLeverage(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := contractLeverageService.DeleteContractLeverage(ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteContractLeverageByIds 批量删除contractLeverage表
// @Tags ContractLeverage
// @Summary 批量删除contractLeverage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /contractLeverage/deleteContractLeverageByIds [delete]
func (contractLeverageApi *ContractLeverageApi) DeleteContractLeverageByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := contractLeverageService.DeleteContractLeverageByIds(IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateContractLeverage 更新contractLeverage表
// @Tags ContractLeverage
// @Summary 更新contractLeverage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body contract.ContractLeverage true "更新contractLeverage表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /contractLeverage/updateContractLeverage [post]
func (contractLeverageApi *ContractLeverageApi) UpdateContractLeverage(c *gin.Context) {
	var contractLeverage contract.ContractLeverage
	err := c.ShouldBindJSON(&contractLeverage)
	if err != nil {
		global.GVA_LOG.Error(i18n.Message(request.GetLanguageTag(c), i18n.ContractLeverageUpdateError, response.ERROR), zap.Error(err))
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ContractLeverageUpdateError, response.ERROR)+err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	err = contractLeverageService.UpdateContractLeverage(&contractLeverage, userID)
	if err != nil {
		global.GVA_LOG.Error(i18n.Message(request.GetLanguageTag(c), i18n.ContractLeverageUpdateError, response.ERROR), zap.Error(err))
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ContractLeverageUpdateError, response.ERROR)+err.Error(), c)
		return
	}
	response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ContractLeverageUpdateSuccess, response.ERROR), c)
}

// FindContractLeverage 用id查询contractLeverage表
// @Tags ContractLeverage
// @Summary 用id查询contractLeverage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query contract.ContractLeverage true "用id查询contractLeverage表"
// @Success 200 {object} response.Response{data=contract.ContractLeverage,msg=string} "查询成功"
// @Router /contractLeverage/findContractLeverage [get]
func (contractLeverageApi *ContractLeverageApi) FindContractLeverage(c *gin.Context) {
	stockId := c.Query("stockId")
	userID := utils.GetUserID(c)
	recontractLeverage, err := contractLeverageService.GetContractLeverage(stockId, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.OkWithData(100, c)
			return
		}
		global.GVA_LOG.Error(i18n.Message(request.GetLanguageTag(c), i18n.ContractQueryError, response.ERROR), zap.Error(err))
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ContractQueryError, response.ERROR)+err.Error(), c)
		return
	}
	response.OkWithData(recontractLeverage.LeverageRatio, c)
}

// GetContractLeverageList 分页获取contractLeverage表列表
// @Tags ContractLeverage
// @Summary 分页获取contractLeverage表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query contractReq.ContractLeverageSearch true "分页获取contractLeverage表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /contractLeverage/getContractLeverageList [get]
func (contractLeverageApi *ContractLeverageApi) GetContractLeverageList(c *gin.Context) {
	var pageInfo contractReq.ContractLeverageSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := contractLeverageService.GetContractLeverageInfoList(pageInfo)
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

// GetContractLeveragePublic 不需要鉴权的contractLeverage表接口
// @Tags ContractLeverage
// @Summary 不需要鉴权的contractLeverage表接口
// @accept application/json
// @Produce application/json
// @Param data query contractReq.ContractLeverageSearch true "分页获取contractLeverage表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /contractLeverage/getContractLeveragePublic [get]
func (contractLeverageApi *ContractLeverageApi) GetContractLeveragePublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	contractLeverageService.GetContractLeveragePublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的contractLeverage表接口信息",
	}, "获取成功", c)
}
