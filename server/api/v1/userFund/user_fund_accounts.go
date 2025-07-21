package userFund

import (
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/userFund/thirdPayment"
	"github.com/flipped-aurora/gin-vue-admin/server/i18n"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cast"

	"github.com/flipped-aurora/gin-vue-admin/server/enums"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/userfund"
	userfundReq "github.com/flipped-aurora/gin-vue-admin/server/model/userfund/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/snowflake"
	"github.com/gin-gonic/gin"
	. "github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type UserFundAccountsApi struct{}

// 自定义返回结构体
type UserAccountInfo struct {
	UserID            int         `json:"user_id"`          // 用户ID
	Balance           Decimal     `json:"balance"`          // 账户余额
	FrozenBalance     Decimal     `json:"frozenBalance"`    // 账户状态（例如：正常、冻结等）
	AvailableBalance  Decimal     `json:"availableBalance"` // 账户币种
	FirstChargeAmount NullDecimal `json:"firstChargeAmount" `
	FirstChargeTime   time.Time   `json:"firstChargeTime"`

	TotalBalance Decimal `json:"totalBalance"`
}

// 自定义返回结构体
type UserAccountFundInfo struct {
	UserID             int         `json:"user_id"`        // 用户ID
	Balance            Decimal     `json:"balance"`        // 账户余额
	RechargeAmount     Decimal     `json:"rechargeAmount"` // 总充值
	WithdrawAmount     Decimal     `json:"withdrawAmount"` // 总提现
	FirstChargeAmount  NullDecimal `json:"firstChargeAmount" `
	FirstChargeTime    *time.Time  `json:"firstChargeTime"`
	FirstChargeTimeInt int64       `json:"firstChargeTimeInt"`
	FrozenAmount       Decimal     `json:"frozenAmount"`     // 总提现
	AvaliableBalance   Decimal     `json:"avaliableBalance"` // 总提现
}

// CreateUserFundAccounts 创建userFundAccounts表
// @Tags UserFundAccounts
// @Summary 创建userFundAccounts表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userfund.UserFundAccounts true "创建userFundAccounts表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /userFundAccounts/createUserFundAccounts [post]
func (userFundAccountsApi *UserFundAccountsApi) CreateUserFundAccounts(c *gin.Context) {
	var userFundAccounts userfund.UserFundAccounts
	err := c.ShouldBindJSON(&userFundAccounts)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userFundAccountsService.CreateUserFundAccounts(&userFundAccounts)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteUserFundAccounts 删除userFundAccounts表
// @Tags UserFundAccounts
// @Summary 删除userFundAccounts表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userfund.UserFundAccounts true "删除userFundAccounts表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /userFundAccounts/deleteUserFundAccounts [delete]
func (userFundAccountsApi *UserFundAccountsApi) DeleteUserFundAccounts(c *gin.Context) {
	id := c.Query("id")
	err := userFundAccountsService.DeleteUserFundAccounts(id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteUserFundAccountsByIds 批量删除userFundAccounts表
// @Tags UserFundAccounts
// @Summary 批量删除userFundAccounts表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /userFundAccounts/deleteUserFundAccountsByIds [delete]
func (userFundAccountsApi *UserFundAccountsApi) DeleteUserFundAccountsByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	err := userFundAccountsService.DeleteUserFundAccountsByIds(ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateUserFundAccounts 更新userFundAccounts表
// @Tags UserFundAccounts
// @Summary 更新userFundAccounts表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userfund.UserFundAccounts true "更新userFundAccounts表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /userFundAccounts/updateUserFundAccounts [put]
func (userFundAccountsApi *UserFundAccountsApi) UpdateUserFundAccounts(c *gin.Context) {
	var userFundAccounts userfund.UserFundAccounts
	err := c.ShouldBindJSON(&userFundAccounts)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userFundAccountsService.UpdateUserFundAccounts(userFundAccounts)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.FundUpdateFailed, response.ERROR), c)
		return
	}
	//response.OkWithMessage("更新成功", c)
	response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.OperateSuccess, response.ERROR), c)
}

