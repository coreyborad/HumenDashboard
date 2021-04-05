package auth

import (
	"encoding/json"
	"errors"
	"stock/config"
	"stock/models"
	"stock/utils"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Error constants
var (
	ErrTokenExpired         = errors.New("Token has expired and can no longer be refreshed")
	ErrTokenInvalid         = errors.New("Token is invalid")
	ErrUnsupportedGrantType = errors.New("Unsupported grant type")
)

// CreateToken Create JWT
func CreateToken(customClaims models.CustomClaims) (token string, expire time.Time, err error) {
	if customClaims.GrantType != "password" && customClaims.GrantType != "refresh_token" && customClaims.GrantType != "client_credentials" {
		return token, expire, ErrUnsupportedGrantType
	}

	expire = time.Now().Add(config.App.TTL * time.Second)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.Claims{
		StandardClaims: jwt.StandardClaims{
			Id:        utils.UUID(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    config.Server.Host,
			NotBefore: time.Now().Unix(),
			ExpiresAt: expire.Unix(),
		},
		CustomClaims: customClaims,
	})

	token, err = claims.SignedString([]byte(config.App.Key))

	return
}

// RefreshToken Refresh JWT
func RefreshToken(tokenString string) (refreshToken string, err error) {
	token, err := ParseToken(tokenString)

	if payload, ok := token.Claims.(*models.Claims); ok {
		refreshTTL := time.Unix(payload.IssuedAt, 0).Add(config.App.RefreshTTL * time.Second).Unix()
		if time.Now().Unix() > refreshTTL || payload.IssuedAt > time.Now().Unix() {
			return "", ErrTokenExpired
		}

		expiresAt := time.Now().Add(config.App.TTL * time.Second).Unix()
		if expiresAt > refreshTTL {
			expiresAt = refreshTTL
		}

		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.Claims{
			StandardClaims: jwt.StandardClaims{
				Id:        utils.UUID(),
				IssuedAt:  payload.IssuedAt,
				Issuer:    payload.Issuer,
				NotBefore: time.Now().Unix(),
				ExpiresAt: expiresAt,
			},
			CustomClaims: models.CustomClaims{
				GrantType: "refresh_token",
				UserID:    payload.UserID,
				ClientID:  payload.ClientID,
			},
		})

		refreshToken, err = claims.SignedString([]byte(config.App.Key))
	}

	return
}

// ParseToken Parse JWT
func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &models.Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.App.Key), nil
	})
}

// Decode Decode
func Decode(tokenString string) (claims *models.Claims, err error) {
	payload := strings.Split(tokenString, ".")
	if len(payload) != 3 {
		return nil, ErrTokenInvalid
	}

	bytes, err := jwt.DecodeSegment(payload[1])
	if err != nil {
		return nil, err
	}

	json.Unmarshal(bytes, &claims)

	return
}
