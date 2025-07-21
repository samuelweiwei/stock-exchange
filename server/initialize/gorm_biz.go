package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

func bizModel() error {
	Db := global.GetGlobalDBByDBName("")
	_ = Db.AutoMigrate()
	return nil
}
