package services

import (
	"context"
	"encoding/json"
	"fmt"
	"stock/models"
	"stock/repositories"
	"stock/websocket"
)

// StockService Stock Service
type StockService struct {
	stockRep *repositories.StockRepository
}

// NewStockService New Stock Service
func NewStockService(
	stockRep *repositories.StockRepository,
) *StockService {
	return &StockService{
		stockRep: stockRep,
	}
}

func (s *StockService) ReadWs(ctx context.Context, client *websocket.Client) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg, ok := <-client.ReadMessage:
			if !ok {
				break
			}
			var data struct {
				Type string          `json:"type"`
				Data json.RawMessage `json:"data"`
			}
			err := json.Unmarshal(msg, &data)
			if err != nil {
				break
			}
			// 先了解是傳甚麼資料過來
			if data.Type == "stock" {
				filter := models.WsStockData{}
				err := json.Unmarshal(data.Data, &filter)
				if err != nil {
					fmt.Println(err)
					break
				}
				// 檢查欄位
				if filter.StockNumber == "" {
					break
				}
				if filter.StartDate == nil {
					break
				}
				if filter.EndDate == nil {
					break
				}
				to_client := &models.WsToClientData{}
				stock, err := s.stockRep.GetStockDataByDate(filter.StockNumber, filter.StartDate, filter.EndDate)
				to_client.Data = stock
				jsondata, err := json.Marshal(to_client)
				client.Send(jsondata)
			}
		}
	}
}
