package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Status      int8
	ResultCode  string `gorm:"size:255;"`
	Amount      float64
	Description string
	Orders []Order `gorm:"foreignKey:TransactionID;"`
}
