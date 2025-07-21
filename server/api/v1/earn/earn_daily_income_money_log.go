package earn

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/i18n"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/earn"
	earnReq "github.com/flipped-aurora/gin-vue-admin/server/model/earn/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EarnDailyIncomeMoneyLogApi struct{}

// CreateEarnDailyIncomeMoneyLog 创建earnDailyIncomeMoneyLog表
// @Tags EarnDailyIncomeMoneyLog
// @Summary 创建earnDailyIncomeMoneyLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body earn.EarnDailyIncomeMoneyLog true "创建earnDailyIncomeMoneyLog表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /earnDailyIncomeMoneyLog/createEarnDailyIncomeMoneyLog [post]
func (earnDailyIncomeMoneyLogApi *EarnDailyIncomeMoneyLogApi) CreateEarnDailyIncomeMoneyLog(c *gin.Context) {
	var earnDailyIncomeMoneyLog earn.EarnDailyIncomeMoneyLog
	err := c.ShouldBindJSON(&earnDailyIncomeMoneyLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = earnDailyIncomeMoneyLogService.CreateEarnDailyIncomeMoneyLog(&earnDailyIncomeMoneyLog)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteEarnDailyIncomeMoneyLog 删除earnDailyIncomeMoneyLog表
// @Tags EarnDailyIncomeMoneyLog
// @Summary 删除earnDailyIncomeMoneyLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body earn.EarnDailyIncomeMoneyLog true "删除earnDailyIncomeMoneyLog表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /earnDailyIncomeMoneyLog/deleteEarnDailyIncomeMoneyLog [delete]
func (earnDailyIncomeMoneyLogApi *EarnDailyIncomeMoneyLogApi) DeleteEarnDailyIncomeMoneyLog(c *gin.Context) {
	id := c.Query("id")
	err := earnDailyIncomeMoneyLogService.DeleteEarnDailyIncomeMoneyLog(id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteEarnDailyIncomeMoneyLogByIds 批量删除earnDailyIncomeMoneyLog表
// @Tags EarnDailyIncomeMoneyLog
// @Summary 批量删除earnDailyIncomeMoneyLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /earnDailyIncomeMoneyLog/deleteEarnDailyIncomeMoneyLogByIds [delete]
func (earnDailyIncomeMoneyLogApi *EarnDailyIncomeMoneyLogApi) DeleteEarnDailyIncomeMoneyLogByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	err := earnDailyIncomeMoneyLogService.DeleteEarnDailyIncomeMoneyLogByIds(ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateEarnDailyIncomeMoneyLog 更新earnDailyIncomeMoneyLog表
// @Tags EarnDailyIncomeMoneyLog
// @Summary 更新earnDailyIncomeMoneyLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body earn.EarnDailyIncomeMoneyLog true "更新earnDailyIncomeMoneyLog表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /earnDailyIncomeMoneyLog/updateEarnDailyIncomeMoneyLog [put]
func (earnDailyIncomeMoneyLogApi *EarnDailyIncomeMoneyLogApi) UpdateEarnDailyIncomeMoneyLog(c *gin.Context) {
	var earnDailyIncomeMoneyLog earn.EarnDailyIncomeMoneyLog
	err := c.ShouldBindJSON(&earnDailyIncomeMoneyLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = earnDailyIncomeMoneyLogService.UpdateEarnDailyIncomeMoneyLog(earnDailyIncomeMoneyLog)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindEarnDailyIncomeMoneyLog 后台查询理财收支情况
// @Tags EarnDailyIncomeMoneyLog
// @Summary 后台查询理财收支情况
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query earn.EarnDailyIncomeMoneyLog true "后台查询理财收支情况"
// @Success 200 {object} response.Response{data=earn.EarnDailyIncomeMoneyLog,msg=string} "查询成功"
// @Router /earnDailyIncomeMoneyLog/findEarnDailyIncomeMoneyLog [get]
func (earnDailyIncomeMoneyLogApi *EarnDailyIncomeMoneyLogApi) FindEarnDailyIncomeMoneyLog(c *gin.Context) {
	id := c.Query("id")
	reearnDailyIncomeMoneyLog, err := earnDailyIncomeMoneyLogService.GetEarnDailyIncomeMoneyLog(id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reearnDailyIncomeMoneyLog, c)
}

type PageResultV struct {
	*earnReq.EarnProductsIncomeSummaryRes
	Total    int64 `json:"total"`
	Page     int   `json:"page"`
	PageSize int   `json:"pageSize"`
}

// GetEarnProductDailyIncomeMoneyLogSummary 分页获取earnDailyIncomeMoneyLog表列表
// @Tags EarnDailyIncomeMoneyLog
// @Summary 分页获取earnDailyIncomeMoneySummary表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query earnReq.EarnDailyIncomeMoneyLogSearch true "分页获取earnDailyIncomeMoneyLog表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /earn/product/daily/income/summary [get]
func (earnDailyIncomeMoneyLogApi *EarnDailyIncomeMoneyLogApi) GetEarnProductDailyIncomeMoneyLogSummary(c *gin.Context) {
	var (
		req        earnReq.EarnDailyIncomeMoneyLogSearch
		res        *earnReq.EarnProductsIncomeSummaryRes
		userList   []*user.FrontendUsers
		incomeList []*earn.EarnDailyIncomeMoneyLog
		total      int64
	)
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userList, err = service.ServiceGroupApp.UserServiceGroup.FrontendUsersService.GetUsers(req.Uid, req.Phone, req.Email, req.ParentId, req.SuperiorId, req.UserType)
	if len(userList) > 0 || req.Uid != "" || req.Email != "" || req.ParentId != "" || req.SuperiorId != "" || req.UserType != 0 {
		var userIdList []uint
		for _, u := range userList {
			userIdList = append(userIdList, u.ID)
		}
		incomeList, total, err = earnDailyIncomeMoneyLogService.GetEarnDailyIncomeMoneyLogInfoList(userIdList, req)
		if err != nil {
			global.GVA_LOG.Error("获取失败!", zap.Error(err), zap.Any("req", req))
			response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Success, response.ERROR), c)
			return
		}
	} else {
		incomeList, total, err = earnDailyIncomeMoneyLogService.GetEarnDailyIncomeMoneyLogInfoList(nil, req)
		if err != nil {
			global.GVA_LOG.Error("获取失败!", zap.Error(err), zap.Any("req", req))
			response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Success, response.ERROR), c)
			return
		}
		var userIdList []uint
		for _, u := range incomeList {
			userIdList = append(userIdList, u.Uid)
		}
		userList, err = service.ServiceGroupApp.UserServiceGroup.FrontendUsersService.GetUsersByIds(userIdList)
		if err != nil {
			global.GVA_LOG.Error("获取失败!", zap.Error(err), zap.Any("req", req))
			response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Success, response.ERROR), c)
			return
		}
	}
	res, total = earnDailyIncomeMoneyLogApi.getEarnProductDailyIncomeMoneyLogSummary(userList, incomeList, req.Page, req.PageSize)
	response.OkWithDetailed(PageResultV{
		EarnProductsIncomeSummaryRes: res,
		Page:                         req.Page,
		PageSize:                     req.PageSize,
		Total:                        total,
	}, i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}

// GetEarnProductDailyIncomeMoneyLogDetail 前端用户查看理财产品明细
// @Tags EarnDailyIncomeMoneyLog
// @Summary 前端用户查看理财产品明细
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query earnReq.EarnDailyIncomeMoneyLogSearch true "分页获取earnDailyIncomeMoneyLog表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /front/earn/product/daily/income/detail [get]
func (earnDailyIncomeMoneyLogApi *EarnDailyIncomeMoneyLogApi) GetEarnProductDailyIncomeMoneyLogDetail(c *gin.Context) {
	var (
		req   earnReq.EarnDailyIncomeMoneyLogDetailSearch
		res   []*earn.EarnDailyIncomeMoneyLog
		total int64
		uid   uint
	)
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid = utils.GetUserID(c)

	res, total, err = earnDailyIncomeMoneyLogService.GetFrontEarnProductDailyIncomeMoneyLogDetail(uid, req.SubscriptionId, req.Page, req.PageSize)
	if err != nil {
		global.GVA_LOG.Error("get earn product daily income money detail", zap.Error(err), zap.Any("req", req))
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ServerError, response.SUCCESS), c)
		return
	}
	for i, v := range res {
		res[i].Earnings, _ = currenciesService.ProcessPriceByCurrency(v.Earnings, earnProductStakeCurrency)
		res[i].BoughtNum, _ = currenciesService.ProcessPriceByCurrency(v.BoughtNum, earnProductStakeCurrency)
		res[i].InterestRates, _ = currenciesService.ProcessPriceByCurrency(v.InterestRates, earnProductStakeCurrency)
	}
	response.OkWithDetailed(response.PageResult{
		List:     res,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}

func (earnDailyIncomeMoneyLogApi *EarnDailyIncomeMoneyLogApi) getEarnProductDailyIncomeMoneyLogSummary(u []*user.FrontendUsers,
	incomes []*earn.EarnDailyIncomeMoneyLog, page, pageSize int) (res *earnReq.EarnProductsIncomeSummaryRes, total int64) {
	res = new(earnReq.EarnProductsIncomeSummaryRes)
	var (
		userMap   = make(map[uint]*user.FrontendUsers)
		incomeMap = make(map[uint][]*earn.EarnDailyIncomeMoneyLog)
	)
	for _, v := range u {
		userMap[v.ID] = v
	}
	for _, v := range incomes {
		incomeMap[v.Uid] = append(incomeMap[v.Uid], v)
	}
	for k, v := range incomeMap {
		u, has := userMap[k]
		if !has {
			continue
		}
		i := &earnReq.EarnProductsIncome{
			Phone:      u.Phone,
			Email:      u.Email,
			Uid:        u.ID,
			ParentId:   u.ParentId,
			SuperiorId: u.RootUserid,
			UserType:   u.UserType,
		}
		res.TotalStakeNum = res.TotalStakeNum + 1
		for _, vv := range v {
			res.TotalStakeAmount = res.TotalStakeAmount.Add(vv.BoughtNum)
			res.TotalStakeProfit = res.TotalStakeProfit.Add(vv.Earnings)
			i.StakeNum = i.StakeNum + 1
			i.StakeAmount = i.StakeAmount.Add(vv.BoughtNum)
			i.StakeEarnings = i.StakeEarnings.Add(vv.Earnings)
		}
		res.List = append(res.List, i)
	}
	total = int64(len(res.List))
	start := (page - 1) * pageSize
	end := start + pageSize
	if end > len(res.List) {
		end = len(res.List)
	}
	res.List = res.List[start:end]
	return res, total
}

// GetEarnDailyIncomeMoneyLogPublic 不需要鉴权的earnDailyIncomeMoneyLog表接口
// @Tags EarnDailyIncomeMoneyLog
// @Summary 不需要鉴权的earnDailyIncomeMoneyLog表接口
// @accept application/json
// @Produce application/json
// @Param data query earnReq.EarnDailyIncomeMoneyLogSearch true "分页获取earnDailyIncomeMoneyLog表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /earnDailyIncomeMoneyLog/getEarnDailyIncomeMoneyLogPublic [get]
func (earnDailyIncomeMoneyLogApi *EarnDailyIncomeMoneyLogApi) GetEarnDailyIncomeMoneyLogPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	earnDailyIncomeMoneyLogService.GetEarnDailyIncomeMoneyLogPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的earnDailyIncomeMoneyLog表接口信息",
	}, "获取成功", c)
}

