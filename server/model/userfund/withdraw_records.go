// 自动生成模板WithdrawRecords
package userfund

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	. "github.com/shopspring/decimal"
	"time"
)

// withdrawRecords表 结构体  WithdrawRecords
type WithdrawRecords struct {
	global.GVA_MODEL
	OrderId      string `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单ID，唯一标识该提现记录;size:19;"`  //订单ID，唯一标识该提现记录
	ThirdOrderId string `json:"thirdOrderId" form:"thirdOrderId" gorm:"column:third_order_id;comment:第三方订单ID;"` //第三方订单ID
	MemberId     int    `json:"memberId" form:"memberId" gorm:"column:member_id;comment:会员ID，标识提现会员;"`          //会员ID，标识提现会员

	ParentId int `json:"parentUserId" form:"parentUserId" gorm:"column:parent_id;comment:父级用户ID;"` //会员ID，标识提现会员
	RootId   int `json:"rootUserId" form:"rootUserId" gorm:"column:root_id;comment:root_id;"`      //根用户ID

	MemberPhone           string     `json:"memberPhone" form:"memberPhone" gorm:"column:member_phone;comment:会员手机号;size:20;"`                              //会员手机号
	MemberEmail           string     `json:"memberEmail" form:"memberEmail" gorm:"column:member_email;comment:会员邮箱;size:30;"`                               //会员邮箱
	WithdrawChannelId     int        `json:"withdrawChannelId" form:"withdrawChannelId" gorm:"column:withdraw_channel_id;comment:提现渠道ID;"`                  //提现渠道ID
	Currency              string     `json:"currency" form:"currency" gorm:"column:currency;comment:提现货币类型，如BTC;size:10;"`                                  //提现货币类型，如BTC
	WithdrawAmount        Decimal    `json:"withdrawAmount" form:"withdrawAmount" gorm:"column:withdraw_amount;comment:从平台转出的USDT金额;"`                      //从平台转出的USDT金额
	RealReceived          Decimal    `json:"realReceived" form:"realReceived" gorm:"column:real_received;comment:实际收到的金额（回调）;"`                             //从平台转出的USDT金额
	ExchangedAmountTarget Decimal    `json:"exchangedAmountTarget" form:"exchangedAmountTarget" gorm:"column:exchanged_amount_target;comment:用户收到的目标货币金额;"` //用户收到的目标货币金额
	Commission            Decimal    `json:"commission" form:"commission" gorm:"column:commission;comment:手续费;"`                                            //手续费
	ThirdCommission       Decimal    `json:"thirdCommission" form:"thirdCommission" gorm:"column:third_commission;comment:第三方手续费;"`                         //手续费
	ChannelType           string     `json:"channelType" form:"channelType" gorm:"column:channel_type;comment:提现渠道1 地址 2 银行卡 3 其他;size:255;"`               //提现渠道1 地址 2 银行卡 3 其他
	WithdrawType          string     `json:"withdrawType" form:"withdrawType" gorm:"column:withdraw_type;comment:提现方式：系统提现，快捷提现;size:191;"`                 //提现方式：系统提现，快捷提现
	Channel               string     `json:"channel" form:"channel" gorm:"column:channel;comment:渠道ERC20,TRC20 其他等等;size:255;"`                             //渠道ERC20,TRC20 其他等等
	FromAddress           string     `json:"fromAddress" form:"fromAddress" gorm:"column:from_address;comment:平台出币地址，提现发起的平台地址;size:255;"`                  //平台出币地址，提现发起的平台地址
	ToAddress             string     `json:"toAddress" form:"toAddress" gorm:"column:to_address;comment:用户提现地址，用户接收提现的目标地址;size:255;"`                      //用户提现地址，用户接收提现的目标地址
	OrderStatus           string     `json:"orderStatus" form:"orderStatus" gorm:"column:order_status;comment:订单状态，待审核、已确认、已完成;size:191;"`                  //订单状态，待审核、已确认、已完成
	WithdrawTime          *time.Time `json:"withdrawTime" form:"withdrawTime" gorm:"column:withdraw_time;comment:用户提现的时间;"`                                 //用户提现的时间
	ApprovalTime          *time.Time `json:"approvalTime" form:"approvalTime" gorm:"column:approval_time;comment:审核通过的时间;"`                                 //审核通过的时间
	UserAction            string     `json:"userAction" form:"userAction" gorm:"column:user_action;comment:用户操作，提交申请或撤销;size:191;"`                         //用户操作，提交申请或撤销
	ReviewStatus          string     `json:"reviewStatus" form:"reviewStatus" gorm:"column:review_status;comment:审核状态，锁定或解锁;size:191;"`                     //审核状态，锁定或解锁
	IsLock                string     `json:"isLock" form:"isLock" gorm:"column:is_lock;comment:1 锁定 0 未锁定;"`                                                //1 锁定 0 未锁定
	Locker                int        `json:"locker" form:"locker" gorm:"column:locker;comment:锁定权限归属;size:20;"`
	WithdrawRate          Decimal    `json:"withdrawRate" form:"withdrawRate" gorm:"column:withdraw_rate;comment:提现的货币汇率;size:18;"`   //充值的货币数量
	NotifyContent         string     `json:"notifyContent" form:"notifyContent" gorm:"column:notify_content;comment:回调内容;size:2550;"` //回调内容
	NoticeContent         string     `json:"noticeContent" form:"noticeContent" gorm:"column:notice_content;comment:通知内容;size:2550;"` //回调内容

	RefusedReason string `json:"refusedReason" form:"refusedReason" gorm:"column:refused_reason;comment:拒绝原因;"` //拒绝原因
	UserType      uint   `json:"userType" form:"userType" gorm:"column:user_type;comment:用户类型：1-普通用户，2-试玩用户;"`  //用户类型：1-普通用户，2-试玩用户

	HashAuth        bool  `json:"hashAuth" gorm:"-"`
	WithdrawTimeInt int64 `json:"withdrawTimeInt" gorm:"-"`
	ApprovalTimeInt int64 `json:"approvalTimeInt" gorm:"-"`
	CreatedAtInt    int64 `json:"createTimeInt" gorm:"-"`
	UpdatedAtInt    int64 `json:"updateTimeInt" gorm:"-"`
}

// TableName withdrawRecords表 WithdrawRecords自定义表名 withdraw_records
func (WithdrawRecords) TableName() string {
	return "withdraw_records"
}
