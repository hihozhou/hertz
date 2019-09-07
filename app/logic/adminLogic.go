package logic

import (
	"crypto/md5"
	"errors"
	"fmt"
	"hertz/app/models"
	"hertz/datebase"
)

type AdminLogic struct {
}

/*
密码加密方法
@author：hihozhou
*/
func (adminLogic *AdminLogic) PasswordEncryption(password string) string {
	data := []byte(password)
	has := md5.Sum(data)
	encryptedPassword := fmt.Sprintf("%x", has)
	return encryptedPassword
}

//验证账号密码是否正确
func (adminLogic *AdminLogic) PasswordCheck(admin *models.Admin, password string) bool {
	return admin.Password == adminLogic.PasswordEncryption(password)
}

/*
通过电话号码获取账号
*/
func (adminLogic *AdminLogic) GetByPhone(phone string) (admin *models.Admin, err error) {
	admin = &models.Admin{}
	//err = nil
	query := datebase.DB.Where("phone = ?", phone).First(admin)
	if err = query.Error; err != nil {
		return admin, err
	}
	if query.RecordNotFound() {
		err = errors.New("无效账号电话号码")
	}
	return admin, err
}
