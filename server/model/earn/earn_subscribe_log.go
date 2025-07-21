// 自动生成模板EarnSubscribeLog
package earn

import "github.com/shopspring/decimal"

type SubscribeStatus int

const (
	Staking         = SubscribeStatus(1)
	Redeemed        = SubscribeStatus(2)
	RedeemPending   = 0
	RedeemInAdvance = 1
	RedeemNormal    = 2
)

// earnSubscribeLog表 结构体  EarnSubscribeLog
type EarnSubscribeLog struct {
	Id              uint            `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:10;"`                                    //id字段
	Uid             uint            `json:"uid" form:"uid" gorm:"column:uid;comment:用户id;size:10;"`                                        //用户id
	ProductId       uint            `json:"productId" form:"productId" gorm:"column:product_id;comment:理财产品ID;size:10;"`                   //理财产品ID
	PenaltyRatio    decimal.Decimal `json:"penaltyRatio" form:"penaltyRatio" gorm:"column:penalty_ratio;comment:违约金比例;size:20;"`           //违约金比例
	UserType        uint            `json:"userType" form:"userType" gorm:"column:user_type;comment:用户类型 1:普通用户 2: 试玩用户;"`                 //用户类型 1:普通用户 2: 试玩用户
	Status          SubscribeStatus `json:"status" form:"status" gorm:"column:status;comment:状态|||1: 质押,2:赎回;"`                            //状态|||1: 质押,2:赎回
	RedeemInAdvance int             `json:"redeemInAdvance" form:"redeemInAdvance" gorm:"column:redeem_in_advance;comment:1提前赎回, 2未提前赎回;"` //1提前赎回, 2未提前赎回
	Fine            decimal.Decimal `json:"fine" form:"fine" gorm:"column:fine;comment:罚金;size:20;"`                                       //罚金
	StartAt         int64           `json:"startAt" form:"startAt" gorm:"column:start_at;comment:活/定期产品开始时间;size:10;"`                     //活/定期产品开始时间
	EndAt           int64           `json:"endAt" form:"endAt" gorm:"column:end_at;comment:活/定期产品到期时间;size:10;"`                           //活/定期产品到期时间
	BoughtNum       decimal.Decimal `json:"boughtNum" form:"boughtNum" gorm:"column:bought_num;comment:买入数量;size:20;"`                     //买入数量
	CreatedAt       int64           `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:;"`                                 //createdAt字段
	UpdatedAtX      int64           `json:"updatedAt" form:"updatedAt" gorm:"column:updated_atx" `                                         //updatedAtX字段
}

// TableName earnSubscribeLog表 EarnSubscribeLog自定义表名 earn_subscribe_log
func (EarnSubscribeLog) TableName() string {
	return "earn_subscribe_log"
}
