package routes

import (
	"github.com/gin-gonic/gin"
	"hertz/app/http/controllers"
	"hertz/app/http/controllers/admin"
	"hertz/app/http/middleware"
	adminMiddleware "hertz/app/http/middleware/admin"
	adminModule "hertz/app/logic/admin"
	"net/http"
)

//中间件，日志，接口鉴权，错误处理
func CreateRoutes(router *gin.Engine) {

	//加载模板页面
	router.LoadHTMLGlob("resources/views/*/*/*")
	//加载静态文件
	router.StaticFS("/static", http.Dir("./resources/static"))

	//接口访问nocache
	router.Use(middleware.NoCache)

	indexController := &controllers.IndexController{}
	router.GET("/", indexController.Index)
	router.POST("/login", indexController.Login)
	router.GET("/home", middleware.Authenticate(), indexController.Home)

	authController := &admin.AuthController{}
	//router.GET(adminModule.LOGIN_PATH, adminMiddleware.RedirectIfAuthenticated(), authController.LoginForm)
	router.GET(adminModule.LOGIN_PATH, authController.LoginForm)
	router.POST("/admin/login", authController.Login)
	router.GET("/admin/logout", authController.Logout)

	//后台路由
	adminGroup := router.Group("/admin", adminMiddleware.Authenticate())
	{
		adminIndex := &admin.IndexController{}

		adminGroup.GET("/", adminIndex.Index)
		adminGroup.GET("/index/console", adminIndex.Console)

		adminController := &admin.AdminController{}
		adminGroup.GET("/admin", adminController.Index)
		adminGroup.GET("/admin/index", adminController.Index)

	}

}
