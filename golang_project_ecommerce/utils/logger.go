// Package utils ...
package utils

import (
	"log"
	"os"
	"time"
)

// Logger is.
var Logger *log.Logger

// InitConfig is to initialize the log.
func InitConfig() error {
	InitLogger()

	if err := LoadEnv(); err != nil {
		return err
	}

	Logger.Printf("Application time, %v\n", time.Now().Format(time.RFC3339Nano))

	return nil
}

// InitLogger is to show log.
func InitLogger() {
	Logger = log.New(os.Stdout, "[LOG]: ", log.Ldate|log.Ltime|log.Llongfile)
}
