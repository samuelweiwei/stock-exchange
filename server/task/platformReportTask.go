package task

import (
	"github.com/flipped-aurora/gin-vue-admin/server/enums/fund"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order"
	"github.com/flipped-aurora/gin-vue-admin/server/model/report"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"sync"
)

type CalculateSummaryFunc func(w *sync.WaitGroup, rootUserId uint, db *gorm.DB, r *report.PlatformReportSummary)

var calculateSummaryFuncs = []CalculateSummaryFunc{
	calculateTotalBalance,
	calculateTotalRecharge,
	calculateTotalWithdraw,
	calculateTotalUserCount,
	calculateTotalFollowUserCount,
	calculatePlatformProfit,
	calculateInviterUserCount,
	calculateEffectiveInviteUserCount,
	calculateTotalActualWithdraw,
}

func GeneratePlatformReportSummary(db *gorm.DB) error {
	var rootUserIds []uint
	err := db.Model(user.FrontendUsers{}).Select("distinct(root_userid)").Find(&rootUserIds).Error
	if err != nil {
		return err
	}

	w := sync.WaitGroup{}
	w.Add(len(rootUserIds))

	for _, rootUserId := range rootUserIds {
		go func() {
			defer w.Done()

			var platformReportSummary report.PlatformReportSummary
			db.Where("root_user_id = ?", rootUserId).First(&platformReportSummary)
			platformReportSummary.RootUserId = rootUserId

			w1 := sync.WaitGroup{}
			w1.Add(len(calculateSummaryFuncs))
			for _, v := range calculateSummaryFuncs {
				go v(&w1, rootUserId, db, &platformReportSummary)
			}

			w1.Wait()
			if platformReportSummary.ID > 0 {
				db.Save(&platformReportSummary)
			} else {
				db.Create(&platformReportSummary)
			}
		}()
	}

	w.Wait()
	return nil
}

func calculateTotalBalance(w *sync.WaitGroup, rootUserId uint, db *gorm.DB, r *report.PlatformReportSummary) {
	defer w.Done()

	var (
		rawSql                  string
		followAvailableSql      string
		followHoldingAmountSql  string
		followAppendAmountSql   string
		totalBalance            float64
		totalFollowAvailable    float64
		totalHoldingAmount      float64
		totalFollowAppendAmount float64
		args                    []interface{}
		args1                   []interface{}
		args2                   []interface{}
		args3                   []interface{}
	)

	if rootUserId == 0 {
		rawSql = "select sum(balance) from user_fund_accounts a inner join frontend_users u on u.id = a.user_id and u.user_type != 2 where a.deleted_at is null"
		followHoldingAmountSql = "select sum(buy_price * stock_num) from user_follow_stock_detail f inner join frontend_users u on u.id = f.user_id and u.user_type != 2 where f.sell_price is null and f.deleted_at is null"
		followAvailableSql = "select sum(follow_available_amount + retrievable_amount) from user_follow_order o inner join frontend_users u on u.id = o.user_id and u.user_type != 2 where o.follow_order_status in (?,?) and o.deleted_at is null"
		followAppendAmountSql = "select sum(append_amount) from user_follow_append_order o inner join frontend_users u on u.id = o.user_id and u.user_type != 2 where o.append_order_status in (?) and o.deleted_at is null"
		args = []interface{}{}
		args1 = []interface{}{}
		args2 = []interface{}{order.FollowOrderStatusAuditing, order.FollowOrderStatusFollowing}
		args3 = []interface{}{order.AppendOrderStatusAuditing}
	} else {
		rawSql = "select sum(balance) from user_fund_accounts a inner join frontend_users u on a.user_id = u.id and u.user_type != 2 where u.root_userid = ? and a.deleted_at is null"
		followHoldingAmountSql = "select sum(buy_price * stock_num) from user_follow_stock_detail f inner join frontend_users u on u.id = f.user_id and u.user_type != 2 where f.sell_price is null and u.root_userid = ? and f.deleted_at is null"
		followAvailableSql = "select sum(follow_available_amount + retrievable_amount) from user_follow_order o inner join frontend_users u on u.id = o.user_id and u.user_type != 2 where o.follow_order_status in (?,?) and u.root_userid = ? and o.deleted_at is null"
		followAppendAmountSql = "select sum(append_amount) from user_follow_append_order o inner join frontend_users u on u.id = o.user_id and u.user_type != 2 where o.append_order_status in (?) and u.root_userid = ? and o.deleted_at is null"
		args = []interface{}{rootUserId}
		args1 = []interface{}{rootUserId}
		args2 = []interface{}{order.FollowOrderStatusAuditing, order.FollowOrderStatusFollowing, rootUserId}
		args3 = []interface{}{order.AppendOrderStatusAuditing, rootUserId}
	}

	db.Raw(rawSql, args...).First(&totalBalance, false)
	db.Raw(followHoldingAmountSql, args1...).First(&totalHoldingAmount, false)
	db.Raw(followAvailableSql, args2...).First(&totalFollowAvailable, false)
	db.Raw(followAppendAmountSql, args3...).First(&totalFollowAppendAmount, false)
	r.TotalBalance = decimal.NewFromFloat(totalBalance).Add(decimal.NewFromFloat(totalHoldingAmount)).Add(decimal.NewFromFloat(totalFollowAvailable)).Add(decimal.NewFromFloat(totalFollowAppendAmount)).RoundFloor(2)
}

