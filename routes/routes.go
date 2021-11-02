package routes

import (
	"BackLight/services"
	"github.com/labstack/echo/v4"
)

func Setup(e *echo.Echo) {
	UsersRoute(e.Group("/users", services.BasicAuth))
	AuthRoute(e.Group("/auth"))
}
