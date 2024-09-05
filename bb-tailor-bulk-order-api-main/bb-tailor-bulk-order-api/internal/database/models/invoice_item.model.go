package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InvoiceItem struct {
	ID          string  `gorm:"type:varchar(36);primaryKey"`
	Description string  `gorm:"type:varchar(255)"`
	Quantity    int     `gorm:"type:int"`
	UnitPrice   float32 `gorm:"type:float"`
	TotalPrice  float32 `gorm:"type:float"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	InvoiceID string `gorm:"type:varchar(36)"`
	ProductID string `gorm:"type:varchar(36)"`
}

func (ii *InvoiceItem) BeforeCreate(tx *gorm.DB) (err error) {
	ii.ID = uuid.New().String()
	return
}
