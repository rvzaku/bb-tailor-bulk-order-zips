package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerProductConfiguration struct {
	ID        string         `gorm:"type:varchar(36);primaryKey"`
	Data      sql.NullString `gorm:"type:json"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	ProductID       string `gorm:"type:varchar(36)"`
	CustomizationID string `gorm:"type:varchar(36)"`
	CustomerID      string `gorm:"type:varchar(36)"`
}

func (cpc *CustomerProductConfiguration) BeforeCreate(tx *gorm.DB) (err error) {
	cpc.ID = uuid.New().String()
	return
}
