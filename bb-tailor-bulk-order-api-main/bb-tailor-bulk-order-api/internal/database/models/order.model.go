package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID          string  `gorm:"type:varchar(36);primaryKey"`
	Subtotal    float32 `gorm:"type:float"`
	TaxRate     float32 `gorm:"type:float"`
	TaxAmount   float32 `gorm:"type:float"`
	Total       float32 `gorm:"type:float"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	PlacedAt    time.Time
	ShippedAt   time.Time
	DeliveredAt time.Time

	Status     string `gorm:"type:varchar(255)"`
	CustomerID string `gorm:"type:varchar(36)"`
	OrderItems []OrderItem
	Invoices   []Invoice
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = uuid.New().String()
	return
}
