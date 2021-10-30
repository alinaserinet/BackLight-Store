package models

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	UserID         uint
	AddresseeName  string `gorm:"size:255;"`
	AddresseePhone string `gorm:"size:20;"`
	ZipCode        uint8
	City           string `gorm:"size:255;"`
	Body           string
	Orders         []Order `gorm:"foreignKey:AddressID;"`
}
