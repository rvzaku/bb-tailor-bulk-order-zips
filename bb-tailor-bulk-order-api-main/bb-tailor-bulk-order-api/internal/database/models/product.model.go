package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID               string `gorm:"type:varchar(36);primaryKey"`
	Name             string `gorm:"type:varchar(255)"`
	Sku              string `gorm:"unique;type:varchar(255)"`
	ShortDescription string `gorm:"type:varchar(255)"`
	LongDescription  string `gorm:"type:varchar(255)"`
	ThumbnailUrl     string `gorm:"type:varchar(255)"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`

	CategoryID                    string `gorm:"type:varchar(36)"`
	OrderItems                    []OrderItem
	CustomerProductConfigurations []CustomerProductConfiguration
	InvoiceItems                  []InvoiceItem
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New().String()
	return
}
