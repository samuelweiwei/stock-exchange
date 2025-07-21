// 自动生成模板EarnDailyIncomeMoneyLog
package earn

import (
	"github.com/shopspring/decimal"
	"time"
)

// earnDailyIncomeMoneyLog表 结构体  EarnDailyIncomeMoneyLog
type EarnDailyIncomeMoneyLog struct {
	Id            uint            `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:10;"`                            //id字段
	Uid           uint            `json:"uid" form:"uid" gorm:"column:uid;comment:用户id;size:10;"`                                //用户id
	ProductId     uint            `json:"productId" form:"productId" gorm:"column:product_id;comment:理财产品ID;size:10;"`           //理财产品ID
	SubscribeId   uint            `json:"subscribeId" form:"subscribeId" gorm:"column:subscribe_id;comment:下单ID;size:10;"`       //下单ID
	Earnings      decimal.Decimal `json:"earnings" form:"earnings" gorm:"column:earnings;comment:当前收益;size:20;"`                 //当前收益
	InterestRates decimal.Decimal `json:"interestRates" form:"interestRates" gorm:"column:interest_rates;comment:当天利率;size:20;"` //当天利率
	BoughtNum     decimal.Decimal `json:"boughtNum" form:"boughtNum" gorm:"column:bought_num;comment:买入数量;size:20;"`             //买入数量
	UserType      uint            `json:"userType" form:"userType" gorm:"column:user_type;comment:用户类型 1:普通用户 2: 试玩用户;"`         //用户类型 1:普通用户 2: 试玩用户;
	OfferedAt     int64           `json:"offeredAt" form:"offeredAt" gorm:"column:offered_at;comment:到帐时间;size:10;"`             //到帐时间
	CreatedAt     int64           `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:;size:10;"`                 //createdAt字段
	UpdatedAt     *time.Time      `json:"updatedAt" form:"updatedAt" gorm:"autoUpdateTime:milli column:updated_at;comment:;"`    //updatedAt字段
}

// TableName earnDailyIncomeMoneyLog表 EarnDailyIncomeMoneyLog自定义表名 earn_daily_income_money_log
func (EarnDailyIncomeMoneyLog) TableName() string {
	return "earn_daily_income_money_log"
}
