package logic

import (
	"crypto/md5"
	"errors"
	"fmt"
	"hertz/app/models"
	"hertz/datebase"
)

type UserLogic struct {
}


/*
密码加密方法
author：hihozhou
 */
func (userLogic *UserLogic) PasswordEncryption(password string) string {
	data := []byte(password)
	has := md5.Sum(data)
	encryptedPassword := fmt.Sprintf("%x", has)
	return encryptedPassword
}

//验证账号密码是否正确
func (userLogic *UserLogic) PasswordCheck(user *models.User, password string) bool {
	return user.Password == userLogic.PasswordEncryption(password)
}


/*
通过电话号码获取账号
 */
func (userLogic *UserLogic) GetByPhone(phone string) (user *models.User, err error) {
	user = &models.User{}
	err = nil
	if datebase.DB.Where("phone = ?", phone).First(user).RecordNotFound() {
		err = errors.New("无效账号电话号码")
	}
	return user, err
}
