package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Invoice struct {
	ID string `gorm:"type:varchar(36);primaryKey"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	IssuedAt  time.Time
	DueDate   time.Time

	Status       string `gorm:"type:varchar(255)"`
	InvoiceItems []InvoiceItem
	OrderID      string `gorm:"type:varchar(36)"`
}

func (i *Invoice) BeforeCreate(tx *gorm.DB) (err error) {
	i.ID = uuid.New().String()
	return
}
