package datebase

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"hertz/pkg/setting"
	"log"
)

var dB *gorm.DB

// 基本模型
type Model struct {
}

// 初始化数据库配置
func init() {
	var (
		err                                                        error
		connection, dbName, username, password, host, port, prefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	connection = sec.Key("CONNECTION").String()
	dbName = sec.Key("NAME").String()
	username = sec.Key("USERNAME").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	port = sec.Key("PORT").String()
	prefix = sec.Key("PREFIX").String()

	dB, err = gorm.Open(connection, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		dbName))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return prefix + defaultTableName;
	}

	//如果是debug模式，输出sql日志
	dB.LogMode(setting.AppDebug)

	dB.SingularTable(true)
	dB.DB().SetMaxIdleConns(10)
	dB.DB().SetMaxOpenConns(100)
}

// 关闭数据库链接
func CloseDB() {
	defer GetDB().Close()
}

func GetDB() *gorm.DB {
	return dB
}
