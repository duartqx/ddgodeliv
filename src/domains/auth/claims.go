package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	SessionUser
	jwt.RegisteredClaims
}

func GetNewClaims() *CustomClaims {
	return &CustomClaims{
		SessionUser:      *GetNewSessionUser(),
		RegisteredClaims: jwt.RegisteredClaims{},
	}
}

func GetNewPopulatedClaims(user *SessionUser, expiresAt time.Time) *CustomClaims {
	return &CustomClaims{
		SessionUser: *user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}
}
