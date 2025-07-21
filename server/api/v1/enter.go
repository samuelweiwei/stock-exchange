package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/contract"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/coupon"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/earn"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/i18n"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/order"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/report"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/settingManage"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/symbol"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/user"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/userFund"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup        system.ApiGroup
	ExampleApiGroup       example.ApiGroup
	OrderApiGroup         order.ApiGroup
	ContractApiGroup      contract.ApiGroup
	UserApiGroup          user.ApiGroup
	UserfundApiGroup      userFund.ApiGroup
	SymbolApiGroup        symbol.ApiGroup
	CouponApiGroup        coupon.ApiGroup
	I18nApiGroup          i18n.ApiGroup
	ReportApiGroup        report.ApiGroup
	EarnApiGroup          earn.ApiGroup
	SettingManageApiGroup settingManage.ApiGroup
}
