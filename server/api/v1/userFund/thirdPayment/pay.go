/**
* @Author: Jackey
* @Date: 1/2/25 4:47 pm
 */

package thirdPayment

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/constants"
	"github.com/gin-gonic/gin"
	. "github.com/shopspring/decimal"
	"io"
	"strconv"
)

/*-------创建充值订单-------*/

type RechargeRequestInfo struct {
	CurrencyMoney string `json:"currencyMoney"`
	CurrencyCode  string `json:"currencyCode"`
	CoinCode      string `json:"coinCode"`
	SyncJumpURL   string `json:"syncJumpURL"`
	UserOrderId   string `json:"userOrderId"`
	Language      int    `json:"language"`
	UserCustomId  string `json:"userCustomId"`
	Remark        string `json:"remark"`
}

type RechargeResponseInfo struct {
	OrderId   string      `json:"orderId"`
	ToAddress string      `json:"toAddress"`
	PayUrl    string      `json:"payUrl"`
	ReqBody   interface{} `json:"reqBody"`
}

/*-------创建提现订单-------*/

type WithdrawCreateRequestInfo struct {
	CurrencyCode    string `json:"currencyCode"`
	CoinCode        string `json:"coinCode"`
	UserOrderId     string `json:"userOrderId"`
	UserCustomId    string `json:"userCustomId"`
	WithdrawAddress string `json:"withdrawAddress"`
	CurrencyAmount  string `json:"currencyAmount"`
}

type WithdrawCreateResponseInfo struct {
	OrderId    string      `json:"orderId"`
	ToAddress  string      `json:"toAddress"`
	Commission Decimal     `json:"commission"`
	ReqBody    interface{} `json:"reqBody"`
}

/*-------查询充值订单-------*/

type RechargeQueryRequestInfo struct {
	OrderId string `json:"orderId"`
}

type RechargeQueryResponseInfo struct {
	OrderStatus      string      `json:"orderStatus"`
	OrderId          string      `json:"orderId"`
	CoinAddress      string      `json:"coinAddress"`
	CoinReceiptMoney string      `json:"coinReceiptMoney"`
	RespInfo         interface{} `json:"respInfo"`
}

type WithdrawQueryRequestInfo struct {
	OrderId string `json:"orderId"`
}

type WithdrawQueryResponseInfo struct {
	OrderStatus      string      `json:"orderStatus"`
	OrderId          string      `json:"orderId"`
	CoinAddress      string      `json:"coinAddress"`
	CoinReceiptMoney Decimal     `json:"coinReceiptMoney"`
	RespInfo         interface{} `json:"respInfo"`
}

func SendRechargeRequest(thirdPaymentId int, reqInfo RechargeRequestInfo) (respInfo RechargeResponseInfo, err error) {

	var (
		pUri, pUserId, pSecretKey = constants.GetThirdPaymentInfoById(thirdPaymentId)
		notifyUrl                 = constants.GetThirdPaymentRechargeNotifyUrlById(thirdPaymentId)
		pUserIdInt, _             = strconv.Atoi(pUserId)
	)
	switch thirdPaymentId {
	case constants.ThirdPaymentIdForAAPay:
		var (
			aaPayResp AAPayRechargeResp
			aaPayReq  = AAPayRechargeReq{
				UserID:         int64(pUserIdInt),
				CurrencyMoney:  reqInfo.CurrencyMoney,
				CurrencyCode:   reqInfo.CurrencyCode,
				CoinCode:       reqInfo.CoinCode,
				AsyncNoticeUrl: notifyUrl,
				SyncJumpURL:    reqInfo.SyncJumpURL,
				UserOrderId:    reqInfo.UserOrderId,
				Language:       reqInfo.Language,
				UserCustomId:   reqInfo.UserCustomId,
				Remark:         reqInfo.Remark,
			}
		)
		aaPayResp, err = aaPaySendRechargeRequest(pUri, constants.AAPayRouter.RechargeCreat, aaPayReq, int64(pUserIdInt), pSecretKey, "zh-CN")
		if err != nil {
			return
		}
		respInfo.OrderId = strconv.Itoa(aaPayResp.Data.OrderId)
		respInfo.ToAddress = aaPayResp.Data.CoinAddress
		respInfo.PayUrl = aaPayResp.Data.PayAddressURL
		respInfo.ReqBody = aaPayReq
	default:
		//没有此类充值通道
		err = errors.New(fmt.Sprintf("无此类三方通道->id:%d", thirdPaymentId))
	}

	return
}

