// Package config ...
package config

import (
	"net/http"

	"golang_project_ecommerce/utils"

	"github.com/gin-gonic/gin"
)

// CorsOrigin is to handle cors policy error between frontend and backend.
func CorsOrigin() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.Writer.Header().Set("Access-Control-Allow-Origin", GetConfig().CorsOrigin)
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

// Config is.
type Config struct {
	Port             string
	GinMode          string
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseHost     string
	DatabasePort     string
	DatabaseSSLMode  string
	CorsOrigin       string
}

// GetConfig is to load the Env variable.
func GetConfig() *Config {
	config := &Config{
		Port:             utils.GetEnv("PORT", "8080"),
		GinMode:          utils.GetEnv("GIN_MODE", "release"), // debug, release, test
		DatabaseName:     utils.GetEnv("DATABASE_NAME", ""),
		DatabaseUser:     utils.GetEnv("DATABASE_USER", ""),
		DatabasePassword: utils.GetEnv("DATABASE_PASSWORD", ""),
		DatabaseHost:     utils.GetEnv("DATABASE_HOST", ""),
		DatabasePort:     utils.GetEnv("DATABASE_PORT", ""),
		DatabaseSSLMode:  utils.GetEnv("DATABASE_SSL_MODE", "prefer"),
		CorsOrigin:       utils.GetEnv("CORS_ORIGIN", ""),
	}

	return config
}
