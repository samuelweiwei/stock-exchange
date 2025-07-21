package earn

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/enums/fund"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/i18n"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/earn"
	earnReq "github.com/flipped-aurora/gin-vue-admin/server/model/earn/request"
	earnRes "github.com/flipped-aurora/gin-vue-admin/server/model/earn/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/service/userfund"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"time"
)

type EarnSubscribeLogApi struct{}

// CreateEarnSubscribeLog 后台增加理财质押记录
// @Tags EarnProductionSubscribeLog
// @Summary 后台增加理财质押记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body earn.EarnSubscribeLog true "创建earnSubscribeLog表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /earn/subscribe/log/create [post]
func (earnSubscribeLogApi *EarnSubscribeLogApi) CreateEarnSubscribeLog(c *gin.Context) {
	var earnSubscribeLog earn.EarnSubscribeLog
	err := c.ShouldBindJSON(&earnSubscribeLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = earnSubscribeLogService.CreateEarnSubscribeLog(nil, &earnSubscribeLog)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteEarnSubscribeLog 后台删除理财质押记录
// @Tags EarnProductionSubscribeLog
// @Summary 后台删除理财质押记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body earn.EarnSubscribeLog true "删除earnSubscribeLog表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /earn/subscribe/log/delete [delete]
func (earnSubscribeLogApi *EarnSubscribeLogApi) DeleteEarnSubscribeLog(c *gin.Context) {
	id := c.Query("id")
	err := earnSubscribeLogService.DeleteEarnSubscribeLog(id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteEarnSubscribeLogByIds 批量删除earnSubscribeLog表
// @Tags EarnProductionSubscribeLog
// @Summary 批量删除earnSubscribeLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /earnSubscribeLog/deleteEarnSubscribeLogByIds [delete]
func (earnSubscribeLogApi *EarnSubscribeLogApi) DeleteEarnSubscribeLogByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	err := earnSubscribeLogService.DeleteEarnSubscribeLogByIds(ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateEarnSubscribeLog 后台更新理财质押记录
// @Tags EarnProductionSubscribeLog
// @Summary 后台更新理财质押记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body earn.EarnSubscribeLog true "更新earnSubscribeLog表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /earn/subscribe/log/edit [put]
func (earnSubscribeLogApi *EarnSubscribeLogApi) UpdateEarnSubscribeLog(c *gin.Context) {
	var earnSubscribeLog earn.EarnSubscribeLog
	err := c.ShouldBindJSON(&earnSubscribeLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = earnSubscribeLogService.UpdateEarnSubscribeLog(earnSubscribeLog)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindEarnSubscribeLog 用id查询earnSubscribeLog表
// @Tags EarnProductionSubscribeLog
// @Summary 用id查询earnSubscribeLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query earn.EarnSubscribeLog true "用id查询earnSubscribeLog表"
// @Success 200 {object} response.Response{data=earn.EarnSubscribeLog,msg=string} "查询成功"
// @Router /earnSubscribeLog/findEarnSubscribeLog [get]
func (earnSubscribeLogApi *EarnSubscribeLogApi) FindEarnSubscribeLog(c *gin.Context) {
	id := c.Query("id")
	reearnSubscribeLog, err := earnSubscribeLogService.GetEarnSubscribeLog(id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reearnSubscribeLog, c)
}

type EarnProductSubscribeLogItem struct {
	*earn.EarnSubscribeLog
	ProductName string `json:"productName"`
}

// GetEarnSubscribeLogList 分页获取earnSubscribeLog表列表
// @Tags EarnProductionSubscribeLog
// @Summary 分页获取earnSubscribeLog表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query earnReq.EarnSubscribeLogSearch true "分页获取earnSubscribeLog表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /earn/subscribe/log/list [get]
func (earnSubscribeLogApi *EarnSubscribeLogApi) GetEarnSubscribeLogList(c *gin.Context) {
	var (
		pageInfo        earnReq.EarnSubscribeLogSearch
		result          []*EarnProductSubscribeLogItem
		earnProductList []*earn.EarnProducts
	)
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := earnSubscribeLogService.GetEarnSubscribeLogInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	earnProductList, err = earnProductsService.GetEarnProductsList(getProductIdList(list))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	for _, v := range list {
		v.BoughtNum, _ = currenciesService.ProcessPriceByCurrency(v.BoughtNum, earnProductStakeCurrency)
		v.PenaltyRatio, _ = currenciesService.ProcessPriceByCurrency(v.PenaltyRatio, earnProductStakeCurrency)
		item := &EarnProductSubscribeLogItem{
			EarnSubscribeLog: &v,
		}
		for _, vv := range earnProductList {
			if v.ProductId == vv.Id {
				item.ProductName = vv.Name
			}
		}
		result = append(result, item)
	}
	response.OkWithDetailed(response.PageResult{
		List:     result,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func getProductIdList(p []earn.EarnSubscribeLog) (res []uint) {
	for _, v := range p {
		res = append(res, v.ProductId)
	}
	return res
}

func (earnSubscribeLogApi *EarnSubscribeLogApi) GetEarnSubscribeLogById(c *gin.Context) {}

// GetEarnSubscribeLogPublic 不需要鉴权的earnSubscribeLog表接口
// @Tags EarnProductionSubscribeLog
// @Summary 不需要鉴权的earnSubscribeLog表接口
// @accept application/json
// @Produce application/json
// @Param data query earnReq.EarnSubscribeLogSearch true "分页获取earnSubscribeLog表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /earnSubscribeLog/getEarnSubscribeLogPublic [get]
func (earnSubscribeLogApi *EarnSubscribeLogApi) GetEarnSubscribeLogPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	earnSubscribeLogService.GetEarnSubscribeLogPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的earnSubscribeLog表接口信息",
	}, "获取成功", c)
}

// Stake  用户质押
// @Tags EarnProductionSubscribeLog
// @Summary  用户质押
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body earnReq.EarnProductStake true "用户质押请求"
// @Success 200 {object} response.Response{msg=string} "申请成功"
// @Router /front/earn/product/subscription/stake [post]
func (earnSubscribeLogApi *EarnSubscribeLogApi) Stake(c *gin.Context) {
	var (
		req              earnReq.EarnProductStake
		err              error
		u                user.FrontendUsers
		product          earn.EarnProducts
		earnSubscribeLog *earn.EarnSubscribeLog
		z                decimal.Decimal
		n                = time.Now()
	)
	err = c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if req.Amount.Cmp(decimal.NewFromFloat(0.0)) <= 0 || req.ProductId == 0 {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ParamError, response.ERROR), c)
		return
	}
	u, err = service.ServiceGroupApp.UserServiceGroup.FrontendUsersService.GetFrontendUsers(fmt.Sprint(utils.GetUserIDFrontUser(c)))
	if err != nil {
		global.GVA_LOG.Error("not found the user", zap.Error(err), zap.String("userId", fmt.Sprint(utils.GetUserIDFrontUser(c))))
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.UserDoesNotExist, response.ERROR), c)
		return
	}
	product, err = service.ServiceGroupApp.EarnServiceGroup.GetEarnProducts(fmt.Sprint(req.ProductId))
	if err != nil {
		global.GVA_LOG.Error("earn product not found ", zap.Error(err), zap.Any("product id", req.ProductId))
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.EarnProductNotFound, response.ERROR), c)
		return
	}
	tx := global.GVA_DB.Begin()
	defer tx.Rollback()
	earnSubscribeLog = &earn.EarnSubscribeLog{
		Uid:             u.ID,
		ProductId:       req.ProductId,
		PenaltyRatio:    product.PenaltyRatio,
		Status:          earn.Staking,
		RedeemInAdvance: earn.RedeemPending,
		Fine:            z,
		StartAt:         n.UnixMilli(),
		EndAt:           n.AddDate(0, 0, product.Duration).UnixMilli(),
		BoughtNum:       req.Amount,
		UpdatedAtX:      n.UnixMilli(),
		CreatedAt:       n.UnixMilli(),
		UserType:        u.UserType,
	}
	err = earnSubscribeLogService.CreateEarnSubscribeLog(tx, earnSubscribeLog)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.SubscribeEarnProductFailed, response.ERROR), c)
		return
	}
	err = userfund.NewUserFundAccountService(tx, false).
		UpdateUserFundAccountsAndNewFlow(int(u.ID), fund.StakeEarnProduct, req.Amount.InexactFloat64(),
			fmt.Sprintf("%v_%v_%v_%v", u.ID, req.ProductId, req.Amount, n.Unix()))
	if err != nil {
		global.GVA_LOG.Error("deduct user balance err", zap.Error(err))
		if errors.Is(err, fund.BalanceDoesNotEnoughError) {
			response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.BalanceDoseNotEnough, response.ERROR), c)
			return
		}
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.SubscribeEarnProductFailed, response.ERROR), c)
		return
	}
	if err = tx.Commit().Error; err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.SubscribeEarnProductFailed, response.ERROR), c)
		return
	}
	response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.SuccessfullyObtained, response.SUCCESS), c)
}

