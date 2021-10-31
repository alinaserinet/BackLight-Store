package services

import (
	"BackLight/database"
	"BackLight/models"
	"fmt"
	"net/http"
)

func CreateOneUser(user *models.User) error {
	result := database.DB.Create(user)
	return result.Error
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := database.DB.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func GetOneUser(id uint) (models.User, error) {
	user := models.User{}
	result := database.DB.Where("id = ?", id).Find(&user)

	return user, result.Error
}

func DeleteOneUser(id uint) (int, string) {
	result := database.DB.Where("id = ?", id).Delete(&models.User{})

	if result.Error != nil {
		return http.StatusInternalServerError,
			fmt.Sprintf("An error has occurred!\n %s", result.Error.Error())
	}

	if result.RowsAffected == 0 {
		return http.StatusNotFound,
			fmt.Sprintf("User by id %d Not Found!", id)
	}

	return http.StatusOK,
		fmt.Sprintf("%d rows affected. User by id %d Deleted.", result.RowsAffected, id)
}

func UpdateOneUser(id uint, user *models.User) (int, string) {
	result := database.DB.Model(models.User{}).Where("id = ?", id).Updates(user)

	if result.Error != nil {
		return http.StatusInternalServerError,
			fmt.Sprintf("An error has occurred: %s", result.Error.Error())
	}

	if result.RowsAffected == 0 {
		return http.StatusNotFound,
			fmt.Sprintf("User by id %d Not Found!", id)
	}

	return http.StatusOK,
		fmt.Sprintf("%d rows affected. User by id %d Updated.", result.RowsAffected, id)
}