func (earnDailyIncomeMoneyLogApi *EarnDailyIncomeMoneyLogApi) getEarnProductDailyIncomeMoneyLogDetail(u []*user.FrontendUsers,
	incomes []*earn.EarnDailyIncomeMoneyLog, page, pageSize int) (res *earnReq.EarnProductsIncomeSummaryRes) {
	res = new(earnReq.EarnProductsIncomeSummaryRes)
	var (
		userMap   = make(map[uint]*user.FrontendUsers)
		incomeMap = make(map[uint][]*earn.EarnDailyIncomeMoneyLog)
	)
	for _, v := range u {
		userMap[v.ID] = v
	}
	for _, v := range incomes {
		incomeMap[v.Uid] = append(incomeMap[v.Uid], v)
	}
	for k, v := range incomeMap {
		u, has := userMap[k]
		if !has {
			continue
		}
		i := &earnReq.EarnProductsIncome{
			Phone:      u.Phone,
			Email:      u.Email,
			Uid:        u.ID,
			ParentId:   u.ParentId,
			SuperiorId: u.RootUserid,
		}
		for _, vv := range v {
			res.TotalStakeNum = res.TotalStakeNum + 1
			res.TotalStakeAmount = res.TotalStakeAmount.Add(vv.BoughtNum)
			res.TotalStakeProfit = res.TotalStakeProfit.Add(vv.Earnings)
			i.StakeNum = i.StakeNum + 1
			i.StakeAmount = i.StakeAmount.Add(vv.Earnings)
			i.StakeEarnings = i.StakeEarnings.Add(vv.Earnings)
		}
		res.List = append(res.List, i)
	}
	start := page * pageSize
	end := page*pageSize + pageSize
	if end > len(res.List) {
		end = len(res.List)
	}
	res.List = res.List[start:end]
	return res
}
