package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {
}

// 首页
func (index IndexController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index/index.html", nil)
}
