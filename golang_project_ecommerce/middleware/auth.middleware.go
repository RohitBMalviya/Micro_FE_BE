// Package middleware...
package middleware

import (
	"golang_project_ecommerce/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	utils.Logger.Println("Auth Middleware")
	ctx.Next()
}
