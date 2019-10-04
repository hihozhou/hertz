package admin

import (
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
func attemptLogin(params *adminValidator.LoginValidator) (admin *models.Admin, err error) {
	al := logic.GetAdminLogic()
	if admin, err = al.GetByPhone(params.Phone); err != nil {
		return nil, err
	}
	if !al.PasswordCheck(admin, params.Password) {
		return nil, logic.ACCOUNT_PASSWORD_WRONG_ERROR
	}
	return admin, nil
}

//登录操作
func (auth AuthController) Login(c *gin.Context) {
	//验证器验证
	params := validateLogin(c)
	//查询账号
	admin, err := attemptLogin(&params);
	if err != nil {
		controllers.Fail(c, err.Error(),nil)
		return
	}
	data := adminLogic.Login(c, admin)
	controllers.Success(c, data)

}
