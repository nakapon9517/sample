package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// UserID testUserId
func UserID(c echo.Context) (err error) {
	id := c.Param("id")
	name := GetName(id)
	email := GetEmail(id)
	return c.String(http.StatusOK, "Name:"+name+"Email:"+email)
}

// Login execLogin
func Login(c echo.Context) (err error) {
	userID := c.Param("userId")
	pwd := c.Param("pwd")
	isLogin := ExecuteLogin(userID, pwd)
	return c.String(http.StatusOK, isLogin)
}

// GetName getName
func GetName(id string) string {
	switch id {
	case "1":
		fmt.Println("aaa")
		return "aaa"
	case "2":
		fmt.Println("bbb")
		return "bbb"
	default:
		fmt.Println("empty name")
		return "empty name"
	}
}

// GetEmail getEmail
func GetEmail(id string) string {
	switch id {
	case "1":
		fmt.Println("XXX")
		return "XXX"
	case "2":
		fmt.Println("YYY")
		return "YYY"
	default:
		fmt.Println("empty email")
		return "empty email"
	}
}

// ExecuteLogin execLogin
func ExecuteLogin(userId string, pwd string) string {
	return userId + pwd
}
