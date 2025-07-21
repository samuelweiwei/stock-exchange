package client

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/constants"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
	"io"
	"net/http"
)

func InitUserSymbols(token string) {
	// 定义请求的 URL
	url := global.GVA_CONFIG.Service.Symbol + constants.InitCustomTradingPairsUrl

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		global.GVA_LOG.Error("创建请求失败", zap.Error(err))
		return
	}

	// 设置请求头 'x-token'
	req.Header.Set("x-token", token)

	// 发送请求
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		global.GVA_LOG.Error("发送请求失败", zap.Error(err))
		return
	}
	defer resp.Body.Close()

	// 读取并记录响应体
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		global.GVA_LOG.Error("读取响应体失败", zap.Error(err))
		return
	}
	global.GVA_LOG.Info("Response",
		zap.Int("statusCode", resp.StatusCode),
		zap.String("body", string(respBody)),
	)

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		global.GVA_LOG.Error("请求失败",
			zap.Int("statusCode", resp.StatusCode),
			zap.String("body", string(respBody)),
		)
		return
	}
	global.GVA_LOG.Info("InitUserSymbols 请求成功")
}

func InitUserSymbolsFromAdmin(token string, userID uint) {
	// 定义请求的 URL
	url := global.GVA_CONFIG.Service.Symbol + fmt.Sprintf(constants.InitCustomTradingPairsAdminUrl, userID)

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		global.GVA_LOG.Error("创建请求失败", zap.Error(err))
		return
	}

	// 设置请求头 'x-token'
	req.Header.Set("x-token", token)

	// 发送请求
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		global.GVA_LOG.Error("发送请求失败", zap.Error(err))
		return
	}
	defer resp.Body.Close()

	// 读取并记录响应体
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		global.GVA_LOG.Error("读取响应体失败", zap.Error(err))
		return
	}
	global.GVA_LOG.Info("Response",
		zap.Int("statusCode", resp.StatusCode),
		zap.String("body", string(respBody)),
	)

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		global.GVA_LOG.Error("请求失败",
			zap.Int("statusCode", resp.StatusCode),
			zap.String("body", string(respBody)),
		)
		return
	}
	global.GVA_LOG.Info("InitUserSymbols 请求成功")
}
