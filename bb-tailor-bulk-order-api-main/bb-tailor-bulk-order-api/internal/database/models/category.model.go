package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID               string `gorm:"type:varchar(36);primaryKey"`
	Name             string `gorm:"unique;type:varchar(255)"`
	ShortDescription string `gorm:"type:varchar(255)"`
	LongDescription  string `gorm:"type:varchar(255)"`
	ThumbnailUrl     string `gorm:"type:varchar(255)"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`

	Products        []Product
	CustomizationID string `gorm:"type:varchar(36)"`
}

func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()
	return
}
