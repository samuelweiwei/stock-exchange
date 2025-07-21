package task

import (
	"github.com/flipped-aurora/gin-vue-admin/server/enums/fund"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order"
	"github.com/flipped-aurora/gin-vue-admin/server/model/report"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"sync"
	"time"
)

var calculateDailyFuncs = []DailyCalculateFunc{
	calculateDailyTotalRecharge,
	calculateDailyRechargeUserCount,
	calculateDailyTotalWithdraw,
	calculateDailyWithdrawUserCount,
	calculateDailyNewUserCount,
	calculateDailyFirstRechargeUserCount,
	calculateDailyFirstRechargeAmount,
	calculateDailyFollowProfit,
	calculateDailyFollowUserCount,
	calculateDailyFollowAmount,
	calculateDailyInviteUserCount,
	calculateDailyActiveUserCount,
	calculateDailyTotalActualWithdraw,
}

type DailyCalculateFunc func(w *sync.WaitGroup, rootUserId uint, beginDate time.Time, endDate time.Time, db *gorm.DB, r *report.PlatformReportDaily)

func GeneratePlatformReportDaily(db *gorm.DB) error {
	var rootUserIds []uint
	err := db.Model(user.FrontendUsers{}).Select("distinct(root_userid)").Find(&rootUserIds).Error
	if err != nil {
		return err
	}

	year, month, day := time.Now().Date()
	reportBeginDate := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	reportEndDate := reportBeginDate.AddDate(0, 0, 1)

	w := sync.WaitGroup{}
	w.Add(len(rootUserIds))

	for _, rootUserId := range rootUserIds {
		go func() {
			defer w.Done()

			var reportDaily report.PlatformReportDaily
			db.Where("report_date = ? and root_user_id = ?", reportBeginDate, rootUserId).First(&reportDaily)
			reportDaily.ReportDate = reportBeginDate
			reportDaily.RootUserId = rootUserId

			w1 := sync.WaitGroup{}
			w1.Add(len(calculateDailyFuncs))

			for _, v := range calculateDailyFuncs {
				go v(&w1, rootUserId, reportBeginDate, reportEndDate, db, &reportDaily)
			}

			w1.Wait()
			if reportDaily.ID > 0 {
				db.Save(&reportDaily)
			} else {
				db.Create(&reportDaily)
			}
		}()
	}

	w.Wait()
	return nil
}

func calculateDailyTotalRecharge(w *sync.WaitGroup, rootUserId uint, beginDate time.Time, endDate time.Time, db *gorm.DB, r *report.PlatformReportDaily) {
	defer w.Done()

	var (
		rawSql        string
		totalRecharge float64
		args          []interface{}
	)

	if rootUserId == 0 {
		rawSql = `select sum(amount) from user_account_flow f inner join frontend_users u on f.user_id = u.id and u.user_type != 2
                   where transaction_type = ? and transaction_date >= ? and transaction_date < ? and f.deleted_at is null `
		args = []interface{}{fund.Recharge, beginDate, endDate}
	} else {
		rawSql = `select sum(amount) from user_account_flow f inner join frontend_users u on f.user_id = u.id and u.user_type != 2
                   where transaction_type = ? and transaction_date >= ? and transaction_date < ? and f.deleted_at is null and u.root_userid = ? `
		args = []interface{}{fund.Recharge, beginDate, endDate, rootUserId}
	}

	db.Raw(rawSql, args...).First(&totalRecharge, false)
	r.TotalRecharge = decimal.NewFromFloat(totalRecharge).RoundFloor(2)
}

func calculateDailyRechargeUserCount(w *sync.WaitGroup, rootUserId uint, beginDate time.Time, endDate time.Time, db *gorm.DB, r *report.PlatformReportDaily) {
	defer w.Done()

	var (
		rawSql            string
		rechargeUserCount int64
		args              []interface{}
	)

	if rootUserId == 0 {
		rawSql = `select count(distinct(user_id)) from user_account_flow f inner join frontend_users u on f.user_id = u.id and u.user_type != 2
                                where transaction_type = ? and transaction_date >= ? and transaction_date < ? and f.deleted_at is null `
		args = []interface{}{fund.Recharge, beginDate, endDate}
	} else {
		rawSql = `select count(distinct(user_id)) from user_account_flow f inner join frontend_users u on f.user_id = u.id  and u.user_type != 2
                   where transaction_type = ? and transaction_date >= ? and transaction_date < ? and f.deleted_at is null and u.root_userid = ?`
		args = []interface{}{fund.Recharge, beginDate, endDate, rootUserId}
	}

	db.Raw(rawSql, args...).First(&rechargeUserCount, false)
	r.RechargeUserCount = rechargeUserCount
}

