package i18n

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	SysI18nLocalizeConfigApi
}

var (
	sysI18nLocalizeConfigService = service.ServiceGroupApp.I18nServiceGroup.SysI18nLocalizeConfigService
)
