package middlewares

import (
	"git.bopconsultancy.com/tejasc/bb-tailor-bulk-order-api/internal/config"
	jwtUtils "git.bopconsultancy.com/tejasc/bb-tailor-bulk-order-api/internal/utils"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JwtMiddleware(cfg *config.Config) echo.MiddlewareFunc {
	config := echojwt.Config{
		SigningKey: []byte(cfg.JwtAuth.AccessTokenSecret),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtUtils.CustomClaims)
		},
		SuccessHandler: func(c echo.Context) {
			claims := c.Get("user").(*jwt.Token).Claims.(*jwtUtils.CustomClaims)
			c.Set("userID", claims.Subject) // Set the user ID in context
		},
	}

	return echojwt.WithConfig(config)
}

func JwtRefreshMiddleware(cfg *config.Config) echo.MiddlewareFunc {
	config := echojwt.Config{
		SigningKey: []byte(cfg.JwtAuth.RefreshTokenSecret),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtUtils.CustomClaims)
		},
	}

	return echojwt.WithConfig(config)
}
