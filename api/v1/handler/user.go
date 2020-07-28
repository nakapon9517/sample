package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func UserID(c echo.Context) (err error) {
	userid := c.Param("id")
	name := GetName(userid)
	return c.String(http.StatusOK, "Hello, World!"+name)
}

func GetName(id string) string {
	switch id {
	case "1":
		fmt.Println("aaa")
		return "aaa"
	case "2":
		fmt.Println("bbb")
		return "bbb"
	default:
		fmt.Println("not found")
		return "not found"
	}
}
