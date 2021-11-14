package tools

import (
	"log"
	"os"
)

var logger *log.Logger

func GetLogger() *log.Logger {
	if logger == nil {
		logger = log.New(os.Stdout, "go-micro ", log.LstdFlags)
	}
	return logger
}
