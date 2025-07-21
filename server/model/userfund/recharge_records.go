// 自动生成模板RechargeRecords
package userfund

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	. "github.com/shopspring/decimal"
	"time"
)

// rechargeRecords表 结构体  RechargeRecords
type RechargeRecords struct {
	global.GVA_MODEL
	OrderId      string `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单ID，唯一标识该充值记录;size:19;"`          //订单ID，唯一标识该充值记录
	ThirdOrderId string `json:"thirdOrderId" form:"thirdOrderId" gorm:"column:third_order_id;comment:第三方订单ID;size:19;"` //第三方订单ID
	MemberId     int    `json:"memberId" form:"memberId" gorm:"column:member_id;comment:会员ID，标识充值会员;size:20;"`          //会员ID，标识充值会员
	ParentId     int    `json:"parentUserId" form:"parentUserId" gorm:"column:parent_id;comment:父级用户ID;"`               //会员ID，标识提现会员
	RootId       int    `json:"rootUserId" form:"rootUserId" gorm:"column:root_id;comment:root_id;"`                    //根用户ID

	MemberPhone       string  `json:"memberPhone" form:"memberPhone" gorm:"column:member_phone;comment:会员手机号;size:20;"`                     //会员手机号
	MemberEmail       string  `json:"memberEmail" form:"memberEmail" gorm:"column:member_email;comment:会员邮箱;size:30;"`                      //会员邮箱
	RechargeChannelId uint    `json:"rechargeChannelId" form:"rechargeChannelId" gorm:"column:recharge_channel_id;comment:充值渠道ID;size:10;"` //充值渠道ID
	Currency          string  `json:"currency" form:"currency" gorm:"column:currency;comment:充值货币类型，如BTC;size:10;"`                         //充值货币类型，如BTC
	RechargeAmount    Decimal `json:"rechargeAmount" form:"rechargeAmount" gorm:"column:recharge_amount;comment:充值的货币数量;size:18;"`          //充值的货币数量

	RechargeRate  Decimal `json:"rechargeRate" form:"rechargeRate" gorm:"column:recharge_rate;comment:充值的货币汇率;size:18;"`   //充值的货币数量
	NotifyContent string  `json:"notifyContent" form:"notifyContent" gorm:"column:notify_content;comment:回调内容;size:2550;"` //回调内容
	NoticeContent string  `json:"noticeContent" form:"noticeContent" gorm:"column:notice_content;comment:通知内容;size:2550;"` //回调内容

	ExchangedAmountUsdt Decimal `json:"exchangedAmountUsdt" form:"exchangedAmountUsdt" gorm:"column:exchanged_amount_usdt;comment:折算为USDT的金额;size:18;"` //折算为USDT的金额

	CoinReceiptMoney string `json:"coinReceiptMoney" form:"coinReceiptMoney" gorm:"column:coin_receipt_money;comment:实际支付金额（第三方）;size:18;"` //第三方实际支付金额

	FromAddress   string     `json:"fromAddress" form:"fromAddress" gorm:"column:from_address;comment:充值来源地址，用户的充值地址;size:255;"`      // 原 UserAddress
	ToAddress     string     `json:"toAddress" form:"toAddress" gorm:"column:to_address;comment:充值目标地址，平台接收地址;size:255;"`             // 原 TargetAddress
	RechargeType  string     `json:"rechargeType" form:"rechargeType" gorm:"column:recharge_type;comment:充值方式：系统充值，快捷充值;size:191;"`   //充值方式：系统充值，快捷充值
	ChannelType   string     `json:"channelType" form:"channelType" gorm:"column:channel_type;comment:充值渠道1 地址 2 银行卡 3 其他;size:255;"` //充值方式：系统充值，快捷充值
	Channel       string     `json:"channel" form:"channel" gorm:"column:channel;comment:渠道ERC20,TRC20 其他等等;size:255;"`               //渠道类型，如银行、第三方、钱包
	RefusedReason string     `json:"refusedReason" form:"refusedReason" gorm:"column:refused_reason;comment:拒绝原因"`                    //渠道类型，如银行、第三方、钱包
	IsLock        string     `json:"isLock" form:"isLock" gorm:"column:is_lock;comment:是否锁定 1 是  0 否;size:1;"`                        //渠道类型，如银行、第三方、钱包
	Locker        int        `json:"locker" form:"locker" gorm:"column:locker;comment:锁定权限归属;size:20;"`
	OrderStatus   string     `json:"orderStatus" form:"orderStatus" gorm:"column:order_status;comment:订单状态，待审核、已确认、已完成;"`       //订单状态，待审核、已确认、已完成
	RechargeTime  *time.Time `json:"rechargeTime" form:"rechargeTime" gorm:"column:recharge_time;comment:用户充值的时间;"`             //用户充值的时间
	ApprovalTime  *time.Time `json:"approvalTime" form:"approvalTime" gorm:"column:approval_time;comment:审核通过的时间;"`             //审核通过的时间
	UserAction    string     `json:"userAction" form:"userAction" gorm:"column:user_action;comment:用户操作，提交申请或撤销;"`              //用户操作，提交申请或撤销
	ReviewStatus  string     `json:"reviewStatus" form:"reviewStatus" gorm:"column:review_status;comment:审核状态，1审核中 2 拒绝 3 通过;"` //审核状态，锁定或解锁
	UserType      uint       `json:"userType" form:"userType" gorm:"column:user_type;comment:用户类型：1-普通用户，2-试玩用户;"`              //用户类型：1-普通用户，2-试玩用户

	HashAuth        bool  `json:"hashAuth" gorm:"-"`
	RechargeTimeInt int64 `json:"rechargeTimeInt" gorm:"-"`
	ApprovalTimeInt int64 `json:"approvalTimeInt" gorm:"-"`
	CreatedAtInt    int64 `json:"createTimeInt" gorm:"-"`
	UpdatedAtInt    int64 `json:"updateTimeInt" gorm:"-"`
}

// TableName rechargeRecords表 RechargeRecords自定义表名 recharge_records
func (RechargeRecords) TableName() string {
	return "recharge_records"
}
