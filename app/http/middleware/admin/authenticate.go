package middleware

import (
	"github.com/gin-gonic/gin"
	"hertz/app/http/controllers"
	"hertz/app/logic/admin"
	"net/http"
)

//登录验证器
//author hihozhou
//todo 判断是否html还是ajax请求，分别处理
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := c.Cookie(admin.AUTH_TOKEN_KEY)
		if token == "" {
			//判断是否html还是ajax请求，分别处理
			if (c.GetHeader("X-Requested-With") == "XMLHttpRequest") {
				//无效token，没有登录，应该跳转去登录
				controllers.Unauthorized(c, 1, "")
			}
			//跳转到登录页面
			c.Redirect(http.StatusMovedPermanently, admin.LOGIN_PATH)
			return
		}
		// parseToken 解析token包含的信息
		jwt := admin.NewJWT()
		claims, err := jwt.ParseToken(token)
		if err != nil {
			if err == admin.TokenExpired {
				//授权过期，
				//todo  应该刷新token
				controllers.Unauthorized(c, 2, "授权已过期")
				return
			}
			controllers.Unauthorized(c, 3, err.Error())
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		//todo 考虑存储用户数据或直接储存
		c.Set("claims", claims)
	}
}
