package i18n

import (
	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"go.uber.org/zap"
	"testing"
)

func _setup() {
	global.GVA_VP = core.Viper() // 初始化Viper
	initialize.OtherInit()
	global.GVA_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.DBList()
	defaultLocalization = NewLocalization()
}
func TestCodeMessage(t *testing.T) {
	_setup()
	t.Log(defaultLocalization.bundle.LanguageTags())
	//按消息ID 来获取语言对应的文本
	t.Log(Message("en", "HelloPerson", 0, "Name", "Jerry"))
	t.Log(Message("es", "HelloPerson", 0, "Name", "Jerry"))
	//按消息错误码 来获取语言对应的文本
	t.Log(Message("en", "HelloPerson", 0, "Name", "tom"))
	t.Log(Message("es", "HelloPerson", 0, "Name", "tom"))
}

func TestLocalizeMultiParam(t *testing.T) {
	_setup()
	//按消息错误码 来获取语言对应的文本
	t.Log(Message("en", "PersonUnreadEmails", 0, "Name", "tom", "UnreadEmailCount", 1))
	t.Log(Message("es", "PersonUnreadEmails", 0, "Name", "tom", "UnreadEmailCount", 2))
}
