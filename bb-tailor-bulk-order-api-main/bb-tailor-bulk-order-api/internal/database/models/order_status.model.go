package models

type OrderStatus struct {
	Name   string  `gorm:"type:varchar(255);primaryKey"`
	Orders []Order `gorm:"foreignKey:Status"`
}
