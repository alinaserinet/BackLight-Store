package models

type Attribute struct {
	ID          uint
	Name        string `gorm:"size:255"`
	Description string
	Values      []AttributeValue `gorm:"foreignKey:AttributeID;"`
	Products    *[]Product       `gorm:"many2many:product_attributes;"`
	Categories  *[]Category      `gorm:"many2many:default_attributes;"`
}
