/**
* @Author: Jackey
* @Date: 12/24/24 4:40 pm
 */

package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"strings"
)

var ServerList []string

const (
	InterfaceServer string = "数据接口"
	TaskServer      string = "定时任务"
	WebsocketServer string = "获取三方数据websocket"
)

func PostServerInfo() {
	//判断启动的服务模式
	tgMsg := "【服务启动】\n[服务类型]："
	if len(ServerList) > 0 {
		tgMsg += strings.Join(ServerList, " | ")
		utils.SendMsgToTgAsync(tgMsg)
	} else {
		tgMsg += "无（即将结束程序）"
		utils.SendMsgToTg(tgMsg)
	}
}
