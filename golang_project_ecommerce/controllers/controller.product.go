// Package controllers ...
package controllers

import (
	"golang_project_ecommerce/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	tagService services.ProductService
}

// NewProductController is
func NewProductController(service services.ProductService) *ProductController {
	return &ProductController{tagService: service}
}

// Create controller
func (controller *ProductController) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "")
}

// Update controller
func (controller *ProductController) Update(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, "")
}

// Delete controller
func (controller *ProductController) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "")

}

// FindById controller
func (controller *ProductController) FindById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "")
}

// FindAll controller
func (controller *ProductController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Product")

}
