package handler

import (
	"com.phpstu/webapp/src/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

//响应json格式
func SendJson(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
	c.Abort()
}

//响应jsonp格式
func SendJsonp(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)
	c.JSONP(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
	c.Abort()
}
