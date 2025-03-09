// Package database ...
package database

import (
	"fmt"
	"strconv"
	"time"

	"golang_project_ecommerce/config"
	"golang_project_ecommerce/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connection is to establish the connection with postgresSql.
func Connection() *gorm.DB {
	var (
		host     = config.GetConfig().DatabaseHost
		port     = config.GetConfig().DatabasePort
		user     = config.GetConfig().DatabaseUser
		password = config.GetConfig().DatabasePassword
		dbName   = config.GetConfig().DatabaseName
		sslMode  = config.GetConfig().DatabaseSSLMode
	)

	portInt, err := strconv.Atoi(port)
	if err != nil {
		utils.Logger.Printf("Error converting port to integer: %v\n", err)

		return nil
	}

	postgresSQL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, portInt, user, password, dbName, sslMode)

	database, err := gorm.Open(postgres.Open(postgresSQL), &gorm.Config{
		NowFunc: time.Now,
	})
	if err != nil {
		utils.Logger.Printf("Error connecting to database: %v\n", err)

		return nil
	}

	utils.Logger.Println("Established a successful DB connection!")

	return database
}
