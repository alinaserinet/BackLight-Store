package models

type AttributeValue struct {
	AttributeID uint
	Value string `gorm:"size:255;"`
}
