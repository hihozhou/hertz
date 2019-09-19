package util

import "github.com/gin-gonic/gin"

//判断是会否是ajax请求
func IsAjax(c *gin.Context) bool {
	if (c.GetHeader("X-Requested-With") == "XMLHttpRequest") {
		return true
	}
	return false
}