func calculateDailyTotalWithdraw(w *sync.WaitGroup, rootUserId uint, beginDate time.Time, endDate time.Time, db *gorm.DB, r *report.PlatformReportDaily) {
	defer w.Done()

	var (
		rawSql        string
		totalWithdraw float64
		args          []interface{}
	)

	if rootUserId == 0 {
		rawSql = `select sum(amount) from user_account_flow f inner join frontend_users u on f.user_id = u.id and u.user_type != 2
                   where transaction_type = ? and transaction_date >= ? and transaction_date < ? and f.deleted_at is null`
		args = []interface{}{fund.Withdraw, beginDate, endDate}
	} else {
		rawSql = `select sum(amount) from user_account_flow f inner join frontend_users u on f.user_id = u.id  and u.user_type != 2
                   where transaction_type = ? and transaction_date >= ? and transaction_date < ? and f.deleted_at is null and u.root_userid = ?`
		args = []interface{}{fund.Withdraw, beginDate, endDate, rootUserId}
	}

	db.Raw(rawSql, args...).First(&totalWithdraw, false)
	r.TotalWithdraw = decimal.NewFromFloat(totalWithdraw).RoundFloor(2)
}

func calculateDailyWithdrawUserCount(w *sync.WaitGroup, rootUserId uint, beginDate time.Time, endDate time.Time, db *gorm.DB, r *report.PlatformReportDaily) {
	defer w.Done()

	var (
		rawSql            string
		withdrawUserCount int64
		args              []interface{}
	)

	if rootUserId == 0 {
		rawSql = `select count(distinct(user_id)) from user_account_flow f inner join frontend_users u on f.user_id = u.id and u.user_type != 2
                              where transaction_type = ? and transaction_date >= ? and transaction_date < ? and f.deleted_at is null`
		args = []interface{}{fund.Withdraw, beginDate, endDate}
	} else {
		rawSql = `select count(distinct(user_id)) from user_account_flow f inner join frontend_users u on f.user_id = u.id and u.user_type != 2
                   where transaction_type = ? and transaction_date >= ? and transaction_date < ? and f.deleted_at is null and u.root_userid = ?`
		args = []interface{}{fund.Withdraw, beginDate, endDate, rootUserId}
	}

	db.Raw(rawSql, args...).First(&withdrawUserCount, false)
	r.WithdrawUserCount = withdrawUserCount
}

func calculateDailyNewUserCount(w *sync.WaitGroup, rootUserId uint, beginDate time.Time, endDate time.Time, db *gorm.DB, r *report.PlatformReportDaily) {
	defer w.Done()

	var (
		rawSql       string
		newUserCount int64
		args         []interface{}
	)

	if rootUserId == 0 {
		rawSql = "select count(*) from frontend_users where created_at >= ? and created_at < ? and user_type != 2 and deleted_at is null"
		args = []interface{}{beginDate, endDate}
	} else {
		rawSql = "select count(*) from frontend_users where created_at >= ? and created_at < ? and root_userid = ? and user_type != 2 and deleted_at is null"
		args = []interface{}{beginDate, endDate, rootUserId}
	}

	db.Raw(rawSql, args...).First(&newUserCount, false)
	r.TotalNewUserCount = newUserCount
}

func calculateDailyFirstRechargeUserCount(w *sync.WaitGroup, rootUserId uint, beginDate time.Time, endDate time.Time, db *gorm.DB, r *report.PlatformReportDaily) {
	defer w.Done()

	var (
		rawSql                 string
		firstRechargeUserCount int64
		args                   []interface{}
	)

	if rootUserId == 0 {
		rawSql = `select count(distinct(user_id)) from user_fund_accounts f inner join frontend_users u on f.user_id = u.id and u.user_type != 2 
                                where first_charge_time >= ? and first_charge_time < ? and f.deleted_at is null`
		args = []interface{}{beginDate, endDate}
	} else {
		rawSql = `select count(distinct(user_id)) from user_fund_accounts f inner join frontend_users u on f.user_id = u.id and u.user_type != 2 
                                where first_charge_time >= ? and first_charge_time < ? and u.root_userid = ? and f.deleted_at is null`
		args = []interface{}{beginDate, endDate, rootUserId}
	}

	db.Raw(rawSql, args...).First(&firstRechargeUserCount, false)
	r.FirstRechargeUserCount = firstRechargeUserCount
}

