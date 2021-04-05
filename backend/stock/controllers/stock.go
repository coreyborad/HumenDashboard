package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"stock/models"
	"stock/services"
	"stock/websocket"

	"github.com/gin-gonic/gin"
)

type StockController struct {
	serv *services.StockService
}

// NewStockController NewStockController
func NewStockController(service *services.StockService) *StockController {
	return &StockController{
		serv: service,
	}
}

// GetInfo GetInfo
func (c *StockController) StockWs(ctx *gin.Context) {
	if user, ok := ctx.MustGet("user").(*models.User); ok {
		client, err := websocket.NewClient(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 傳送歡迎訊息
		msg := &models.WsToClientData{}
		msg.Data = fmt.Sprintf("Welcome, %s", user.Name)
		message, _ := json.Marshal(&msg)
		client.Send(message)

		// 設定接收Channel
		go c.serv.ReadWs(client.Context(), client)
	}
}
