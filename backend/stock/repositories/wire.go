// +build wireinject

package repositories

import (
	"stock/database"
	"stock/mongodb"

	"github.com/google/wire"
)

var (
	UserRepositorySet  = wire.NewSet(NewUserRepository, database.GetDB)
	StockRepositorySet = wire.NewSet(NewStockRepository, database.GetDB, mongodb.GetMongoDB)
)

// CreateUserRepository CreateUserRepository
func CreateUserRepository() *UserRepository {
	wire.Build(UserRepositorySet)

	return nil
}

// CreateStockRepository CreateStockRepository
func CreateStockRepository() *StockRepository {
	wire.Build(StockRepositorySet)

	return nil
}