// FindUserFundAccounts 用id查询userFundAccounts表
// @Tags UserFundAccounts
// @Summary 用id查询userFundAccounts表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userfund.UserFundAccounts true "用id查询userFundAccounts表"
// @Success 200 {object} response.Response{data=userfund.UserFundAccounts,msg=string} "查询成功"
// @Router /userFundAccounts/findUserFundAccounts [get]
func (userFundAccountsApi *UserFundAccountsApi) FindUserFundAccounts(c *gin.Context) {
	id := c.Query("id")
	reuserFundAccounts, err := userFundAccountsService.GetUserFundAccounts(id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reuserFundAccounts, c)
}

// GetUserFundAccountsList 分页获取userFundAccounts表列表
// @Tags UserFundAccounts
// @Summary 分页获取userFundAccounts表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userfundReq.UserFundAccountsSearch true "分页获取userFundAccounts表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /userFundAccounts/getUserFundAccountsList [get]
func (userFundAccountsApi *UserFundAccountsApi) GetUserFundAccountsList(c *gin.Context) {
	var pageInfo userfundReq.UserFundAccountsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userFundAccountsService.GetUserFundAccountsInfoList(pageInfo)
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

// GetUserFundAccountsPublic 不需要鉴权的userFundAccounts表接口
// @Tags UserFundAccounts
// @Summary 不需要鉴权的userFundAccounts表接口
// @accept application/json
// @Produce application/json
// @Param data query userfundReq.UserFundAccountsSearch true "分页获取userFundAccounts表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /userFundAccounts/getUserFundAccountsPublic [get]
func (userFundAccountsApi *UserFundAccountsApi) GetUserFundAccountsPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	userFundAccountsService.GetUserFundAccountsPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的userFundAccounts表接口信息",
	}, "获取成功", c)
}

// GetUserAccountInfo 查询用户账户信息
// @Tags UserFundAccounts
// @Summary 查询用户账户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=UserAccountInfo,msg=string} "获取成功"
// @Router /userFundAccounts/getUserAccountInfo [get]
func (userFundAccountsApi *UserFundAccountsApi) GetUserAccountInfo(c *gin.Context) {
	//userIDStr := c.Query("userId")
	//userID := utils.GetUserID(c)
	//if userIDStr == "" {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
	//	return
	//}
	//
	//// 将字符串转换为 int 类型
	//userID, err := strconv.Atoi(userIDStr)
	userID := utils.GetUserID(c)
	//if err != nil {
	//	global.GVA_LOG.Error("参数格式不正确!", zap.Error(err))
	//	response.FailWithMessage("参数格式不正确:"+err.Error(), c)
	//	return
	//}
	totalFollowAmount := userFollowOrderService.GetTotalAmount(userID)
	recontactAccount, err := contractAccountService.GetContractAccountByUserId(userID)
	unrealizedProfitLoss := Zero
	if err != nil {
		global.GVA_LOG.Error("查询合约账户失败!", zap.Error(err))
		//response.FailWithMessage("查询合约账户失败:"+err.Error(), c)
		//return
		recontactAccount.TotalMargin = Zero
	} else {
		unrealizedProfitLoss, err = contractPositionService.GetUnrealizedProfitLoss(userID)
		if err != nil {
			global.GVA_LOG.Error("计算未实现盈亏失败!", zap.Error(err))
			response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.OperateFail, response.ERROR), c)
			return
		}
	}

	userFoundAccount, err := userFundAccountsService.GetUserFundAccountsByUserId(int(userID))
	if err != nil {
		global.GVA_LOG.Error("没有获取到账户信息!", zap.Error(err))
		//response.FailWithMessage("没有获取到账户信息:"+err.Error(), c)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.OperateFail, response.ERROR), c)
		return
	}
	totalAmount := NewFromFloat(totalFollowAmount).Add(userFoundAccount.Balance).Add(recontactAccount.TotalMargin).Add(unrealizedProfitLoss)

	accountInfo := UserAccountInfo{
		UserID:            userFoundAccount.UserId,
		Balance:           userFoundAccount.Balance,
		FrozenBalance:     userFoundAccount.FrozenBalance,
		AvailableBalance:  userFoundAccount.AvailableBalance,
		FirstChargeAmount: userFoundAccount.FirstChargeAmount,
		TotalBalance:      totalAmount,
	}
	if userFoundAccount.FirstChargeTime != nil {
		accountInfo.FirstChargeTime = *userFoundAccount.FirstChargeTime
	}
	response.OkWithData(accountInfo, c)
}

