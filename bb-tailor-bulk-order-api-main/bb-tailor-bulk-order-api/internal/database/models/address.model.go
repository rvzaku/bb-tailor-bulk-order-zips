package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Address struct {
	ID             string `gorm:"type:varchar(36);primaryKey"`
	Line1          string `gorm:"type:varchar(255)"`
	Line2          string `gorm:"type:varchar(255)"`
	NearbyLandmark string `gorm:"type:varchar(255)"`
	PinCode        string `gorm:"type:varchar(255)"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`

	UserProfileID     string `gorm:"type:varchar(36)"`
	CustomerProfileID string `gorm:"type:varchar(36)"`
	CityID            string `gorm:"type:varchar(36)"`
	City              City
}

func (a *Address) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New().String()
	return
}
