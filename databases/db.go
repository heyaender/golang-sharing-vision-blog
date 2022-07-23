package databases

import (
	"log"
	"sv-article/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func MySQLConnect() {

	DB, err = gorm.Open(mysql.Open(configs.DB_USER+":"+configs.DB_PASSWORD+"@tcp("+configs.DB_HOST+")/"+configs.DB_NAME+"?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		log.Println("DB Connection Error: ")
	} else {
		log.Println("DB Connection Success")
	}

	DBMigrate()
}
