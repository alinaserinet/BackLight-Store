package models

type DeliveryMethod struct {
	ID        uint `gorm:"primarykey"`
	Title string `gorm:"size:255;"`
	Orders []Order `gorm:"foreignKey:DeliveryMethodID;"`
}
