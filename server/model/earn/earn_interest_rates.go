// 自动生成模板EarnInterestRates
package earn

import (
	"github.com/shopspring/decimal"
	"time"
)

// earnInterestRates表 结构体  EarnInterestRates
type EarnInterestRates struct {
	Id            uint            `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:10;"`                          //id字段
	ProductId     uint            `json:"productId" form:"productId" gorm:"column:product_id;comment:理财产品ID;size:10;"`         //理财产品ID
	InterestRates decimal.Decimal `json:"interestRates" form:"interestRates" gorm:"column:interest_rates;comment:利率;size:20;"` //利率
	Period        time.Time       `json:"period" form:"period" gorm:"column:period;comment:;"`                                 //period字段
	CreatedAt     int64           `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:;size:10;"`               //createdAt字段
	UpdatedAt     time.Time       `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;comment:;"`                       //updatedAt字段
}

// TableName earnInterestRates表 EarnInterestRates自定义表名 earn_interest_rates
func (EarnInterestRates) TableName() string {
	return "earn_interest_rates"
}
