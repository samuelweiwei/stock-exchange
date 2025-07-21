/**
* @Author: Jackey
* @Date: 1/1/25 7:39 pm
 */

package thirdPayment

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	. "github.com/shopspring/decimal"
	"io"
	"net/http"
	"strings"
	"time"
)

/*-------创建充值订单-------*/

type AAPayRechargeReq struct {
	UserID         int64  `json:"user_id"`
	CurrencyMoney  string `json:"currency_money"`
	CurrencyCode   string `json:"currency_code"`
	CoinCode       string `json:"coin_code"`
	AsyncNoticeUrl string `json:"asyn_notice_url"`
	SyncJumpURL    string `json:"sync_jump_url"` // 可选
	UserOrderId    string `json:"user_order_id"`
	Language       int    `json:"language"`
	UserCustomId   string `json:"user_custom_id"`
	Remark         string `json:"remark"` // 可选
}

type AAPayRechargeResp struct {
	Code      int                   `json:"code"`
	Message   string                `json:"message"`
	Sign      string                `json:"sign"`
	Timestamp int64                 `json:"timestamp"`
	Data      AAPayRechargeRespData `json:"data"`
}

type AAPayRechargeRespData struct {
	OrderId         int    `json:"order_id"`
	CurrencyCode    string `json:"currency_code"`
	CurrencyMoney   string `json:"currency_money"`
	CoinCode        string `json:"coin_code"`
	CoinMoney       string `json:"coin_money"`
	CoinAddress     string `json:"coin_address"`
	PayAddressURL   string `json:"pay_address_url"`
	OrderExpireTime int64  `json:"order_expire_time"`
}

/*-------创建提现订单-------*/

type AAPayWithdrawReq struct {
	UserID           int64  `json:"user_id"`
	UserWithdrawalId string `json:"user_withdrawal_id"`
	UserCustomId     string `json:"user_custom_id"`
	WithdrawAddress  string `json:"withdrawal_address"`
	CurrencyCode     string `json:"currency_code"`
	CoinCode         string `json:"coin_code"`
	CurrencyAmount   string `json:"currency_amount"`
	AsyncNoticeUrl   string `json:"asyn_notice_url"`
	Remark           string `json:"remark"` // 可选
}

type AAPayWithdrawResp struct {
	Code      int                   `json:"code"`
	Message   string                `json:"message"`
	Sign      string                `json:"sign"`
	Timestamp int64                 `json:"timestamp"`
	Data      AAPayWithdrawRespData `json:"data"`
}

type AAPayWithdrawRespData struct {
	WithdrawId        int     `json:"withdrawal_id"`
	UserWithdrawalId  string  `json:"user_withdrawal_id"`
	WithdrawalAddress string  `json:"withdrawal_address"`
	CoinMoney         Decimal `json:"coin_money"`
	Commission        Decimal `json:"commission"`
	CoinCode          string  `json:"coin_code"`
	CurrencyCode      string  `json:"currency_code"`
}

/*-------查询充值订单-------*/

type AAPayRechargeQueryReq struct {
	UserId  int64 `json:"user_id"`
	OrderId int64 `json:"order_id"`
}
type AAPayRechargeQueryResp struct {
	Code      int                        `json:"code"`
	Message   string                     `json:"message"`
	Sign      string                     `json:"sign"`
	Timestamp int64                      `json:"timestamp"`
	Data      AAPayRechargeQueryRespData `json:"data"`
}

type AAPayRechargeQueryRespData struct {
	OrderId              int    `json:"order_id"`
	OrderStatus          int    `json:"order_status"`
	CurrencyCode         string `json:"currency_code"`
	CoinCode             string `json:"coin_code"`
	CurrencyReceiptMoney string `json:"currency_receipt_money"`
	CoinAddress          string `json:"coin_address"`
	UserOrderId          string `json:"user_order_id"`
	CoinReceiptMoney     string `json:"coin_receipt_money"`
}

