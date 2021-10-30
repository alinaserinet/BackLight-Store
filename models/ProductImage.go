package models

type ProductImage struct {
	ProductID uint
	Url       string `gorm:"size:255;"`
	Alt       string `gorm:"size:255;"`
}
