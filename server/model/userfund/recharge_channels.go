// 自动生成模板RechargeChannels
package userfund

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	. "github.com/shopspring/decimal"
)

// rechargeChannels表 结构体  RechargeChannels
type RechargeChannels struct {
	global.GVA_MODEL
	CoinId                 *int   `json:"coinId" form:"coinId" gorm:"column:coin_id;comment:关联的货币ID;"`                                                            //关联的货币ID
	ThirdRechargeChannel   string `json:"thirdRechargeChannel" form:"thirdRechargeChannel" gorm:"column:third_recharge_channel;comment:第三方厂商名称：aa pay;size:200;"` //第三方厂商名称：aa pay
	ThirdCurrencyCode      string `json:"thirdCurrencyCode" form:"thirdCurrencyCode" gorm:"column:third_currency_code;comment:;size:20;"`                         //thirdCurrencyCode字段
	ThirdCoinCode          string `json:"thirdCoinCode" form:"thirdCoinCode" gorm:"column:third_coin_code;comment:;size:20;"`                                     //thirdCoinCode字段
	RequireExchangeRate    *bool  `json:"requireExchangeRate" form:"requireExchangeRate" gorm:"column:require_exchange_rate;comment:是否需要查询第三方汇率;"`                //是否需要查询第三方汇率
	RechargeType           string `json:"rechargeType" form:"rechargeType" gorm:"column:recharge_type;comment:充值方式：系统充值，快捷充值;size:191;"`                          //充值方式：系统充值，快捷充值
	ChannelType            string `json:"channelType" form:"channelType" gorm:"column:channel_type;comment:充值渠道1 地址 2 银行卡 3 其他;size:255;"`                        //充值方式：系统充值，快捷充值
	Channel                string `json:"channel" form:"channel" gorm:"column:channel;comment:渠道ERC20,TRC20 其他等等;size:255;"`                                      //渠道1 地址  2 银行卡 3 其他
	Address                string `json:"address" form:"address" gorm:"column:address;comment:对应渠道的地址;size:255;"`                                                 //对应渠道的地址
	SortOrder              *int   `json:"sortOrder" form:"sortOrder" gorm:"column:sort_order;comment:排序权重，默认为0;size:19;"`                                         //排序权重，默认为0
	STATUS                 string `json:"STATUS" form:"STATUS" gorm:"column:STATUS;comment:状态，开启或关闭;size:1;"`                                                     //状态，开启或关闭
	CreatedBy              *int   `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者;"`                                                       //创建者
	UpdatedBy              *int   `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者;"`                                                       //更新者
	DeletedBy              *int   `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:删除者;"`                                                       //删除者
	NoticeUrl              string `json:"noticeUrl" form:"noticeUrl" gorm:"column:notice_url;comment:通知;size:19;"`                                                //第三方接口通知地址
	ThirdPayId             int    `json:"thirdPayId" form:"thirdPayId" gorm:"column:third_pay_id;comment:第三方ID;size:19;"`                                         //第三方支付ID
	NotifyUrl              string `json:"notifyUrl" form:"notifyUrl" gorm:"column:notify_url;comment:回调;size:19;"`
	RechargeRouter         string `json:"rechargeRouter" form:"rechargeRouter" gorm:"column:recharge_router;comment:回调;size:19;"`
	RechargeCallbackRouter string `json:"rechargeCallbackRouter" form:"rechargeCallbackRouter" gorm:"column:recharge_callback_router;comment:回调;size:19;"`
	QueryOrderRouter       string `json:"queryOrderRouter" form:"queryOrderRouter" gorm:"column:query_order_router;comment:回调;size:19;"`

	//第三方回调地址
	// 关联字段，不存入数据库
	Currency       string  `json:"currency" gorm:"-"`
	CurrencyIcon   string  `json:"currencyIcon" gorm:"-"`
	PriceUsdt      Decimal `json:"priceUsdt" gorm:"-"`
	CoinType       *int    `json:"coinType" gorm:"-"`
	TicketSize     Decimal `json:"ticketSize" gorm:"-"`     // 价格进度
	TicketNumSize  Decimal `json:"ticketNumSize" gorm:"-"`  // 数量精度
	MinRechargeNum Decimal `json:"minRechargeNum" gorm:"-"` // 可充值最小数量

	CreateAtInt int64 `json:"createTimeInt" gorm:"-"`
	UpdateAtInt int64 `json:"updateTimeInt" gorm:"-"`
}

// TableName rechargeChannels表 RechargeChannels自定义表名 recharge_channels
func (RechargeChannels) TableName() string {
	return "recharge_channels"
}
