package server

import (
	"BackLight/routes"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

var e *echo.Echo

func Init()  {
	e = echo.New()
	routes.Init(e)

	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))

}


func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	errorPage := fmt.Sprintf("%d.html", code)
	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}
	c.Logger().Error(err)
}
