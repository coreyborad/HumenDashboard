package middleware

import (
	"net/http"
	"stock/config"
	"stock/models"
	"stock/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

// AuthRequired 認證JWT
func AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, token, err := GetUserFromRequest(ctx.Request)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.Set("user", user)
		ctx.Set("token", token)
		ctx.Next()
	}
}

// GetUserFromRequest GetUserFromRequest
func GetUserFromRequest(req *http.Request) (user *models.User, token *jwt.Token, err error) {
	userServ := services.CreateUserService()
	token, err = request.ParseFromRequestWithClaims(
		req,
		request.OAuth2Extractor,
		&models.Claims{},
		func(t *jwt.Token) (interface{}, error) {
			user, err = userServ.GetUserFromToken(t.Raw)
			if err != nil {
				return nil, err
			}

			return []byte(config.App.Key), nil
		},
	)
	if err != nil {
		return nil, nil, err
	}

	return
}
