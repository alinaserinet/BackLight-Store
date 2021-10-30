package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title         string `gorm:"size:255;"`
	Description   string
	CategoryID    uint
	Image         string         `gorm:"size:255;"`
	Images        []ProductImage `gorm:"foreignKey:ProductID;"`
	CartProducts  []CartProduct  `gorm:"foreignKey:ProductID"`
	OrderProducts []OrderProduct `gorm:"foreignKey:ProductID;"`
	Comments      []Comment      `gorm:"foreignKey:ProductID;"`
	Suppliers     []SupplierProduct    `gorm:"foreignKey:ProductID;"`
	Attributes    *[]Attribute   `gorm:"many2many:product_attributes;"`
}
