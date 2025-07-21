package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/captcha"
	"github.com/flipped-aurora/gin-vue-admin/server/service/contract"
	"github.com/flipped-aurora/gin-vue-admin/server/service/coupon"
	"github.com/flipped-aurora/gin-vue-admin/server/service/earn"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/i18n"
	"github.com/flipped-aurora/gin-vue-admin/server/service/order"
	"github.com/flipped-aurora/gin-vue-admin/server/service/report"
	"github.com/flipped-aurora/gin-vue-admin/server/service/settingManage"
	"github.com/flipped-aurora/gin-vue-admin/server/service/symbol"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/user"
	"github.com/flipped-aurora/gin-vue-admin/server/service/userfund"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup         system.ServiceGroup
	ExampleServiceGroup        example.ServiceGroup
	OrderServiceGroup          order.ServiceGroup
	ContractServiceGroup       contract.ServiceGroup
	UserServiceGroup           user.ServiceGroup
	UserfundServiceGroup       userfund.ServiceGroup
	SymbolServiceGroup         symbol.ServiceGroup
	CouponServiceGroup         coupon.ServiceGroup
	I18nServiceGroup           i18n.ServiceGroup
	PlatformReportServiceGroup report.ServiceGroup
	EarnServiceGroup           earn.ServiceGroup
	SettingManageServiceGroup  settingManage.ServiceGroup
	CaptchaServiceGroup        captcha.ServiceGroup
}
