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

type RechargeRecordsApi struct{}

// RequestBody 定义了请求体的结构体
type QueryRequestBody struct {
	UserID  int64 `json:"user_id"`
	OrderID int64 `json:"order_id"`
}

// RequestBody 定义了请求体的结构体
type WithdrawQueryRequestBody struct {
	UserID     int64 `json:"user_id"`
	WithdrawId int64 `json:"withdrawal_id"`
}

// CreateRechargeRecords 创建rechargeRecords表
// @Tags RechargeRecords
// @Summary 创建rechargeRecords表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userfund.RechargeRecords true "创建rechargeRecords表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /rechargeRecords/createRechargeRecords [post]
func (rechargeRecordsApi *RechargeRecordsApi) CreateRechargeRecords(c *gin.Context) {
	var rechargeRecords userfund.RechargeRecords
	err := c.ShouldBindJSON(&rechargeRecords)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = rechargeRecordsService.CreateRechargeRecords(&rechargeRecords)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteRechargeRecords 删除rechargeRecords表
// @Tags RechargeRecords
// @Summary 删除rechargeRecords表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userfund.RechargeRecords true "删除rechargeRecords表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /rechargeRecords/deleteRechargeRecords [delete]
func (rechargeRecordsApi *RechargeRecordsApi) DeleteRechargeRecords(c *gin.Context) {
	ID := c.Query("ID")
	err := rechargeRecordsService.DeleteRechargeRecords(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteRechargeRecordsByIds 批量删除rechargeRecords表
// @Tags RechargeRecords
// @Summary 批量删除rechargeRecords表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /rechargeRecords/deleteRechargeRecordsByIds [delete]
func (rechargeRecordsApi *RechargeRecordsApi) DeleteRechargeRecordsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := rechargeRecordsService.DeleteRechargeRecordsByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateRechargeRecords 更新rechargeRecords表
// @Tags RechargeRecords
// @Summary 更新rechargeRecords表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userfund.RechargeRecords true "更新rechargeRecords表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /rechargeRecords/updateRechargeRecords [put]
func (rechargeRecordsApi *RechargeRecordsApi) UpdateRechargeRecords(c *gin.Context) {
	var rechargeRecords userfund.RechargeRecords
	err := c.ShouldBindJSON(&rechargeRecords)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = rechargeRecordsService.UpdateRechargeRecords(rechargeRecords)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindRechargeRecords 用id查询rechargeRecords表
// @Tags RechargeRecords
// @Summary 用id查询rechargeRecords表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userfund.RechargeRecords true "用id查询rechargeRecords表"
// @Success 200 {object} response.Response{data=userfund.RechargeRecords,msg=string} "查询成功"
// @Router /rechargeRecords/findRechargeRecords [get]
func (rechargeRecordsApi *RechargeRecordsApi) FindRechargeRecords(c *gin.Context) {
	ID := c.Query("ID")
	rechargeRecords, err := rechargeRecordsService.GetRechargeRecords(ID)
	currentUserId := utils.GetUserID(c)

	if rechargeRecords.Locker == int(currentUserId) {
		// 如果找到匹配的记录，设置 hashAuth 为 true
		rechargeRecords.HashAuth = true
	}
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(rechargeRecords, c)
}

// GetRechargeRecordsList 分页获取rechargeRecords表列表
// @Tags RechargeRecords
// @Summary 分页获取rechargeRecords表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userfundReq.RechargeRecordsSearch true "分页获取rechargeRecords表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /rechargeRecords/getRechargeRecordsList [get]
func (rechargeRecordsApi *RechargeRecordsApi) GetRechargeRecordsList(c *gin.Context) {
	var pageInfo userfundReq.RechargeRecordsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	memberId := int(utils.GetUserInfo(c).FrontUserId)
	var list []userfund.RechargeRecords // 定义一个空的切片
	total := int64(0)                   // 初始化 total 为 0
	// 扩展 list 切片为指定长度
	list = make([]userfund.RechargeRecords, pageInfo.PageSize)
	if memberId > 0 { //判断存在
		pageInfo.MemberId = int(utils.GetUserInfo(c).FrontUserId)
		list, total, err = rechargeRecordsService.GetRechargeRecordsInfoListByRootUser(pageInfo)

	} else {
		list, total, err = rechargeRecordsService.GetRechargeRecordsInfoList(pageInfo)
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
		if record.RechargeTime == nil || record.RechargeTime.IsZero() {
			record.RechargeTimeInt = 0
		} else {
			record.RechargeTimeInt = record.RechargeTime.UnixMilli()
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

// GetRechargeRecordsPublic 不需要鉴权的rechargeRecords表接口
// @Tags RechargeRecords
// @Summary 不需要鉴权的rechargeRecords表接口
// @accept application/json
// @Produce application/json
// @Param data query userfundReq.RechargeRecordsSearch true "分页获取rechargeRecords表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /rechargeRecords/getRechargeRecordsPublic [get]
func (rechargeRecordsApi *RechargeRecordsApi) GetRechargeRecordsPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	rechargeRecordsService.GetRechargeRecordsPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的rechargeRecords表接口信息",
	}, "获取成功", c)
}

// PayNotify 支付回调接口
// @Tags RechargeRecords
// @Summary 支付回调接口
// @accept application/json
// @Produce application/json
// @Param data body AAPayResponse true "支付回调数据"
// @Success 200 {string} string "ok"
// @Router /rechargeRecords/payNotify [post]
func (rechargeRecordsApi *RechargeRecordsApi) PayNotify(c *gin.Context) {

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
		order userfund.RechargeRecords
	)
	defer func() {
		if err == nil {
			c.String(http.StatusOK, succeedString)
			succeedMsg := fmt.Sprintf("【%s充值回调】->成功"+
				"\n[订单号]：%s"+
				"\n[原状态]：%s->%s"+
				"\n[处理结果]：%s",
				paymentName, order.OrderId, originStatus,
				enums.GetRechargeStatusString(originStatus),
				notifyResultMsg)
			utils.SendMsgToTgAsync(succeedMsg)
		} else {
			_ = c.Error(errors.New("充值回调失败喽！！！"))
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
	orderId, userOrderId, err = thirdPayment.GetRechargeNotifyReqInfo(paymentId, c)
	if err != nil {
		err = errors.New(fmt.Sprintf("获取订单号信息错误:" + err.Error()))
		return
	}

	//获取原始订单数据
	order, err = rechargeRecordsService.GetRechargeRecordsByOrderId(userOrderId)
	if err != nil {
		err = errors.New(fmt.Sprintf("获取原始订单数据失败:" + err.Error()))
		return
	}
	originStatus = order.OrderStatus

	//判断订单状态
	//1  订单已成功，重复调用，返回成功
	if order.OrderStatus == enums.SUCCESS {
		notifyResultMsg = "订单已是<成功>状态，无任何更新操作"
		return
	}
	//订单不在支付中，也不是超时
	if order.OrderStatus != enums.PENDING && order.OrderStatus != enums.OUTTIME {
		notifyResultMsg = "订单不在<支付中>或<超时>状态，无任何更新操作"
		return
	}

	//查询三方订单状态
	var (
		thirdRespInfo thirdPayment.RechargeQueryResponseInfo
		thirdReqInfo  = thirdPayment.RechargeQueryRequestInfo{
			OrderId: orderId,
		}
	)
	thirdRespInfo, err = thirdPayment.SendRechargeQueryRequest(paymentId, thirdReqInfo)
	if err != nil {
		err = errors.New(fmt.Sprintf("查询订单失败:%s " + err.Error()))
		return
	}

	//判断订单状态 如果订单状态不正确，直接返回
	if !constants.IsThirdRechargeSucceed(paymentId, thirdRespInfo.OrderStatus) {
		notifyResultMsg = fmt.Sprintf("查询订单状态为：%s<非成功>，无任何更新操作", thirdRespInfo.OrderStatus)
		return
	}
	order.ThirdOrderId = thirdRespInfo.OrderId
	order.OrderStatus = enums.SUCCESS
	order.ToAddress = thirdRespInfo.CoinAddress
	order.CoinReceiptMoney = thirdRespInfo.CoinReceiptMoney
	order.ExchangedAmountUsdt, err = getUsdtAmount(thirdRespInfo.CoinReceiptMoney, order.RechargeRate)
	if err != nil {
		err = errors.New(fmt.Sprintf("折算USDT金额失败，金额:%s "+err.Error(), order.ExchangedAmountUsdt))
		return
	}
	order.UpdatedAt = time.Now()
	jsonData, err = json.Marshal(thirdRespInfo.RespInfo)
	if err != nil {
		err = errors.New(fmt.Sprintf("将请求参数解析成json错误:" + err.Error()))
		return
	}
	order.NotifyContent = string(jsonData)

	//用户资金表-余额更新
	err = userFundAccountsService.UpdateUserFundAccountsWithFlowAndRechargeRecords(order)
	if err != nil {
		err = errors.New(fmt.Sprintf("支付回调更新资金失败:" + err.Error()))
		return
	}
	notifyResultMsg = fmt.Sprintf("[成功] -> 金额: %s", thirdRespInfo.CoinReceiptMoney)
	return
}

// ReviewRecord 审核充值订单
// @Tags RechargeRecords
// @Summary 审核充值订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userfund.RechargeRecords true "审核充值订单"
// @Success 200 {object} response.Response{msg=string} "审核成功"
// @Router /rechargeRecords/reviewRecord [post]
func (rechargeRecordsApi *RechargeRecordsApi) ReviewRecord(c *gin.Context) {
	var rechargeRecords userfund.RechargeRecords
	err := c.ShouldBindJSON(&rechargeRecords)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	rechargeRecord, err := rechargeRecordsService.GetRechargeRecords(strconv.Itoa(int(rechargeRecords.ID)))
	if rechargeRecord.IsLock == enums.UNLOCK {
		//global.GVA_LOG.Error("订单处于锁定状态!", zap.Error(err))
		response.FailWithMessage("订单处于未锁定状态，必须先锁定才可以操作！", c)
		return
	}
	if rechargeRecord.Locker != int(utils.GetUserID(c)) {
		response.FailWithMessage("当前用户没有审核权限!", c)
		return
	}
	if rechargeRecord.ReviewStatus != "1" {
		response.FailWithMessage("订单状态不正确！", c)
		return
	}
	rechargeRecord.ReviewStatus = rechargeRecords.ReviewStatus
	rechargeRecord.RefusedReason = rechargeRecords.RefusedReason
	rechargeRecord.UpdatedAt = time.Now()
	if rechargeRecord.ReviewStatus == enums.PASSED {
		rechargeRecord.OrderStatus = enums.SUCCESS
	}
	if rechargeRecord.ReviewStatus == enums.REFUSED {
		rechargeRecord.OrderStatus = enums.FAILED
	}

	//审核通过
	if rechargeRecord.ReviewStatus == enums.PASSED {
		err = userFundAccountsService.UpdateUserFundAccountsWithFlowAndRechargeRecords(rechargeRecord)
		if err != nil {
			global.GVA_LOG.Error("更新用户资金失败!", zap.Error(err))
			response.FailWithMessage("更新用户资金失败:"+err.Error(), c)
			return
		}

	} else {
		err = rechargeRecordsService.UpdateRechargeRecords(rechargeRecord)
		if err != nil {
			global.GVA_LOG.Error("更新失败!", zap.Error(err))
			response.FailWithMessage("更新失败:"+err.Error(), c)
			return
		}
	}
	response.OkWithMessage("更新成功", c)
}

// LockRecord 锁定/解锁充值记录
// @Tags RechargeRecords
// @Summary 锁定/解锁充值记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userfund.RechargeRecords true "锁定/解锁充值记录"
// @Success 200 {object} response.Response{msg=string} "操作成功"
// @Router /rechargeRecords/lock [post]
func (rechargeRecordsApi *RechargeRecordsApi) LockRecord(c *gin.Context) {
	var rechargeRecords userfund.RechargeRecords
	err := c.ShouldBindJSON(&rechargeRecords)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	rechargeRecord, err := rechargeRecordsService.GetRechargeRecords(strconv.Itoa(int(rechargeRecords.ID)))
	rechargeRecord.IsLock = rechargeRecords.IsLock
	rechargeRecord.UpdatedAt = time.Now()
	if rechargeRecords.IsLock == enums.UNLOCK && rechargeRecord.Locker == int(utils.GetUserID(c)) { //解锁  解除绑定权限归属
		rechargeRecord.Locker = -1
	} else { //锁定 判断是否有权限锁定
		if rechargeRecord.Locker > 0 {
			response.FailWithMessage("该订单已经拥有权限归属!", c)
			return
		}
		rechargeRecord.Locker = int(utils.GetUserID(c))
	}
	err = rechargeRecordsService.UpdateRechargeRecords(rechargeRecord)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// getUsdtAmount 将指定币种金额转换为USDT金额
// params:
//   - coincode: 币种代码，例如"CNY"/"USDT"等
//   - amount: 原始金额
//   - priceusdt: USDT汇率，当coincode为CNY时使用
//
// returns:
//   - *float64: 转换后的USDT金额，保留2位小数
//   - error: 转换过程中的错误信息
func getUsdtAmount(amount string, priceusdt decimal.Decimal) (decimal.Decimal, error) {
	// 将 amount 转换为 float64
	amountFloat, err := decimal.NewFromString(amount)
	if err != nil {
		return decimal.Zero, fmt.Errorf("无法将 amount 转换为 decimal: %w", err)
	}
	totalAmount := amountFloat.Mul(priceusdt)
	return totalAmount, nil

}

// GetUserRecordsList 前端-根据用户获取充值记录
// @Tags RechargeRecords
// @Summary 根据用户ID获取充值记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userfundReq.UserRechargeRecordsSearch true "分页参数"
// @Description 获取当前登录用户的所有充值记录
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /rechargeRecords/getUserRecordsList [get]
func (rechargeRecordsApi *RechargeRecordsApi) GetUserRecordsList(c *gin.Context) {
	userId := utils.GetUserIDFrontUser(c)
	if userId == 0 {
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ErrGetTokenFail, response.ERROR), c)
		return
	}

	var pageInfo userfundReq.UserRechargeRecordsSearch
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 将 uint 类型转换为 int64
	userIdInt64 := int64(userId)

	records, total, err := rechargeRecordsService.GetRechargeRecordsByUserIdWithPagination(userIdInt64, pageInfo.Page, pageInfo.PageSize)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.OperateFail, response.ERROR), c)
		return
	}
	// 遍历记录列表
	for i := range records {
		record := &records[i] // 使用指针
		if record.RechargeTime == nil || record.RechargeTime.IsZero() {
			record.RechargeTimeInt = 0
		} else {
			record.RechargeTimeInt = record.RechargeTime.UnixMilli()
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
	}, i18n.Message(request.GetLanguageTag(c), i18n.OperateSuccess, 0), c)
}

// GetUserRecordDetail 前端-获取用户充值记录详情
// @Tags RechargeRecords
// @Summary 获取用户充值记录详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id path string true "记录ID"
// @Description 获取当前登录用户的指定充值记录详情
// @Success 200 {object} response.Response{data=userfund.RechargeRecords,msg=string} "获取成功"
// @Router /rechargeRecords/getUserRecordDetail/{id} [get]
func (rechargeRecordsApi *RechargeRecordsApi) GetUserRecordDetail(c *gin.Context) {
	userId := utils.GetUserIDFrontUser(c)
	if userId == 0 {
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ErrGetTokenFail, response.ERROR), c)
		return
	}

	recordId := c.Param("id")
	if recordId == "" {
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.OperateFail, response.ERROR), c)
		return
	}

	// 将 uint 类型转换为 int64
	userIdInt64 := int64(userId)

	record, err := rechargeRecordsService.GetRechargeRecordDetail(userIdInt64, recordId)
	if record.RechargeTime == nil || record.RechargeTime.IsZero() {
		record.RechargeTimeInt = 0
	} else {
		record.RechargeTimeInt = record.RechargeTime.UnixMilli()
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
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.OperateFail, response.ERROR), c)
		return
	}
	response.OkWithData(record, c)
}
