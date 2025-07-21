package task

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/enums"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	userfundReq "github.com/flipped-aurora/gin-vue-admin/server/model/userfund/request"
	"time"

	recordService "github.com/flipped-aurora/gin-vue-admin/server/service/userfund"
	"go.uber.org/zap"
)

// UpdateSymbolsKline 更新所有symbols的K线数据
func UpdateRecordStatus() error {
	defer func() {
		if err := recover(); err != nil {
			global.GVA_LOG.Error(fmt.Sprintf("Recover from UpdateSymbolsKline panic: %v", err))
		}
	}()
	rechargeRecordsService := &recordService.RechargeRecordsService{}
	var pageInfo userfundReq.RechargeRecordsSearch
	pageInfo.OrderStatus = enums.PENDING
	pageInfo.RechargeType = enums.RechargeTypeQuick
	//查询所有支付中的快捷充值订单
	list, _, err := rechargeRecordsService.GetRechargeRecordsInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取支付中的快捷充值订单失败", zap.Error(err))
		return err
	}
	timeoutDuration := 2 * time.Hour
	now := time.Now()
	for _, order := range list {
		if order.OrderStatus == enums.PENDING && now.Sub(order.CreatedAt) > timeoutDuration {
			// 更新订单状态为支付超时（实际应更新数据库）
			order.OrderStatus = enums.OUTTIME
			fmt.Printf("Order %d timed out\n", order.ID)
			rechargeRecordsService.UpdateRechargeRecords(order)
		}
	}
	return nil
}