func calculateTotalRecharge(w *sync.WaitGroup, rootUserId uint, db *gorm.DB, r *report.PlatformReportSummary) {
	defer w.Done()

	var (
		rawSql        string
		totalRecharge float64
		args          []interface{}
	)

	if rootUserId == 0 {
		rawSql = "select sum(amount) from user_account_flow f inner join frontend_users u on f.user_id = u.id and u.user_type != 2 where transaction_type = ? and f.deleted_at is null"
		args = []interface{}{fund.Recharge}
	} else {
		rawSql = "select sum(amount) from user_account_flow f inner join  frontend_users u on f.user_id = u.id and u.user_type != 2 where transaction_type = ? and u.root_userid=? and f.deleted_at is null"
		args = []interface{}{fund.Recharge, rootUserId}
	}

	db.Raw(rawSql, args...).First(&totalRecharge, false)
	r.TotalRecharge = decimal.NewFromFloat(totalRecharge).RoundFloor(2)
}

func calculateTotalWithdraw(w *sync.WaitGroup, rootUserId uint, db *gorm.DB, r *report.PlatformReportSummary) {
	defer w.Done()

	var (
		rawSql        string
		totalWithdraw float64
		args          []interface{}
	)
	if rootUserId == 0 {
		rawSql = "select sum(amount) from user_account_flow f inner join frontend_users u on f.user_id = u.id and u.user_type != 2 where transaction_type = ? and f.deleted_at is null"
		args = []interface{}{fund.Withdraw}
	} else {
		rawSql = "select sum(amount) from user_account_flow f inner join  frontend_users u on f.user_id = u.id and u.user_type != 2 where f.transaction_type = ? and u.root_userid=? and f.deleted_at is null"
		args = []interface{}{fund.Withdraw, rootUserId}
	}

	db.Raw(rawSql, args...).First(&totalWithdraw, false)
	r.TotalWithdraw = decimal.NewFromFloat(totalWithdraw).RoundFloor(2)
}

func calculateTotalUserCount(w *sync.WaitGroup, rootUserId uint, db *gorm.DB, r *report.PlatformReportSummary) {
	defer w.Done()

	var (
		rawSql         string
		totalUserCount int64
		args           []interface{}
	)
	if rootUserId == 0 {
		rawSql = "select count(*) from frontend_users where user_type != 2 and  deleted_at is null"
		args = []interface{}{}
	} else {
		rawSql = "select count(*) from frontend_users where root_userid = ? and user_type != 2 and deleted_at is null"
		args = []interface{}{rootUserId}
	}

	db.Raw(rawSql, args...).First(&totalUserCount, false)
	r.TotalUserCount = totalUserCount
}

func calculateInviterUserCount(w *sync.WaitGroup, rootUserId uint, db *gorm.DB, r *report.PlatformReportSummary) {
	defer w.Done()

	var (
		rawSql          string
		inviteUserCount int64
		args            []interface{}
	)

	if rootUserId == 0 {
		rawSql = `select count(*) from frontend_users where user_type != 2 and root_userid > 0 and deleted_at is null`
		args = []interface{}{}
	} else {
		rawSql = `select count(*) from frontend_users where user_type != 2 and root_userid = ? and deleted_at is null`
		args = []interface{}{rootUserId}
	}
	db.Raw(rawSql, args...).First(&inviteUserCount, false)
	r.InviteUserCount = inviteUserCount
}

