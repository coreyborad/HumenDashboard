package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// CustomClaims CustomClaims
type CustomClaims struct {
	GrantType string    `json:"grant_type"`
	UserID    uint64    `json:"user_id"`
	ClientID  uuid.UUID `json:"client_id"`
	ReSub     uint64    `json:"sub"`
}

// Claims Claims
type Claims struct {
	jwt.StandardClaims

	CustomClaims
}
