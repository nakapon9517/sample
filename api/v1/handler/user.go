package handler

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// Login execLogin
func Login(c echo.Context) (err error) {
	userID := c.Param("userId")
	pwd := c.Param("pwd")
	isLogin := execLogin(db, userID, pwd)
	return c.String(http.StatusOK, isLogin)
}

// UserID testUserId
func UserID(c echo.Context) (err error) {
	id := c.Param("id")
	name := GetName(id)
	email := GetEmail(id)
	return c.String(http.StatusOK, "Name:"+name+"Email:"+email)
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

// executeLogin execLogin
func execLogin(db *gorm.DB, userID string, pwd string) string {
	user := getUser(db, userID, pwd)
	fmt.Println(user)
	return user.Email
}

func getUser(db *gorm.DB, userID string, pwd string) User {
	user := User{}
	// db.Table("User").First(&user).Where("id = ?", userID, "password = ?", pwd)
	db.Where("id = ?", userID, "password = ?", pwd).First(&user)
	fmt.Println(user)
	return user
}
