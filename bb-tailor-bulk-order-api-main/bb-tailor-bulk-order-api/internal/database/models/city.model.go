package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type City struct {
	ID   string `gorm:"type:varchar(36);primary_key"`
	Name string `gorm:"type:varchar(255)"`

	StateID string `gorm:"type:varchar(36)"`
}

func (c *City) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()
	return
}
