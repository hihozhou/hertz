package middleware

import (
	"github.com/gin-gonic/gin"
	"hertz/app/http/controllers"
	"hertz/app/logic"
)

//登录验证器
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			//无效token，没有登录，应该跳转去登录
			controllers.Unauthorized(c, 1, "")
			return
		}
		// parseToken 解析token包含的信息
		jwt := logic.NewJWT()
		claims, err := jwt.ParseToken(token)
		if err != nil {
			if err == logic.TokenExpired {
				//授权过期，应该刷新token
				controllers.Unauthorized(c, 2, "授权已过期")
				return
			}
			controllers.Unauthorized(c, 3, err.Error())
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
	}
}
