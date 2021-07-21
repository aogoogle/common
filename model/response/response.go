package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"status"`
	Data interface{} `json:"content"`
	Msg  string      `json:"msg"`
}

const (
	ERROR     		= 500
	SUCCESS   		= 0
	TOKENEXPIRED 	= 6001
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Success(c *gin.Context) {
	Result(SUCCESS, "", "操作成功", c)
}

func SuccessWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, "", message, c)
}

func SuccessWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func SuccessWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, "", "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, "", message, c)
	c.Abort()
}

func FailWithParamsInvalid(c *gin.Context) {
	Result(ERROR, "", "参数验证失败", c)
	c.Abort()
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}
