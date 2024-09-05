package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Currency struct {
	ID            string `gorm:"type:varchar(36);primaryKey"`
	Code          string `gorm:"unique;type:varchar(5)"`
	Name          string `gorm:"type:varchar(255)"`
	NamePlural    string `gorm:"type:varchar(255)"`
	Symbol        string `gorm:"type:varchar(10)"`
	SymbolNative  string `gorm:"type:varchar(10)"`
	DecimalDigits int    `gorm:"type:int"`
	Rounding      int    `gorm:"type:int"`
}

func (c *Currency) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()
	return
}
