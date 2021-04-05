package middleware

import (
	"net/http"
	"time"

	"github.com/didip/tollbooth/v6"
	"github.com/didip/tollbooth/v6/limiter"
	"github.com/gin-gonic/gin"
)

var rateLimiter = tollbooth.NewLimiter(10, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour})

// RateLimit RateLimit
func RateLimit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := tollbooth.LimitByRequest(rateLimiter, ctx.Writer, ctx.Request); err != nil {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": err.Message,
			})
			return
		}

		ctx.Next()
		return
	}
}
