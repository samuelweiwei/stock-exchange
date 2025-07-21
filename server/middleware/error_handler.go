package middleware

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/errorx"
	"github.com/flipped-aurora/gin-vue-admin/server/i18n"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		var (
			lan = request.GetLanguageTag(c)
		)
		for _, err := range c.Errors {
			var codeError errorx.CodeError
			if errors.As(err.Err, &codeError) {
				response.FailWithCodeAndMessage(codeError.Code, i18n.Message(lan, err.Err.Error(), uint64(codeError.Code), codeError.Args...), c)
				return
			} else {
				response.FailWithCodeAndMessage(errorx.InternalServerError, i18n.Message(lan, "", uint64(errorx.InternalServerError)), c)
			}
		}
	}
}
