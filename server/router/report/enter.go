package report

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	PlatformReportRouter
}

var platformReportApi = api.ApiGroupApp.ReportApiGroup.PlatformReportApi
