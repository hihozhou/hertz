package logic

import (
	"crypto/md5"
	"fmt"
	"hertz/app/models"
)

type UserLogic struct {
}

//密码加密方法
//password 密码
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
