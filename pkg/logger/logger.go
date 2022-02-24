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

	/* defer func() {
		err := preLogger.Sync()
		if err != nil {
			log.Fatal(err)
		}
	}() */

	logger := preLogger.Sugar()
	return logger
}
