package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CountriesRouter struct {}

// InitCountriesRouter 初始化 countries表 路由信息
func (s *CountriesRouter) InitCountriesRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	countriesRouter := Router.Group("countries").Use(middleware.OperationRecord())
	countriesRouterWithoutRecord := Router.Group("countries")
	countriesRouterWithoutAuth := PublicRouter.Group("countries")
	{
		countriesRouter.POST("createCountries", countriesApi.CreateCountries)   // 新建countries表
		countriesRouter.DELETE("deleteCountries", countriesApi.DeleteCountries) // 删除countries表
		countriesRouter.DELETE("deleteCountriesByIds", countriesApi.DeleteCountriesByIds) // 批量删除countries表
		countriesRouter.PUT("updateCountries", countriesApi.UpdateCountries)    // 更新countries表
	}
	{
		countriesRouterWithoutRecord.GET("findCountries", countriesApi.FindCountries)        // 根据ID获取countries表
		countriesRouterWithoutRecord.GET("getCountriesList", countriesApi.GetCountriesList)  // 获取countries表列表
	}
	{
	    countriesRouterWithoutAuth.GET("getCountriesPublic", countriesApi.GetCountriesPublic)  // countries表开放接口
	}
}
