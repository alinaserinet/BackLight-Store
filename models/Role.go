package models

type Role struct {
	ID        uint `gorm:"primarykey;"`
	Title string `gorm:"unique;"`
	Users []User `gorm:"foreignKey:Role;"`
}