// Redeem 用户主动赎回
// @Tags EarnProductionSubscribeLog
// @Summary 用户主动赎回
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query  earnReq.EarnProductRedeem  true "用户质押赎回请求"
// @Success 200 {object} response.Response{msg=string} "申请成功"
// @Router /front/earn/product/subscription/redeem [post]
func (earnSubscribeLogApi *EarnSubscribeLogApi) Redeem(c *gin.Context) {
	var (
		req         earnReq.EarnProductRedeem
		subscribe   earn.EarnSubscribeLog
		earnProduct earn.EarnProducts
	)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if req.SubscriptionID == 0 {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ParamError, response.ERROR), c)
		return
	}

	subscribe, err = earnSubscribeLogService.GetUserEarnSubscribeLog(fmt.Sprint(req.SubscriptionID), utils.GetUserIDFrontUser(c))
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.SubscribeLogNotFound, response.ERROR), c)
		return
	}
	if subscribe.Id == 0 {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.SubscribeLogNotFound, response.ERROR), c)
		return
	}
	earnProduct, err = earnProductsService.GetEarnProducts(fmt.Sprint(subscribe.ProductId))
	if err != nil {
		global.GVA_LOG.Error("not found the product error", zap.Error(err), zap.Any("product id", subscribe.ProductId))
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.QueryError, response.ERROR), c)
		return
	}
	if earnProduct.Id == 0 {
		global.GVA_LOG.Error("not found the product id", zap.Error(err), zap.Any("product id", subscribe.ProductId))
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.EarnProductNotFound, response.ERROR), c)
		return
	}
	subscribe.RedeemInAdvance = earn.RedeemNormal
	if earnProduct.Type == earn.Fixed {
		if subscribe.EndAt > time.Now().UnixMilli() {
			subscribe.RedeemInAdvance = earn.RedeemInAdvance
		}
	}
	subscribe.EndAt = time.Now().UnixMilli()
	err = earnSubscribeLogService.UpdateEarnSubscribeLog(subscribe)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ServerError, response.SUCCESS), c)
		return
	}
	response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}

