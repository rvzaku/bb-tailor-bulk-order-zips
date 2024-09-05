package models

type InvoiceStatus struct {
	Name     string    `gorm:"type:varchar(255);primaryKey"`
	Invoices []Invoice `gorm:"foreignKey:Status"`
}
