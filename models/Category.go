package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ParentID    uint
	Title       string `gorm:"size:255;"`
	Description string
	Products    []Product  `gorm:"foreignKey:CategoryID;"`
	Categories  []Category `gorm:"foreignKey:ParentID;"`
	DefaultAttributes *[]Attribute `gorm:"many2many:default_attributes;"`
}
