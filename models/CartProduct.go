package models

type CartProduct struct {
	ProductID uint
	UserID uint
	Count uint
	Color string `gorm:"size:100;"`
}