func calculateTotalFollowUserCount(w *sync.WaitGroup, rootUserId uint, db *gorm.DB, r *report.PlatformReportSummary) {
	defer w.Done()

	var (
		rawSql               string
		totalFollowUserCount int64
		args                 []interface{}
	)
	if rootUserId == 0 {
		rawSql = "select count(distinct(user_id)) from user_follow_order f inner join frontend_users u on f.user_id = u.id and u.user_type != 2 where f.follow_order_status not in (?,?) and f.deleted_at is null"
		args = []interface{}{order.FollowOrderStatusRejected, order.FollowOrderStatusCancelled}
	} else {
		rawSql = "select count(distinct(user_id)) from user_follow_order f inner join frontend_users u on f.user_id = u.id and u.user_type != 2 where u.root_userid = ? and f.follow_order_status not in (?,?) and f.deleted_at is null"
		args = []interface{}{rootUserId, order.FollowOrderStatusRejected, order.FollowOrderStatusCancelled}
	}

	db.Raw(rawSql, args...).First(&totalFollowUserCount, false)
	r.TotalFollowUserCount = totalFollowUserCount
}

func calculatePlatformProfit(w *sync.WaitGroup, rootUserId uint, db *gorm.DB, r *report.PlatformReportSummary) {
	defer w.Done()

	var (
		profitRawSql string
		lossRawSql   string
		profit       float64
		loss         float64
		args         []interface{}
	)

	if rootUserId == 0 {
		profitRawSql = `select sum((buy_price - sell_price)*stock_num) 
						from user_follow_stock_detail d inner join frontend_users u on d.user_id = u.id and u.user_type != 2
                        where sell_price <= buy_price and d.deleted_at is null`

		lossRawSql = `select sum((sell_price - buy_price)*stock_num - advisor_commission - platform_commission) 
					  from user_follow_stock_detail d inner join frontend_users u on d.user_id = u.id and u.user_type != 2
                      where sell_price > buy_price and d.deleted_at is null`

		args = []interface{}{}
	} else {
		profitRawSql = `select sum((buy_price - sell_price)*stock_num) 
						from user_follow_stock_detail d inner join frontend_users u on d.user_id = u.id and u.user_type != 2
						where d.sell_price <= d.buy_price and u.root_userid = ? and d.deleted_at is null`

		lossRawSql = `select sum((sell_price - buy_price)*stock_num - advisor_commission - platform_commission) 
					  from user_follow_stock_detail d inner join frontend_users u on d.user_id = u.id and u.user_type != 2
					  where d.sell_price > d.buy_price and u.root_userid = ? and d.deleted_at is null`

		args = []interface{}{rootUserId}
	}

	db.Raw(profitRawSql, args...).First(&profit, false)
	db.Raw(lossRawSql, args...).First(&loss, false)

	r.PlatformProfit = decimal.NewFromFloat(profit).Sub(decimal.NewFromFloat(loss)).RoundFloor(2)
}

func calculateEffectiveInviteUserCount(w *sync.WaitGroup, rootUserId uint, db *gorm.DB, r *report.PlatformReportSummary) {
	defer w.Done()

	var (
		rawSql             string
		effectiveUserCount int64
		args               []interface{}
	)

	if rootUserId == 0 {
		rawSql = `select count(distinct(a.user_id)) from user_fund_accounts a inner join frontend_users u on u.id = a.user_id and u.user_type != 2 where u.root_userid > 0 and a.first_charge_time is not null and u.deleted_at is null and a.deleted_at is null`
		args = []interface{}{}
	} else {
		rawSql = `select count(distinct(a.user_id)) from user_fund_accounts a inner join frontend_users u on u.id = a.user_id and u.user_type != 2 where u.root_userid = ? and a.first_charge_time is not null and u.deleted_at is null and a.deleted_at is null`
		args = []interface{}{rootUserId}
	}

	db.Raw(rawSql, args...).First(&effectiveUserCount, false)
	r.EffectiveInviteUserCount = effectiveUserCount
}

func calculateTotalActualWithdraw(w *sync.WaitGroup, rootUserId uint, db *gorm.DB, r *report.PlatformReportSummary) {
	defer w.Done()

	var (
		rawSql              string
		totalActualWithdraw float64
		args                []interface{}
	)
	if rootUserId == 0 {
		rawSql = "select sum(amount - coalesce(total_commission, 0)) from user_account_flow f inner join frontend_users u on f.user_id = u.id and u.user_type != 2 where transaction_type = ? and f.deleted_at is null"
		args = []interface{}{fund.Withdraw}
	} else {
		rawSql = "select sum(amount - coalesce(total_commission, 0)) from user_account_flow f inner join  frontend_users u on f.user_id = u.id and u.user_type != 2 where f.transaction_type = ? and u.root_userid=? and f.deleted_at is null"
		args = []interface{}{fund.Withdraw, rootUserId}
	}

	db.Raw(rawSql, args...).First(&totalActualWithdraw, false)
	r.TotalActualWithdraw = decimal.NewFromFloat(totalActualWithdraw)
}
