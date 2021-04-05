package models

import (
	"errors"
	"time"
)

// Error constants
var (
	ErrStockNotExist = errors.New("Stock Not Exist")
)

type Stock struct {
	StockNumber string     `json:"stock_number" gorm:"primaryKey"`
	StockName   string     `json:"stock_name"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type StockData struct {
	StockNumber    string     `json:"stock_number" bson:"stock_number"`
	DealDate       *time.Time `json:"deal_date" bson:"deal_date"`
	DealCount      uint64     `json:"deal_count" bson:"deal_count"`
	PriceOnOpen    float64    `json:"price_on_open" bson:"price_on_open"`
	PriceOnHighest float64    `json:"price_on_highest" bson:"price_on_highest"`
	PriceOnLowest  float64    `json:"price_on_lowest" bson:"price_on_lowest"`
	PriceOnClose   float64    `json:"price_on_close" bson:"price_on_close"`
	CreatedAt      *time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at" bson:"updated_at"`
}
