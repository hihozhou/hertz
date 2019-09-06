package models

import (
	"hertz/datebase"
	"time"
)

// 用户表
type User struct {
	datebase.Model
	ID        int       `gorm:"primary_key;AUTO_INCREMENT" json:"id"`       //主键id
	Phone     string    `gorm:"type:varchar(15);unique_index" json:"phone"` //手机号码
	Nickname  string    `gorm:"type:varchar(80)" json:"nickname"`           //昵称
	Password  string    `gorm:"type:varchar(32)" json:"password"`           //密码，加密后
	CreatedAt time.Time `json:"created_at"`                                 //创建时间
	UpdatedAt time.Time `json:"updated_at"`                                 //删除时间
	datebase.SoftDeletes
}
