package interceptor

import (
	"fmt"
	"github.com/aogoogle/common/model/response"
	"github.com/aogoogle/common/utils"
	"github.com/aogoogle/common/utils/encrypt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)

type SignHeader struct {
	Sign     string `header:"g-sign" binding:"required"`		//签名
	Time     string `header:"g-time" binding:"required"`		//访问端当前时间 (2分钟内)
	Version  string `header:"g-ver" binding:"required"`			//版本
	ClientOS string `header:"g-clientOs" binding:"required"`	//访问者操作系统，android,ios,web
}

// SignInterceptor
// @Description: 参数签名拦截校验
// @Date 2021-05-23 21:54:10
// @param interval
// @param signKey
// @return gin.HandlerFunc
func SignInterceptor(interval int64, signKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//fmt.Println("检查头参数......")
		//for k, v := range c.Request.Header {
		//	fmt.Println("header key->", k, "value->", v)
		//}
		//校验头公共参数
		var header SignHeader
		if err := c.ShouldBindHeader(&header); err != nil {
			response.FailWithParamsInvalid(c)
			c.Abort()
			return
		}

		//校验请求时间是否过期
		reqTime, err := strconv.ParseInt(header.Time, 10, 64)
		if err != nil {
			response.FailWithMessage("时间参数格式异常，无法通过校验", c)
			c.Abort()
			return
		}

		if time.Now().Unix()-reqTime > interval*1000 {
			response.FailWithMessage("请求时效已过，请检查本地时间", c)
			c.Abort()
			return
		}

		//校验签名
		_ = c.Request.ParseForm()
		params := make(map[string]interface{})
		params["g-time"] = reqTime
		for k, v := range c.Request.Form {
			params[k] = v[0]
			//fmt.Println("key->", k, "value->", v)
		}
		path := c.Request.RequestURI
		pos := strings.Index(path, "?")
		if pos >= 0 {
			path = path[0 : pos]
		}
		path += "?" + utils.BuildMapToSortString(params)
		sign := encrypt.HmacSha1(path, signKey)
		fmt.Println("appSign:", header.Sign, "\nsign:", sign)
		if sign != header.Sign {
			response.FailWithMessage("签名校验失败", c)
			c.Abort()
			return
		}
		c.Next()
	}
}
