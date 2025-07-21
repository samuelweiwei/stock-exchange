package initialize

import (
	"context"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/task"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

func Timer() {
	go func() {
		var option []cron.Option
		option = append(option, cron.WithSeconds())
		// 清理DB定时任务
		_, err := global.GVA_Timer.AddTaskByFunc("ClearDB", "@daily", func() {
			err := task.ClearTable(global.GVA_DB) // 定时任务方法定在task文件包中
			if err != nil {
				fmt.Println("timer error:", err)
			}
		}, "定时清理数据库【日志，黑名单】内容", option...)
		if err != nil {
			fmt.Println("add timer error:", err)
		}

		// 添加更新股票信息的定时任务
		_, err = global.GVA_Timer.AddTaskByFunc("UpdateSymbolsInfo", "@daily", func() {
			err := task.UpdateSymbolsInfo()
			if err != nil {
				global.GVA_LOG.Error("更新股票信息失败", zap.Error(err))
			}
		}, "更新股票基本信息", option...)
		if err != nil {
			global.GVA_LOG.Error("添加股票更新定时任务失败", zap.Error(err))
		}

		// 添加用户跟单自动续期功能
		addAutoRenewUserFollowOrderTask(option)
		// 添加平台基础设置缓存刷新
		addRefreshPlatformBaseConfigCache(option)
		// 同步市场节假日
		addSyncMarketHolidayTask(option)
		// 添加K线最新一个元素数据更新任务
		addUpdateSymbolsKlineLastTask(option)
		// 添加平台固定报表更新任务
		addGeneratePlatformReportSummary(option)
		// 添加平台日报表更新任务
		addGeneratePlatformReportDaily(option)

		// 添加订单状态支付超时 更新
		addUpdateOrderStatusTask(option)
		// 理财质押相关任务
		addRandomEarnProductRate(option)
		doEarnProductDailyIncome(option)
		earnProductExpiration(option)

		// 添加定时根据登录ip获取地区的任务
		addFetchIpToRegionTask(option)
		// addUpdateWebSocketSubscriptionsTask(option)

		// 添加历史K线数据更新任务
		addUpdateSymbolsKlineHistoryTask(option)
	}()
}

func addAutoRenewUserFollowOrderTask(option []cron.Option) {
	_, err := global.GVA_Timer.AddTaskByFunc("AutoRenewUserFollowOrder", "@daily", func() {
		err := task.AutoRenewUserFollowOrder(global.GVA_DB)
		if err != nil {
			fmt.Println("autoRenewUserFollowOrder error:", err)
		}
	}, "用户跟单自动续期定时任务", option...)

	if err != nil {
		fmt.Println("add timer error:", err)
	}
}

func addRefreshPlatformBaseConfigCache(option []cron.Option) {
	_, err := global.GVA_Timer.AddTaskByFunc("RefreshPlatformBaseConfigCache", "@every 10s", func() {
		err := task.RefreshPlatformBaseConfigCache(global.GVA_DB)
		if err != nil {
			fmt.Println("refreshPlatformBaseConfigCache error:", err)
		}
	}, "自动刷新平台基础设置缓存", option...)

	if err != nil {
		fmt.Println("add timer error:", err)
	}
}

func addSyncMarketHolidayTask(option []cron.Option) {
	_, err := global.GVA_Timer.AddTaskByFunc("SyncMarketHolidayTask", "@daily", func() {
		err := task.SyncExchangeHoliday(global.GVA_DB)
		if err != nil {
			fmt.Println("syncMarketHolidayTask error:", err)
		}
	}, "同步市场节假日", option...)
	if err != nil {
		fmt.Println("add timer error:", err)
	}
}

func addUpdateOrderStatusTask(option []cron.Option) {
	_, err := global.GVA_Timer.AddTaskByFunc("UpdateOrderStatus", "@every 5m", func() {
		err := task.UpdateRecordStatus()
		if err != nil {
			global.GVA_LOG.Error("更新订单支付超时状态失败", zap.Error(err))
		}
	}, "更新订单支付超时状态", option...)
	if err != nil {
		global.GVA_LOG.Error("更新订单支付超时状态失败", zap.Error(err))
	}
}

func addUpdateSymbolsKlineLastTask(option []cron.Option) {
	_, err := global.GVA_Timer.AddTaskByFunc("UpdateSymbolsKlineLast", "@every 2s", func() {
		err := task.UpdateSymbolsKlineLast()
		if err != nil {
			global.GVA_LOG.Error("更新K线数据失败", zap.Error(err))
		}
	}, "更新symbols的K线数据", option...)

	if err != nil {
		global.GVA_LOG.Error("添加K线数据更新定时任务失败", zap.Error(err))
	}
}

func addGeneratePlatformReportSummary(option []cron.Option) {
	_, err := global.GVA_Timer.AddTaskByFunc("GeneratePlatformReportSummary", "@every 1m", func() {
		err := task.GeneratePlatformReportSummary(global.GVA_DB)
		if err != nil {
			global.GVA_LOG.Error("更新平台固定报表数据失败", zap.Error(err))
		}
	}, "更新平台固定报表数据", option...)

	if err != nil {
		global.GVA_LOG.Error("更新平台固定报表数据任务添加失败", zap.Error(err))
	}
}

func addGeneratePlatformReportDaily(option []cron.Option) {
	_, err := global.GVA_Timer.AddTaskByFunc("GeneratePlatformReportDaily", "@every 1m", func() {
		err := task.GeneratePlatformReportDaily(global.GVA_DB)
		if err != nil {
			global.GVA_LOG.Error("更新平台日报表失败", zap.Error(err))
		}
	}, "更新平台日报表任务", option...)

	if err != nil {
		global.GVA_LOG.Error("更新平台日报表任务添加失败", zap.Error(err))
	}
}

func addRandomEarnProductRate(option []cron.Option) {
	_, err := global.GVA_Timer.AddTaskByFunc("RandomEarnProductRate", "@every 5m", func() {
		err := task.RandomProductRate(global.GVA_DB)
		if err != nil {
			global.GVA_LOG.Error("generate earn product daily rate err", zap.Error(err))
		}
	}, "RandomEarnProductRate", option...)

	if err != nil {
		global.GVA_LOG.Error("add random earn product rate failed", zap.Error(err))
	}
}

func doEarnProductDailyIncome(option []cron.Option) {
	_, err := global.GVA_Timer.AddTaskByFunc("DoEarnProductDailyIncome", "@every 5m", func() {
		err := task.DoEarnProductDailyIncome(context.Background())
		if err != nil {
			global.GVA_LOG.Error("generate DoEarnProductDailyIncome err", zap.Error(err))
		}
	}, "DoEarnProductDailyIncome", option...)

	if err != nil {
		global.GVA_LOG.Error("add do earn product daily income failed", zap.Error(err))
	}
}

func earnProductExpiration(option []cron.Option) {
	_, err := global.GVA_Timer.AddTaskByFunc("EarnProductExpiration", "@every 1s", func() {
		err := task.EarnProductExpiration(context.Background())
		if err != nil {
			global.GVA_LOG.Error("generate EarnProductExpiration err", zap.Error(err))
		}
	}, "EarnProductExpiration", option...)

	if err != nil {
		global.GVA_LOG.Error("add earn product expiration failed", zap.Error(err))
	}
}

// func addUpdateWebSocketSubscriptionsTask(option []cron.Option) {
// 	_, err := global.GVA_Timer.AddTaskByFunc("UpdateWebSocketSubscriptions", "@every 1m", func() {
// 		err := task.UpdateWebSocketSubscriptions(context.Background())
// 		if err != nil {
// 			global.GVA_LOG.Error("更新WebSocket订阅失败", zap.Error(err))
// 		}
// 	}, "更新WebSocket订阅列表", option...)

// 	if err != nil {
// 		global.GVA_LOG.Error("添加WebSocket订阅更新定时任务失败", zap.Error(err))
// 	}
// }

func addUpdateSymbolsKlineHistoryTask(option []cron.Option) {
	_, err := global.GVA_Timer.AddTaskByFunc("UpdateSymbolsKlineHistory", "@every 1m", func() {
		err := task.UpdateSymbolsKlineHistory()
		if err != nil {
			global.GVA_LOG.Error("更新历史K线数据失败", zap.Error(err))
		}
	}, "更新symbols的历史K线数据", option...)

	if err != nil {
		global.GVA_LOG.Error("添加历史K线数据更新定时任务失败", zap.Error(err))
	}
}

func addFetchIpToRegionTask(option []cron.Option) {
	_, err := global.GVA_Timer.AddTaskByFunc("FetchIpToRegion", "@every 10s", func() {
		_ = task.FetchIpToRegion(context.Background())
	}, "获取ip地址信息", option...)

	if err != nil {
		global.GVA_LOG.Error("添加获取ip地址信息任务失败", zap.Error(err))
	}
}
