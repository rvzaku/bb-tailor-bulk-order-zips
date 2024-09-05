package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderItem struct {
	ID         string  `gorm:"type:varchar(36);primaryKey"`
	Quantity   int     `gorm:"type:int"`
	UnitPrice  float32 `gorm:"type:float"`
	TotalPrice float32 `gorm:"type:float"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`

	OrderID                        string `gorm:"type:varchar(36)"`
	ProductID                      string `gorm:"type:varchar(36)"`
	CustomerProductConfigurationID string `gorm:"type:varchar(36)"`
}

func (oi *OrderItem) BeforeCreate(tx *gorm.DB) (err error) {
	oi.ID = uuid.New().String()
	return
}