/*-------查询提现订单-------*/

type AAPayWithdrawQueryReq struct {
	UserId     int64 `json:"user_id"`
	WithdrawId int64 `json:"withdrawal_id"`
}
type AAPayWithdrawQueryResp struct {
	Code      int                        `json:"code"`
	Message   string                     `json:"message"`
	Sign      string                     `json:"sign"`
	Timestamp int64                      `json:"timestamp"`
	Data      AAPayWithdrawQueryRespData `json:"data"`
}

type AAPayWithdrawQueryRespData struct {
	WithdrawId      int     `json:"withdrawal_id"`
	UserWithdrawId  string  `json:"user_withdrawal_id"`
	WithdrawAddress string  `json:"withdrawal_address"`
	CoinCode        string  `json:"coin_code"`
	CoinMoney       Decimal `json:"coin_money"`
	Commission      Decimal `json:"commission"`
	TxId            string  `json:"tx_id"`
	CurrencyCode    string  `json:"currency_code"`
	CurrencyAmount  Decimal `json:"currency_amount"`
	WithdrawStatus  int     `json:"withdrawal_status"`
}

/*-------充值订单回调-------*/

type AAPayRechargeNotifyReq struct {
	Code      int                     `json:"code"`
	Message   string                  `json:"message"`
	Sign      string                  `json:"sign"`
	Timestamp int64                   `json:"timestamp"`
	Data      AAPayRechargeNotifyData `json:"data"`
}

type AAPayRechargeNotifyData struct {
	OrderId              int    `json:"order_id"`
	OrderStatus          int    `json:"order_status"`
	CurrencyCode         string `json:"currency_code"`
	CoinCode             string `json:"coin_code"`
	CoinAddress          string `json:"coin_address"`
	UserOrderId          string `json:"user_order_id"`
	CoinReceiptMoney     string `json:"coin_receipt_money"`
	CurrencyReceiptMoney string `json:"currency_receipt_money"`
}

/*-------提现订单回调-------*/

type AAPayWithdrawNotifyReq struct {
	Code      int                     `json:"code"`
	Message   string                  `json:"message"`
	Sign      string                  `json:"sign"`
	Timestamp int64                   `json:"timestamp"`
	Data      AAPayWithdrawNotifyData `json:"data"`
}

type AAPayWithdrawNotifyData struct {
	WithdrawId        int     `json:"withdrawal_id"`
	UserWithdrawalId  string  `json:"user_withdrawal_id"`
	WithdrawalAddress string  `json:"withdrawal_address"`
	CoinMoney         Decimal `json:"coin_money"`
	Commission        Decimal `json:"commission"`
	CoinCode          string  `json:"coin_code"`
	CurrencyCode      string  `json:"currency_code"`
}

func aaPayGenerateAccessSign(key string, timestamp string, userId int64, method string, router string, body interface{}) string {
	// 将结构体转换为 JSON 字符串
	var (
		jsonData, err = json.Marshal(body)
	)
	if err != nil {
		return ""
	}
	// 将字节数组转换为字符串
	var (
		jsonString = string(jsonData)
		usdIdStr   = fmt.Sprintf("%d", userId)
		message    = timestamp + usdIdStr + method + router + jsonString
		h          = hmac.New(sha256.New, []byte(key)) // 使用 HMAC-SHA256 生成签名
		signature  []byte
	)
	h.Write([]byte(message))
	signature = h.Sum(nil)
	return base64.URLEncoding.EncodeToString(signature)
}

func aaPaySendRechargeRequest(payUrl string, router string, body AAPayRechargeReq, userID int64, secret string, language string) (resp AAPayRechargeResp, err error) {
	var (
		succeedCode = 200
	)
	// 发送请求
	err = sendAAPayRequest(payUrl, router, body, userID, secret, language, &resp)
	if err != nil {
		err = errors.New(fmt.Sprintf("请求三方失败->%s", err.Error()))
		return
	}
	// 判断状态是否正常
	if resp.Code != succeedCode {
		err = errors.New(fmt.Sprintf("状态非成功：%d-> 原因：%s", resp.Code, resp.Message))
		return
	}
	return
}

