package logger

import "go.uber.org/zap"

func NewLogger() *zap.SugaredLogger {
	preLogger, _ := zap.NewProduction()
	defer preLogger.Sync()
	logger := preLogger.Sugar()
	return logger
}
