package admin

import (
	"github.com/gin-gonic/gin"
	"hertz/app/http/controllers"
	"hertz/pkg/util"
	"net/http"
)

type AdminController struct {
}

func (adminController *AdminController) Index(c *gin.Context) {
	if util.IsAjax(c) {
		controllers.Success(c, gin.H{})
	} else {
		c.HTML(http.StatusOK, "admin/admin/index.html", nil)
	}
}
