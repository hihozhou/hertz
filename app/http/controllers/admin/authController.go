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
func (auth *AuthController) LoginForm(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/auth/login.html", nil)
}

// 登录验证
func validateLogin(c *gin.Context) (params *adminValidator.LoginValidator, err error) {
	params = &adminValidator.LoginValidator{}
	err = c.ShouldBind(params);
	return params, err
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
//author hihozhou
func (auth *AuthController) Login(c *gin.Context) {
	//验证器验证
	params, dataErr := validateLogin(c)
	if dataErr != nil {
		controllers.Fail(c, dataErr.Error(), nil)
		return
	}
	//查询账号
	admin, err := attemptLogin(params);
	if err != nil {
		controllers.Fail(c, err.Error(), nil)
		return
	}

	//登录操作
	data, loginErr := adminLogic.Login(c, admin)
	if loginErr != nil {
		controllers.Fail(c, loginErr.Error(), nil)
		return
	}
	controllers.Success(c, data)

}

func (auth *AuthController) Logout(c *gin.Context) {
	adminLogic.Logout(c)
	c.Redirect(http.StatusMovedPermanently, adminLogic.LOGIN_PATH)
}
