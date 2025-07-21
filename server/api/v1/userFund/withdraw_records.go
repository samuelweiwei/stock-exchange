package userFund

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/userFund/thirdPayment"
	"github.com/flipped-aurora/gin-vue-admin/server/constants"
	"github.com/flipped-aurora/gin-vue-admin/server/i18n"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/shopspring/decimal"
	"net/http"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"github.com/flipped-aurora/gin-vue-admin/server/enums"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/userfund"
	userfundReq "github.com/flipped-aurora/gin-vue-admin/server/model/userfund/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WithdrawRecordsApi struct{}

// CreateWithdrawRecords 创建withdrawRecords表
// @Tags WithdrawRecords
// @Summary 创建withdrawRecords表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userfund.WithdrawRecords true "创建withdrawRecords表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /withdrawRecords/createWithdrawRecords [post]
func (withdrawRecordsApi *WithdrawRecordsApi) CreateWithdrawRecords(c *gin.Context) {
	var withdrawRecords userfund.WithdrawRecords
	err := c.ShouldBindJSON(&withdrawRecords)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = withdrawRecordsService.CreateWithdrawRecords(&withdrawRecords)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteWithdrawRecords 删除withdrawRecords表
// @Tags WithdrawRecords
// @Summary 删除withdrawRecords表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userfund.WithdrawRecords true "删除withdrawRecords表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /withdrawRecords/deleteWithdrawRecords [delete]
func (withdrawRecordsApi *WithdrawRecordsApi) DeleteWithdrawRecords(c *gin.Context) {
	ID := c.Query("ID")
	err := withdrawRecordsService.DeleteWithdrawRecords(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteWithdrawRecordsByIds 批量删除withdrawRecords表
// @Tags WithdrawRecords
// @Summary 批量删除withdrawRecords表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /withdrawRecords/deleteWithdrawRecordsByIds [delete]
func (withdrawRecordsApi *WithdrawRecordsApi) DeleteWithdrawRecordsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := withdrawRecordsService.DeleteWithdrawRecordsByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateWithdrawRecords 更新withdrawRecords表
// @Tags WithdrawRecords
// @Summary 更新withdrawRecords表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userfund.WithdrawRecords true "更新withdrawRecords表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /withdrawRecords/updateWithdrawRecords [put]
func (withdrawRecordsApi *WithdrawRecordsApi) UpdateWithdrawRecords(c *gin.Context) {
	var withdrawRecords userfund.WithdrawRecords
	err := c.ShouldBindJSON(&withdrawRecords)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = withdrawRecordsService.UpdateWithdrawRecords(withdrawRecords)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindWithdrawRecords 用id查询withdrawRecords表
// @Tags WithdrawRecords
// @Summary 用id查询withdrawRecords表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userfund.WithdrawRecords true "用id查询withdrawRecords表"
// @Success 200 {object} response.Response{data=userfund.WithdrawRecords,msg=string} "查询成功"
// @Router /withdrawRecords/findWithdrawRecords [get]
func (withdrawRecordsApi *WithdrawRecordsApi) FindWithdrawRecords(c *gin.Context) {
	ID := c.Query("ID")
	rewithdrawRecords, err := withdrawRecordsService.GetWithdrawRecords(ID)
	currentUserId := utils.GetUserID(c)
	if rewithdrawRecords.Locker > 0 && rewithdrawRecords.Locker == int(currentUserId) {
		// 如果找到匹配的记录，设置 hashAuth 为 true
		rewithdrawRecords.HashAuth = true
	} else {
		rewithdrawRecords.HashAuth = false
	}
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(rewithdrawRecords, c)
}

// GetWithdrawRecordsList 分页获取withdrawRecords表列表
// @Tags WithdrawRecords
// @Summary 分页获取withdrawRecords表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userfundReq.WithdrawRecordsSearch true "分页获取withdrawRecords表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /withdrawRecords/getWithdrawRecordsList [get]
func (withdrawRecordsApi *WithdrawRecordsApi) GetWithdrawRecordsList(c *gin.Context) {
	var pageInfo userfundReq.WithdrawRecordsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	memberId := int(utils.GetUserInfo(c).FrontUserId)
	//list, total, err := withdrawRecordsService.GetWithdrawRecordsInfoList(pageInfo)
	var list []userfund.WithdrawRecords // 定义一个空的切片
	total := int64(0)                   // 初始化 total 为 0

	// 扩展 list 切片为指定长度
	list = make([]userfund.WithdrawRecords, pageInfo.PageSize)
	if memberId > 0 { //判断存在
		pageInfo.MemberId = int(utils.GetUserInfo(c).FrontUserId)
		list, total, err = withdrawRecordsService.GetWithdrawRecordsInfoListByRootUser(pageInfo)
	} else {
		list, total, err = withdrawRecordsService.GetWithdrawRecordsInfoList(pageInfo)
	}
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	currentUserId := utils.GetUserID(c)
	// 遍历记录列表
	for i := range list {
		record := &list[i] // 使用指针
		// 检查 locker 属性是否等于当前登录用户
		if record.Locker > 0 && record.Locker == int(currentUserId) {
			// 如果找到匹配的记录，设置 hashAuth 为 true
			record.HashAuth = true
		} else {
			record.HashAuth = false
		}
		if record.WithdrawTime == nil || record.WithdrawTime.IsZero() {
			record.WithdrawTimeInt = 0
		} else {
			record.WithdrawTimeInt = record.WithdrawTime.UnixMilli()
		}

		if record.ApprovalTime == nil || record.ApprovalTime.IsZero() {
			record.ApprovalTimeInt = 0
		} else {
			record.ApprovalTimeInt = record.ApprovalTime.UnixMilli()
		}

		if record.CreatedAt.IsZero() {
			record.CreatedAtInt = 0
		} else {
			record.CreatedAtInt = record.CreatedAt.UnixMilli()
		}

		if record.UpdatedAt.IsZero() {
			record.UpdatedAtInt = 0
		} else {
			record.UpdatedAtInt = record.UpdatedAt.UnixMilli()
		}

	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetWithdrawRecordsPublic 不需要鉴权的withdrawRecords表接口
// @Tags WithdrawRecords
// @Summary 不需要鉴权的withdrawRecords表接口
// @accept application/json
// @Produce application/json
// @Param data query userfundReq.WithdrawRecordsSearch true "分页获取withdrawRecords表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /withdrawRecords/getWithdrawRecordsPublic [get]
func (withdrawRecordsApi *WithdrawRecordsApi) GetWithdrawRecordsPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	withdrawRecordsService.GetWithdrawRecordsPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的withdrawRecords表接口信息",
	}, "获取成功", c)
}

// WithdrawNotify 提现回调接口
// @Tags WithdrawRecords
// @Summary 提现回调接口
// @accept application/json
// @Produce application/json
// @Param data body WithdrawResponse true "提现回调数据"
// @Success 200 {string} string "ok"
// @Router /withdrawRecords/withdrawNotify [post]
func (withdrawRecordsApi *WithdrawRecordsApi) WithdrawNotify(c *gin.Context) {

	var (
		paymentId int
		paymentName,
		originStatus,
		notifyResultMsg string
		err           error
		succeedString = "success"
		jsonData      []byte
		orderId,
		userOrderId string
		order userfund.WithdrawRecords
	)
	defer func() {
		if err == nil {
			c.String(http.StatusOK, succeedString)
			succeedMsg := fmt.Sprintf("【%s提现回调】->成功"+
				"\n[订单号]：%s"+
				"\n[原状态]：%s->%s"+
				"\n[处理结果]：%s",
				paymentName, order.OrderId, originStatus,
				enums.GetWithdrawStatusString(originStatus),
				notifyResultMsg)
			utils.SendMsgToTgAsync(succeedMsg)
		} else {
			_ = c.Error(errors.New("提现回调失败喽！！！"))
			response.FailWithMessage(err.Error(), c)
		}
	}()

	//先获取三方id
	paymentId, err = strconv.Atoi(c.Param("paymentId"))
	if err != nil {
		err = errors.New(fmt.Sprintf("paymentId错误:" + err.Error()))
		return
	}
	paymentName = constants.GetThirdPaymentNameById(paymentId)

	//根据三方id获取必要请求参数
	orderId, userOrderId, err = thirdPayment.GetWithdrawNotifyReqInfo(paymentId, c)
	if err != nil {
		err = errors.New(fmt.Sprintf("获取订单号信息错误:" + err.Error()))
		return
	}

	//获取原始订单数据
	order, err = withdrawRecordsService.GetWithdrawRecordsByOrderId(userOrderId)
	if err != nil {
		err = errors.New(fmt.Sprintf("获取原始订单数据失败:" + err.Error()))
		return
	}
	originStatus = order.OrderStatus

	//如果是已经提现完成的 第三方仍然调用 返回OK
	if order.OrderStatus == enums.WITHDRAWED {
		notifyResultMsg = "订单已是<完成>状态，无任何更新操作"
		return
	}

	//查询三方订单状态
	var (
		thirdRespInfo thirdPayment.WithdrawQueryResponseInfo
		thirdReqInfo  = thirdPayment.WithdrawQueryRequestInfo{
			OrderId: orderId,
		}
	)
	thirdRespInfo, err = thirdPayment.SendWithdrawQueryRequest(paymentId, thirdReqInfo)
	if err != nil {
		err = errors.New(fmt.Sprintf("查询订单失败:%s " + err.Error()))
		return
	}

	//判断订单状态 如果订单状态不正确，直接返回
	var (
		isSucceed = constants.IsThirdWithdrawSucceed(paymentId, thirdRespInfo.OrderStatus)
		isFailed  = constants.IsThirdWithdrawFailed(paymentId, thirdRespInfo.OrderStatus)
	)
	if !isSucceed && !isFailed {
		notifyResultMsg = fmt.Sprintf("查询订单状态为：%s<非成功或失败>，无任何更新操作", thirdRespInfo.OrderStatus)
		return
	}

	//接下去要判断成功或者失败，成功和失败需要做不同的处理
	if isFailed {
		userAccount, err := userFundAccountsService.GetUserFundAccountsByUserId(order.MemberId)
		if err != nil {
			global.GVA_LOG.Error("获取用户账户失败!", zap.Error(err))
			response.FailWithMessage("获取用户账户失败:"+err.Error(), c)
			return
		}
		order.ReviewStatus = enums.INVALID //如果失败 标记提现订单完成，审核状态为无效
		err = userFundAccountsService.UpdateUserFundAccountsWithFlowAndRefusedWithdrawRecords(order, userAccount)
		if err != nil {
			global.GVA_LOG.Error("提现更新用户资金失败!", zap.Error(err))
			response.FailWithMessage("提现更新用户资金失败:"+err.Error(), c)
			return
		}
		return
	}

	order.ThirdOrderId = thirdRespInfo.OrderId
	order.OrderStatus = enums.WITHDRAWED
	order.ToAddress = thirdRespInfo.CoinAddress
	order.RealReceived = thirdRespInfo.CoinReceiptMoney
	order.UpdatedAt = time.Now()
	jsonData, err = json.Marshal(thirdRespInfo.RespInfo)
	if err != nil {
		err = errors.New(fmt.Sprintf("将请求参数解析成json错误:" + err.Error()))
		return
	}
	order.NotifyContent = string(jsonData)

	//userAccount, err := userFundAccountsService.GetUserFundAccountsByUserId(order.MemberId)
	//用户资金表-余额更新
	err = userFundAccountsService.UpdateUserFundAccountsWithFlowAndWithdrawRecords(order)
	if err != nil {
		err = errors.New(fmt.Sprintf("提现回调更新资金失败:" + err.Error()))
		return
	}
	notifyResultMsg = fmt.Sprintf("[成功] -> 金额: %s", thirdRespInfo.CoinReceiptMoney)
	return
}

// ReviewRecord 审核提现订单
// @Tags WithdrawRecords
// @Summary 审核提现订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userfund.WithdrawRecords true "审核提现订单"
// @Success 200 {object} response.Response{msg=string} "审核成功"
// @Router /withdrawRecords/reviewRecord [post]
func (withdrawRecordsApi *WithdrawRecordsApi) ReviewRecord(c *gin.Context) {
	// 接收请求参数
	var reviewRequest userfund.WithdrawRecords
	err := c.ShouldBindJSON(&reviewRequest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 获取原记录
	record, err := withdrawRecordsService.GetWithdrawRecords(strconv.Itoa(int(reviewRequest.ID)))
	if err != nil {
		response.FailWithMessage("获取记录失败:"+err.Error(), c)
		return
	}

	// 检查锁定状态
	if record.IsLock == enums.UNLOCK {
		//global.GVA_LOG.Error("订单处于锁定状态!", zap.Error(err))
		response.FailWithMessage("订单处于未锁定状态，必须先锁定才可以操作！", c)
		return
	}

	if record.Locker != int(utils.GetUserID(c)) {
		response.FailWithMessage("当前用户没有审核权限!", c)
		return
	}

	// 检查审核状态
	if record.ReviewStatus != "1" {
		response.FailWithMessage("订单状态不正确！", c)
		return
	}

	// 更新审核相关字段
	record.ReviewStatus = reviewRequest.ReviewStatus
	record.RefusedReason = reviewRequest.RefusedReason
	record.UpdatedAt = time.Now()
	// 审核通过的处理逻辑
	userAccount, err := userFundAccountsService.GetUserFundAccountsByUserId(record.MemberId)
	if err != nil {
		global.GVA_LOG.Error("获取用户账户失败!", zap.Error(err))
		response.FailWithMessage("获取用户账户失败:"+err.Error(), c)
		return
	}
	if reviewRequest.ReviewStatus == enums.PASSED {

		// 检查冻结余额
		if userAccount.FrozenBalance == decimal.Zero || userAccount.FrozenBalance.LessThan(record.WithdrawAmount) {
			global.GVA_LOG.Error("冻结余额异常!")
			response.FailWithMessage("冻结余额异常", c)
			return
		}
		if record.WithdrawType == enums.WithDrawTypeQuick {
			record.OrderStatus = enums.WITHDRAW_CHECKED
			var (
				targetStr       = record.ExchangedAmountTarget.String()
				withDrawChannel userfund.WithdrawChannels
			)
			// 根据id查询提现渠道接口
			withDrawChannel, err = withdrawChannelsService.GetWithdrawChannels(strconv.Itoa(record.WithdrawChannelId))
			if err != nil {
				global.GVA_LOG.Error("查询失败!", zap.Error(err))
				response.FailWithMessage("查询失败:"+err.Error(), c)
				return
			}

			var (
				thirdPaymentId = withDrawChannel.ThirdPayId
				thirdRespInfo  thirdPayment.WithdrawCreateResponseInfo
				thirdReqInfo   = thirdPayment.WithdrawCreateRequestInfo{
					CurrencyCode:    withDrawChannel.ThirdCurrencyCode,
					CoinCode:        withDrawChannel.ThirdCoinCode,
					UserOrderId:     record.OrderId,
					UserCustomId:    strconv.Itoa(int(utils.GetUserID(c))),
					WithdrawAddress: record.ToAddress,
					CurrencyAmount:  targetStr,
				}
			)
			//三方订单号、收款地址、支付地址、传给三方的所有参数
			thirdRespInfo, err = thirdPayment.SendWithdrawRequest(thirdPaymentId, thirdReqInfo)
			if err != nil {
				response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.OperateFail, response.ERROR)+err.Error(), c)
				return
			}

			record.ThirdOrderId = thirdRespInfo.OrderId
			record.ToAddress = thirdRespInfo.ToAddress
			record.ThirdCommission = thirdRespInfo.Commission
			// 更新提现记录状态
			record.UpdatedAt = time.Now()
			//更新订单
			err = withdrawRecordsService.UpdateWithdrawRecords(record)
			if err != nil {
				global.GVA_LOG.Error("更新失败!", zap.Error(err))
				response.FailWithMessage("更新失败:"+err.Error(), c)
				return
			}
		} else {
			// 更新提现记录状态
			approvalTime := time.Now()
			record.ApprovalTime = &approvalTime
			record.OrderStatus = enums.WITHDRAWED
			err = userFundAccountsService.UpdateUserFundAccountsWithFlowAndWithdrawRecords(record)
			if err != nil {
				global.GVA_LOG.Error("提现更新用户资金失败!", zap.Error(err))
				response.FailWithMessage("提现更新用户资金失败:"+err.Error(), c)
				return
			}
		}

	} else if reviewRequest.ReviewStatus == enums.REFUSED {
		//record.OrderStatus = enums.FAILED
		// 审核拒绝的处理逻辑
		err = userFundAccountsService.UpdateUserFundAccountsWithFlowAndRefusedWithdrawRecords(record, userAccount)
		if err != nil {
			global.GVA_LOG.Error("提现更新用户资金失败!", zap.Error(err))
			response.FailWithMessage("提现更新用户资金失败:"+err.Error(), c)
			return
		}

	}

	response.OkWithMessage("更新成功", c)
}

// LockRecord 锁定/解锁提现记录
// @Tags WithdrawRecords
// @Summary 锁定/解锁提现记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userfund.WithdrawRecords true "锁定/解锁提现记录"
// @Success 200 {object} response.Response{msg=string} "操作成功"
// @Router /withdrawRecords/lock [post]
func (withdrawRecordsApi *WithdrawRecordsApi) LockRecord(c *gin.Context) {
	var withdrawRecords userfund.RechargeRecords
	err := c.ShouldBindJSON(&withdrawRecords)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	withdrawRecord, err := withdrawRecordsService.GetWithdrawRecords(strconv.Itoa(int(withdrawRecords.ID)))
	withdrawRecord.IsLock = withdrawRecords.IsLock
	withdrawRecord.UpdatedAt = time.Now()
	if withdrawRecord.IsLock == enums.UNLOCK && withdrawRecord.Locker == int(utils.GetUserID(c)) { //解锁  解除绑定权限归属
		withdrawRecord.Locker = -1
	} else { //锁定 判断是否有权限锁定
		if withdrawRecord.Locker > 0 {
			response.FailWithMessage("该订单已经拥有权限归属!", c)
			return
		}
		withdrawRecord.Locker = int(utils.GetUserID(c))
	}
	err = withdrawRecordsService.UpdateWithdrawRecords(withdrawRecord)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetUserRecordsList 前端-根据用户获取提现记录
// @Tags WithdrawRecords
// @Summary 根据用户ID获取提现记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userfundReq.UserWithdrawRecordsSearch true "分页参数"
// @Description 获取当前登录用户的所有提现记录
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /withdrawRecords/getUserRecordsList [get]
func (withdrawRecordsApi *WithdrawRecordsApi) GetUserRecordsList(c *gin.Context) {
	userId := utils.GetUserIDFrontUser(c)
	if userId == 0 {
		//response.FailWithMessage("用户未登录", c)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ErrGetTokenFail, response.ERROR), c)
		return
	}

	var pageInfo userfundReq.UserWithdrawRecordsSearch
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 将 uint 类型转换为 int64
	userIdInt64 := int64(userId)

	records, total, err := withdrawRecordsService.GetWithdrawRecordsByUserIdWithPagination(userIdInt64, pageInfo.Page, pageInfo.PageSize)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.OperateFail, response.ERROR), c)
		return
	}
	// 遍历记录列表
	for i := range records {
		record := &records[i] // 使用指针
		if record.WithdrawTime == nil || record.WithdrawTime.IsZero() {
			record.WithdrawTimeInt = 0
		} else {
			record.WithdrawTimeInt = record.WithdrawTime.UnixMilli()
		}

		if record.ApprovalTime == nil || record.ApprovalTime.IsZero() {
			record.ApprovalTimeInt = 0
		} else {
			record.ApprovalTimeInt = record.ApprovalTime.UnixMilli()
		}

		if record.CreatedAt.IsZero() {
			record.CreatedAtInt = 0
		} else {
			record.CreatedAtInt = record.CreatedAt.UnixMilli()
		}

		if record.UpdatedAt.IsZero() {
			record.UpdatedAtInt = 0
		} else {
			record.UpdatedAtInt = record.UpdatedAt.UnixMilli()
		}

	}
	response.OkWithDetailed(response.PageResult{
		List:     records,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetUserRecordDetail 前端-获取用户提现记录详情
// @Tags WithdrawRecords
// @Summary 获取用户提现记录详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id path string true "记录ID"
// @Description 获取当前登录用户的指定提现记录详情
// @Success 200 {object} response.Response{data=userfund.WithdrawRecords,msg=string} "获取成功"
// @Router /withdrawRecords/getUserRecordDetail/{id} [get]
func (withdrawRecordsApi *WithdrawRecordsApi) GetUserRecordDetail(c *gin.Context) {
	userId := utils.GetUserIDFrontUser(c)
	if userId == 0 {
		response.FailWithMessage("用户未登录", c)
		return
	}

	recordId := c.Param("id")
	if recordId == "" {
		response.FailWithMessage("记录ID不能为空", c)
		return
	}

	// 将 uint 类型转换为 int64
	userIdInt64 := int64(userId)

	record, err := withdrawRecordsService.GetWithdrawRecordDetail(userIdInt64, recordId)
	if record.WithdrawTime == nil || record.WithdrawTime.IsZero() {
		record.WithdrawTimeInt = 0
	} else {
		record.WithdrawTimeInt = record.WithdrawTime.UnixMilli()
	}

	if record.ApprovalTime == nil || record.ApprovalTime.IsZero() {
		record.ApprovalTimeInt = 0
	} else {
		record.ApprovalTimeInt = record.ApprovalTime.UnixMilli()
	}

	if record.CreatedAt.IsZero() {
		record.CreatedAtInt = 0
	} else {
		record.CreatedAtInt = record.CreatedAt.UnixMilli()
	}

	if record.UpdatedAt.IsZero() {
		record.UpdatedAtInt = 0
	} else {
		record.UpdatedAtInt = record.UpdatedAt.UnixMilli()
	}
	if err != nil {
		global.GVA_LOG.Error("获取提现记录详情失败!", zap.Error(err))
		response.FailWithMessage("获取提现记录详情失败:"+err.Error(), c)
		return
	}

	response.OkWithData(record, c)
}
