package interceptor

import (
	"github.com/aogoogle/common/model/response"
	"github.com/aogoogle/common/utils"
	"github.com/aogoogle/common/utils/jstring"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)

type ApiHeader struct {
	CheckSign 	string `header:"g-checkSign" binding:"required"`
	AppSign 	string `header:"g-sign" binding:"required"`
	AppKey 		string `header:"g-key" binding:"required"`
	ReqTime 	string `header:"g-time" binding:"required"`
	Verify 		string `header:"g-ver" binding:"required"`
}

func OpenApiInterceptor(accounts string, interval int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		//校验头参数
		var header ApiHeader
		if err := c.ShouldBindHeader(&header); err != nil {
			response.FailWithDetailed("", "请求参数校验错误，请重新登陆", c)
			c.Abort()
			return
		}

		//检验appKey是否合法
		keys := strings.Split(accounts, ",")

		if !jstring.IsContainStr(keys, header.AppKey) {
			response.FailWithMessage("无经授权的非法请求", c)
			c.Abort()
			return
		}

		//判断g-time是否超时(当前-g_time > time.interval)
		reqTime, err := strconv.ParseInt(header.ReqTime, 10, 64)
		if err != nil {
			response.FailWithMessage("请求参数校验错误，请重新登陆", c)
			c.Abort()
			return
		}

		if err == nil {
			if time.Now().Unix()-reqTime > interval*1000 {
				response.FailWithMessage("请求时效已过，请检查本地时间", c)
				c.Abort()
				return
			}
		}

		//判断g-sign签名串是否正确
		params := make(map[string]interface{})
		str := c.Request.RequestURI + "?"
		params["g-key"] = header.AppKey
		params["g-time"] = header.ReqTime
		params["g-ver"] = header.Verify

		str += utils.BuildMapToSortString(params)
		if str != header.AppSign {
			response.FailWithMessage("签名校验失败", c)
			c.Abort()
			return
		}
		c.Next()
	}
}