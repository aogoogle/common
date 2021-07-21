package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponsePage struct {
	Response
	Total      int  `json:"total"`
	Pages      int  `json:"pages"`
	PageSize   int  `json:"pageSize"`
	PageIndex  int  `json:"pageIndex"`
	IsLastPage bool `json:"lastPage"`
}

func ResultPage(code int, content interface{}, total, pageIndex, pageSize int, msg string, c *gin.Context) {
	pages := 0
	if pageSize > 0 {
		pages = total/pageSize
		if (total % pageSize) != 0 {
			pages += 1
		}
	}
	c.JSON(http.StatusOK, ResponsePage{
		Total:      total,
		PageSize:   pageSize,
		PageIndex:  pageIndex,
		Pages:      pages,
		IsLastPage: pageIndex >= pages,
		Response: Response {
			Data: content, Code: code, Msg:  msg,
		},
	})
}

func ResultSuccessWithData(data interface{}, total, pageIndex, pageSize int, c *gin.Context) {
	ResultPage(SUCCESS, data, total, pageIndex, pageSize, "操作成功", c)
}

func ResultFailedWithMessage(msg string, c *gin.Context) {
	ResultPage(ERROR, nil, 0, 0, 0, msg, c)
	c.Abort()
}

func ResultFailedWithData(data interface{}, msg string, c *gin.Context) {
	ResultPage(ERROR, data, 0, 0, 0, msg, c)
	c.Abort()
}

func PageFailedWithParamsInvalid(c *gin.Context) {
	ResultFailedWithData([][]int{}, "参数校验失败", c)
}

func PageFailedWithMessage(msg string, c *gin.Context) {
	ResultFailedWithData([][]int{}, msg, c)
}

func PageSuccessWithData(data interface{}, total, pageIndex, pageSize int, c *gin.Context) {
	if data == nil || total == 0 {
		data = [][]int{}
	}
	ResultSuccessWithData(data, total, pageIndex, pageSize, c)
}