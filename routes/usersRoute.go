package routes

import (
	"BackLight/controllers"
	"github.com/labstack/echo/v4"
)

func UsersRoute(usersRoute *echo.Group)   {
	usersRoute.GET("", controllers.GetAllUsers)
	usersRoute.GET("/:id", controllers.GetOneUser)
	usersRoute.POST("", controllers.SaveUser)
	usersRoute.DELETE("/:id", controllers.DeleteUser)
	usersRoute.PUT("/:id", controllers.PutUser)
}
