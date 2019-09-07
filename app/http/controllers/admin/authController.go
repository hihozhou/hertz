package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hertz/app/http/controllers"
	adminValidator "hertz/app/http/validator/admin"
	"hertz/app/logic"
	adminLogic "hertz/app/logic/admin"
	"hertz/app/models"
	"net/http"
)

type AuthController struct {
}

// 首页
func (auth AuthController) LoginForm(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/auth/login.html", nil)
}

// 登录验证
func validateLogin(c *gin.Context) adminValidator.LoginValidator {
	var params adminValidator.LoginValidator
	if err := c.ShouldBind(&params); err != nil {
		panic(err.Error())
	}
	return params
}

// 尝试登录
func attemptLogin(params *adminValidator.LoginValidator) (admin *models.Admin) {
	var al logic.AdminLogic
	var err error
	if admin, err = al.GetByPhone(params.Phone); err != nil {
		panic(err.Error())
	}
	if !al.PasswordCheck(admin, params.Password) {
		panic("账号或密码错误")
	}
	return admin
}

//登录操作
func (auth AuthController) Login(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			controllers.Fail(c, fmt.Sprintf("%s", err), gin.H{})
		}
	}()
	//验证器验证
	params := validateLogin(c)
	//查询账号
	admin := attemptLogin(&params)

	data := adminLogic.Login(c,admin)
	controllers.Success(c, data)
}
