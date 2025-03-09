package utils

import (
	"log"
	"os"
)

func GetCirrLogger() *log.Logger {
	logger := log.New(os.Stdout, "cirr: ", log.LstdFlags)

	return logger
}
