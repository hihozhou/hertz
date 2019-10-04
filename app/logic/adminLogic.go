package logic

import (
	"crypto/md5"
	"errors"
	"fmt"
	"hertz/app/models"
	"hertz/datebase"
	"math"
	"sync"
)

type AdminLogic struct {
}

var adminLogicSingleton *AdminLogic

var (
	ACCOUNT_INVALID_ERROR error = errors.New("无效账号电话号码")
	ACCOUNT_PASSWORD_WRONG_ERROR error = errors.New("账号或密码错误")
)

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
	query := datebase.GetDB().Where("phone = ?", phone).First(admin)
	if err = query.Error; err != nil {
		return nil, err
	}
	if query.RecordNotFound() {
		return nil, ACCOUNT_INVALID_ERROR
	}
	return admin, nil
}

//统计数量
func (adminLogic *AdminLogic) GetAdminTotal(maps interface{}) (count int) {
	query := datebase.GetDB().Model(&models.Admin{})
	if maps != nil {
		query = query.Where(maps)
	}
	query.Count(&count)
	return
}

//分页查询
func (adminLogic *AdminLogic) GetAdminsByOffset(offset int, limit int, maps interface{}) (admins []models.Admin) {
	query := datebase.GetDB()
	if maps != nil {
		query = query.Where(maps)
	}
	query.Offset(offset).Limit(limit).Find(&admins)
	return admins
}

//获取多条admin记录
func (adminLogic *AdminLogic) GetAdmins(page int, limit int, maps interface{}) (admins []models.Admin) {
	offset := (page - 1) * limit
	admins = adminLogic.GetAdminsByOffset(offset, limit, maps)
	return admins
}

//分页
func (adminLogic *AdminLogic) Paginate(page int, limit int, maps interface{}) (admins []models.Admin, count int) {
	if limit <= 0 {
		limit = 10
	}
	//计算出offset
	offset := (page - 1) * limit
	count = adminLogic.GetAdminTotal(maps)
	totalPages := int(math.Ceil(float64(count) / float64(limit))) //page总数
	if page <= 0 {
		page = 1
	} else if page > totalPages {
		page = totalPages
	}
	admins = adminLogic.GetAdminsByOffset(offset, limit, maps)
	return
}

//单利模式
func GetAdminLogic() *AdminLogic {
	once := sync.Once{}
	once.Do(func() {
		adminLogicSingleton = &AdminLogic{}
	})

	return adminLogicSingleton
}
