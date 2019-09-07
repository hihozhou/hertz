package routes

import (
	"github.com/gin-gonic/gin"
	"hertz/app/http/controllers"
	"hertz/app/http/controllers/admin"
	"hertz/app/http/middleware"
	"net/http"
)

//中间件，日志，接口鉴权，错误处理
func CreateRoutes(router *gin.Engine) {

	//加载模板页面
	router.LoadHTMLGlob("resources/views/*/*/*")
	//加载静态文件
	router.StaticFS("/static", http.Dir("./resources/static"))

	indexController := &controllers.IndexController{}
	router.GET("/", indexController.Index)
	router.POST("/login", indexController.Login)
	router.GET("/home", middleware.Authenticate(), indexController.Home)

	//后台路由
	adminGroup := router.Group("/admin")
	{
		adminIndex := &admin.IndexController{}
		authController := &admin.AuthController{}
		adminGroup.GET("/login", authController.LoginForm)
		adminGroup.POST("/login", authController.Login)
		adminGroup.GET("/", adminIndex.Index)
	}

}