func calculateDailyFirstRechargeAmount(w *sync.WaitGroup, rootUserId uint, beginDate time.Time, endDate time.Time, db *gorm.DB, r *report.PlatformReportDaily) {
	defer w.Done()

	var (
		rawSql              string
		firstRechargeAmount float64
		args                []interface{}
	)

	if rootUserId == 0 {
		rawSql = `select sum(first_charge_amount) from user_fund_accounts f inner join frontend_users u on f.user_id = u.id and u.user_type != 2 
                                where first_charge_time >= ? and first_charge_time < ? and f.deleted_at is null`
		args = []interface{}{beginDate, endDate}
	} else {
		rawSql = `select sum(first_charge_amount) from user_fund_accounts f inner join frontend_users u on f.user_id = u.id and u.user_type != 2 
                                where first_charge_time >= ? and first_charge_time < ? and u.root_userid = ? and f.deleted_at is null`
		args = []interface{}{beginDate, endDate, rootUserId}
	}

	db.Raw(rawSql, args...).First(&firstRechargeAmount, false)
	r.FirstRechargeAmount = decimal.NewFromFloat(firstRechargeAmount)
}

func calculateDailyFollowProfit(w *sync.WaitGroup, rootUserId uint, beginDate time.Time, endDate time.Time, db *gorm.DB, r *report.PlatformReportDaily) {
	defer w.Done()

	var (
		profitRawSql string
		lossRawSql   string
		profit       float64
		loss         float64
		args         []interface{}
	)

	if rootUserId == 0 {
		profitRawSql = `select sum((buy_price - sell_price)*stock_num) from user_follow_stock_detail d inner join frontend_users u on d.user_id = u.id and u.user_type != 2
                                               where sell_price <= buy_price and d.sell_time >= ? and d.sell_time < ? and d.deleted_at is null`

		lossRawSql = `select sum((sell_price - buy_price)*stock_num - advisor_commission - platform_commission) from user_follow_stock_detail d inner join frontend_users u on d.user_id = u.id and u.user_type != 2
                                               where sell_price > buy_price and d.sell_time >= ? and d.sell_time < ? and d.deleted_at is null`
		args = []interface{}{beginDate, endDate}
	} else {
		profitRawSql = `select sum((buy_price - sell_price)*stock_num) from user_follow_stock_detail d inner join frontend_users u on d.user_id = u.id and u.user_type != 2
                                               where sell_price <= buy_price and d.sell_time >= ? and d.sell_time < ? and u.root_userid = ? and d.deleted_at is null`

		lossRawSql = `select sum((sell_price - buy_price)*stock_num - advisor_commission - platform_commission) from user_follow_stock_detail d inner join frontend_users u on d.user_id = u.id and u.user_type != 2
                                               where sell_price > buy_price and d.sell_time >= ? and d.sell_time < ? and u.root_userid = ? and d.deleted_at is null`

		args = []interface{}{beginDate, endDate, rootUserId}
	}

	db.Raw(profitRawSql, args...).First(&profit, false)
	db.Raw(lossRawSql, args...).First(&loss, false)

	r.FollowProfit = decimal.NewFromFloat(profit).Sub(decimal.NewFromFloat(loss)).RoundFloor(2)
}

func calculateDailyFollowAmount(w *sync.WaitGroup, rootUserId uint, beginDate time.Time, endDate time.Time, db *gorm.DB, r *report.PlatformReportDaily) {
	defer w.Done()

	var (
		followAmountRawSql string
		appendAmountRawSql string
		followAmount       float64
		appendAmount       float64
		args               []interface{}
		args2              []interface{}
	)

	if rootUserId == 0 {
		followAmountRawSql = `select sum(follow_amount) from user_follow_order o inner join frontend_users u on o.user_id = u.id and u.user_type != 2 
                          where o.created_at >= ? and o.created_at < ? and o.follow_order_status not in (?,?) and o.deleted_at is null`
		appendAmountRawSql = `select sum(append_amount) from user_follow_append_order o inner join frontend_users u on o.user_id = u.id and u.user_type != 2
							where o.created_at >= ? and o.created_at < ? and o.append_order_status not in (?)  and o.deleted_at is null`
		args = []interface{}{beginDate, endDate, order.FollowOrderStatusRejected, order.FollowOrderStatusCancelled}
		args2 = []interface{}{beginDate, endDate, order.AppendOrderStatusRejected}
	} else {
		followAmountRawSql = `select sum(follow_amount) from user_follow_order o inner join frontend_users u on o.user_id = u.id and u.user_type != 2 
                          where o.created_at >= ? and o.created_at < ? and u.root_userid = ? and o.follow_order_status not in (?,?) and o.deleted_at is null`
		appendAmountRawSql = `select sum(append_amount) from user_follow_append_order o inner join frontend_users u on o.user_id = u.id and u.user_type != 2
							where o.created_at >= ? and o.created_at < ? and u.root_userid = ? and o.append_order_status not in (?) and o.deleted_at is null`
		args = []interface{}{beginDate, endDate, rootUserId, order.FollowOrderStatusRejected, order.FollowOrderStatusCancelled}
		args2 = []interface{}{beginDate, endDate, rootUserId, order.AppendOrderStatusRejected}
	}

	db.Raw(followAmountRawSql, args...).First(&followAmount, false)
	db.Raw(appendAmountRawSql, args2...).First(&appendAmount, false)

	r.FollowAmount = decimal.NewFromFloat(followAmount).Add(decimal.NewFromFloat(appendAmount)).RoundFloor(2)
}

