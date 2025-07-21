package report

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/report"
	"github.com/flipped-aurora/gin-vue-admin/server/model/report/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/report/response"
	"gorm.io/gorm"
	"sync"
	"time"
)

type PlatformReportService struct{}

func (s *PlatformReportService) GetPlatformReportSummary(rootUserId uint) (*response.PlatformReportSummary, error) {
	var data report.PlatformReportSummary
	err := global.GVA_DB.Model(&data).Where("root_user_id = ?", rootUserId).First(&data).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &response.PlatformReportSummary{}, nil
	}
	if err != nil {
		return nil, err
	}

	resp := &response.PlatformReportSummary{
		TotalBalance:             data.TotalBalance.InexactFloat64(),
		TotalRecharge:            data.TotalRecharge.InexactFloat64(),
		TotalWithdraw:            data.TotalWithdraw.InexactFloat64(),
		TotalActualWithdraw:      data.TotalActualWithdraw.InexactFloat64(),
		TotalFollowUserCount:     data.TotalFollowUserCount,
		TotalUserCount:           data.TotalUserCount,
		PlatformProfit:           data.PlatformProfit.InexactFloat64(),
		TodayActiveUserCount:     data.TodayActiveUserCount,
		YesterdayActiveUserCount: data.YesterdayActiveUserCount,
		InviteUserCount:          data.InviteUserCount,
		EffectiveInviteUserCount: data.EffectiveInviteUserCount,
	}
	return resp, nil
}

func (s *PlatformReportService) PageQueryPlatformReportDaily(req *request.PlatformReportDailyPageQueryReq) ([]response.PlatformReportDaily, int64, error) {
	db := global.GVA_DB.Model(&report.PlatformReportDaily{})
	if req.ReportDateStart != nil && *req.ReportDateStart > 0 {
		db = db.Where("report_date >= ?", time.UnixMilli(*req.ReportDateStart))
	}
	if req.ReportDateEnd != nil && *req.ReportDateEnd > 0 {
		db = db.Where("report_date <= ?", time.UnixMilli(*req.ReportDateEnd))
	}
	if req.RootUserId > 0 {
		db = db.Where("root_user_id = ?", req.RootUserId)
	}
	if req.RootUserIdSearch != nil {
		db = db.Where("root_user_id = ?", *req.RootUserIdSearch)
	}

	var total int64
	if err := db.Count(&total).Error; err != nil || total == 0 {
		return nil, 0, err
	}

	var reportList []report.PlatformReportDaily
	if err := db.Order("report_date desc").Find(&reportList).Error; err != nil {
		return nil, 0, err
	}

	resp := make([]response.PlatformReportDaily, len(reportList))
	for i, v := range reportList {
		resp[i] = response.PlatformReportDaily{
			TotalRecharge:          v.TotalRecharge.InexactFloat64(),
			RechargeUserCount:      v.RechargeUserCount,
			TotalWithdraw:          v.TotalWithdraw.InexactFloat64(),
			TotalActualWithdraw:    v.TotalActualWithdraw.InexactFloat64(),
			WithdrawUserCount:      v.WithdrawUserCount,
			TotalNewUserCount:      v.TotalNewUserCount,
			ActiveUserCount:        v.ActiveUserCount,
			InvitedUserCount:       v.InviteUserCount,
			FirstRechargeUserCount: v.FirstRechargeUserCount,
			FirstRechargeAmount:    v.FirstRechargeAmount.InexactFloat64(),
			FollowProfit:           v.FollowProfit.InexactFloat64(),
			FollowAmount:           v.FollowAmount.InexactFloat64(),
			FollowUserCount:        v.FollowUserCount,
			ReportDate:             v.ReportDate.UnixMilli(),
			RootUserId:             v.RootUserId,
		}
	}
	return resp, total, nil
}

