package middlewares

import (
	"net/http"

	"git.bopconsultancy.com/tejasc/bb-tailor-bulk-order-api/internal/database/models"
	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CasbinRBACMiddlewareConfig struct {
	DB       *gorm.DB
	Enforcer *casbin.Enforcer
}

func CasbinRBACMiddleware(config *CasbinRBACMiddlewareConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Ensure userID is set in the context by JWT middleware
			userID := c.Get("userID").(string)

			var user models.User
			if err := config.DB.Preload("Roles").First(&user, "id = ?", userID).Error; err != nil {
				return echo.NewHTTPError(
					http.StatusInternalServerError,
					"Error fetching user roles",
				)
			}

			// Check permissions for each role
			obj := c.Request().URL.Path // object (endpoint)
			act := c.Request().Method   // action (HTTP method)

			for _, role := range user.Roles {
				if ok, err := config.Enforcer.Enforce(role.Name, obj, act); err != nil {
					return echo.NewHTTPError(
						http.StatusInternalServerError,
						"Error enforcing policy",
					)
				} else if ok {
					return next(c)
				}
			}

			return echo.NewHTTPError(http.StatusForbidden, "Permission denied")
		}
	}
}
