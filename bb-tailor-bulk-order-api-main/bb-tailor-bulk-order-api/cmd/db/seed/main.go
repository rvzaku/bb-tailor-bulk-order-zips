package main

import (
	"fmt"
	"log"

	"git.bopconsultancy.com/tejasc/bb-tailor-bulk-order-api/internal/config"
	"git.bopconsultancy.com/tejasc/bb-tailor-bulk-order-api/internal/database/seeders"
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

	seeders.SeedRoles(db)
	seeders.SeedGenders(db)
	seeders.SeedCustomizationTypes(db)
	seeders.SeedOrderStatuses(db)
	seeders.SeedInvoiceStatuses(db)
	seeders.SeedCountriesStatesCitiesTimezonesCurrencies(db)
	seeders.SeedUsers(db)
	seeders.SeedCasbinPolicies(db)

	fmt.Println("complete database seeding done successfully!")
}
