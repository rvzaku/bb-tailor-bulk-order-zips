package seeders

import (
	"fmt"
	"log"

	"git.bopconsultancy.com/tejasc/bb-tailor-bulk-order-api/internal/database/models"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

func SeedCasbinPolicies(db *gorm.DB) {
	// Initialize GORM adapter for Casbin
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		log.Fatalf("Error creating adapter: %v", err)
	}

	// Initialize Casbin enforcer with model and policy from database
	enforcer, err := casbin.NewEnforcer("internal/config/rbac_model.conf", adapter)
	if err != nil {
		log.Fatalf("Error creating enforcer: %v", err)
	}

	// Load existing policies
	if err := enforcer.LoadPolicy(); err != nil {
		log.Fatalf("Error loading policy: %v", err)
	}

	// Fetch all users with their roles
	var users []models.User
	if err := db.Preload("Roles").Find(&users).Error; err != nil {
		log.Fatalf("Error fetching users: %v", err)
	}

	// Add policies based on roles
	for _, user := range users {
		for _, role := range user.Roles {
			_, err := enforcer.AddGroupingPolicy(user.ID, role.Name)
			if err != nil {
				log.Fatalf("Error adding grouping policy for user %s: %v", user.ID, err)
			}
		}
	}

	// Define your role-based policies here (this should be done once)
	// Example: Define policies for the SUPERADMIN role
	_, err = enforcer.AddPolicy("SUPERADMIN", "/users", "GET")
	if err != nil {
		log.Fatalf("Error adding policy for SUPERADMIN: %v", err)
	}

	// Save policies to database
	if err := enforcer.SavePolicy(); err != nil {
		log.Fatalf("Error saving policy: %v", err)
	}

	fmt.Println("casbin policies seeding done successfully!")
}
