package contract

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/i18n"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/contract"
	contractReq "github.com/flipped-aurora/gin-vue-admin/server/model/contract/request"
	contractRes "github.com/flipped-aurora/gin-vue-admin/server/model/contract/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	. "github.com/shopspring/decimal"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ContractAccountApi struct{}

// CreateContractAccount 创建contractAccount表
// @Tags ContractAccount
// @Summary 创建contractAccount表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body contract.ContractAccount true "创建contractAccount表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /contractAccount/createContractAccount [post]
func (contractAccountApi *ContractAccountApi) CreateContractAccount(c *gin.Context) {
	//var contractAccount contract.ContractAccount
	//err := c.ShouldBindJSON(&contractAccount)
	//if err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	userID := utils.GetUserID(c)
	_, err := contractAccountService.GetContractAccountByUserId(userID)
	if err == nil {
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ContractAccountExist, response.ERROR), c)
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error(i18n.Message(request.GetLanguageTag(c), i18n.Failed, response.ERROR), zap.Error(err))
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Failed, response.ERROR)+err.Error(), c)
		return
	}
	// 初始化合约账户
	contractAccount := contract.ContractAccount{
		GVA_MODEL: global.GVA_MODEL{
			ID: uint(global.Snowflake.Generate()),
		},
		UserId:             userID,
		TotalMargin:        Zero,
		AvailableMargin:    Zero,
		FrozenMargin:       Zero,
		UsedMargin:         Zero,
		RealizedProfitLoss: Zero,
		AccountStatus:      contract.Unreviewed,
		CreatedBy:          userID,
		UpdatedBy:          0,
		DeletedBy:          0,
	}
	err = contractAccountService.CreateContractAccount(&contractAccount)
	if err != nil {
		global.GVA_LOG.Error(i18n.Message(request.GetLanguageTag(c), i18n.Failed, response.ERROR), zap.Error(err))
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Failed, response.ERROR)+err.Error(), c)
		return
	}

	response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}

// DeleteContractAccount 删除contractAccount表
// @Tags ContractAccount
// @Summary 删除contractAccount表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body contract.ContractAccount true "删除contractAccount表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /contractAccount/deleteContractAccount [delete]
func (contractAccountApi *ContractAccountApi) DeleteContractAccount(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := contractAccountService.DeleteContractAccount(ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteContractAccountByIds 批量删除contractAccount表
// @Tags ContractAccount
// @Summary 批量删除contractAccount表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /contractAccount/deleteContractAccountByIds [delete]
func (contractAccountApi *ContractAccountApi) DeleteContractAccountByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := contractAccountService.DeleteContractAccountByIds(IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateContractAccount 更新contractAccount表
// @Tags ContractAccount
// @Summary 更新contractAccount表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body contract.ContractAccount true "更新contractAccount表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /contractAccount/updateContractAccount [put]
func (contractAccountApi *ContractAccountApi) UpdateContractAccount(c *gin.Context) {
	var contractAccount contract.ContractAccount
	err := c.ShouldBindJSON(&contractAccount)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	contractAccount.UpdatedBy = utils.GetUserID(c)
	err = contractAccountService.UpdateContractAccount(contractAccount)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindContractAccount 用id查询contractAccount表
// @Tags ContractAccount
// @Summary 用id查询contractAccount表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query contract.ContractAccount true "用id查询contractAccount表"
// @Success 200 {object} response.Response{data=contract.ContractAccount,msg=string} "查询成功"
// @Router /contractAccount/findContractAccount [get]
func (contractAccountApi *ContractAccountApi) FindContractAccount(c *gin.Context) {
	// ID := c.Query("ID")
	userID := utils.GetUserID(c)
	recontactAccount, err := contractAccountService.GetContractAccountByUserId(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			contractAccount := contractRes.ContractAccountRes{
				TotalAmount:          Zero,
				Balance:              Zero,
				RealizedProfitLoss:   Zero,
				UnrealizedProfitLoss: Zero,
				AccountStatus:        0,
			}
			response.OkWithData(contractAccount, c)
			return
		} else {
			global.GVA_LOG.Error(i18n.Message(request.GetLanguageTag(c), i18n.ContractQueryError, response.ERROR), zap.Error(err))
			response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ContractQueryError, response.ERROR)+err.Error(), c)
			return
		}
	}
	unrealizedProfitLoss, err := contractPositionService.GetUnrealizedProfitLoss(userID)
	if err != nil {
		global.GVA_LOG.Error(i18n.Message(request.GetLanguageTag(c), i18n.ContractQueryError, response.ERROR), zap.Error(err))
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ContractQueryError, response.ERROR)+err.Error(), c)
		return
	}
	// 开仓只能用保证金和未实现盈亏的80%
	availableMargin := recontactAccount.AvailableMargin.Add(unrealizedProfitLoss).Mul(NewFromFloat(0.8)).RoundFloor(2)
	// 可划转金额
	transferableAmount := contractAccountService.GetTransferableAmount(recontactAccount.TotalMargin, recontactAccount.AvailableMargin, unrealizedProfitLoss)
	// 返回数据
	contractAccount := contractRes.ContractAccountRes{
		TotalAmount:          recontactAccount.TotalMargin.Add(unrealizedProfitLoss),
		Balance:              recontactAccount.TotalMargin,
		RealizedProfitLoss:   recontactAccount.RealizedProfitLoss,
		UnrealizedProfitLoss: unrealizedProfitLoss,
		AvailableMargin:      availableMargin,
		TransferableAmount:   transferableAmount,
		AccountStatus:        recontactAccount.AccountStatus,
	}
	response.OkWithData(contractAccount, c)
}

