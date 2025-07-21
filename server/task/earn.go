package task

import (
	"context"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/enums/fund"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/earn"
	earnReq "github.com/flipped-aurora/gin-vue-admin/server/model/earn/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service/userfund"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"math/rand"
	"strings"
	"time"
)

func RandomProductRate(db *gorm.DB) error {
	var (
		err           error
		info          earnReq.EarnProductsSearch
		earnProducts  []earn.EarnProducts
		period        = time.Now().Truncate(24 * time.Hour)
		n             = time.Now().UnixMilli()
		interestRates []*earn.EarnInterestRates
	)
	info.PageSize = 99999
	info.Page = 1
	earnProducts, _, err = earnProductService.GetEarnProductsInfoList(info)
	if err != nil {
		global.GVA_LOG.Error("query earn product info err", zap.Error(err))
		return err
	}

	for _, v := range earnProducts {
		rate := decimal.NewFromFloat(rand.Float64() * v.MaxInterestRates.Sub(v.MinInterestRates).InexactFloat64()).
			Add(v.MinInterestRates).Round(2)
		interestRate := &earn.EarnInterestRates{
			ProductId:     v.Id,
			Period:        period,
			InterestRates: rate,
			CreatedAt:     n,
			UpdatedAt:     time.Now(),
		}
		interestRates = append(interestRates, interestRate)
	}
	if err = earnInterestRateService.BatchCreateEarnInterestRates(interestRates); err != nil && !strings.Contains(err.Error(), "Duplicate") {
		global.GVA_LOG.Error("batch create interest rate err", zap.Error(err))
		return err
	}
	return nil
}

func DoEarnProductDailyIncome(ctx context.Context) (err error) {
	period := time.Now().Add(-24 * time.Hour).Truncate(24 * time.Hour)
	tx := global.GVA_DB.Begin()
	defer tx.Rollback()
	var (
		subscribeLogList []*earn.EarnSubscribeLog
		dailyIncomes     []*earn.EarnDailyIncomeMoneyLog
		n                = time.Now()
	)
	subscribeLogList, err = earnProductSubscriptionService.FindStakingSubscribeLog(tx, period.UnixMilli())
	for _, r := range subscribeLogList {
		interestRate, err := earnInterestRateService.GetPeriodEarnInterestRates(r.ProductId, period)
		if err != nil {
			global.GVA_LOG.Error("get rate err", zap.Error(err),
				zap.Time("period", period), zap.Any("product id", r.ProductId))
			continue
		}
		dailyIncome := &earn.EarnDailyIncomeMoneyLog{
			Uid:           r.Uid,
			ProductId:     r.ProductId,
			SubscribeId:   r.Id,
			Earnings:      r.BoughtNum.Mul(interestRate.InterestRates).Div(decimal.NewFromFloat(365.0)),
			InterestRates: interestRate.InterestRates,
			OfferedAt:     period.UnixMilli(),
			CreatedAt:     n.UnixMilli(),
			UpdatedAt:     &n,
			BoughtNum:     r.BoughtNum,
		}
		dailyIncomes = append(dailyIncomes, dailyIncome)
	}

	if len(dailyIncomes) > 0 {
		if err = earnProductDailyIncomeService.BatchSave(tx, dailyIncomes); err != nil {
			global.GVA_LOG.Error("batch save daily income err", zap.Error(err), zap.Any("period", period))
			return err
		}
	}
	if err = tx.Commit().Error; err != nil {
		global.GVA_LOG.Error("commit tx err", zap.Error(err), zap.Any("period", period))
	}
	return err
}

func EarnProductExpiration(ctx context.Context) (err error) {
	tx := global.GVA_DB.Begin()
	defer tx.Rollback()
	var (
		subscribeLogList []*earn.EarnSubscribeLog
		n                = time.Now()
	)
	subscribeLogList, err = earnProductSubscriptionService.FindExpiredSubscribeLog(tx, n.UnixMilli())

	for _, v := range subscribeLogList {
		var (
			product earn.EarnProducts
			amount  decimal.Decimal
		)
		product, err = earnProductService.GetEarnProducts(fmt.Sprint(v.ProductId))
		if err != nil {
			global.GVA_LOG.Error("EarnProductExpiration get earn product err",
				zap.Error(err), zap.Any("uid", v.Uid), zap.Any("product id", v.ProductId))
			continue
		}

		var (
			earned decimal.Decimal
			fine   decimal.Decimal
		)
		earned, err = earnProductDailyIncomeService.GetUserProductionEarnings(tx, v.Id, v.Uid)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			global.GVA_LOG.Error("get earn product earnings err", zap.Error(err), zap.Any("uid", v.Uid), zap.Any("product id", v.ProductId))
			continue
		}
		if v.RedeemInAdvance == earn.RedeemInAdvance && product.Type == earn.Fixed {
			fine = v.BoughtNum.Mul(v.PenaltyRatio).Neg()
		}
		v.Status = earn.Redeemed
		v.Fine = fine
		v.UpdatedAtX = time.Now().UnixMilli()
		err = earnProductSubscriptionService.Fine(tx, *v)
		if err != nil {
			global.GVA_LOG.Error("EarnProductExpiration update earn products err:", zap.Error(err),
				zap.Any("uid", v.Uid), zap.Any("subscription id", v.Id))
			continue
		}
		earnedWithFine := fine.Add(earned)
		amount = v.BoughtNum.Add(earnedWithFine)
		global.GVA_LOG.Error("EarnProductExpiration account change info",
			zap.Error(err), zap.Any("uid", v.Uid), zap.Any("product id", product.Id), zap.Any("earned", earned),
			zap.Any("fine", fine), zap.Any("uid", v.Uid), zap.Any("amount", amount))
		err = userfund.NewUserFundAccountService(tx, false, userfund.WithAvailableChange(amount),
			userfund.WithFrozenChange(v.BoughtNum.Neg()), userfund.WithTotalBalanceChange(earnedWithFine)).
			UpdateUserFundAccountsAndNewFlow(int(v.Uid), fund.RedeemEarnProduct, amount.InexactFloat64(),
				fmt.Sprintf("%v_%v_%v_%v", v.Uid, v.ProductId, amount, v.Id))

	}
	if err = tx.Commit().Error; err != nil {
		global.GVA_LOG.Error("EarnProductExpiration commit tx err", zap.Error(err), zap.Any("period", n))
		return err
	}
	return nil
}
