package admin

import (
	"github.com/gin-gonic/gin"
	"hertz/config"
	"net/http"
)

type IndexController struct {
}

// 首页
func (index IndexController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index/index.html", gin.H{
		"menus": config.Menus,
	})
}
