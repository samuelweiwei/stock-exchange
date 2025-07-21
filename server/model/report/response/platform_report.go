package response

// PlatformReportSummary 平台固定报表
// @Description 平台固定报表
type PlatformReportSummary struct {
	TotalBalance             float64 `json:"totalBalance"`             //盘内结余
	TotalRecharge            float64 `json:"totalRecharge"`            //历史总充值
	TotalWithdraw            float64 `json:"totalWithdraw"`            //历史总提现
	TotalActualWithdraw      float64 `json:"totalActualWithdraw"`      //历史实际总提现
	TotalUserCount           int64   `json:"totalUserCount"`           //用户总数
	TotalFollowUserCount     int64   `json:"totalFollowUserCount"`     //跟单用户总数
	PlatformProfit           float64 `json:"platformProfit"`           //平台总盈亏
	TodayActiveUserCount     int64   `json:"todayActiveUserCount"`     //今日活跃用户
	YesterdayActiveUserCount int64   `json:"yesterdayActiveUserCount"` //昨日活跃用户
	InviteUserCount          int64   `json:"inviteUserCount"`          //邀请用户数
	EffectiveInviteUserCount int64   `json:"effectiveInviteUserCount"` //有效邀请用户数
}

// PlatformReportDaily 平台每日报表
// @Description 平台每日报表
type PlatformReportDaily struct {
	TotalRecharge          float64 `json:"totalRecharge"`          //充值总额
	RechargeUserCount      int64   `json:"rechargeUserCount"`      //充值任务
	TotalWithdraw          float64 `json:"totalWithdraw"`          //提现总额
	TotalActualWithdraw    float64 `json:"totalActualWithdraw"`    //实际提现总额
	WithdrawUserCount      int64   `json:"withdrawUserCount"`      //提现人数
	TotalNewUserCount      int64   `json:"totalNewUserCount"`      //注册人数
	ActiveUserCount        int64   `json:"activeUserCount"`        //活跃人数
	InvitedUserCount       int64   `json:"invitedUserCount"`       //邀请用户人数
	FirstRechargeUserCount int64   `json:"firstRechargeUserCount"` //首冲人数
	FirstRechargeAmount    float64 `json:"firstRechargeAmount"`    //首冲总额
	FollowProfit           float64 `json:"followProfit"`           //跟单总盈亏
	FollowAmount           float64 `json:"followAmount"`           //跟单总金额
	FollowUserCount        int64   `json:"followUserCount"`        //跟单总人数
	ReportDate             int64   `json:"reportDate"`             //报告日期
	RootUserId             uint    `json:"rootUserId"`             //根用户ID
}

// PlatformDynamicReport 平台动态报表
// @Description 平台动态报表
type PlatformDynamicReport struct {
	CurrentDayTotalRecharge          float64 `json:"currentDayTotalRecharge"`          //当日总充值
	CurrentDayTotalWithdraw          float64 `json:"currentDayTotalWithdraw"`          //当日总提现
	CurrentDayTotalActualWithdraw    float64 `json:"currentDayTotalActualWithdraw"`    //当日实际总提现
	CurrentDayNewUserCount           int64   `json:"currentDayNewUserCount"`           //当日新增用户数
	CurrentDayFollowProfit           float64 `json:"currentDayFollowProfit"`           //当日跟单总盈亏
	CurrentDayFollowUserCount        int64   `json:"currentDayFollowUserCount"`        //当日跟单人数
	CurrentDayFirstRechargeUserCount int64   `json:"currentDayFirstRechargeUserCount"` //当日首冲用户数
	CurrentDayActiveUserCount        int64   `json:"currentDayActiveUserCount"`        //当日活跃用户数

	CurrentMonthTotalRecharge       float64 `json:"currentMonthTotalRecharge"`       //当月总充值
	CurrentMonthTotalWithdraw       float64 `json:"currentMonthTotalWithdraw"`       //当月总提现
	CurrentMonthTotalActualWithdraw float64 `json:"currentMonthTotalActualWithdraw"` //当月实际总提现
	CurrentMonthNewUserCount        int64   `json:"currentMonthNewUserCount"`        //当月新增用户数
	CurrentMonthFollowProfit        float64 `json:"currentMonthFollowProfit"`        //当月跟单总盈亏
	CurrentMonthFollowUserCount     int64   `json:"currentMonthFollowUserCount"`     //当月跟单人数

	LastMonthTotalRecharge       float64 `json:"lastMonthTotalRecharge"`       //上月总充值
	LastMonthTotalWithdraw       float64 `json:"lastMonthTotalWithdraw"`       //上月总提现
	LastMonthTotalActualWithdraw float64 `json:"lastMonthTotalActualWithdraw"` //上月实际总提现
	LastMonthFollowUserCount     int64   `json:"lastMonthFollowUserCount"`     //上月跟单人数

	CurrentYearNewUserCount int64   `json:"currentYearNewUserCount"` //全年新增用户数
	CurrentYearFollowProfit float64 `json:"currentYearFollowProfit"` //全年跟单总盈亏

	YesterdayFollowUserCount int64 `json:"yesterdayFollowUserCount"` //昨日跟单人数
	YesterdayActiveUserCount int64 `json:"yesterdayActiveUserCount"` //昨日活跃用户数

	CurrentWeekFollowUserCount int64 `json:"currentWeekFollowUserCount"` //本周跟单人数

	CurrentTotalRecharge float64 `json:"currentTotalRecharge"` //当前总充值
}
