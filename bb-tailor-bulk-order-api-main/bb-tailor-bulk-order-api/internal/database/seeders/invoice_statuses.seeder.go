package seeders

import (
	"fmt"
	"log"

	"git.bopconsultancy.com/tejasc/bb-tailor-bulk-order-api/internal/database/models"
	"gorm.io/gorm"
)

func SeedInvoiceStatuses(db *gorm.DB) {
	invoice_statuses := []models.InvoiceStatus{
		{Name: "DRAFT"},
		{Name: "SENT"},
		{Name: "PAID"},
		{Name: "PARTIALLY_PAID"},
		{Name: "OVERDUE"},
		{Name: "CANCELLED"},
		{Name: "REFUNDED"},
		{Name: "VOID"},
		{Name: "PENDING"},
	}

	for _, invoice_status := range invoice_statuses {
		if err := db.FirstOrCreate(&invoice_status, models.InvoiceStatus{Name: invoice_status.Name}).Error; err != nil {
			log.Fatalf("failed to seed invoice status %s: %v", invoice_status.Name, err)
		}
	}

	fmt.Println("invoice statuses table seeding done successfully!")
}
