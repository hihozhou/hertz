package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hertz/app/http/controllers"
)

//登录验证器
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Request.Cookie("session_id"); err == nil {
			value := cookie.Value
			fmt.Println(value)
			if value == "123" {
				c.Next()
				return
			}
		}
		controllers.Unauthorized(c)
	}
}
