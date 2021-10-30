package models

import (
	"time"
)

type SupplierProduct struct {
	SupplierID uint
	ProductID  uint
	Quantity   int
	Price      float64
	Colors string `gorm:"size:100;"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
