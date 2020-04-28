package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func New() *gorm.DB {
	var err error

	DBMS     := "mysql"
	USER     := "root"
	PASS     := "password"
	PROTOCOL := "tcp(mysql:3306)"
	DBNAME   := "test"

	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME
	db, err := gorm.Open(DBMS, CONNECT+"?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true")
	if err != nil {
		panic("failed to connect database!!")
	}

	db.LogMode(true)
	return db
}
