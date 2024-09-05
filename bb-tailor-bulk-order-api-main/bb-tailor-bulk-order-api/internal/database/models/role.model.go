package models

type Role struct {
	Name string `gorm:"type:varchar(255);primaryKey"`
}