func SendWithdrawRequest(thirdPaymentId int, reqInfo WithdrawCreateRequestInfo) (respInfo WithdrawCreateResponseInfo, err error) {

	var (
		pUri, pUserId, pSecretKey = constants.GetThirdPaymentInfoById(thirdPaymentId)
		notifyUrl                 = constants.GetThirdPaymentWithdrawNotifyUrlById(thirdPaymentId)
		pUserIdInt, _             = strconv.Atoi(pUserId)
	)
	switch thirdPaymentId {
	case constants.ThirdPaymentIdForAAPay:
		var (
			aaPayResp AAPayWithdrawResp
			aaPayReq  = AAPayWithdrawReq{
				UserID:           int64(pUserIdInt),
				UserWithdrawalId: reqInfo.UserOrderId,
				UserCustomId:     reqInfo.UserCustomId,
				WithdrawAddress:  reqInfo.WithdrawAddress,
				CurrencyCode:     reqInfo.CurrencyCode,
				CoinCode:         reqInfo.CoinCode,
				CurrencyAmount:   reqInfo.CurrencyAmount,
				AsyncNoticeUrl:   notifyUrl,
			}
		)
		aaPayResp, err = aaPaySendWithdrawRequest(pUri, constants.AAPayRouter.WithdrawCreat, aaPayReq, int64(pUserIdInt), pSecretKey, "zh-CN")
		if err != nil {
			return
		}
		respInfo.OrderId = strconv.Itoa(aaPayResp.Data.WithdrawId)
		respInfo.ToAddress = aaPayResp.Data.WithdrawalAddress
		respInfo.Commission = aaPayResp.Data.Commission
		respInfo.ReqBody = aaPayReq
	default:
		//没有此类充值通道
		err = errors.New(fmt.Sprintf("无此类三方通道->id:%d", thirdPaymentId))
	}

	return
}

func SendRechargeQueryRequest(thirdPaymentId int, reqInfo RechargeQueryRequestInfo) (respInfo RechargeQueryResponseInfo, err error) {

	var (
		pUri, pUserId, pSecretKey = constants.GetThirdPaymentInfoById(thirdPaymentId)
		pUserIdInt, _             = strconv.Atoi(pUserId)
	)
	switch thirdPaymentId {
	case constants.ThirdPaymentIdForAAPay:
		var (
			aaPayResp       AAPayRechargeQueryResp
			queryOrderId, _ = strconv.Atoi(reqInfo.OrderId)
			aaPayReq        = AAPayRechargeQueryReq{
				UserId:  int64(pUserIdInt),
				OrderId: int64(queryOrderId),
			}
		)
		aaPayResp, err = aaPaySendRechargeQueryRequest(pUri, constants.AAPayRouter.RechargeQuery, aaPayReq, int64(pUserIdInt), pSecretKey, "zh-CN")
		if err != nil {
			return
		}
		respInfo.OrderId = strconv.Itoa(aaPayResp.Data.OrderId)
		respInfo.OrderStatus = strconv.Itoa(aaPayResp.Data.OrderStatus)
		respInfo.CoinAddress = aaPayResp.Data.CoinAddress
		respInfo.CoinReceiptMoney = aaPayResp.Data.CoinReceiptMoney
		respInfo.RespInfo = aaPayResp
	default:
		//没有此类充值通道
		err = errors.New(fmt.Sprintf("无此类三方通道->id:%d", thirdPaymentId))
	}

	return
}

