package controllers

import (
	"BackLight/models"
	"BackLight/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetAllUsers(c echo.Context) error {
	return nil
}

func SaveUser(c echo.Context) error {
	firstName := c.FormValue("first_name")
	lastName := c.FormValue("last_name")
	email := c.FormValue("email")
	mobile := c.FormValue("mobile")
	password := c.FormValue("password")

	user := &models.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     &email,
		Mobile:    mobile,
		Password:  password,
	}

	err := services.UserCreator(user)

	if err != nil {
		return c.JSON(http.StatusConflict, &response{
			StatusCode: http.StatusConflict,
			Type:       "Error",
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, &response{
		StatusCode: http.StatusCreated,
		Type:       "Success",
		Message:    "User Created",
		User:       user,
	})
}

type response struct {
	StatusCode uint
	Type       string
	Message    string
	User       *models.User
}
