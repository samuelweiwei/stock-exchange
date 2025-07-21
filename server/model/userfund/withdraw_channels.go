// 自动生成模板WithdrawChannels
package userfund

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
)

// withdrawChannels表 结构体  WithdrawChannels
type WithdrawChannels struct {
	global.GVA_MODEL
	CoinId                 *int   `json:"coinId" form:"coinId" gorm:"column:coin_id;comment:关联的货币ID;"`                                                            //关联的货币ID
	ThirdRechargeChannel   string `json:"thirdRechargeChannel" form:"thirdRechargeChannel" gorm:"column:third_recharge_channel;comment:第三方厂商名称：aa pay;size:200;"` //第三方厂商名称：aa pay
	ThirdCurrencyCode      string `json:"thirdCurrencyCode" form:"thirdCurrencyCode" gorm:"column:third_currency_code;comment:;size:20;"`                         //thirdCurrencyCode字段
	ThirdCoinCode          string `json:"thirdCoinCode" form:"thirdCoinCode" gorm:"column:third_coin_code;comment:;size:20;"`                                     //thirdCoinCode字段
	RequireExchangeRate    *bool  `json:"requireExchangeRate" form:"requireExchangeRate" gorm:"column:require_exchange_rate;comment:是否需要查询第三方汇率;"`                //是否需要查询第三方汇率
	RechargeType           string `json:"rechargeType" form:"rechargeType" gorm:"column:recharge_type;comment:提现方式：系统充值，快捷充值;size:191;"`                          //提现方式：系统充值，快捷充值
	ChannelType            string `json:"channelType" form:"channelType" gorm:"column:channel_type;comment:提现渠道1 地址  2 银行卡 3 其他;size:255;"`                       //提现渠道1 地址  2 银行卡 3 其他
	Channel                string `json:"channel" form:"channel" gorm:"column:channel;comment:渠道：TRC20,ERC20 或者第三方名称;size:255;"`                                  //渠道：TRC20,ERC20 或者第三方名称
	Address                string `json:"address" form:"address" gorm:"column:address;comment:对应渠道的地址;size:255;"`                                                 //对应渠道的地址
	SortOrder              *int   `json:"sortOrder" form:"sortOrder" gorm:"column:sort_order;comment:排序权重，默认为0;"`                                                 //排序权重，默认为0
	STATUS                 string `json:"STATUS" form:"STATUS" gorm:"column:STATUS;comment:状态，开启或关闭;size:1;"`                                                     //状态，开启或关闭
	CreatedBy              *int   `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者;size:19;"`                                               //创建者
	UpdatedBy              *int   `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者;size:19;"`                                               //更新者
	DeletedBy              *int   `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:删除者;size:19;"`
	NoticeUrl              string `json:"noticeUrl" form:"noticeUrl" gorm:"column:notice_url;comment:更新者;size:19;"` //第三方接口通知地址
	ThirdPayId             int    `json:"thirdPayId" form:"thirdPayId" gorm:"column:third_pay_id;comment:删除者;size:19;"`
	NotifyUrl              string `json:"notifyUrl" form:"notifyUrl" gorm:"column:notify_url;comment:回调;size:19;"`
	WithdrawRouter         string `json:"withdrawRouter" form:"withdrawRouter" gorm:"column:withdraw_router;comment:回调;size:19;"`
	WithdrawCallbackRouter string `json:"withdrawCallbackRouter" form:"withdrawCallbackRouter" gorm:"column:withdraw_callback_router;comment:回调;size:19;"`
	QueryOrderRouter       string `json:"queryOrderRouter" form:"queryOrderRouter" gorm:"column:query_order_router;comment:回调;size:19;"`

	// 关联字段，不存入数据库
	Currency       string          `json:"currency" gorm:"-"`
	CurrencyIcon   string          `json:"currencyIcon" gorm:"-"`
	PriceUsdt      decimal.Decimal `json:"priceUsdt" gorm:"-"`
	CoinType       *int            `json:"coinType" gorm:"-"`
	TicketSize     decimal.Decimal `json:"ticketSize" gorm:"-"`     // 价格进度
	TicketNumSize  decimal.Decimal `json:"ticketNumSize" gorm:"-"`  // 数量精度
	MinWithdrawNum decimal.Decimal `json:"minWithdrawNum" gorm:"-"` // 可提现最小数量

	CreatedAtInt int64 `json:"createTimeInt" gorm:"-"`
	UpdatedAtInt int64 `json:"updateTimeInt" gorm:"-"`
}

// TableName withdrawChannels表 WithdrawChannels自定义表名 withdraw_channels
func (WithdrawChannels) TableName() string {
	return "withdraw_channels"
}
