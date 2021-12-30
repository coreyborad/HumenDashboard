//go:build wireinject
// +build wireinject

package controllers

import (
	"stock/services"

	"github.com/google/wire"
)

var (
	StockControllerSet = wire.NewSet(NewStockController, services.CreateStockService)
	XlsxControllerSet  = wire.NewSet(NewXlsxController, services.CreateXlsxService)
)

// CreateStockController CreateStockController
func CreateStockController() *StockController {
	wire.Build(StockControllerSet)

	return nil
}

// CreateXlsxController CreateXlsxController
func CreateXlsxController() *XlsxController {
	wire.Build(XlsxControllerSet)

	return nil
}
