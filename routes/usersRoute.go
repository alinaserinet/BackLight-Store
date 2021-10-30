package routes

import (
	"BackLight/controllers"
	"github.com/labstack/echo/v4"
)

func UsersRoute(usersRoute *echo.Group)   {
	usersRoute.GET("", controllers.GetAllUsers)
	usersRoute.POST("", controllers.SaveUser)
}
