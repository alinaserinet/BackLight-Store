package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID uint
	AddressID uint
	TransactionID uint
	DeliveryMethodID uint
	Amount float64
	State string `gorm:"size:100;"`
	Products []OrderProduct `gorm:"foreignKey:OrderID;"`
}