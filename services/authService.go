package services

import (
	"BackLight/database"
	"BackLight/models"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func BasicAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		token := context.Request().Header.Get("Authorization")
		_, err := CheckToken(token)
		if err == nil {
			return next(context)
		}
		return context.JSON(http.StatusUnauthorized, "Your request can't be processed: " + err.Error())
	}
}

func CheckUser(email string, password string) (models.User, error) {
	var user = models.User{}
	result := database.DB.Where("email = ?", email).Find(&user)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	if result.RowsAffected == 0 {
		return models.User{}, errors.New("email or password incorrect")
	}

	checkPassword := Compare(user.Password, password)
	if !checkPassword {
		return models.User{}, errors.New("email or password incorrect")
	}

	return user, nil
}
