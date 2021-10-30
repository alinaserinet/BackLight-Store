package models

import (
	"gorm.io/gorm"
	"time"
)

type Supplier struct {
	gorm.Model
	Name     string `gorm:"size:255;"`
	Username string `gorm:"size:255;unique;"`
	Email    string `gorm:"size:255;unique;"`
	Mobile   string `gorm:"size:20;unique;"`
	VerifyAt *time.Time
	Password string    `gorm:"size:255;"`
	Products []SupplierProduct `gorm:"ForeignKey:SupplierID;"`
}
