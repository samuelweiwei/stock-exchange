package system

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
)

type Config struct {
	PlatformCommissionRate  float64         `json:"platformCommissionRate"`
	FirstGradeShareRate     float64         `json:"firstGradeShareRate"`
	SecondGradeShareRate    float64         `json:"secondGradeShareRate"`
	ThirdGradeShareRate     float64         `json:"thirdGradeShareRate"`
	WithdrawCommissionType  int64           `json:"withdrawCommissionType"`
	WithdrawCommissionQuota decimal.Decimal `json:"withdrawCommissionQuota"`
	WithdrawCommissionRate  decimal.Decimal `json:"withdrawCommissionRate"`
	IosAppDownloadUrl       string          `json:"iosAppDownloadUrl"`
	AndroidAppDownloadUrl   string          `json:"androidAppDownloadUrl"`
	BackWhiteIpStrings      string          `json:"backWhiteIpStrings"`
}

type SystemConfig struct {
	global.GVA_MODEL
	Config Config `gorm:"column:config;TYPE:json"`
}

func (SystemConfig) TableName() string {
	return "system_config"
}

func (c *Config) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal Config value:", value))
	}

	var config = Config{}
	err := json.Unmarshal(bytes, &config)
	if err != nil {
		return err
	}

	*c = config
	return nil
}

func (c Config) Value() (driver.Value, error) {
	return json.Marshal(c)
}
