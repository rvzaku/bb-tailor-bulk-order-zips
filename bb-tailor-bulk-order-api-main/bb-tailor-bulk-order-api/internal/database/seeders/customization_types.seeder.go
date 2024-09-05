package seeders

import (
	"fmt"
	"log"

	"git.bopconsultancy.com/tejasc/bb-tailor-bulk-order-api/internal/database/models"
	"gorm.io/gorm"
)

func SeedCustomizationTypes(db *gorm.DB) {
	customization_types := []models.CustomizationType{
		{Name: "STYLE"},
		{Name: "MATERIAL"},
	}

	for _, customization_type := range customization_types {
		if err := db.FirstOrCreate(&customization_type, models.CustomizationType{Name: customization_type.Name}).Error; err != nil {
			log.Fatalf("failed to seed customization type %s: %v", customization_type.Name, err)
		}
	}

	fmt.Println("customization types table seeding done successfully!")
}
