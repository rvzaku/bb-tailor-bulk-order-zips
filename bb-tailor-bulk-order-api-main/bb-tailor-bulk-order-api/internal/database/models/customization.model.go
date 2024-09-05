package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customization struct {
	ID        string `gorm:"type:varchar(36);primaryKey"`
	Name      string `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Type                          string `gorm:"type:varchar(255)"`
	Categories                    []Category
	CustomizationItem             []CustomizationItem
	CustomerProductConfigurations []CustomerProductConfiguration
}

func (c *Customization) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()
	return
}
