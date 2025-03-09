// Package utils ...
package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CorsOrigin is to handle cors policy error between frontend and backend.
func CorsOrigin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// ctx.Writer.Header().Set("Access-Control-Allow-Origin",)
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusNoContent)

			return
		}

		ctx.Next()
	}
}
