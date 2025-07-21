package task

import (
	"fmt"
	"net/http"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
)

// KlineData 定义K线数据结构
type KlineData struct {
	Open   float64 `json:"o"`
	High   float64 `json:"h"`
	Low    float64 `json:"l"`
	Close  float64 `json:"c"`
	Volume float64 `json:"v"`
	Time   int64   `json:"t"`
}

// parseInterval 将时间间隔转换为Polygon支持的格式
func parseInterval(interval string) (multiplier int, timespan string, err error) {
	switch interval {
	case "1m":
		return 1, "minute", nil
	case "1h":
		return 1, "hour", nil
	case "4h":
		return 4, "hour", nil
	case "1d":
		return 1, "day", nil
	case "1w":
		return 1, "week", nil
	case "1M":
		return 1, "month", nil
	default:
		return 0, "", fmt.Errorf("unsupported interval: %s", interval)
	}
}

func makeRequest(client *http.Client, url string, maxRetries int) (*http.Response, error) {
	var resp *http.Response
	var err error

	for i := 0; i < maxRetries; i++ {
		resp, err = client.Get(url)
		if err == nil {
			return resp, nil
		}

		global.GVA_LOG.Warn("请求失败，准备重试",
			zap.String("url", url),
			zap.Error(err),
			zap.Int("retry", i+1))

		// 在重试之前等待一段时间，时间随重试次数增加
		time.Sleep(time.Second * time.Duration(i+1))
	}

	return nil, err
}
