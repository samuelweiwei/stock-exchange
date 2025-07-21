/**
* @Author: Jackey
* @Date: 12/25/24 1:45 pm
 */

package utils

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"net/http"
	"net/url"
	"time"
)

func SendMsgToTgAsync(msg string) {
	go func() {
		SendMsgToTg(msg)
	}()
}
func SendMsgToTg(msg string) {

	var (
		host        = global.GVA_CONFIG.Tg.Host
		apiToken    = global.GVA_CONFIG.Tg.Token
		path        = global.GVA_CONFIG.Tg.Path
		nodeName    = global.GVA_CONFIG.Tg.NodeName
		nodeVersion = global.GVA_CONFIG.Tg.NodeVersion
		chatId      = global.GVA_CONFIG.Tg.ChatId
		charIdKey   = "chat_id"
		textKey     = "text"
		params      = url.Values{}
		err         error
		urlModel    *url.URL
	)

	defer func() {
		if err != nil {
			global.GVA_LOG.Error("【tg发送群消息功能出错】：" + err.Error())
			global.GVA_LOG.Error(msg)
		}
	}()

	if len(chatId) == 0 {
		err = errors.New("没配置tg:chat-id")
		return
	}

	if len(nodeName) > 1 {
		msg += fmt.Sprintf("\n[节点信息]：%s", nodeName)
		if len(nodeVersion) > 0 {
			msg += fmt.Sprintf("->%s", nodeVersion)
		}
	}
	msg = msg + "\n[北京时间]：" + getBeijingTimeStringByUnix(time.Now().Unix())
	params.Set(charIdKey, chatId)
	params.Set(textKey, msg)

	urlModel, err = url.Parse(host + apiToken + path)
	if err != nil {
		return
	}
	urlModel.RawQuery = params.Encode()

	var (
		urlPath    = urlModel.String()
		httpClient = http.Client{Timeout: 10 * time.Second}
	)
	_, err = httpClient.Get(urlPath)
}

func getBeijingTimeStringByUnix(unix int64) string {
	var (
		location, _ = time.LoadLocation("Asia/Shanghai")
		beijingTime = time.Unix(unix, 0).In(location)
		timeString  = beijingTime.Format(time.DateTime)
	)
	return timeString
}