func SendWithdrawQueryRequest(thirdPaymentId int, reqInfo WithdrawQueryRequestInfo) (respInfo WithdrawQueryResponseInfo, err error) {

	var (
		pUri, pUserId, pSecretKey = constants.GetThirdPaymentInfoById(thirdPaymentId)
		pUserIdInt, _             = strconv.Atoi(pUserId)
	)
	switch thirdPaymentId {
	case constants.ThirdPaymentIdForAAPay:
		var (
			aaPayResp       AAPayWithdrawQueryResp
			queryOrderId, _ = strconv.Atoi(reqInfo.OrderId)
			aaPayReq        = AAPayWithdrawQueryReq{
				UserId:     int64(pUserIdInt),
				WithdrawId: int64(queryOrderId),
			}
		)
		aaPayResp, err = aaPaySendWithdrawQueryRequest(pUri, constants.AAPayRouter.WithdrawQuery, aaPayReq, int64(pUserIdInt), pSecretKey, "zh-CN")
		if err != nil {
			return
		}
		respInfo.OrderId = strconv.Itoa(aaPayResp.Data.WithdrawId)
		respInfo.OrderStatus = strconv.Itoa(aaPayResp.Data.WithdrawStatus)
		respInfo.CoinAddress = aaPayResp.Data.WithdrawAddress
		respInfo.CoinReceiptMoney = aaPayResp.Data.CoinMoney
		respInfo.RespInfo = aaPayResp
	default:
		//没有此类充值通道
		err = errors.New(fmt.Sprintf("无此类三方通道->id:%d", thirdPaymentId))
	}

	return
}

func GetRechargeNotifyReqInfo(paymentId int, c *gin.Context) (orderId, userOrderId string, err error) {

	var (
		notifyReqInfo = getRechargeNotifyInfoType(paymentId)
		requestBody   []byte
	)
	//先解析请求参数
	if err = c.ShouldBindJSON(notifyReqInfo); err != nil {
		requestBody, _ = io.ReadAll(c.Request.Body)
		err = errors.New(fmt.Sprintf("解析请求参数错误:%s | %s", err.Error(), string(requestBody)))
		return
	}

	switch paymentId {
	case constants.ThirdPaymentIdForAAPay:
		//断言请求参数为对应的结构体
		notifyReq, isOk := notifyReqInfo.(*AAPayRechargeNotifyReq)
		if !isOk {
			err = errors.New(fmt.Sprintf("绑定请求参数错误"))
			return
		}
		// 赋值
		orderId = strconv.Itoa(notifyReq.Data.OrderId)
		userOrderId = notifyReq.Data.UserOrderId

	default:
		err = errors.New(fmt.Sprintf("paymentId:%d 不存在", paymentId))
		return
	}
	return
}

func GetWithdrawNotifyReqInfo(paymentId int, c *gin.Context) (orderId, userOrderId string, err error) {

	var (
		notifyReqInfo = getWithdrawNotifyInfoType(paymentId)
		requestBody   []byte
	)
	//先解析请求参数
	if err = c.ShouldBindJSON(notifyReqInfo); err != nil {
		requestBody, _ = io.ReadAll(c.Request.Body)
		err = errors.New(fmt.Sprintf("解析请求参数错误:%s | %s", err.Error(), string(requestBody)))
		return
	}

	switch paymentId {
	case constants.ThirdPaymentIdForAAPay:
		//断言请求参数为对应的结构体
		notifyReq, isOk := notifyReqInfo.(*AAPayWithdrawNotifyReq)
		if !isOk {
			err = errors.New(fmt.Sprintf("绑定请求参数错误"))
			return
		}
		// 赋值
		orderId = strconv.Itoa(notifyReq.Data.WithdrawId)
		userOrderId = notifyReq.Data.UserWithdrawalId

	default:
		err = errors.New(fmt.Sprintf("paymentId:%d 不存在", paymentId))
		return
	}
	return
}

func getRechargeNotifyInfoType(paymentId int) (info interface{}) {
	switch paymentId {
	case constants.ThirdPaymentIdForAAPay:
		return &AAPayRechargeNotifyReq{}
	}
	return
}

func getWithdrawNotifyInfoType(paymentId int) (info interface{}) {
	switch paymentId {
	case constants.ThirdPaymentIdForAAPay:
		return &AAPayWithdrawNotifyReq{}
	}
	return
}
