package response

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/errorx"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR         = 7
	SUCCESS       = 0
	UPDATE_FAILED = 10
)

const (
	ResKey           = "ResKey"
	ResInnerErrorKey = "ResInnerErrorKey"
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	var (
		response = Response{
			code,
			data,
			msg,
		}
	)
	c.JSON(http.StatusOK, response)
	c.Set(ResKey, response)
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithCodeAndMessage(code int, message string, c *gin.Context) {
	Result(code, map[string]interface{}{}, message, c)
}

func NoAuth(message string, c *gin.Context) {
	response := Response{
		7,
		nil,
		message,
	}
	c.JSON(http.StatusUnauthorized, response)
	c.Set(ResKey, response)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}

func GetInnerErrorMsg(c *gin.Context) (errorMsg string) {

	for _, err := range c.Errors {
		if !errors.As(err.Err, &errorx.CodeError{}) {
			errorMsg += err.Error()
			if !errors.Is(err, c.Errors.Last()) {
				errorMsg += " | "
			}
		}
	}
	return
}
