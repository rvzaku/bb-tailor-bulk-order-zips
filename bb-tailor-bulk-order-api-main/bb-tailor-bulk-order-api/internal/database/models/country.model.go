package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Country struct {
	ID          string `gorm:"type:varchar(36);primaryKey"`
	Iso2        string `gorm:"unique;type:varchar(2)"`
	Iso3        string `gorm:"unique;type:varchar(3)"`
	Name        string `gorm:"unique;type:varchar(255)"`
	NumericCode string `gorm:"unique;type:varchar(3)"`
	PhoneCode   string `gorm:"type:varchar(5)"`
	Tld         string `gorm:"type:varchar(5)"`
	Native      string `gorm:"type:varchar(255)"`
	Region      string `gorm:"type:varchar(255)"`
	SubRegion   string `gorm:"type:varchar(255)"`
	RegionId    string `gorm:"type:varchar(5)"`
	SubRegionId string `gorm:"type:varchar(5)"`
	Nationality string `gorm:"type:varchar(255)"`

	States     []State
	Timezones  []*Timezone `gorm:"many2many:countries_timezones;"`
	Currencies []*Currency `gorm:"many2many:countries_currencies;"`
}

func (c *Country) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()
	return
}
