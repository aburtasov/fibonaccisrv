package logger

import (
	"log"

	"go.uber.org/zap"
)

func NewLogger() *zap.SugaredLogger {
	preLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	defer preLogger.Sync()

	logger := preLogger.Sugar()
	return logger
}
