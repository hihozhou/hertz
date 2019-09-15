package admin

import (
	"github.com/gin-gonic/gin"
	"hertz/app/http/controllers"
	"hertz/app/logic"
	"hertz/pkg/util"
	"net/http"
)

//admin控制器
type AdminController struct {
}

func (adminController *AdminController) Index(c *gin.Context) {
	if util.IsAjax(c) {
		//admin := &models.Admin{}
		adminLogic := logic.GetAdminLogic()
		count := adminLogic.GetAdminTotal(nil)
		//datebase.GetDB().Find(&admin)
		//fmt.Println(admin)
		admins := adminLogic.GetAdmins(0, 1, nil)
		controllers.Success(c, gin.H{"count": count, "list": admins})
	} else {
		c.HTML(http.StatusOK, "admin/admin/index.html", nil)
	}
}
