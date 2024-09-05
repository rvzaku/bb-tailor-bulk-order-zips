package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomizationItemOption struct {
	ID               string `gorm:"type:varchar(36);primaryKey"`
	Name             string `gorm:"type:varchar(255)"`
	Sku              string `gorm:"unique;type:varchar(255)"`
	ShortDescription string `gorm:"type:varchar(255)"`
	LongDescription  string `gorm:"type:varchar(255)"`
	ThumbnailUrl     string `gorm:"type:varchar(255)"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`

	CustomizationItemID string `gorm:"type:varchar(36)"`
}

func (cio *CustomizationItemOption) BeforeCreate(tx *gorm.DB) (err error) {
	cio.ID = uuid.New().String()
	return
}
