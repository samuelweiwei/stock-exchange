package user

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	CountriesRouter
	FrontendUsersRouter
	FrontendUserLoginLogRouter
}

var (
	countriesApi            = api.ApiGroupApp.UserApiGroup.CountriesApi
	frontendUsersApi        = api.ApiGroupApp.UserApiGroup.FrontendUsersApi
	frontendUserLoginLogApi = api.ApiGroupApp.UserApiGroup.FrontendUserLoginLogApi
)
