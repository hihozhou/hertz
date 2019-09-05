package routes

import (
	"github.com/gin-gonic/gin"
)

func CreateRoutes(route *gin.Engine) {

	route.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"err_code":    0,
			"err_message": "index page！",
		})
	})

	route.GET("/index", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"err_code":    0,
			"err_message": "index page！",
		})
	})

}
