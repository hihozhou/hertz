package routes

import (
	"github.com/gin-gonic/gin"
	"hertz/app/http/controllers"
	"hertz/app/http/middleware"
)

//中间件，日志，接口鉴权，错误处理
func CreateRoutes(route *gin.Engine) {

	index := controllers.Index{}
	route.GET("/", index.Index)
	route.POST("/login", index.Login)
	route.GET("/home", middleware.Authenticate(), index.Home)

}
