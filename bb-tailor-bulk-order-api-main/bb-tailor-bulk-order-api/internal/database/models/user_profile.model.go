package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserProfile struct {
	ID        string `gorm:"type:varchar(36);primaryKey"`
	FirstName string `gorm:"type:varchar(255)"`
	LastName  string `gorm:"type:varchar(255)"`
	Phone     string `gorm:"type:varchar(255)"`
	Age       int    `gorm:"type:integer"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	UserID    string `gorm:"type:varchar(36)"`
	Gender    string `gorm:"type:varchar(255)"`
	Addresses []Address
}

func (up *UserProfile) BeforeCreate(tx *gorm.DB) (err error) {
	up.ID = uuid.New().String()
	return
}
