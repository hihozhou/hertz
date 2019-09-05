package models

import "hertz/datebase"

type User struct {
	datebase.Model
	ID         int `gorm:"primary_key" json:"id"`
	CreateTime int `json:"created_time"`
	UpdateTime int `json:"update_time"`
}
