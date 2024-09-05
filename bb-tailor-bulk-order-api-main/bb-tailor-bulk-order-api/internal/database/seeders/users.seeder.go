package seeders

import (
	"fmt"
	"log"

	"git.bopconsultancy.com/tejasc/bb-tailor-bulk-order-api/internal/database/models"
	"git.bopconsultancy.com/tejasc/bb-tailor-bulk-order-api/internal/utils"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) {
	users := []struct {
		Email     string
		Password  string
		FirstName string
		LastName  string
		Phone     string
		Age       int
		Gender    string
		Role      string
	}{
		{
			"tejasc@consultbop.com",
			"admin1234",
			"Tejas",
			"Chari",
			"+919511705719",
			28,
			"MALE",
			"SUPERADMIN",
		},
		{
			"employee@example.com",
			"password123",
			"Jane",
			"Doe",
			"+919876543210",
			35,
			"FEMALE",
			"EMPLOYEE",
		},
	}

	for _, user := range users {
		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			log.Fatalf("failed to hash password for user %s: %v", user.Email, err)
		}

		profile := models.UserProfile{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Gender:    user.Gender,
			Age:       user.Age,
			Phone:     user.Phone,
		}

		newUser := models.User{
			Email:        user.Email,
			PasswordHash: string(hashedPassword),
			Profile:      profile,
		}

		if err := db.FirstOrCreate(&newUser, models.User{Email: user.Email}).Error; err != nil {
			log.Fatalf("failed to seed user %s: %v", user.Email, err)
		}

		query := `
			INSERT INTO users_roles(user_id, role_name)
			VALUES (?, ?)
		`
		if err := db.Exec(query, newUser.ID, user.Role).Error; err != nil {
			log.Fatalf("failed to insert into users_roles: %v", err)
		}
	}

	fmt.Println("users table seeding done successfully!")
}
