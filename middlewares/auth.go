package middlewares

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func BasicAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		token := context.Request().Header.Get("Authorization")
		if token == "token" {
			return next(context)
		}
		return context.JSON(http.StatusUnauthorized, "Your request can't be processed!")
	}
}