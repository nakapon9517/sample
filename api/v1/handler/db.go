package handler

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "user2"
	PASS := "user2"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "go_sample"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

var db *gorm.DB

func init() {
	db = gormConnect()
	defer db.Close()
}
