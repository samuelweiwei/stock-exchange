package user

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	CountriesApi
	FrontendUsersApi
	FrontendUserLoginLogApi
}

var (
	countriesService            = service.ServiceGroupApp.UserServiceGroup.CountriesService
	frontendUsersService        = service.ServiceGroupApp.UserServiceGroup.FrontendUsersService
	jwtService                  = service.ServiceGroupApp.SystemServiceGroup.JwtService
	userfundAccountService      = service.ServiceGroupApp.UserfundServiceGroup.UserFundAccountsService
	frontendUserLoginLogService = service.ServiceGroupApp.UserServiceGroup.FrontendUserLoginLogService
)
