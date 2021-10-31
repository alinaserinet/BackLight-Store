package routes

import (
	"github.com/labstack/echo/v4"
)

func Setup(e *echo.Echo) {
	UsersRoute(e.Group("/users"))
}
