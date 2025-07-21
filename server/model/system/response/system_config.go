package response

import "github.com/shopspring/decimal"

type SystemConfig struct {
	BackWhiteIpStrings string `json:"backWhiteIpStrings"`
	DomainInfo
}

type DomainInfo struct {
	PlatformCommissionRate  float64         `json:"platformCommissionRate"`
	FirstGradeShareRate     float64         `json:"firstGradeShareRate"`
	SecondGradeShareRate    float64         `json:"secondGradeShareRate"`
	ThirdGradeShareRate     float64         `json:"thirdGradeShareRate"`
	WithdrawCommissionType  int64           `json:"withdrawCommissionType"`
	WithdrawCommissionQuota decimal.Decimal `json:"withdrawCommissionQuota"`
	WithdrawCommissionRate  decimal.Decimal `json:"withdrawCommissionRate"`
	IosAppDownloadUrl       string          `json:"iosAppDownloadUrl"`
	AndroidAppDownloadUrl   string          `json:"androidAppDownloadUrl"`
}
