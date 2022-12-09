package main

import (
	"errors"

	"github.com/tiena2cva/go-log"
)

func main() {
	log.Debug("debug")
	log.Error("error")

	logger := log.New(log.DebugLevel)
	logger.Debug("debug")
	logger.Error("error")
	logger.Err(errors.New("test error"))

	log.UseProduction()
	logger = log.New(log.InfoLevel)
	logger.Debug("debug")
	logger.Error("error")
	logger.Err(errors.New("test error"))

	log.LogFile("log.txt")
	log.Debug("debug")
	log.Error("error")
	log.Err(errors.New("test error"))
}
