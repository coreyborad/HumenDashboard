//go:build wireinject
// +build wireinject

package concrete

import (
	"stock/services"

	"github.com/google/wire"
)

var (
	StockConcreteSet = wire.NewSet(NewStockConcrete, services.CreateStockService)
)

// CreateStockService CreateStockService
func CreateStockConcrete() *StockConcrete {
	wire.Build(StockConcreteSet)

	return nil
}
