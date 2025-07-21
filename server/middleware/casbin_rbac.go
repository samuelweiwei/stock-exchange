package middleware

import (
	"strconv"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
)

var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService

// CasbinHandler 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			waitUse, _ = utils.GetClaims(c)
			path       = c.Request.URL.Path //获取请求的PATH
			obj        = strings.TrimPrefix(path, global.GVA_CONFIG.System.RouterPrefix)
			act        = c.Request.Method                       // 获取请求方法
			sub        = strconv.Itoa(int(waitUse.AuthorityId)) // 获取用户的角色
			e          = casbinService.Casbin()                 // 判断策略中是否存在
			success, _ = e.Enforce(sub, obj, act)
		)
		if !success {
			response.FailWithDetailed(gin.H{}, "权限不足", c)
			c.Abort()
			return
		}
		c.Next()
	}
}
