package logger

import (
	"sync"

	"go.uber.org/zap"
)

var (
	log  *zap.Logger
	once sync.Once
)

// InitLogger initializes the zap logger (called once)
func InitLogger() {
	once.Do(func() {
		logger, err := zap.NewProduction() // Change to zap.NewDevelopment() for pretty logs
		if err != nil {
			panic("Failed to initialize logger: " + err.Error())
		}
		log = logger
	})
}

// GetLogger returns the shared logger instance
func GetLogger() *zap.Logger {
	if log == nil {
		InitLogger()
	}
	return log
}
