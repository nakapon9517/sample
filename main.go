package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/hello:id", func(c echo.Context) error {
		userid := c.Param("id")
		name := getName(userid)
		return c.String(http.StatusOK, "Hello, World! "+name)
	})
	e.GET("/mail:id", func(c echo.Context) error {
		userid := c.Param("id")
		name := getName(userid)
		return c.String(http.StatusOK, "Hello, World! "+name)
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func getName(id string) string {
	fmt.Println("id:" + id)
	switch id {
	case "1":
		return "aaa"
	case "2":
		return "bbb"
	default:
		return "not found"
	}
}
