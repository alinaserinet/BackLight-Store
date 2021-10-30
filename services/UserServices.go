package services

import (
	"BackLight/database"
	"BackLight/models"
)

func UserCreator(user *models.User) error {
	result := database.DB.Create(user)

	return result.Error
}
