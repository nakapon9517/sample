package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/hello:userid", func(c echo.Context) error {
		var userid string = c.Param("userid")
		name := getName(userid)
		return c.String(http.StatusOK, "Hello, World!"+name)
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func getName(id string) string {
	switch id {
	case "1":
		return "aaa"
		fmt.Println("aaa")
	case "2":
		return "bbb"
		fmt.Println("bbb")
	default:
		return "not found"
		fmt.Println("not found")
	}
	return ""
}
