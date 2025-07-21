/**
* @Author: Jackey
* @Date: 12/31/24 8:53 pm
 */

package constants

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"strconv"
)

type ThirdPaymentRouter struct {
	RechargeCreat string
	RechargeQuery string
	WithdrawCreat string
	WithdrawQuery string
}

var (
	AAPayRouter = ThirdPaymentRouter{
		RechargeCreat: "/api/v1/order/create",
		RechargeQuery: "/api/v1/order/find",
		WithdrawCreat: "/api/v1/withdrawal/create",
		WithdrawQuery: "/api/v1/withdrawal/find",
	}
)

const (
	ThirdPaymentIdForAAPay int = 101
)

const (
	ThirdPaymentNameForAAPay = "AAPay"
)

func GetThirdPaymentNameById(paymentId int) string {
	switch paymentId {
	case ThirdPaymentIdForAAPay:
		return ThirdPaymentNameForAAPay
	}
	return ""
}

func IsThirdRechargeSucceed(paymentId int, status string) bool {

	switch paymentId {
	case ThirdPaymentIdForAAPay:
		return status == "2"
	}
	return false
}

func IsThirdWithdrawSucceed(paymentId int, status string) bool {

	switch paymentId {
	case ThirdPaymentIdForAAPay:
		return status == "3"
	}
	return false
}

func IsThirdWithdrawFailed(paymentId int, status string) bool {

	switch paymentId {
	case ThirdPaymentIdForAAPay:
		return status == "4" || status == "5"
	}
	return false
}

func GetThirdPaymentInfoById(paymentId int) (uri, userId, secretKey string) {
	switch paymentId {
	case ThirdPaymentIdForAAPay:
		uri = global.GVA_CONFIG.ThirdPay.AAPayUri
		userId = global.GVA_CONFIG.ThirdPay.AAPayUserId
		secretKey = global.GVA_CONFIG.ThirdPay.AAPaySecretKey
	}
	return
}

func GetThirdPaymentRechargeNotifyUrlById(paymentId int) (notifyUrl string) {
	notifyUrl = global.GVA_CONFIG.ThirdPay.NotifyUri + "/api/rechargeRecords/payNotify/" + strconv.Itoa(paymentId)
	return
}

func GetThirdPaymentWithdrawNotifyUrlById(paymentId int) (notifyUrl string) {
	notifyUrl = global.GVA_CONFIG.ThirdPay.NotifyUri + "/api/withdrawRecords/withdrawNotify/" + strconv.Itoa(paymentId)
	return
}
