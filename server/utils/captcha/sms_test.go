package captcha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/constants"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
)

func Test_sms_Send(t *testing.T) {
	levels := global.GVA_CONFIG.Zap.Levels()
	length := len(levels)
	cores := make([]zapcore.Core, 0, length)
	global.GVA_LOG = zap.New(zapcore.NewTee(cores...))
	s := &sms{
		AppKey:    constants.SMSAppKey,
		AppSecret: constants.SMSAppSecret,
		AppCode:   constants.SMSAppCode,
	}
	target := ""
	msg := "23345478"
	err := s.Send(target, msg)
	t.Logf("%v", err)
}
