package response

// MyUserProfitShareRecord 我的分成记录
//
// @Description 我的分成记录
type MyUserProfitShareRecord struct {
	FromUserName string  `json:"fromUserName"` //来源用户名
	Amount       float64 `json:"amount"`       //分成金额
	Date         int64   `json:"date"`         //日期
}
