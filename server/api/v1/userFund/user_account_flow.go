package userFund

import (
	"github.com/flipped-aurora/gin-vue-admin/server/enums/fund"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/i18n"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/userfund"
	userfundReq "github.com/flipped-aurora/gin-vue-admin/server/model/userfund/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	. "github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type UserAccountFlowApi struct{}

// CreateUserAccountFlow 创建userAccountFlow表
// @Tags UserAccountFlow
// @Summary 创建userAccountFlow表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userfund.UserAccountFlow true "创建userAccountFlow表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /userAccountFlow/createUserAccountFlow [post]
func (userAccountFlowApi *UserAccountFlowApi) CreateUserAccountFlow(c *gin.Context) {
	var userAccountFlow userfund.UserAccountFlow
	err := c.ShouldBindJSON(&userAccountFlow)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userAccountFlowService.CreateUserAccountFlow(&userAccountFlow)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteUserAccountFlow 删除userAccountFlow表
// @Tags UserAccountFlow
// @Summary 删除userAccountFlow表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userfund.UserAccountFlow true "删除userAccountFlow表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /userAccountFlow/deleteUserAccountFlow [delete]
func (userAccountFlowApi *UserAccountFlowApi) DeleteUserAccountFlow(c *gin.Context) {
	ID := c.Query("ID")
	err := userAccountFlowService.DeleteUserAccountFlow(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteUserAccountFlowByIds 批量删除userAccountFlow表
// @Tags UserAccountFlow
// @Summary 批量删除userAccountFlow表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /userAccountFlow/deleteUserAccountFlowByIds [delete]
func (userAccountFlowApi *UserAccountFlowApi) DeleteUserAccountFlowByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := userAccountFlowService.DeleteUserAccountFlowByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateUserAccountFlow 更新userAccountFlow表
// @Tags UserAccountFlow
// @Summary 更新userAccountFlow表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userfund.UserAccountFlow true "更新userAccountFlow表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /userAccountFlow/updateUserAccountFlow [put]
func (userAccountFlowApi *UserAccountFlowApi) UpdateUserAccountFlow(c *gin.Context) {
	var userAccountFlow userfund.UserAccountFlow
	err := c.ShouldBindJSON(&userAccountFlow)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userAccountFlowService.UpdateUserAccountFlow(userAccountFlow)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindUserAccountFlow 用id查询userAccountFlow表
// @Tags UserAccountFlow
// @Summary 用id查询userAccountFlow表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userfund.UserAccountFlow true "用id查询userAccountFlow表"
// @Success 200 {object} response.Response{data=userfund.UserAccountFlow,msg=string} "查询成功"
// @Router /userAccountFlow/findUserAccountFlow [get]
func (userAccountFlowApi *UserAccountFlowApi) FindUserAccountFlow(c *gin.Context) {
	ID := c.Query("ID")
	reuserAccountFlow, err := userAccountFlowService.GetUserAccountFlow(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reuserAccountFlow, c)
}

// GetAccountFlowList 获取账户流水列表
// @Tags UserAccountFlow
// @Summary 获取账户流水列表
// @Description 根据交易类型获取用户账户流水列表，transactionType 为必填参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userfundReq.UserAccountFlowSearch true "查询参数(transactionType: 交易类型[必填], username: 用户名[选填])"
// @Success 200 {object} response.Response{data=response.PageResult{list=[]userfund.UserAccountFlowUnion},msg=string} "获取成功"
// @Router /userAccountFlow/getFlowList [get]
func (userAccountFlowApi *UserAccountFlowApi) GetAccountFlowList(c *gin.Context) {
	var pageInfo userfundReq.UserAccountFlowSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	pageInfo.UserId = int(utils.GetUserIDFrontUser(c))
	list, total, err := userAccountFlowService.GetUserAccountFlowInfoList(pageInfo)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.OperateFail, response.ERROR), c)
		return
	}

	lan := request.GetLanguageTag(c)
	for i := range list {
		// 判断是否是数字货币类型，
		flow := &list[i]
		action, _ := fund.GetActionTypeFromString(flow.TransactionType)
		flow.Amount = getChangeAmount(action, flow.Amount)
		flow.TransactionDateInt = flow.TransactionDate.UnixMilli()
		flow.TransactionTypeI18n = i18n.Message(lan, i18n.FundChangeType+flow.TransactionType, 0)
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, i18n.Message(lan, i18n.OperateSuccess, 0), c)
}

func getChangeAmount(actionType fund.ActionType, amount Decimal) Decimal {
	// 根据 actionType 判断变动金额
	switch actionType {
	case fund.Recharge:
		return amount
	case fund.Withdraw:
		return amount.Neg()
	case fund.TransferToContract:
		return amount.Neg()
	case fund.TransferFromContract:
		return amount
	case fund.TradeFollow:
		return amount.Neg()
	case fund.CancelTradeFollow:
		return amount
	case fund.OperationRefused:
		return amount
	case fund.SettleProfit:
		return amount
	case fund.AutoSettle:
		return amount
	case fund.ProfitSharing:
		return amount
	case fund.SysSend:
		return amount
	case fund.ApplyOrderFollow:
		return amount.Neg()
	case fund.RefusedApplyOrderFollow:
		return amount
	case fund.StakeEarnProduct:
		return amount.Neg()
	case fund.RedeemEarnProduct:
		return amount
	case fund.ApplyWithdraw:
		return amount.Neg()
	default:
		return Zero // 未知操作返回 0
	}
}

// GetUserAccountFlowPublic 不需要鉴权的userAccountFlow表接口
// @Tags UserAccountFlow
// @Summary 不需要鉴权的userAccountFlow表接口
// @accept application/json
// @Produce application/json
// @Param data query userfundReq.UserAccountFlowSearch true "分页获取userAccountFlow表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /userAccountFlow/getUserAccountFlowPublic [get]
func (userAccountFlowApi *UserAccountFlowApi) GetUserAccountFlowPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	userAccountFlowService.GetUserAccountFlowPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的userAccountFlow表接口信息",
	}, "获取成功", c)
}

