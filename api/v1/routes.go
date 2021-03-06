package v1

import (
	"github.com/labstack/echo"
	"github.com/nakapon9517/sample/api/v1/handler"
)

// SetupRoutes setup
func SetupRoutes(ver *echo.Group) {
	ver.GET("/users:id", handler.UserID)
	ver.GET("/login:email", handler.Login)
}
