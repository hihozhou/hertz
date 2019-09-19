package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hertz/app/http/validator"
	"hertz/app/logic"
	"hertz/app/models"
)

//首页控制器
type IndexController struct {
}

// 首页
func (index IndexController) Index(c *gin.Context) {
	Success(c, gin.H{"number": 1})
}

// 登录验证
func validateLogin(c *gin.Context) validator.LoginValidator {
	var params validator.LoginValidator
	if err := c.ShouldBind(&params); err != nil {
		panic(err.Error())
	}
	return params
}

// 尝试登录
func attemptLogin(params *validator.LoginValidator) (user *models.User) {
	var userLogic logic.UserLogic
	var err error
	if user, err = userLogic.GetByPhone(params.Phone); err != nil {
		panic(err.Error())
	}
	if !userLogic.PasswordCheck(user, params.Password) {
		panic("账号或密码错误")
	}
	return user
}

// 登录方法
// @author hihozhou
func (index IndexController) Login(c *gin.Context) () {
	defer func() {
		if err := recover(); err != nil {
			Fail(c, fmt.Sprintf("%s", err), gin.H{})
		}
	}()
	//验证器验证
	params := validateLogin(c)
	//查询账号
	user := attemptLogin(&params)

	data := logic.Login(user)
	//todo 使用jwt登录
	Success(c, data)
}

// 个人中心
func (index IndexController) Home(c *gin.Context) {
	//value, exists := c.Get("claims")
	Success(c, gin.H{"info": logic.User(c)})
}
