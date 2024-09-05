package handlers

import (
	"net/http"

	"git.bopconsultancy.com/tejasc/bb-tailor-bulk-order-api/api/dtos"
	"github.com/labstack/echo/v4"
)

// HealthHandler struct holds the necessary dependencies for the health handler
type HealthHandler struct {
	// You can inject dependencies here if needed
}

// NewHealthHandler creates a new instance of HealthHandler
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) GetHealth(c echo.Context) error {
	response := dtos.HealthGetResponseDTO{
		Message: "server is running fine.",
		Status:  "OK",
	}

	return c.JSON(http.StatusOK, response)
}
