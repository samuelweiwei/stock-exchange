package request

import (
	"github.com/guregu/null/v5"
	"github.com/shopspring/decimal"
)

// SystemConfigSaveReq 平台基础配置保存请求
// @Description 平台基础配置保存请求
type SystemConfigSaveReq struct {
	PlatformCommissionRate  null.Float          `json:"platformCommissionRate"`  //平台佣金比例
	FirstGradeShareRate     null.Float          `json:"firstGradeShareRate"`     //一级分成比例
	SecondGradeShareRate    null.Float          `json:"secondGradeShareRate"`    //二级分成比例
	ThirdGradeShareRate     null.Float          `json:"thirdGradeShareRate"`     //三级分成比例
	WithdrawCommissionType  null.Int            `json:"withdrawCommissionType"`  //提现手续费类型
	WithdrawCommissionQuota decimal.NullDecimal `json:"withdrawCommissionQuota"` //定额手续费额度
	WithdrawCommissionRate  decimal.NullDecimal `json:"withdrawCommissionRate"`  //手续费比例
	IosAppDownloadUrl       null.String         `json:"iosAppDownloadUrl"`       //ios下载链接
	AndroidAppDownloadUrl   null.String         `json:"androidAppDownloadUrl"`   //android下载链接
	BackWhiteIpStrings      null.String         `json:"backWhiteIpStrings"`      //后台登录ip白名单
}
