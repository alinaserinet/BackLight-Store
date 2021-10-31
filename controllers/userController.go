package controllers

import (
	"BackLight/models"
	"BackLight/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetAllUsers(c echo.Context) error {
	users, err := services.GetAllUsers()

	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func GetOneUser(c echo.Context) error {
	id, err := services.ParseParam(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, "id is not valid!")
	}

	user, userFetchErr := services.GetOneUser(id)

	if userFetchErr != nil {
		return c.JSON(http.StatusServiceUnavailable, err.Error())
	}

	return c.JSON(http.StatusOK, user)
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

	err := services.CreateOneUser(user)

	if err != nil {
		return c.JSON(http.StatusConflict, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}

func DeleteUser(c echo.Context) error {
	id, err := services.ParseParam(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, "id is not valid!")
	}

	statusCode, massage := services.DeleteOneUser(uint(id))

	return c.JSON(statusCode, massage)
}

func PutUser(c echo.Context) error {
	id, err := services.ParseParam(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, "id is not valid!")
	}

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

	statusCode, message := services.UpdateOneUser(id, user)

	return c.JSON(statusCode, message)
}