// GetUserAccountFlowList 分页获取userAccountFlow表列表
// @Tags UserAccountFlow
// @Summary 分页获取账户流水列表
// @Description 管理端获取账户流水列表，支持多条件查询
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userfundReq.UserAccountFlowSearch true "查询参数(transactionType: 交易类型, username: 用户名)"
// @Success 200 {object} response.Response{data=response.PageResult{list=[]userfund.UserAccountFlowUnion},msg=string} "获取成功"
// @Router /userAccountFlow/getUserAccountFlowList [get]
func (userAccountFlowApi *UserAccountFlowApi) GetUserAccountFlowList(c *gin.Context) {
	var pageInfo userfundReq.UserAccountFlowSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userAccountFlowService.GetUserAccountFlowInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	for i := range list {
		// 判断是否是数字货币类型，
		flow := &list[i]
		action, _ := fund.GetActionTypeFromString(flow.TransactionType)
		flow.Amount = getChangeAmount(action, flow.Amount)
		flow.TransactionDateInt = flow.TransactionDate.UnixMilli()
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetFundChangeTypes 获取所有资金变动类型的国际化值
// @Tags UserAccountFlow
// @Summary 获取所有资金变动类型的国际化值
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=[]map[string]interface{},msg=string} "获取成功"
// @Router /userAccountFlow/getFundChangeTypes [get]
func (userAccountFlowApi *UserAccountFlowApi) GetFundChangeTypes(c *gin.Context) {
	lan := request.GetLanguageTag(c)
	result := make([]map[string]interface{}, 0)

	actionTypes := []fund.ActionType{
		fund.Recharge, fund.Withdraw, fund.TransferToContract, fund.TransferFromContract,
		fund.TradeFollow, fund.CancelTradeFollow, fund.OperationRefused, fund.SettleProfit,
		fund.AutoSettle, fund.ProfitSharing, fund.SysSend, fund.ApplyOrderFollow,
		fund.RefusedApplyOrderFollow, fund.StakeEarnProduct, fund.RedeemEarnProduct, fund.ApplyWithdraw,
	}

	for i, actionType := range actionTypes {
		result = append(result, map[string]interface{}{
			"value": i + 1,
			"label": i18n.Message(lan, i18n.FundChangeType+string(actionType), 0),
		})
	}

	response.OkWithData(result, c)
}
