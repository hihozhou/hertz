package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hertz/routes"
	"net/http"
	"hertz/pkg/setting"
)

func main() {
	//设置框架启动模式
	gin.SetMode(setting.RunMode)

	//获取路由
	route := gin.Default()

	routes.CreateRoutes(route)

	//配置服务
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        route,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	//服务启动
	server.ListenAndServe()
}
