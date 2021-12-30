//go:build wireinject
// +build wireinject

package services

import (
	"github.com/google/wire"
	"stock/repositories"
)

var (
	UserServiceSet  = wire.NewSet(NewUserService, repositories.CreateUserRepository)
	StockServiceSet = wire.NewSet(NewStockService, repositories.CreateStockRepository)
	XlsxServiceSet  = wire.NewSet(NewXlsxService)
)

// CreateUserService CreateUserService
func CreateUserService() *UserService {
	wire.Build(UserServiceSet)

	return nil
}

// CreateStockService CreateStockService
func CreateStockService() *StockService {
	wire.Build(StockServiceSet)

	return nil
}

// CreateXlsxService CreateXlsxService
func CreateXlsxService() *XlsxService {
	wire.Build(XlsxServiceSet)

	return nil
}
