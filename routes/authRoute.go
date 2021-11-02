package routes

import (
	"BackLight/controllers"
	"github.com/labstack/echo/v4"
)

func AuthRoute(authRoute *echo.Group)  {
	authRoute.POST("/get-token", controllers.GetToken)
	authRoute.POST("/register", controllers.GetToken)
}
