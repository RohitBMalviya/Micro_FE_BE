// Package config ...
package config

import "golang_project_ecommerce/utils"

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
