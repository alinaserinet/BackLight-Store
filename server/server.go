package server

import (
	"BackLight/routes"
	"github.com/labstack/echo/v4"
	"os"
)


func Init()  {
	e := echo.New()
	routes.Setup(e)
	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
