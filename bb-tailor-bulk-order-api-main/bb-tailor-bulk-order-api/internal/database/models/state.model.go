package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type State struct {
	ID        string `gorm:"type:varchar(36);primaryKey"`
	Name      string `gorm:"type:varchar(255)"`
	StateCode string `gorm:"type:varchar(10)"`

	CountryID string `gorm:"type:varchar(36)"`
	Cities    []City
}

func (s *State) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New().String()
	return
}
