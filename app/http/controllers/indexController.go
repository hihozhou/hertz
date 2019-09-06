package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hertz/app/http/validator"
	"hertz/app/logic"
	"hertz/app/models"
)

type IndexController struct {
}

func (index IndexController) Index(context *gin.Context) {
	Success(context, gin.H{"number": 1})
}

//登录验证
func validateLogin(context *gin.Context) validator.LoginValidator {
	var params validator.LoginValidator
	if err := context.ShouldBind(&params); err != nil {
		panic(err.Error())
	}
	return params
}

//尝试登录
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

//登录方法
//@author hihozhou
func (index IndexController) Login(context *gin.Context) () {
	defer func() {
		if err := recover(); err != nil {
			Fail(context, fmt.Sprintf("%s", err), gin.H{})
		}
	}()
	//验证器验证
	params := validateLogin(context)
	//查询账号
	user := attemptLogin(&params)
	//todo 使用jwt登录
	Success(context, gin.H{"data": user, "message": "登录成功"})
}

//个人中心
func (index IndexController) Home(context *gin.Context) {
	Success(context, gin.H{"number": 1})
}
