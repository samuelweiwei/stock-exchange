package middleware

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/i18n"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/cast"
	"strconv"
	"time"
)

func JWTAuthFrontUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := utils.GetTokenFrontUser(c)
		if token == "" {
			response.NoAuth(i18n.Message(request.GetLanguageTag(c), i18n.JWTNotLogin, response.ERROR), c)
			c.Abort()
			return
		}
		if jwtService.IsBlacklist(token) {
			response.NoAuth(i18n.Message(request.GetLanguageTag(c), i18n.JWTRepeat, response.ERROR), c)
			utils.ClearTokenFrontUser(c)
			c.Abort()
			return
		}
		j := utils.NewJWT(&global.GVA_CONFIG.JWTUser)
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, utils.TokenExpired) {
				response.NoAuth(i18n.Message(request.GetLanguageTag(c), i18n.JWTExpire, response.ERROR), c)
				utils.ClearTokenFrontUser(c)
				c.Abort()
				return
			}
			response.NoAuth(i18n.Message(request.GetLanguageTag(c), i18n.JWTExpire, response.ERROR), c)
			utils.ClearTokenFrontUser(c)
			c.Abort()
			return
		}

		// 已登录用户被管理员禁用 需要使该用户的jwt失效 此处比较消耗性能 如果需要 请自行打开
		// 用户被删除的逻辑 需要优化 此处比较消耗性能 如果需要 请自行打开

		if user, err := frontUserService.GetFrontendUsers(cast.ToString(claims.BaseClaims.ID)); err != nil || user.Enable == 2 {
			_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: token})
			err = fmt.Errorf("forbid")
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
		}
		c.Set("claims", claims)
		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
			dr, _ := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(dr))
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
			utils.SetToken(c, newToken, int(dr.Seconds()))
			if global.GVA_CONFIG.System.UseMultipoint {
				// 记录新的活跃jwt
				_ = jwtService.SetRedisJWT(newToken, newClaims.Username)
			}
		}
		c.Next()

		if newToken, exists := c.Get("new-token"); exists {
			c.Header("new-token", newToken.(string))
		}
		if newExpiresAt, exists := c.Get("new-expires-at"); exists {
			c.Header("new-expires-at", newExpiresAt.(string))
		}
	}
}
