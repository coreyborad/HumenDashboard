package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InternalCORS() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOriginFunc = func(origin string) bool {
		return origin == "https://*.cabbageattic.com"
	}
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization")

	return cors.New(config)
}
