package utils

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type CustomClaims struct {
	Id        string   `json:"id"`
	Email     string   `json:"email"`
	Roles     []string `json:"roles"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(
	id string,
	email string,
	roles []string,
	firstName string,
	lastName string,
	secretKey string,
	expDays int,
) (string, error) {
	accessTokenExp := time.Duration(expDays) * 24 * time.Hour

	claims := CustomClaims{
		id,
		email,
		roles,
		firstName,
		lastName,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokenExp)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "bb-tailor-bulk-orders-api",
			Subject:   id,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func GenerateRefreshToken(
	id string,
	email string,
	roles []string,
	firstName string,
	lastName string,
	secretKey string,
	expMonths int,
) (string, error) {
	refreshTokenExp := time.Duration(expMonths) * 30 * 24 * time.Hour
	claims := CustomClaims{
		id,
		email,
		roles,
		firstName,
		lastName,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshTokenExp)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "bb-tailor-bulk-orders-api",
			Subject:   id,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func ExtractToken(c echo.Context) (string, error) {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("missing authorization header")
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid authorization format")
	}

	return parts[1], nil
}
