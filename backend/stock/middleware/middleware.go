package middleware

import (
	"stock/config"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

// RouteMiddleware 設定 Middleware
func RouteMiddleware(e *gin.Engine) *gin.Engine {
	// Recovery middleware
	e.Use(
		gin.Recovery(),
		RateLimit(), //流量限制
		// InternalCORS(),

		gzip.Gzip(gzip.DefaultCompression),
	)

	// Logger middleware
	if config.App.Debug {
		e.Use(gin.Logger())
	}

	return e
}
