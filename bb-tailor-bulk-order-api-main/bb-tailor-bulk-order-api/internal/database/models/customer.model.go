package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	ID        string `gorm:"type:varchar(36);primaryKey"`
	Email     string `gorm:"unique;type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Profile                       CustomerProfile
	Orders                        []Order
	CustomerProductConfigurations []CustomerProductConfiguration
	CustomMeasurements            []CustomMeasurement
}

func (c *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()
	return
}