func (s *PlatformReportService) GetPlatformDynamicReport(req *request.PlatformReportDailyGetReq) (*response.PlatformDynamicReport, error) {
	var (
		currentDate time.Time
		w           sync.WaitGroup
		rootUserId  uint
	)
	if req.ReportDate != nil && *req.ReportDate > 0 {
		year, month, day := time.UnixMilli(*req.ReportDate).Date()
		currentDate = time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	} else {
		year, month, day := time.Now().Date()
		currentDate = time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	}

	if req.RootUserId > 0 {
		rootUserId = req.RootUserId
	} else if req.RootUserIdSearch != nil {
		rootUserId = *req.RootUserIdSearch
	}

	resp := &response.PlatformDynamicReport{}
	w.Add(7)

	go func() {
		defer w.Done()
		var dailyReport report.PlatformReportDaily
		err := global.GVA_DB.Model(&dailyReport).Where("report_date = ? and root_user_id = ?", currentDate, rootUserId).First(&dailyReport).Error
		if err == nil {
			resp.CurrentDayTotalRecharge = dailyReport.TotalRecharge.InexactFloat64()
			resp.CurrentDayTotalWithdraw = dailyReport.TotalWithdraw.InexactFloat64()
			resp.CurrentDayNewUserCount = dailyReport.TotalNewUserCount
			resp.CurrentDayFollowProfit = dailyReport.FollowProfit.InexactFloat64()
			resp.CurrentDayFollowUserCount = dailyReport.FollowUserCount
			resp.CurrentDayFirstRechargeUserCount = dailyReport.FirstRechargeUserCount
			resp.CurrentDayActiveUserCount = dailyReport.ActiveUserCount
			resp.CurrentDayTotalActualWithdraw = dailyReport.TotalActualWithdraw.InexactFloat64()
		}
	}()

	currentMonthBegin := currentDate.AddDate(0, 0, -currentDate.Day()+1)
	currentMonthEnd := currentMonthBegin.AddDate(0, 1, -1)
	go func() {
		defer w.Done()
		var currentMonthReport report.PlatformReportDaily
		global.GVA_DB.Model(&currentMonthReport).
			Select("sum(total_recharge) as total_recharge,"+
				"sum(total_withdraw) as total_withdraw,"+
				"sum(total_new_user_count) as total_new_user_count,"+
				"sum(follow_profit) as follow_profit,"+
				"sum(follow_user_count) as follow_user_count,"+
				"sum(first_recharge_user_count) as first_recharge_user_count,"+
				"sum(total_actual_withdraw) as total_actual_withdraw").
			Where("report_date >= ? and report_date <= ? and root_user_id = ?", currentMonthBegin, currentMonthEnd, rootUserId).
			First(&currentMonthReport)
		resp.CurrentMonthTotalRecharge = currentMonthReport.TotalRecharge.InexactFloat64()
		resp.CurrentMonthTotalWithdraw = currentMonthReport.TotalWithdraw.InexactFloat64()
		resp.CurrentMonthNewUserCount = currentMonthReport.TotalNewUserCount
		resp.CurrentMonthFollowProfit = currentMonthReport.FollowProfit.InexactFloat64()
		resp.CurrentMonthFollowUserCount = currentMonthReport.FollowUserCount
		resp.CurrentMonthTotalActualWithdraw = currentMonthReport.TotalActualWithdraw.InexactFloat64()
	}()

	go func() {
		defer w.Done()
		lastMonthBegin := currentMonthBegin.AddDate(0, -1, 0)
		lastMonthEnd := currentMonthBegin.AddDate(0, 0, -1)
		var lastMonthReport report.PlatformReportDaily
		global.GVA_DB.Model(&lastMonthReport).
			Select("sum(total_recharge) as total_recharge,"+
				"sum(total_withdraw) as total_withdraw,"+
				"sum(follow_user_count) as follow_user_count,"+
				"sum(total_actual_withdraw) as total_actual_withdraw").
			Where("report_date >= ? and report_date <= ? and root_user_id = ?", lastMonthBegin, lastMonthEnd, rootUserId).
			First(&lastMonthReport)
		resp.LastMonthTotalRecharge = lastMonthReport.TotalRecharge.InexactFloat64()
		resp.LastMonthTotalWithdraw = lastMonthReport.TotalWithdraw.InexactFloat64()
		resp.LastMonthFollowUserCount = lastMonthReport.FollowUserCount
		resp.LastMonthTotalActualWithdraw = lastMonthReport.TotalActualWithdraw.InexactFloat64()
	}()

	go func() {
		defer w.Done()
		currentYearBegin := time.Date(currentDate.Year(), time.January, 1, 0, 0, 0, 0, time.Local)
		currentYearEnd := currentYearBegin.AddDate(1, 0, -1)
		var currentYearReport report.PlatformReportDaily
		global.GVA_DB.Model(&currentYearReport).
			Select("sum(total_new_user_count) as total_new_user_count,"+
				"sum(follow_profit) as follow_profit").
			Where("report_date >= ? and report_date <= ? and root_user_id = ?", currentYearBegin, currentYearEnd, rootUserId).
			First(&currentYearReport)
		resp.CurrentYearNewUserCount = currentYearReport.TotalNewUserCount
		resp.CurrentYearFollowProfit = currentYearReport.FollowProfit.InexactFloat64()
	}()

	go func() {
		defer w.Done()
		yesterday := currentDate.AddDate(0, 0, -1)
		var yesterdayReport report.PlatformReportDaily
		global.GVA_DB.Model(&yesterdayReport).Where("report_date = ? and root_user_id = ?", yesterday, rootUserId).
			First(&yesterdayReport)
		resp.YesterdayFollowUserCount = yesterdayReport.FollowUserCount
		resp.YesterdayActiveUserCount = yesterdayReport.ActiveUserCount
	}()

	go func() {
		defer w.Done()
		offset := int(time.Monday - currentDate.Weekday())
		if offset > 0 {
			offset = -6
		}
		currentWeekBegin := currentDate.AddDate(0, 0, offset)
		currentWeekEnd := currentWeekBegin.AddDate(0, 0, 6)
		var currentWeekFollowUserCount int64
		global.GVA_DB.Model(&report.PlatformReportDaily{}).Select("sum(follow_user_count)").
			Where("report_date >= ? and report_date <= ? and root_user_id = ?", currentWeekBegin, currentWeekEnd, rootUserId).
			First(&currentWeekFollowUserCount)
		resp.CurrentWeekFollowUserCount = currentWeekFollowUserCount
	}()

	go func() {
		defer w.Done()
		var currentTotalRecharge float64
		global.GVA_DB.Model(&report.PlatformReportDaily{}).Select("sum(total_recharge)").
			Where("report_date <= ? and root_user_id = ?", currentDate, rootUserId).First(&currentTotalRecharge)
		resp.CurrentTotalRecharge = currentTotalRecharge
	}()

	w.Wait()
	return resp, nil
}
