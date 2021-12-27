//go:build wireinject
// +build wireinject

package routes

import (
	"github.com/google/wire"
	"net/http"
	"stock/controllers"
)

var (
	RouterSet = wire.NewSet(wire.Struct(new(Router), "*"), wire.Bind(new(IRouter), new(*Router)))
	RouteSet  = wire.NewSet(Load, RouterSet)
)

// InitRoute InitRoute
func InitRoute() http.Handler {
	wire.Build(
		controllers.CreateStockController,
		controllers.CreateXlsxController,
		RouteSet,
	)

	return nil
}
