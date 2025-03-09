// Package main ...
package main

import (
	"fmt"
	"net/http"
	"time"

	"golang_project_ecommerce/config"
	"golang_project_ecommerce/database"
	"golang_project_ecommerce/utils"

	"github.com/gin-gonic/gin"
)

// Main Function Heart of Application
func main() {
	initConfigErr := utils.InitConfig()
	if initConfigErr != nil {
		utils.Logger.Println("Error init config fail.")
	}

	router := gin.Default()

	router.Use(config.CorsOrigin())

	fmt.Println("Server is Running on http://localhost:" + config.GetConfig().Port)

	database.Connection()

	server := &http.Server{
		Addr:           ":" + config.GetConfig().Port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
