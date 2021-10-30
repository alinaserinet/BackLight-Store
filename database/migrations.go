package database

import (
	"BackLight/models"
	"log"
)

func InitMigrations() {
	err := DB.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.Transaction{},
		&models.DeliveryMethod{},
		&models.Order{},
		&models.Supplier{},
		&models.Product{},
		&models.SupplierProduct{},
		&models.Attribute{},
		&models.AttributeValue{},
		&models.CartProduct{},
		&models.Category{},
		&models.Comment{},
		&models.OrderProduct{},
		&models.ProductImage{},
	)

	if err != nil {
		log.Fatal("Error in Migration!")
	}
}
