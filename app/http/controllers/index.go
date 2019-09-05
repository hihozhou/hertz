package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hertz/app/http/validator"
	"hertz/app/logic"
	"hertz/app/models"
	"hertz/datebase"
)

type Index struct {
}

func (index Index) Index(context *gin.Context) {
	Success(context, gin.H{"number": 1})
}

//登录验证
func validateLogin(context *gin.Context) validator.Login {
	var params validator.Login
	if err := context.ShouldBind(&params); err != nil {
		panic(err.Error())
	}
	return params
}

//登录方法
func (index Index) Login(context *gin.Context) () {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获异常:", err)
			Fail(context, fmt.Sprintf("%s", err), gin.H{})
		}
	}()
	//验证器验证
	params := validateLogin(context)
	//查询账号
	user := models.User{}
	if datebase.DB.Where("phone = ?", params.Phone).First(&user).RecordNotFound() {
		panic("账号或密码错误")
	}
	var userLogic logic.UserLogic
	if !userLogic.PasswordCheck(&user, params.Password) {
		panic("账号或密码错误")
	}
	//todo 使用jwt登录
	Success(context, gin.H{"data": "账号密码正确"})
}

//个人中心
func (index Index) Home(context *gin.Context) {
	Success(context, gin.H{"number": 1})
}
