// +build wireinject

package controllers

import (
	"stock/services"

	"github.com/google/wire"
)

var (
	StockControllerSet = wire.NewSet(NewStockController, services.CreateStockService)
)

// CreateStockController CreateStockController
func CreateStockController() *StockController {
	wire.Build(StockControllerSet)

	return nil
}
