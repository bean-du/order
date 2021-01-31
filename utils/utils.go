package utils

import "github.com/gin-gonic/gin"

func httpBaseResponse(code int, msg string, data interface{}) gin.H {
	return gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	}
}

func HttpResponseOk(c *gin.Context) {
	c.JSON(200, httpBaseResponse(0, "ok", nil))
}

func HttpResponse(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(200, httpBaseResponse(code, msg, data))
}
