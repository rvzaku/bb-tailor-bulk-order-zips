package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerProfile struct {
	ID            string `gorm:"type:varchar(36);primaryKey"`
	FirstName     string `gorm:"type:varchar(255)"`
	LastName      string `gorm:"type:varchar(255)"`
	Phone         string `gorm:"type:varchar(255)"`
	Age           int    `gorm:"type:integer"`
	ProfilePicUrl string `gorm:"type:varchar(255)"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`

	CustomerID string `gorm:"type:varchar(36)"`
	Gender     string `gorm:"type:varchar(255)"`
	Addresses  []Address
}

func (cp *CustomerProfile) BeforeCreate(tx *gorm.DB) (err error) {
	cp.ID = uuid.New().String()
	return
}
