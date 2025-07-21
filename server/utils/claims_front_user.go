package utils

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
	"net"
	"time"
)

func ClearTokenFrontUser(c *gin.Context) {
	// 增加cookie x-token-front 向来源的web添加
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}

	if net.ParseIP(host) != nil {
		c.SetCookie("x-token-front", "", -1, "/", "", false, false)
	} else {
		c.SetCookie("x-token-front", "", -1, "/", host, false, false)
	}
}

func SetTokenFrontUser(c *gin.Context, token string, maxAge int) {
	// 增加cookie x-token-front 向来源的web添加
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}

	if net.ParseIP(host) != nil {
		c.SetCookie("x-token-front", token, maxAge, "/", "", false, false)
	} else {
		c.SetCookie("x-token-front", token, maxAge, "/", host, false, false)
	}
}

func GetTokenFrontUser(c *gin.Context) string {
	token, _ := c.Cookie("x-token-front")
	if token == "" {
		j := NewJWT(&global.GVA_CONFIG.JWTUser)
		token = c.Request.Header.Get("x-token-front")
		claims, err := j.ParseToken(token)
		if err != nil {
			global.GVA_LOG.Error("重新写入cookie token失败,未能成功解析token,请检查请求头是否存在x-token-front且claims是否为规定结构")
			return token
		}
		SetTokenFrontUser(c, token, int((claims.ExpiresAt.Unix()-time.Now().Unix())/60))
	}
	return token
}

func GetClaimsFrontUser(c *gin.Context) (*systemReq.CustomClaims, error) {
	token := GetTokenFrontUser(c)
	j := NewJWT(&global.GVA_CONFIG.JWTUser)
	claims, err := j.ParseToken(token)
	if err != nil {
		//global.GVA_LOG.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token-front且claims是否为规定结构")
	}
	return claims, err
}

// GetUserIDFrontUser 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserIDFrontUser(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaimsFrontUser(c); err != nil {
			return 0
		} else {
			return cl.BaseClaims.ID
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.BaseClaims.ID
	}
}

// GetUserUuidFrontUser 从Gin的Context中获取从jwt解析出来的用户UUID
func GetUserUuidFrontUser(c *gin.Context) uuid.UUID {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaimsFrontUser(c); err != nil {
			return uuid.UUID{}
		} else {
			return cl.UUID
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.UUID
	}
}

// GetUserInfoFrontUser 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserInfoFrontUser(c *gin.Context) *systemReq.CustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaimsFrontUser(c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse
	}
}

// GetUserNameFrontUser 从Gin的Context中获取从jwt解析出来的用户名
func GetUserNameFrontUser(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaimsFrontUser(c); err != nil {
			return ""
		} else {
			return cl.Username
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.Username
	}
}

func LoginTokenFrontUser(user system.Login) (token string, claims systemReq.CustomClaims, err error) {
	j := &JWT{SigningKey: []byte(global.GVA_CONFIG.JWTUser.SigningKey)} // 唯一签名
	claims = j.CreateClaims(systemReq.BaseClaims{
		UUID:        user.GetUUID(),
		ID:          user.GetUserId(),
		NickName:    user.GetNickname(),
		Username:    user.GetUsername(),
		AuthorityId: user.GetAuthorityId(),
		UserType:    user.GetUserType(),
	})
	token, err = j.CreateToken(claims)
	if err != nil {
		return
	}
	return
}
