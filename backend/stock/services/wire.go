// +build wireinject

package services

import (
	"github.com/google/wire"
	"stock/repositories"
)

var (
	UserServiceSet  = wire.NewSet(NewUserService, repositories.CreateUserRepository)
	StockServiceSet = wire.NewSet(NewStockService, repositories.CreateStockRepository)
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
