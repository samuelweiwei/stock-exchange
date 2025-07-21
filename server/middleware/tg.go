/**
* @Author: Jackey
* @Date: 12/16/24 12:32 pm
 */

package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type TgRouterType int

const (
	TgRouterPublic TgRouterType = 0
	TgRouterAdmin  TgRouterType = 1
	TgRouterFront  TgRouterType = 2
)

func TgInterfaceNotify(routerType TgRouterType) gin.HandlerFunc {

	return func(c *gin.Context) {

		bodyBytes, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		c.Next()

		var (
			value, _           = c.Get(response.ResKey)
			responseInfo, isOk = value.(response.Response)
			innerErrMsg        = response.GetInnerErrorMsg(c)
		)
		if !isOk || responseInfo.Code == response.SUCCESS {
			return
		}
		var (
			clientIp           = c.ClientIP()
			host               = c.Request.Host
			path               = c.Request.URL.Path
			method             = c.Request.Method
			errMsg, userString string
			bodyString         = string(bodyBytes)
		)
		if method == http.MethodGet {
			queryBytes, err := json.Marshal(c.Request.URL.Query())
			if err != nil {
				path = c.Request.RequestURI
			} else {
				bodyString = string(queryBytes)
			}
		}
		if len(innerErrMsg) > 0 {
			innerErrMsg = fmt.Sprintf("\n[内部错误]：%s", innerErrMsg)
		}
		switch routerType {
		case TgRouterPublic:
			errMsg = "\n【公共接口】->报错"
		case TgRouterAdmin:
			errMsg = "\n【后台接口】->报错"
			userString = fmt.Sprintf("\n[用户id]：%d", utils.GetUserID(c))
		case TgRouterFront:
			errMsg = "\n【前台接口】->报错"
			userString = fmt.Sprintf("\n[会员id]：%d", utils.GetUserIDFrontUser(c))
		}
		errMsg = errMsg +
			"\n[报错信息]：" + responseInfo.Msg + innerErrMsg +
			"\n[错误码]：" + strconv.Itoa(responseInfo.Code) +
			"\n[客户端ip]：" + clientIp +
			"\n[主机]：" + host +
			"\n[路由]：" + path + userString +
			"\n[请求参数]：" + bodyString
		asyncSendInterfaceError(errMsg)
	}
}

func asyncSendInterfaceError(msg string) {

	go func() {
		if global.GVA_CONFIG.Tg.IsSend == false {
			return
		}

		var (
			routerList     = global.GVA_CONFIG.Tg.IgnoreRouter
			isNeedControls = false
			router         string
		)

		if len(routerList) > 0 {
			for _, r := range routerList {
				if strings.Contains(msg, r) {
					router = r
					isNeedControls = true
					break
				}
			}
		}

		if isNeedControls && !controlsRouterCanSend(router) {
			return
		}

		utils.SendMsgToTg(msg)
	}()
}

var (
	tgSendInterfaceTimes sync.Map
	queueMutex           sync.Mutex
)

func controlsRouterCanSend(router string) (canSend bool) {
	queueMutex.Lock()
	defer queueMutex.Unlock()
	var (
		now              = time.Now()
		lastCallTime, ok = tgSendInterfaceTimes.Load(router)
		minSeconds       = global.GVA_CONFIG.Tg.RouterSendDuration
	)

	if ok && now.Sub(lastCallTime.(time.Time)).Seconds() < minSeconds {
		canSend = false
	} else {
		canSend = true
		tgSendInterfaceTimes.Store(router, now)
	}
	return canSend
}
