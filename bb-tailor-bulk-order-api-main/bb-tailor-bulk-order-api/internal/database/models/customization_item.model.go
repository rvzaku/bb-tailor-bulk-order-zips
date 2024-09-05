package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomizationItem struct {
	ID               string `gorm:"type:varchar(36);primaryKey"`
	Name             string `gorm:"type:varchar(255)"`
	ShortDescription string `gorm:"type:varchar(255)"`
	LongDescription  string `gorm:"type:varchar(255)"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`

	CustomizationID          string `gorm:"type:varchar(36)"`
	CustomizationItemOptions []CustomizationItemOption
}

func (ci *CustomizationItem) BeforeCreate(tx *gorm.DB) (err error) {
	ci.ID = uuid.New().String()
	return
}
