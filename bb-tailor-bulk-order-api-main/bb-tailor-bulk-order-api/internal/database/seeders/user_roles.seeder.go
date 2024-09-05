package seeders

import (
	"fmt"
	"log"

	"git.bopconsultancy.com/tejasc/bb-tailor-bulk-order-api/internal/database/models"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {
	roles := []models.Role{
		{Name: "SUPERADMIN"},
		{Name: "EMPLOYEE"},
	}

	for _, role := range roles {
		if err := db.FirstOrCreate(&role, models.Role{Name: role.Name}).Error; err != nil {
			log.Fatalf("failed to seed role %s: %v", role.Name, err)
		}
	}

	fmt.Println("user roles table seeding done successfully!")
}
