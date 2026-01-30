package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// kita membuat variabel yang di koneksikan dengan gorm db
var DB *gorm.DB

func ConnectDB() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root@tcp(127.0.0.1:3306)/latihan?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("konesi ke database gagal")
	}
	DB = db
}
