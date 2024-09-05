package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomMeasurement struct {
	ID        string         `gorm:"type:varchar(36);primaryKey"`
	Data      sql.NullString `gorm:"type:json"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	CustomerID string `gorm:"type:varchar(36)"`
}

func (cm *CustomMeasurement) BeforeCreate(tx *gorm.DB) (err error) {
	cm.ID = uuid.New().String()
	return
}
