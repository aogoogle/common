package interceptor

import (
	"github.com/aogoogle/common/model/response"
	"github.com/gin-gonic/gin"
	"sort"
)

/*
	域之间调用拦截
*/
func DomainInterceptor(allowIps []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIp := c.Request.RemoteAddr
		index := sort.SearchStrings(allowIps, clientIp)
		if index < len(allowIps) && allowIps[index] == clientIp {
			c.Next()
		} else {
			response.FailWithMessage("您被限制访问", c)
			c.Abort()
		}
	}
}