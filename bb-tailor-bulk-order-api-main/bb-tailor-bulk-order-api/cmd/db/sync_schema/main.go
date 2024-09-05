package main

import (
	"fmt"
	"log"

	"git.bopconsultancy.com/tejasc/bb-tailor-bulk-order-api/internal/config"
	"git.bopconsultancy.com/tejasc/bb-tailor-bulk-order-api/internal/database/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %s", err)
	}

	dsn := cfg.MySqlDb.DbUser + ":" + cfg.MySqlDb.DbUserPassword + "@tcp(" + cfg.MySqlDb.DbHost + ":" + cfg.MySqlDb.DbPort + ")/" + cfg.MySqlDb.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Gender{},
		&models.UserProfile{},
		&models.Role{},
		&models.Customer{},
		&models.CustomerProfile{},
		&models.Timezone{},
		&models.Currency{},
		&models.Country{},
		&models.State{},
		&models.City{},
		&models.Address{},
		&models.CustomizationType{},
		&models.Customization{},
		&models.CustomizationItem{},
		&models.CustomizationItemOption{},
		&models.Category{},
		&models.Product{},
		&models.CustomMeasurement{},
		&models.CustomerProductConfiguration{},
		&models.OrderStatus{},
		&models.Order{},
		&models.OrderItem{},
		&models.InvoiceStatus{},
		&models.Invoice{},
		&models.InvoiceItem{},
	)

	if err != nil {
		fmt.Println("error creating schema:", err)
		return
	}

	fmt.Println("schema successfully created using gorm auto migrate")
}
