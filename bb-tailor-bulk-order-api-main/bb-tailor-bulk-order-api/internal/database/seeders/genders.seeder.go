package seeders

import (
	"fmt"
	"log"

	"git.bopconsultancy.com/tejasc/bb-tailor-bulk-order-api/internal/database/models"
	"gorm.io/gorm"
)

func SeedGenders(db *gorm.DB) {
	genders := []models.Gender{
		{Name: "MALE"},
		{Name: "FEMALE"},
		{Name: "OTHERS"},
	}

	for _, gender := range genders {
		if err := db.FirstOrCreate(&gender, models.Gender{Name: gender.Name}).Error; err != nil {
			log.Fatalf("failed to seed gender %s: %v", gender.Name, err)
		}
	}

	fmt.Println("gender table seeding done successfully!")
}
