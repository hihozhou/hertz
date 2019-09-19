package middleware

import (
	"github.com/gin-gonic/gin"
	"hertz/app/logic/admin"
	"net/http"
)

//验证后自动跳转中间件
func RedirectIfAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := c.Cookie(admin.AUTH_TOKEN_KEY)
		if token == "" {
			return
		}
		jwt := admin.NewJWT()
		_, err := jwt.ParseToken(token)
		if err != nil {
			if err == admin.TokenExpired {
				return
			}
			return
		}
		c.Redirect(http.StatusMovedPermanently, admin.HOME_PATH)
	}
}
