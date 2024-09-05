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

	db.Migrator().DropTable(&models.Address{})
	db.Migrator().DropTable(&models.City{})
	db.Migrator().DropTable(&models.State{})
	db.Migrator().DropTable(&models.Country{})
	db.Migrator().DropTable(&models.Currency{})
	db.Migrator().DropTable(&models.Timezone{})
	db.Migrator().DropTable(&models.UserProfile{})
	db.Migrator().DropTable(&models.User{})
	db.Migrator().DropTable(&models.Role{})
	db.Migrator().DropTable(&models.Gender{})
	db.Migrator().DropTable(&models.Category{})
	db.Migrator().DropTable(&models.Product{})
	db.Migrator().DropTable(&models.Customer{})
	db.Migrator().DropTable(&models.CustomerProfile{})
	db.Migrator().DropTable(&models.CustomerProductConfiguration{})
	db.Migrator().DropTable(&models.CustomMeasurement{})
	db.Migrator().DropTable(&models.Order{})
	db.Migrator().DropTable(&models.OrderItem{})
	db.Migrator().DropTable(&models.OrderStatus{})
	db.Migrator().DropTable(&models.Invoice{})
	db.Migrator().DropTable(&models.InvoiceItem{})
	db.Migrator().DropTable(&models.InvoiceStatus{})
	db.Migrator().DropTable(&models.Customization{})
	db.Migrator().DropTable(&models.CustomizationItem{})
	db.Migrator().DropTable(&models.CustomizationItemOption{})
	db.Migrator().DropTable(&models.CustomizationType{})
	db.Migrator().DropTable("users_roles")
	db.Migrator().DropTable("countries_timezones")
	db.Migrator().DropTable("countries_currencies")
	db.Migrator().DropTable("casbin_rule")

	fmt.Println("schema dropped successfully!")
}
