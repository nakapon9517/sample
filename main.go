package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	v1 "github.com/nakapon9517/sample/api/v1"

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

func main() {

	db := gormConnect()
	defer db.Close()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	api := e.Group("/api")
	ver := api.Group("/v1")
	v1.SetupRoutes(ver)
	e.Logger.Fatal(e.Start(":1323"))
}
