package main

import (
	"hertz/app/models"
	"hertz/datebase"
)

func main() {
	datebase.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{})
	datebase.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Admin{})
}
