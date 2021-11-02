package controllers

import (
	"BackLight/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetToken(context echo.Context) error {
	email := context.FormValue("email")
	password := context.FormValue("password")

	user, userError := services.CheckUser(email, password)

	if userError != nil {
		return context.JSON(http.StatusUnauthorized, userError.Error())
	}

	token, err := services.GenerateToken(&user)

	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, token)
}
