package admin

import (
	"github.com/gin-gonic/gin"
	"hertz/app/http/controllers"
	"hertz/app/logic"
	"hertz/pkg/util"
	"net/http"
	"strconv"
)

//admin控制器
type AdminController struct {
}

//管理员管理首页
func (adminController *AdminController) Index(c *gin.Context) {
	if util.IsAjax(c) {
		page, _ := strconv.Atoi(c.Query("page"));
		limit, _ := strconv.Atoi(c.Query("limit"));
		admins, count := logic.GetAdminLogic().Paginate(page, limit, nil)
		controllers.Success(c, gin.H{"count": count, "list": admins})
	} else {
		c.HTML(http.StatusOK, "admin/admin/index.html", nil)
	}
}

func (adminController *AdminController) Store(c *gin.Context) {
	//接受参数
	//验证器验证
	//保存信息
}
