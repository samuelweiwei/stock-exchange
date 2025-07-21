// 自动生成模板EarnProducts
package earn

import (
	"github.com/shopspring/decimal"
	"time"
)

type Type int

const (
	Flexible = Type(0)
	Fixed    = Type(1)
)

type EarnProductsStatus int

const (
	Available = EarnProductsStatus(1)
	Frozen    = EarnProductsStatus(2)
)

// earnProducts表 结构体  EarnProducts
type EarnProducts struct {
	Id                   uint               `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:10;"`                                                  //id字段
	Wid                  int                `json:"wid" form:"wid" gorm:"column:wid;comment:the ID of the stock;size:10;"`                                       //the ID of the stock
	WidCode              string             `json:"widCode" form:"widCode" gorm:"column:wid_code;comment:the name of the stock;size:20;"`                        //the name of the stock
	Type                 Type               `json:"type" form:"type" gorm:"column:type;comment:0 Flexible 活期, 1 Fixed 定期;"`                                      //0 Flexible 活期, 1 Fixed 定期
	Name                 string             `json:"name" form:"name" gorm:"column:name;comment:产品名称;size:100;"`                                                  //产品名称
	CurrentInterestRates decimal.Decimal    `json:"currentInterestRates" form:"currentInterestRates" gorm:"column:current_interest_rates;comment:当前利率;size:20;"` //当前利率
	MinInterestRates     decimal.Decimal    `json:"minInterestRates" form:"minInterestRates" gorm:"column:min_interest_rates;comment:最小利率;size:20;"`             //最小利率
	MaxInterestRates     decimal.Decimal    `json:"maxInterestRates" form:"maxInterestRates" gorm:"column:max_interest_rates;comment:最大利率;size:20;"`             //最大利率
	PenaltyRatio         decimal.Decimal    `json:"penaltyRatio" form:"penaltyRatio" gorm:"column:penalty_ratio;comment:违约金比例;size:20;"`                         //违约金比例
	Mark                 string             `json:"mark" form:"mark" gorm:"column:mark;comment:product marks;size:1001;"`                                        //product marks
	Stock                int                `json:"stock" gorm:"column:stock;comment:the stock of the product -1 无限库存;size:10;"`                                 //the stock of the product -1 无限库存
	Duration             int                `json:"duration" form:"duration" gorm:"column:duration;comment:the days of the projects 理财定期时长;size:10;"`            //the days of the projects 理财定期时长
	Status               EarnProductsStatus `json:"status" form:"status" gorm:"column:status;comment:状态|||1:可用,2:冻结;"`                                           //状态|||1:可用,2:冻结
	CreatedAt            int64              `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:;size:10;"`                                       //createdAt字段
	UpdatedAt            time.Time          `json:"-" gorm:"column:updated_at;comment:;"`                                                                        //updatedAt字段
	AdminRemark          string             `json:"adminRemark" form:"adminRemark" gorm:"column:admin_remark;comment:后台管理员备注;size:200;"`                         //后台管理员备注
}

// TableName earnProducts表 EarnProducts自定义表名 earn_products
func (EarnProducts) TableName() string {
	return "earn_products"
}