// GetContractAccountList 分页获取contractAccount表列表
// @Tags ContractAccount
// @Summary 分页获取contractAccount表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query contractReq.ContractAccountSearch true "分页获取contractAccount表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /contractAccount/getContractAccountList [get]
func (contractAccountApi *ContractAccountApi) GetContractAccountList(c *gin.Context) {
	var pageInfo contractReq.ContractAccountSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := contractAccountService.GetContractAccountInfoList(pageInfo)
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

// GetContractAccountPublic 不需要鉴权的contractAccount表接口
// @Tags ContractAccount
// @Summary 不需要鉴权的contractAccount表接口
// @accept application/json
// @Produce application/json
// @Param data query contractReq.ContractAccountSearch true "分页获取contractAccount表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /contractAccount/getContractAccountPublic [get]
func (contractAccountApi *ContractAccountApi) GetContractAccountPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	contractAccountService.GetContractAccountPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的contractAccount表接口信息",
	}, "获取成功", c)
}

// ChangeAccountMargin 合约资金划转
// @Tags ContractAccount
// @Summary 合约资金划转
// @accept application/json
// @Produce application/json
// @Param data body contractReq.ChangeMarginReq true "成功"
// @Success 200 {object} response.Response{msg=string} "成功"
// @Router /contractAccount/changeAccountMargin [POST]
func (contractAccountApi *ContractAccountApi) ChangeAccountMargin(c *gin.Context) {
	var ChangeMarginReq contractReq.ChangeMarginReq
	err := c.ShouldBindJSON(&ChangeMarginReq)
	if err != nil {
		global.GVA_LOG.Error(i18n.Message(request.GetLanguageTag(c), i18n.ContractMarginChangeError, response.ERROR), zap.Error(err))
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ContractMarginChangeError, response.ERROR)+err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	// 请添加自己的业务逻辑
	err = contractAccountService.ChangeAccountMargin(ChangeMarginReq, userID)
	if err != nil {
		global.GVA_LOG.Error(i18n.Message(request.GetLanguageTag(c), i18n.ContractMarginChangeError, response.ERROR), zap.Error(err))
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ContractMarginChangeError, response.ERROR)+err.Error(), c)
		return
	}
	response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ContractMarginChangeSuccess, response.SUCCESS), c)
}