// Recharge 用户充值接口
// @Tags UserFundAccounts
// @Summary 用户充值接口
// @accept application/json
// @Produce application/json
// @Param data query userfundReq.UserFundAccountsSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /userFundAccounts/recharge [POST]
func (userFundAccountsApi *UserFundAccountsApi) Recharge(c *gin.Context) {
	// 参数解析
	var rechargeRequest userfund.RechargeRequest
	err := c.ShouldBindJSON(&rechargeRequest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//根据id查询充值渠道接口
	rechargeChannel, err := rechargeChannelsService.GetRechargeChannels(strconv.Itoa(rechargeRequest.RechargeChannelId))
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		//response.FailWithMessage("查询失败:"+err.Error(), c)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.OperateFail, response.ERROR), c)
		return
	}
	decimalPlaces := utils.GetDecimalPlaces2(rechargeChannel.TicketNumSize)
	// 对USDT金额进行精度处理（保留2位小数，向下取整）
	rechargeRequest.RechargeAmount, err = utils.FloorDecimal(rechargeRequest.RechargeAmount, int32(decimalPlaces))
	if err != nil {
		global.GVA_LOG.Error("精度传递错误!", zap.Error(err))
		//response.FailWithMessage("当前数量不满足后台设置的精度要求:"+err.Error(), c)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.FloatNotPatchError, response.ERROR), c)
		return
	}
	// 验证最小充值金额
	if rechargeRequest.RechargeAmount.LessThan(rechargeChannel.MinRechargeNum) {
		//response.FailWithMessage("充值金额不能小于"+rechargeChannel.MinRechargeNum.String(), c)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.RechargeMoutMinFaild, response.ERROR), c)
		return
	}

	//新增订单
	sf, err := snowflake.NewSnowflake(1)
	if err != nil {
		panic(err)
	}
	// 生成雪花 ID
	orderID := sf.Generate()
	userId := utils.GetUserID(c)
	userInfo, err := frontendUsersService.GetFrontendUsers(cast.ToString(userId))

	if err != nil {
		return
	}
	rechargeTime := time.Now()
	rechargeRecords := userfund.RechargeRecords{
		GVA_MODEL:           global.GVA_MODEL{},
		OrderId:             strconv.FormatInt(orderID, 10),
		MemberId:            int(userId),
		ParentId:            int(userInfo.ParentId),
		RootId:              int(userInfo.RootUserid),
		MemberPhone:         userInfo.Phone,
		MemberEmail:         userInfo.Email,
		RechargeChannelId:   rechargeChannel.ID,
		Currency:            rechargeChannel.Currency,
		RechargeAmount:      rechargeRequest.RechargeAmount,
		ExchangedAmountUsdt: Zero,
		RechargeType:        rechargeChannel.RechargeType,
		RechargeRate:        rechargeChannel.PriceUsdt,
		Channel:             rechargeChannel.Channel,
		FromAddress:         rechargeRequest.FromAddress,
		ToAddress:           rechargeChannel.Address,
		ChannelType:         rechargeChannel.ChannelType,
		OrderStatus:         enums.PENDING,
		RechargeTime:        &rechargeTime,
		ApprovalTime:        nil,
		UserAction:          "",
		ReviewStatus:        "",
	}
	// 获取货币信息
	currency, err := currenciesService.GetCurrencies(strconv.Itoa(*rechargeChannel.CoinId))
	if err != nil {
		global.GVA_LOG.Error("取货币信息失败!", zap.Error(err))
		//response.FailWithMessage("获取货币信息失败:"+err.Error(), c)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.CurrencyQueryFailed, response.ERROR), c)
		return
	}

	// 使用 CurrenciesService 计算 USDT 金额
	usdtAmount, err := currenciesService.CalculateUsdtAmount(rechargeRequest.RechargeAmount, currency)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.CalculateUsdtFaild, response.ERROR), c)
		return
	}

	// 对USDT金额进行精度处理（保留2位小数，向下取整）
	usdtAmount = utils.FloorDecimalNoValidate(usdtAmount, 2)
	//如果是快捷充值需要调用第三方接口
	var directUrl string
	if rechargeChannel.RechargeType == enums.RechargeTypeQuick {
		rechargeRecords.OrderStatus = enums.PENDING
		var (
			thirdPaymentId = rechargeChannel.ThirdPayId
			thirdRespInfo  thirdPayment.RechargeResponseInfo
			thirdReqInfo   = thirdPayment.RechargeRequestInfo{
				CurrencyMoney: rechargeRequest.RechargeAmount.String(),
				CurrencyCode:  rechargeChannel.ThirdCurrencyCode,
				CoinCode:      rechargeChannel.ThirdCoinCode,
				SyncJumpURL:   rechargeRequest.JumpUrl,
				UserOrderId:   strconv.FormatInt(orderID, 10),
				Language:      2,
				UserCustomId:  strconv.Itoa(int(utils.GetUserID(c))),
				Remark:        rechargeRequest.Remark,
			}
		)
		//三方订单号、收款地址、支付地址、传给三方的所有参数
		thirdRespInfo, err = thirdPayment.SendRechargeRequest(thirdPaymentId, thirdReqInfo)
		if err != nil {
			_ = c.Error(err)
			response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.OperateFail, response.ERROR), c)
			return
		}
		directUrl = thirdRespInfo.PayUrl
		rechargeRecords.ThirdOrderId = thirdRespInfo.OrderId
		rechargeRecords.ToAddress = thirdRespInfo.ToAddress
		//记录调用第三方的参数
		jsonData, err := json.Marshal(thirdRespInfo.ReqBody)
		if err != nil {
			log.Fatal(err)
		}
		rechargeRecords.NoticeContent = string(jsonData)
	} else {

		rechargeRecords.UserAction = enums.UserAction_SUBMIT //已提交
		rechargeRecords.ReviewStatus = enums.UN_CHECK        //待审核
		rechargeRecords.IsLock = enums.UNLOCK                //未锁定
	}
	rechargeRecords.ExchangedAmountUsdt = usdtAmount
	rechargeRecords.UserType = utils.GetUserInfo(c).UserType
	//创建订单记录
	err = rechargeRecordsService.CreateRechargeRecords(&rechargeRecords)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.OperateFail, response.ERROR), c)
		return
	}
	response.OkWithData(directUrl, c)
}

