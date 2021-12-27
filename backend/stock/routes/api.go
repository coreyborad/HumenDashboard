package routes

import (
	"net/http"
	"stock/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterAPI RegisterAPI
func (r *Router) RegisterAPI(api *gin.RouterGroup) error {
	v1 := api.Group("/v1")
	{
		v1.GET("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "API v1",
			})
		})
		v1.POST("/xlsx", r.Xlsx.XlsxAppendRecord)
		authorized := v1.Use(middleware.AuthRequired())
		{
			authorized.GET("/testToken", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "HaveToken",
				})
			})
		}

		// v1.POST("/mqtt/acl", r.Certificate.PostACL)
		// v1.GET("/mqtt/acl", r.Certificate.GetACL)

		// authorized := v1.Use(middleware.AuthRequired(), middleware.RateLimit())
		// {
		// 	authorized.GET("/", func(ctx *gin.Context) {
		// 		ctx.JSON(http.StatusOK, gin.H{
		// 			"message": "API v1",
		// 		})
		// 	})
		// }
	}

	return nil
}
