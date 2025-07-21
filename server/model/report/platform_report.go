package report

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
	"time"
)

type PlatformReportSummary struct {
	global.GVA_MODEL
	RootUserId               uint            `gorm:"column:root_user_id"`
	TotalBalance             decimal.Decimal `gorm:"column:total_balance"`
	TotalRecharge            decimal.Decimal `gorm:"column:total_recharge"`
	TotalWithdraw            decimal.Decimal `gorm:"column:total_withdraw"`
	TotalActualWithdraw      decimal.Decimal `gorm:"column:total_actual_withdraw"`
	TotalUserCount           int64           `gorm:"column:total_user_count"`
	TotalFollowUserCount     int64           `gorm:"column:total_follow_user_count"`
	PlatformProfit           decimal.Decimal `gorm:"column:platform_profit"`
	TodayActiveUserCount     int64           `gorm:"column:today_active_user_count"`
	YesterdayActiveUserCount int64           `gorm:"column:yesterday_active_user_count"`
	InviteUserCount          int64           `gorm:"column:invite_user_count"`
	EffectiveInviteUserCount int64           `gorm:"column:effective_invite_user_count"`
}

func (PlatformReportSummary) TableName() string {
	return "platform_report_summary"
}

type PlatformReportDaily struct {
	global.GVA_MODEL
	RootUserId             uint            `gorm:"column:root_user_id"`
	TotalRecharge          decimal.Decimal `gorm:"column:total_recharge"`
	RechargeUserCount      int64           `gorm:"column:recharge_user_count"`
	TotalWithdraw          decimal.Decimal `gorm:"column:total_withdraw"`
	TotalActualWithdraw    decimal.Decimal `gorm:"column:total_actual_withdraw"`
	WithdrawUserCount      int64           `gorm:"column:withdraw_user_count"`
	TotalNewUserCount      int64           `gorm:"column:total_new_user_count"`
	ActiveUserCount        int64           `gorm:"column:active_user_count"`
	InviteUserCount        int64           `gorm:"column:invite_user_count"`
	FirstRechargeUserCount int64           `gorm:"column:first_recharge_user_count"`
	FirstRechargeAmount    decimal.Decimal `gorm:"column:first_recharge_amount"`
	FollowProfit           decimal.Decimal `gorm:"column:follow_profit"`
	FollowAmount           decimal.Decimal `gorm:"column:follow_amount"`
	FollowUserCount        int64           `gorm:"column:follow_user_count"`
	ReportDate             time.Time       `gorm:"column:report_date"`
}

func (PlatformReportDaily) TableName() string {
	return "platform_report_daily"
}
