package handlers

import (
	"errors"
	"net/http"

	"git.bopconsultancy.com/tejasc/bb-tailor-bulk-order-api/api/dtos"
	"git.bopconsultancy.com/tejasc/bb-tailor-bulk-order-api/internal/config"
	"git.bopconsultancy.com/tejasc/bb-tailor-bulk-order-api/internal/database/models"
	"git.bopconsultancy.com/tejasc/bb-tailor-bulk-order-api/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthHandler struct {
	db        *gorm.DB
	cfg       *config.Config
	validator *validator.Validate
}

func NewAuthHandler(
	db *gorm.DB,
	cfg *config.Config,
	validator *validator.Validate,
) *AuthHandler {
	return &AuthHandler{
		db:        db,
		cfg:       cfg,
		validator: validator,
	}
}

func (h *AuthHandler) Login(c echo.Context) error {
	var loginPayload dtos.LoginPayloadDTO
	err := c.Bind(&loginPayload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid login payload or malformed json",
		})
	}

	err = h.validator.Struct(loginPayload)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return c.JSON(http.StatusBadRequest, utils.FormatValidationErrors(validationErrors))
	}

	var user models.User
	result := h.db.Preload("Roles").Preload("Profile").First(&user, "email = ?", loginPayload.Email)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "unauthorized, user with email: " + loginPayload.Email + " does not exist",
		})
	}

	if utils.CheckPasswordHash(loginPayload.Password, user.PasswordHash) == false {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "unauthorized, user with email: " + loginPayload.Email + " has entered the wrong password",
		})
	}

	accessToken, err := utils.GenerateAccessToken(
		user.ID,
		user.Email,
		utils.TransformRolesToStringArray(user.Roles),
		user.Profile.FirstName,
		user.Profile.LastName,
		h.cfg.JwtAuth.AccessTokenSecret,
		h.cfg.JwtAuth.AccessTokenExpDays,
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "internal server error, failed to generate access token",
		})
	}

	refreshToken, err := utils.GenerateRefreshToken(
		user.ID,
		user.Email,
		utils.TransformRolesToStringArray(user.Roles),
		user.Profile.FirstName,
		user.Profile.LastName,
		h.cfg.JwtAuth.RefreshTokenSecret,
		h.cfg.JwtAuth.RefreshTokenExpMonths,
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "internal server error, failed to generate refresh token",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}
