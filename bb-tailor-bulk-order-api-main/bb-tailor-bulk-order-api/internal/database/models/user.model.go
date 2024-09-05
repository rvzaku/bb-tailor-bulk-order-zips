package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           string `gorm:"type:varchar(36);primaryKey"`
	Email        string `gorm:"unique;type:varchar(255)"`
	PasswordHash string `gorm:"not null;type:varchar(255)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`

	Roles   []*Role `gorm:"many2many:users_roles;"`
	Profile UserProfile
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}
