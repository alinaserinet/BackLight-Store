package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"size:255;"`
	LastName  string `gorm:"size:255;"`
	Email     *string `gorm:"size:255;unique;"`
	Mobile    string `gorm:"size:20;unique;"`
	Password  string `gorm:"size:255;"`
	Role      uint `gorm:"default:1;"`
	Orders []Order `gorm:"foreignKey:UserID;"`
	CartProducts []CartProduct `gorm:"foreignKey:UserID;"`
	Comments []Comment `gorm:"foreignKey:UserID;"`
}