func calculateDailyFollowUserCount(w *sync.WaitGroup, rootUserId uint, beginDate time.Time, endDate time.Time, db *gorm.DB, r *report.PlatformReportDaily) {
	defer w.Done()

	var (
		rawSql          string
		followUserCount int64
		args            []interface{}
	)

	if rootUserId == 0 {
		rawSql = `select count(distinct(user_id)) from user_follow_order o inner join frontend_users u on o.user_id = u.id and u.user_type != 2 
                                where o.created_at >= ? and o.created_at < ? and o.follow_order_status not in (?,?) and o.deleted_at is null`
		args = []interface{}{beginDate, endDate, order.FollowOrderStatusRejected, order.FollowOrderStatusCancelled}
	} else {
		rawSql = `select count(distinct(user_id)) from user_follow_order o inner join frontend_users u on o.user_id = u.id and u.user_type != 2 
                                where o.created_at >= ? and o.created_at < ? and u.root_userid = ? and o.follow_order_status not in (?,?) and o.deleted_at is null`
		args = []interface{}{beginDate, endDate, rootUserId, order.FollowOrderStatusRejected, order.FollowOrderStatusCancelled}
	}

	db.Raw(rawSql, args...).First(&followUserCount, false)
	r.FollowUserCount = followUserCount
}

func calculateDailyInviteUserCount(w *sync.WaitGroup, rootUserId uint, beginDate time.Time, endDate time.Time, db *gorm.DB, r *report.PlatformReportDaily) {
	defer w.Done()

	var (
		rawSql          string
		inviteUserCount int64
		args            []interface{}
	)

	if rootUserId == 0 {
		rawSql = `select count(*) from frontend_users where user_type != 2 and created_at >= ? and created_at < ? and root_userid > 0 and deleted_at is null`
		args = []interface{}{beginDate, endDate}
	} else {
		rawSql = `select count(*) from frontend_users where user_type != 2 and created_at >= ? and created_at < ? and root_userid = ? and deleted_at is null`
		args = []interface{}{beginDate, endDate, rootUserId}
	}

	db.Raw(rawSql, args...).First(&inviteUserCount, false)
	r.InviteUserCount = inviteUserCount
}

func calculateDailyActiveUserCount(w *sync.WaitGroup, rootUserId uint, beginDate time.Time, endDate time.Time, db *gorm.DB, r *report.PlatformReportDaily) {
	defer w.Done()

	var (
		rawSql          string
		activeUserCount int64
		args            []interface{}
	)

	if rootUserId == 0 {
		rawSql = `select count(*) from frontend_users where user_type != 2 and last_login_time >= ? and last_login_time < ? and deleted_at is null`
		args = []interface{}{beginDate.UnixMilli(), endDate.UnixMilli()}
	} else {
		rawSql = `select count(*) from frontend_users where user_type != 2 and last_login_time >= ? and last_login_time < ? and root_userid = ? and deleted_at is null`
		args = []interface{}{beginDate.UnixMilli(), endDate.UnixMilli(), rootUserId}
	}

	db.Raw(rawSql, args...).First(&activeUserCount, false)
	r.ActiveUserCount = activeUserCount
}

func calculateDailyTotalActualWithdraw(w *sync.WaitGroup, rootUserId uint, beginDate time.Time, endDate time.Time, db *gorm.DB, r *report.PlatformReportDaily) {
	defer w.Done()

	var (
		rawSql              string
		totalActualWithdraw float64
		args                []interface{}
	)

	if rootUserId == 0 {
		rawSql = `select sum(amount - coalesce(total_commission, 0)) from user_account_flow f inner join frontend_users u on f.user_id = u.id and u.user_type != 2
                   where transaction_type = ? and transaction_date >= ? and transaction_date < ? and f.deleted_at is null`
		args = []interface{}{fund.Withdraw, beginDate, endDate}
	} else {
		rawSql = `select sum(amount -coalesce(total_commission, 0)) from user_account_flow f inner join frontend_users u on f.user_id = u.id  and u.user_type != 2
                   where transaction_type = ? and transaction_date >= ? and transaction_date < ? and f.deleted_at is null and u.root_userid = ?`
		args = []interface{}{fund.Withdraw, beginDate, endDate, rootUserId}
	}

	db.Raw(rawSql, args...).First(&totalActualWithdraw, false)
	r.TotalActualWithdraw = decimal.NewFromFloat(totalActualWithdraw)
}
