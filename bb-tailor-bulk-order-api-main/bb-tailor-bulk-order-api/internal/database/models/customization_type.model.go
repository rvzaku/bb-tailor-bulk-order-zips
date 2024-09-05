package models

type CustomizationType struct {
	Name           string          `gorm:"type:varchar(255);primaryKey"`
	Customizations []Customization `gorm:"foreignKey:Type"`
}
