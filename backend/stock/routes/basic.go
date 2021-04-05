package routes

import (
	"net/http"
	"stock/config"
	"stock/middleware"

	"github.com/gin-gonic/gin"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
	// _ "github.com/swaggo/gin-swagger/example/basic/docs"
)

// Load initializes the routing of the application.
// @title API
// @version 1.0
// @description This is a api for Service.
// @BasePath /api/v1
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func Load(r IRouter) http.Handler {
	if config.App.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	e := gin.New()

	// url := ginSwagger.URL(config.Swagger.URL) // The url pointing to API definition
	// e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// middleware
	middleware.RouteMiddleware(e)

	serviceRoute := e.Group("/")
	// api
	api := serviceRoute.Group("/api")
	{
		r.RegisterAPI(api)
	}

	// websocket
	ws := serviceRoute.Group("/ws")
	{
		r.RegisterWebSocket(ws)
	}

	e.NoMethod(NotFound)

	return e
}

// NotFound represents the 404 page.
func NotFound(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusNotFound, "PAGE NOT FOUND")
	return
}