// GetFrontEarnSubscribeLogList
// @Tags EarnProductionSubscribeLog
// @Summary 后台查询质押记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query earnReq.EarnSubscribeLogSearch true "后台查询质押记录"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /front/earn/product/subscription/list [get]
func (earnSubscribeLogApi *EarnSubscribeLogApi) GetFrontEarnSubscribeLogList(c *gin.Context) {
	var req earnReq.EarnSubscribeLogSearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	req.Uid = utils.GetUserIDFrontUser(c)
	list, total, err := earnSubscribeLogService.GetEarnSubscribeLogInfoList(req)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err), zap.Any("req", req))
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ServerError, response.ERROR), c)
		return
	}
	for i, v := range list {
		list[i].Fine, _ = currenciesService.ProcessPriceByCurrency(v.Fine, earnProductStakeCurrency)
		list[i].PenaltyRatio, _ = currenciesService.ProcessPriceByCurrency(v.PenaltyRatio, earnProductStakeCurrency)
		list[i].BoughtNum, _ = currenciesService.ProcessPriceByCurrency(v.BoughtNum, earnProductStakeCurrency)
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}

// GetFrontEarnSubscribeSummary
// @Tags EarnProductionSubscribeLog
// @Summary 前端查询质押记录统计
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query earnReq.EarnSubscribeLogSearch true "前端查询质押记录统计"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /front/earn/product/subscription/summary [get]
func (earnSubscribeLogApi *EarnSubscribeLogApi) GetFrontEarnSubscribeSummary(c *gin.Context) {
	var (
		req earnReq.PurchasedEarnProductsSearch
		uid uint
	)
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid = utils.GetUserIDFrontUser(c)
	list, err := earnSubscribeLogService.GetUserEarnSubscribeLogList(uid)
	if err != nil {
		global.GVA_LOG.Error("get user earn subscribe log list err", zap.Error(err), zap.Any("req", req))
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ServerError, response.ERROR), c)
		return
	}
	res, totalEarned, total, err := getEarnProductPurchasedRes(list, &req)
	response.OkWithDetailed(PageResultVS{
		List:              res,
		Total:             total,
		AllProductsEarned: totalEarned,
		Page:              req.Page,
		PageSize:          req.PageSize,
	}, i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}

