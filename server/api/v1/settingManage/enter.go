package settingManage

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ ServiceLinkApi }

var serviceLinkService = service.ServiceGroupApp.SettingManageServiceGroup.ServiceLinkService
