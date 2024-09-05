package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Timezone struct {
	ID            string `gorm:"type:varchar(36);primaryKey"`
	Name          string `gorm:"unique;type:varchar(255)"`
	GmtOffset     int    `gorm:"type:int"`
	GmtOffsetName string `gorm:"type:varchar(255)"`
	Abbreviation  string `gorm:"type:varchar(10)"`
	TzName        string `gorm:"type:varchar(255)"`
}

func (t *Timezone) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New().String()
	return
}
