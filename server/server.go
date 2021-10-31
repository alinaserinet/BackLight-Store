package server

import (
	"BackLight/routes"
	"github.com/labstack/echo/v4"
	"os"
)

var e *echo.Echo

func Init()  {
	e = echo.New()
	routes.Setup(e)
	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
