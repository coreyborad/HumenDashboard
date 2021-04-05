package routes

import (
	"stock/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterWebSocket RegisterWebSocket
func (r *Router) RegisterWebSocket(ws *gin.RouterGroup) error {
	ws.Use(middleware.AuthRequired())
	{
		ws.GET("", r.Stock.StockWs)
	}

	return nil
}
