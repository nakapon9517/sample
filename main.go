package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	v1 "github.com/nakapon9517/sample/api/v1"
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	api := e.Group("/api")
	ver := api.Group("/v1")
	v1.SetupRoutes(ver)
	e.Logger.Fatal(e.Start(":1323"))
}
