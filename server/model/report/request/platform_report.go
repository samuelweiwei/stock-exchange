package request

import "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"

// PlatformReportDailyPageQueryReq 每日报表分页查询参数
// @Description 每日报表分页查询参数
type PlatformReportDailyPageQueryReq struct {
	request.PageInfo
	ReportDateStart  *int64 `form:"reportDateStart"` //报表开始日期
	ReportDateEnd    *int64 `form:"reportDateEnd"`   //报表结束日期
	RootUserIdSearch *uint  `form:"rootUserId"`      //跟用户ID搜索
	RootUserId       uint   `form:"-"`               //根用户ID
}

// PlatformReportDailyGetReq 平台动态报表查询参数
// @Description 平台动态报表查询参数
type PlatformReportDailyGetReq struct {
	ReportDate       *int64 `form:"reportDate"` //查询日期
	RootUserIdSearch *uint  `form:"rootUserId"` //跟用户ID
	RootUserId       uint   `form:"-"`          //根用户ID
}
