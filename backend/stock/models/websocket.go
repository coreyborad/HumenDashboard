package models

import "time"

type WsToClientData struct {
	Data interface{} `json:"data"`
}

type WsStockData struct {
	StockNumber string     `json:"stock_number"`
	StartDate   *time.Time `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
}