type PageResultVS struct {
	List              []*earnRes.PurchasedEarnProduct `json:"list"`
	Total             int64                           `json:"total"`
	Page              int                             `json:"page"`
	PageSize          int                             `json:"pageSize"`
	AllProductsEarned decimal.Decimal                 `json:"allProductsEarned"`
}

func getEarnProductPurchasedRes(sbs []*earn.EarnSubscribeLog, req *earnReq.PurchasedEarnProductsSearch) (res []*earnRes.PurchasedEarnProduct, allProductsTotalEarned decimal.Decimal, total int64, err error) {
	res = make([]*earnRes.PurchasedEarnProduct, 0)
	var (
		period                   = time.Now().Truncate(24 * time.Hour)
		productIdList            []uint
		subscriptionIdList       []uint
		earnProductInterestRates []*earn.EarnInterestRates
		earnProductList          []*earn.EarnProducts
	)

	for _, v := range sbs {
		productIdList = append(productIdList, v.ProductId)
		subscriptionIdList = append(subscriptionIdList, v.Id)
	}
	earnProductInterestRates, err = earnInterestRatesService.GetCurrentPeriodEarnInterestRates(productIdList, period)
	if err != nil {
		return res, allProductsTotalEarned, total, err
	}

	earnProductList, err = earnProductsService.GetEarnProductsList(productIdList)

	if err != nil {
		return res, allProductsTotalEarned, total, err
	}

	totalEarnings, err := earnDailyIncomeMoneyLogService.GetEarnDailyIncomeSummaryListGroupBySubscribe(subscriptionIdList)
	if err != nil {
		return res, allProductsTotalEarned, total, err
	}

	for _, v := range sbs {
		var r earnRes.PurchasedEarnProduct
		if v.RedeemInAdvance == earn.RedeemInAdvance {
			v.Status = earn.Redeemed
		}
		r.Subscription = v
		for _, vv := range earnProductInterestRates {
			if v.ProductId == vv.ProductId {
				r.TodayInterestRate = vv.InterestRates
			}
		}
		allProductsTotalEarned = allProductsTotalEarned.Add(v.Fine)
		for _, vv := range totalEarnings {
			if v.Id == vv.SubscribeId {
				r.TotalEarned = decimal.NewFromFloat(vv.N)
				allProductsTotalEarned = allProductsTotalEarned.Add(decimal.NewFromFloat(vv.N))
			}
		}
		for _, vv := range earnProductList {
			if v.ProductId == vv.Id {
				r.EarnProduct = vv
			}
		}
		r.TotalEarned, _ = currenciesService.ProcessPriceByCurrency(r.TotalEarned, earnProductStakeCurrency)
		r.Subscription.Fine, _ = currenciesService.ProcessPriceByCurrency(r.Subscription.Fine, earnProductStakeCurrency)
		r.Subscription.BoughtNum, _ = currenciesService.ProcessPriceByCurrency(r.Subscription.BoughtNum, earnProductStakeCurrency)
		r.Subscription.PenaltyRatio, _ = currenciesService.ProcessPriceByCurrency(r.Subscription.PenaltyRatio, earnProductStakeCurrency)
		res = append(res, &r)
	}
	total = int64(len(res))
	start := (req.Page - 1) * req.PageSize
	end := start + req.PageSize
	if end > len(res) {
		end = len(res)
	}
	if start < 0 {
		start = 0
	}
	res = res[start:end]
	allProductsTotalEarned, _ = currenciesService.ProcessPriceByCurrency(allProductsTotalEarned, earnProductStakeCurrency)

	return res, allProductsTotalEarned, total, err
}
