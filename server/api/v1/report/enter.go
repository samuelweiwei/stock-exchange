package report

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	PlatformReportApi
}

var platformReportService = service.ServiceGroupApp.PlatformReportServiceGroup.PlatformReportService