// Withdraw 用户提现接口
// @Tags UserFundAccounts
// @Summary 用户提现接口
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userfund.WithdrawRequest true "提现请求参数"
// @Success 200 {object} response.Response{msg=string} "提现申请成功"
// @Router /userFundAccounts/withdraw [post]
func (userFundAccountsApi *UserFundAccountsApi) Withdraw(c *gin.Context) {
	// 参数解析
	var withdrawRequest userfund.WithdrawRequest
	err := c.ShouldBindJSON(&withdrawRequest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 验证支付密码
	userId := utils.GetUserIDFrontUser(c)
	ok, err := frontendUsersService.VerifyPaymentPassword(uint(userId), withdrawRequest.PaymentPassword)
	if err != nil {
		//response.FailWithMessage("支付密码验证失败", c)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.PayPwdVerifyFaild, response.ERROR), c)
		return
	}
	if !ok {
		//response.FailWithMessage("支付密码错误", c)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.PayPwdMismatchError, response.ERROR), c)
		return
	}

	// 根据id查询提现渠道接口
	withDrawChannel, err := withdrawChannelsService.GetWithdrawChannels(strconv.Itoa(withdrawRequest.WithdrawChannelId))
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.OperateFail, response.ERROR), c)
		return
	}

	//新增订单
	sf, err := snowflake.NewSnowflake(1)
	if err != nil {
		panic(err)
	}
	// 生成雪花 ID
	orderID := sf.Generate()
	userInfo, err := frontendUsersService.GetFrontendUsers(cast.ToString(userId))
	if err != nil {
		return
	}
	userfundAccount, err := userFundAccountsService.GetUserFundAccountsByUserId(int(userId))
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.OperateFail, response.ERROR), c)
		return
	}
	decimalPlaces := utils.GetDecimalPlaces2(withDrawChannel.TicketNumSize)
	// 对USDT金额进行精度处理（保留2位小数，向下取整）
	withdrawRequest.WithdrawAmount, err = utils.FloorDecimal(withdrawRequest.WithdrawAmount, int32(decimalPlaces))
	if err != nil {
		global.GVA_LOG.Error("精度传递错误!", zap.Error(err))
		//response.FailWithMessage("当前数量不满足后台设置的精度要求:"+err.Error(), c)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.FloatNotPatchError, response.ERROR), c)
		return
	}

	// 检查可用余额是否足够
	if userfundAccount.AvailableBalance.LessThan(withdrawRequest.WithdrawAmount) {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.FundAmountNotEnough, response.ERROR), c)
		return
	}

	//计算提现手续费花费的金额
	systemConfig, _ := systemConfigService.GetPlatformSystemConfig()
	commission := Zero
	if systemConfig.WithdrawCommissionType == enums.WithdrawCommissionQuota {
		commission = systemConfig.WithdrawCommissionQuota
	} else if systemConfig.WithdrawCommissionType == enums.WithdrawCommissionRate {
		rate := systemConfig.WithdrawCommissionRate
		commission = withdrawRequest.WithdrawAmount.Mul(rate)
	}
	withdrawTime := time.Now()
	exchangedAmountTargetBefore := (withdrawRequest.WithdrawAmount).Div(withDrawChannel.PriceUsdt) //减去手续费之前目标货币的数量
	// 验证最小提现金额
	if exchangedAmountTargetBefore.LessThan(withDrawChannel.MinWithdrawNum) {
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.WithdrawAmoutCheckFaild, response.ERROR), c)
		return
	}
	exchangedAmountTarget := (withdrawRequest.WithdrawAmount.Sub(commission)).Div(withDrawChannel.PriceUsdt) //减去手续费后目标货币的数量

	if !exchangedAmountTarget.IsPositive() {
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.WithdrawAmoutCheckFaild, response.ERROR), c)
		return
	}
	// 根据TicketNumSize处理数量精度
	targetCoinAmount := Zero
	if withDrawChannel.TicketNumSize != Zero {
		decimalPlaces = utils.GetDecimalPlaces2(withDrawChannel.TicketNumSize)
		targetCoinAmount = utils.TruncateDecimal(exchangedAmountTarget, int32(decimalPlaces))
	} else {
		targetCoinAmount = utils.TruncateDecimal(exchangedAmountTarget, int32(2))
		//targetCoinAmount = exchangedAmountTarget
	}
	withdrawRecords := userfund.WithdrawRecords{
		GVA_MODEL:             global.GVA_MODEL{},
		OrderId:               strconv.FormatInt(orderID, 10),
		MemberId:              int(userId),
		ParentId:              int(userInfo.ParentId),
		RootId:                int(userInfo.RootUserid),
		MemberPhone:           userInfo.Phone,
		MemberEmail:           userInfo.Email,
		WithdrawChannelId:     int(withDrawChannel.ID),
		Currency:              withDrawChannel.Currency,
		WithdrawAmount:        withdrawRequest.WithdrawAmount,
		WithdrawRate:          withDrawChannel.PriceUsdt,
		ExchangedAmountTarget: targetCoinAmount,
		RealReceived:          Zero,
		WithdrawType:          withDrawChannel.RechargeType,
		Channel:               withDrawChannel.Channel,
		FromAddress:           withDrawChannel.Address,
		ToAddress:             withdrawRequest.ToAddress,
		ChannelType:           withDrawChannel.ChannelType,
		OrderStatus:           enums.WITHDRAWING,
		WithdrawTime:          &withdrawTime,
		ApprovalTime:          nil,
		UserAction:            "",
		ReviewStatus:          "",
		Commission:            commission,
		UserType:              utils.GetUserInfo(c).UserType,
	}
	withdrawRecords.UserAction = enums.UserAction_SUBMIT // 已提交
	withdrawRecords.ReviewStatus = enums.UN_CHECK        // 待审核
	withdrawRecords.IsLock = enums.UNLOCK                // 未锁定
	withdrawRecords.Commission = commission              //手续费
	//先创建提现记录,并且更新余额
	err = withdrawRecordsService.CreateWithdrawRecordsAndFrozenBalance(withdrawRecords, userfundAccount)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.OperateFail, response.ERROR), c)
		return
	}
	response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.OperateSuccess, response.ERROR), c)
}

