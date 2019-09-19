package migrate

import (
	"hertz/app/models"
	"hertz/datebase"
)

func main() {
	datebase.GetDB().Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{})
	datebase.GetDB().Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Admin{})
}
