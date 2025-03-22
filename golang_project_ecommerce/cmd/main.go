// Package main ...
package main

import (
	"fmt"
	"net/http"
	"time"

	"golang_project_ecommerce/config"
	"golang_project_ecommerce/controllers"
	"golang_project_ecommerce/database"
	implRepository "golang_project_ecommerce/repository/impl"
	"golang_project_ecommerce/routers"
	implService "golang_project_ecommerce/services/Impl"
	"golang_project_ecommerce/utils"

	"github.com/go-playground/validator/v10"
)

// Main Function Heart of Application
func main() {
	initConfigErr := utils.InitConfig()
	if initConfigErr != nil {
		utils.Logger.Println("Error init config fail.")
	}

	db := database.Connection()
	validate := validator.New()

	//Init Repository
	repository := implRepository.NewProductRepositoryImpl(db)

	//Init Service
	service := implService.NewProductServiceImpl(repository, validate)

	//Init controller
	controller := controllers.NewProductController(service)

	//Router
	router := routers.NewRouter(controller)

	fmt.Println("Server is Running on http://localhost:" + config.GetConfig().Port)

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
