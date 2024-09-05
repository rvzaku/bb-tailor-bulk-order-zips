package seeders

import (
	"fmt"
	"log"

	"git.bopconsultancy.com/tejasc/bb-tailor-bulk-order-api/internal/database/models"
	"gorm.io/gorm"
)

func SeedOrderStatuses(db *gorm.DB) {
	order_statuses := []models.OrderStatus{
		{Name: "PENDING"},
		{Name: "PROCESSING"},
		{Name: "ON_HOLD"},
		{Name: "COMPLETED"},
		{Name: "CANCELLED"},
		{Name: "REFUNDED"},
		{Name: "FAILED"},
		{Name: "AWAITING_PAYMENT"},
		{Name: "AWAITING_FULFILLMENT"},
		{Name: "AWAITING_SHIPMENT"},
		{Name: "SHIPPED"},
		{Name: "DELIVERED"},
		{Name: "AWAITING_PICKUP"},
	}

	for _, order_status := range order_statuses {
		if err := db.FirstOrCreate(&order_status, models.OrderStatus{Name: order_status.Name}).Error; err != nil {
			log.Fatalf("failed to seed order status %s: %v", order_status.Name, err)
		}
	}

	fmt.Println("order statuses table seeding done successfully!")
}
