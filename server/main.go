package main

import (
	"context"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/i18n"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title                       Gin-Vue-Admin Swagger API接口文档
// @version                     v2.7.6
// @description                 使用gin+vue进行极速开发的全栈开发基础平台
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	global.GVA_VP = core.Viper() // 初始化Viper
	initialize.OtherInit()
	global.GVA_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库

	// 1. 先初始化 Redis
	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		initialize.Redis()
		initialize.RedisList()
	}

	// 2. 再初始化定时任务
	if global.GVA_CONFIG.System.UseTask {
		initialize.ServerList = append(initialize.ServerList, initialize.TaskServer)
		initialize.Timer()
	}

	initialize.DBList()
	initialize.BusinessCalendar() //初始化商业日历
	i18n.StartUp()
	if global.GVA_DB != nil {
		//initialize.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}

	if global.GVA_CONFIG.System.UseWs {
		initialize.ServerList = append(initialize.ServerList, initialize.WebsocketServer)
		// 在一个新的 goroutine 中初始化 WebSocket
		go func() {
			// 等待其他服务启动
			time.Sleep(5 * time.Second)
			initialize.InitSymbolsPricesData()

			global.GVA_LOG.Info("Starting WebSocket initialization...")
			if err := initialize.WebsocketInit(context.Background()); err != nil {
				global.GVA_LOG.Error("Failed to initialize websocket", zap.Error(err))
			}
		}()
	}

	if global.GVA_CONFIG.System.UseInterface {
		//有启动http服务
		initialize.ServerList = append(initialize.ServerList, initialize.InterfaceServer)
		initialize.PostServerInfo()
		core.RunWindowsServer()
	} else {
		//没启动http服务，需要判断是否有启动其他服务，有启动的话hold住进程
		initialize.PostServerInfo()
		if len(initialize.ServerList) > 0 {
			select {}
		}
	}
}
