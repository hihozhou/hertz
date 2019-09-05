package controllers

import (
	"github.com/gin-gonic/gin"
	"hertz/app/http/validator"
	"hertz/app/logic"
	"hertz/app/models"
	"hertz/datebase"
	"net/http"
)

type Index struct {
}

func (index Index) Index(context *gin.Context) {
	Success(context, gin.H{"number": 1})
}

func (index Index) Login(context *gin.Context) () {
	//验证器验证
	var params validator.Login
	// This will infer what binder to use depending on the content-type header.
	if err := context.ShouldBind(&params); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//查询账号
	user := models.User{}
	if datebase.DB.Where("phone = ?", params.Phone).First(&user).RecordNotFound() {
		Fail(context, "账号或密码错误1", gin.H{})
		return
	}
	var userLogic logic.UserLogic
	if !userLogic.PasswordCheck(&user, params.Password) {
		Fail(context, "账号或密码错误2", gin.H{})
		return
	}
	//使用jwt登录
	Success(context, gin.H{"data": "账号密码正确"})
}

func (index Index) Home(context *gin.Context) {
	Success(context, gin.H{"number": 1})
}
