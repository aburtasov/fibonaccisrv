package main

import (
	"github.com/aburtasov/fibonaccisrv/pkg/config"
	"github.com/aburtasov/fibonaccisrv/pkg/handler"
	"github.com/aburtasov/fibonaccisrv/pkg/storage"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)

	cfg, err := config.NewConfig()
	if err != nil {
		logger.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	redisStorage := storage.NewRedisStorage(cfg.DBAddr)
	defer func() {
		if err := redisStorage.Close(); err != nil {
			logger.Errorf("Ошибка при закрытии подключения к Redis: %v", err)
		}
	}()

	h := handler.NewHandler(redisStorage, logger)

	router := gin.Default()

	router.POST("/fibonacci/:len", h.CreateFibonacci)

	router.GET("/fibonacci", h.GetFibonacci)

	logger.Infof("Запуск сервера на %s", cfg.HTTPAddr)
	if err := router.Run(cfg.HTTPAddr); err != nil {
		logger.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
