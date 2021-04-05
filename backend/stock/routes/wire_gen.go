// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package routes

import (
	"github.com/google/wire"
	"net/http"
	"stock/controllers"
)

// Injectors from wire.go:

// InitRoute InitRoute
func InitRoute() http.Handler {
	stockController := controllers.CreateStockController()
	router := &Router{
		Stock: stockController,
	}
	handler := Load(router)
	return handler
}

// wire.go:

var (
	RouterSet = wire.NewSet(wire.Struct(new(Router), "*"), wire.Bind(new(IRouter), new(*Router)))
	RouteSet  = wire.NewSet(Load, RouterSet)
)