func aaPaySendRechargeQueryRequest(payUrl string, router string, body AAPayRechargeQueryReq, userID int64, secret string, language string) (resp AAPayRechargeQueryResp, err error) {
	var (
		succeedCode = 200
	)
	// 发送请求
	err = sendAAPayRequest(payUrl, router, body, userID, secret, language, &resp)
	if err != nil {
		err = errors.New(fmt.Sprintf("请求三方失败->%s", err.Error()))
		return
	}
	// 判断状态是否正常
	if resp.Code != succeedCode {
		err = errors.New(fmt.Sprintf("状态非成功：%d-> 原因：%s", resp.Code, resp.Message))
		return
	}
	return
}

func aaPaySendWithdrawRequest(payUrl string, router string, body AAPayWithdrawReq, userID int64, secret string, language string) (resp AAPayWithdrawResp, err error) {
	var (
		succeedCode = 200
	)
	// 发送请求
	err = sendAAPayRequest(payUrl, router, body, userID, secret, language, &resp)
	if err != nil {
		err = errors.New(fmt.Sprintf("请求三方失败->%s", err.Error()))
		return
	}
	// 判断状态是否正常
	if resp.Code != succeedCode {
		err = errors.New(fmt.Sprintf("状态非成功：%d-> 原因：%s", resp.Code, resp.Message))
		return
	}
	return
}

func aaPaySendWithdrawQueryRequest(payUrl string, router string, body AAPayWithdrawQueryReq, userID int64, secret string, language string) (resp AAPayWithdrawQueryResp, err error) {
	var (
		succeedCode = 200
	)
	// 发送请求
	err = sendAAPayRequest(payUrl, router, body, userID, secret, language, &resp)
	if err != nil {
		err = errors.New(fmt.Sprintf("请求三方失败->%s", err.Error()))
		return
	}
	// 判断状态是否正常
	if resp.Code != succeedCode {
		err = errors.New(fmt.Sprintf("状态非成功：%d-> 原因：%s", resp.Code, resp.Message))
		return
	}
	return
}
func sendAAPayRequest(payUrl string, router string, body interface{}, userID int64, secret string, language string, result any) (err error) {

	var (
		timestamp  = fmt.Sprintf("%d", time.Now().Unix())
		accessSign = aaPayGenerateAccessSign(secret, timestamp, userID, "POST", router, body)
		bodyJson   []byte
		usdIdStr   = fmt.Sprintf("%d", userID)
		req        *http.Request
		resp       *http.Response
		client     *http.Client
		respBody   []byte
	)
	bodyJson, err = json.Marshal(body) // 将请求体序列化为 JSON
	if err != nil {
		err = errors.New(fmt.Sprintf("序列化请求体错误->%s", err.Error()))
		return
	}

	// 设置请求头
	req, err = http.NewRequest("POST", payUrl+router, strings.NewReader(string(bodyJson)))
	if err != nil {
		err = errors.New(fmt.Sprintf("组装请求错误->%s", err.Error()))
		return
	}
	req.Header.Set("Accept-Language", language)
	req.Header.Set("X-ACCESS-SIGN", accessSign)
	req.Header.Set("X-ACCESS-TIMESTAMP", timestamp)
	req.Header.Set("X-ACCESS-USERID", usdIdStr)

	// 发送请求
	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		err = errors.New(fmt.Sprintf("发送请求错误->%s", err.Error()))
		return
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	// 处理响应
	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		err = errors.New(fmt.Sprintf("读取响应结果错误->%s", err.Error()))
		return
	}

	// 解析 JSON 响应
	if err = json.Unmarshal(respBody, &result); err != nil {
		err = errors.New(fmt.Sprintf("解析响应结果错误->" + err.Error()))
		return
	}
	return
}