// ProcessAccount 处理账户操作
// @Tags UserFundAccounts
// @Summary 处理账户操作
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userfundReq.AccountOperation true "账户操作参数"
// @Success 200 {object} response.Response{msg=string} "操作成功"
// @Router /userFundAccounts/processAccount [post]
func (userFundAccountsApi *UserFundAccountsApi) ProcessAccount(c *gin.Context) {
	// 参数解析
	var accountOperation userfundReq.AccountOperation
	err := c.ShouldBindJSON(&accountOperation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//更新流水 account flow
	userFundAccountsService.UpdateUserFundAccountsAndNewFlow(int(accountOperation.UserId), accountOperation.Action, accountOperation.Amount, accountOperation.OrderId)
	response.OkWithData("success", c)
}

// GetUserFundStatic 获取用户资金统计
// @Tags UserFundAccounts
// @Summary 获取用户资金计
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param userId path string true "用户ID"
// @Success 200 {object} response.Response{data=UserAccountFundInfo,msg=string} "获取成功"
// @Router /userFundAccounts/getUserFundStatic/{userId} [get]
func (userFundAccountsApi *UserFundAccountsApi) GetUserFundStatic(c *gin.Context) {
	// 参数解析
	userIdStr := c.Param("userId")
	if userIdStr == "" {
		response.FailWithMessage("UserID不能为空", c)
		return
	}

	// 将 userIdStr 转换为 int
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		global.GVA_LOG.Error("用户ID格式错误!", zap.Error(err))
		response.FailWithMessage("用户ID格式错误", c)
		return
	}

	userFoundAccount, err := userFundAccountsService.GetUserFundAccountsByUserId(userId)
	if err != nil {
		global.GVA_LOG.Error("没有获取到账户信息!", zap.Error(err))
		response.FailWithMessage("没有获取到账户信息:"+err.Error(), c)
		return
	}

	// 转换为 int64
	userIdInt64, _ := strconv.ParseInt(userIdStr, 10, 64)
	rechargeAmount := rechargeRecordsService.CountAllByUserId(userIdInt64)
	withdrawAmount := withdrawRecordsService.CountAllByUserId(userIdInt64)

	var firstChargeTimeInt int64 = 0
	if userFoundAccount.FirstChargeTime != nil {
		firstChargeTimeInt = userFoundAccount.FirstChargeTime.UnixMilli()
	}

	accountFundInfo := UserAccountFundInfo{
		UserID:             userFoundAccount.UserId,
		Balance:            userFoundAccount.Balance,
		FirstChargeTime:    userFoundAccount.FirstChargeTime,
		FirstChargeTimeInt: firstChargeTimeInt,
		FirstChargeAmount:  userFoundAccount.FirstChargeAmount,
		RechargeAmount:     rechargeAmount,
		WithdrawAmount:     withdrawAmount,
		FrozenAmount:       userFoundAccount.FrozenBalance,
		AvaliableBalance:   userFoundAccount.AvailableBalance,
	}
	response.OkWithData(accountFundInfo, c)
}

// BatchSendFund 批量发送资金
// @Tags UserFundAccounts
// @Summary 批量发送资金
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userfundReq.UserFundAccountsSearch true "发送参数"
// @Success 200 {object} response.Response{msg=string} "发送成功"
// @Router /userFundAccounts/batchSendFund [post]
func (userFundAccountsApi *UserFundAccountsApi) BatchSendFund(c *gin.Context) {
	// 参数解析
	var userFund userfundReq.UserFundAccountsSearch
	err := c.ShouldBindJSON(&userFund)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 使用strings.Split函数将字符串分割成数组
	strArray := strings.Split(userFund.UserIds, ",")

	// 使用for循环遍历数组
	for _, userId := range strArray {
		userFundAccountsService.AmountSend(userId, userFund.SendAmount)
	}
	response.OkWithData("success", c)
}
