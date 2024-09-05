package handlers

import (
	"net/http"

	"git.bopconsultancy.com/tejasc/bb-tailor-bulk-order-api/api/dtos"
	"git.bopconsultancy.com/tejasc/bb-tailor-bulk-order-api/internal/config"
	"git.bopconsultancy.com/tejasc/bb-tailor-bulk-order-api/internal/database/models"
	"git.bopconsultancy.com/tejasc/bb-tailor-bulk-order-api/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserHandler struct {
	db        *gorm.DB
	cfg       *config.Config
	validator *validator.Validate
}

func NewUserHandler(
	db *gorm.DB,
	cfg *config.Config,
	validator *validator.Validate,
) *UserHandler {
	return &UserHandler{
		db:        db,
		cfg:       cfg,
		validator: validator,
	}
}

func (h *UserHandler) GetAll(c echo.Context) error {
	var users []models.User

	if err := h.db.Preload("Roles").Preload("Profile").Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Error fetching users from database",
		})
	}

	var userResponses []dtos.UserResponseDTO
	for _, user := range users {
		userResponse := dtos.UserResponseDTO{
			ID:        user.ID,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Roles:     utils.TransformRolesToStringArray(user.Roles),
			Profile: dtos.UserProfileResponseDTO{
				ID:        user.Profile.ID,
				FirstName: user.Profile.FirstName,
				LastName:  user.Profile.LastName,
				Phone:     user.Profile.Phone,
				Age:       user.Profile.Age,
				Gender:    user.Profile.Gender,
			},
		}
		userResponses = append(userResponses, userResponse)
	}

	return c.JSON(http.StatusOK, userResponses)
}
