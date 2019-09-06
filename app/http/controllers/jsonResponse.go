package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//todo data考虑用指针
//json返回
func Response(c *gin.Context, status int, errorCode int, errorMsg string, data gin.H) {

	result := gin.H{
		"error_code": errorCode,
		"error_msg":  errorMsg,
	}
	//合并附加数据
	for k, v := range data {
		result[k] = v;
	}

	c.JSON(status, result)
}

//成功返回
func Success(c *gin.Context, data gin.H) {
	Response(c, http.StatusOK, 0, "", data)
}

//失败返回
func Fail(c *gin.Context, errorMsg string, data gin.H) {
	Response(c, http.StatusOK, 1, errorMsg, data)
	c.Abort()
}

//
func Unauthorized(c *gin.Context, errorCode int, extraMessage string) {
	Response(c, http.StatusUnauthorized, errorCode, "Request Unauthorized", gin.H{"extra_message": extraMessage})
	c.Abort()
}

//todo catch 500 错误，写入日志
