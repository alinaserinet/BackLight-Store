package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID uint
	ProductID uint
	Title string `gorm:"size:255;"`
	Body string
}
