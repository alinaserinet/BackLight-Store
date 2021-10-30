package models

type OrderProduct struct {
	OrderID uint
	ProductID uint
	SupplierID uint
	Price float64
	Count uint
	Color string `gorm:"size:100;"`
}
