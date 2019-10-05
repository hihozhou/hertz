package middleware

import (
	"github.com/gin-gonic/gin"
	"hertz/app/http/controllers"
	"hertz/app/logic/admin"
	"hertz/pkg/util"
	"net/http"
	"time"
)

//登录验证器
//author hihozhou
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := c.Cookie(admin.AUTH_TOKEN_KEY)
		if token == "" {
			//无效token，没有登录，应该跳转去登录
			UnauthorizedHandler(c, 1, "")
			return
		}
		// parseToken 解析token包含的信息
		jwt := admin.NewJWT()
		claims, err := jwt.ParseToken(token)
		if err != nil {
			if err == admin.TokenExpired {
				UnauthorizedHandler(c, 2, "授权已过期")
				return
			} else {
				UnauthorizedHandler(c, 3, err.Error())
				return
			}
		} else {
			//判断时间是否大于一小时
			if (time.Now().Unix() - claims.NotBefore) >= admin.TokenVerifyEffective {
				newToken, err := jwt.RefreshClaimsToken(claims)
				if err != nil {
					UnauthorizedHandler(c, 1, "刷新token失败："+err.Error())
				}
				//保存到cookie中
				admin.LoginSign(c, newToken)
			}
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		//todo 考虑存储用户数据或直接储存
		c.Set("claims", claims)
	}
}

//ajax返回或跳转到登录页
func UnauthorizedHandler(c *gin.Context, errorCode int, errorMsg string) {
	//判断是否html还是ajax请求，分别处理
	if util.IsAjax(c) {
		//判断是否html还是ajax请求，分别处理
		controllers.Unauthorized(c, errorCode, errorMsg)
	}
	//跳转到登录页面
	c.Redirect(http.StatusMovedPermanently, admin.LOGIN_PATH)
}
