// Package routers ...
package routers

import (
	"golang_project_ecommerce/controllers"
	"golang_project_ecommerce/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(productController *controllers.ProductController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	productRouter := router.Group("/api/v1")
	tagRouter := productRouter.Group("/product", middleware.AuthMiddleware)
	tagRouter.GET("", productController.FindAll)
	tagRouter.GET("/:id", productController.FindById)
	tagRouter.POST("", productController.Create)
	tagRouter.PATCH("/:id", productController.Update)
	tagRouter.DELETE("/:id", productController.Delete)

	return router
}
