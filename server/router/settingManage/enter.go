package settingManage

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ ServiceLinkRouter }

var serviceLinkApi = api.ApiGroupApp.SettingManageApiGroup.ServiceLinkApi
