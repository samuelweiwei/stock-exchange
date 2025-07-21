package task

import (
	"database/sql"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/enums/fund"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order"
	"github.com/flipped-aurora/gin-vue-admin/server/service/userfund"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"strconv"
	"time"
)

var userFundService userfund.UserFundAccountsService

func AutoRenewUserFollowOrder(db *gorm.DB) error {
	var followOrders []order.UserFollowOrder
	year, month, day := time.Now().Date()
	err := db.Preload("AdvisorProd").Model(order.UserFollowOrder{}).
		Where("period_end is not null and period_end < ? and follow_order_status != ?",
			time.Date(year, month, day, 0, 0, 0, 0, time.Local),
			order.FollowOrderStatusFinished).
		Find(&followOrders).Error
	if err != nil {
		return err
	}

	for _, v := range followOrders {
		if v.AutoRenew == order.DisableAutoRenew {
			settleAmount := v.RetrievableAmount.Add(v.FollowAvailableAmount).InexactFloat64()
			err = userFundService.UpdateUserFundAccountsAndNewFlow(int(v.UserId), fund.AutoSettle, settleAmount, strconv.Itoa(int(v.ID)))
			if err != nil {
				continue
			}

			v.FollowOrderStatus = order.FollowOrderStatusFinished
			if v.StockStatus != order.FollowOrderStockStatusOpened {
				v.StockStatus = order.FollowOrderStockStatusExpired
			}
			v.FollowAvailableAmount = decimal.Zero
			v.RetrievableAmount = decimal.Zero
			v.EndTime = sql.NullTime{Time: time.Now(), Valid: true}
			err = db.Model(v).Select("FollowOrderStatus", "StockStatus", "EndTime", "FollowAvailableAmount", "RetrievableAmount").Updates(v).Error
		} else {
			v.PeriodEnd = sql.NullTime{Time: global.Calendar.WorkdaysFrom(v.PeriodEnd.Time, v.AdvisorProd.FollowPeriod), Valid: true}
			err = db.Model(v).Select("PeriodEnd").Updates(v).Error
		}

		if err != nil {
			fmt.Println("自动续期错误：" + err.Error())
		}
	}
	return nil
}
