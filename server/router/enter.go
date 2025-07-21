package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/contract"
	"github.com/flipped-aurora/gin-vue-admin/server/router/coupon"
	"github.com/flipped-aurora/gin-vue-admin/server/router/earn"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/i18n"
	"github.com/flipped-aurora/gin-vue-admin/server/router/order"
	"github.com/flipped-aurora/gin-vue-admin/server/router/report"
	"github.com/flipped-aurora/gin-vue-admin/server/router/settingManage"
	"github.com/flipped-aurora/gin-vue-admin/server/router/symbol"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/user"
	"github.com/flipped-aurora/gin-vue-admin/server/router/userfund"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System        system.RouterGroup
	Example       example.RouterGroup
	Order         order.RouterGroup
	Contract      contract.RouterGroup
	User          user.RouterGroup
	Userfund      userfund.RouterGroup
	Coupon        coupon.RouterGroup
	I18n          i18n.RouterGroup
	Symbol        symbol.RouterGroup
	Report        report.RouterGroup
	Earn          earn.RouterGroup
	SettingManage settingManage.RouterGroup
}
