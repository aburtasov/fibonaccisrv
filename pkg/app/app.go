package app

import (
	"github.com/aburtasov/fibonaccisrv/pkg/config"
	"github.com/aburtasov/fibonaccisrv/pkg/handler"
	"github.com/aburtasov/fibonaccisrv/pkg/logger"
	"github.com/aburtasov/fibonaccisrv/pkg/storage"
	"github.com/gin-gonic/gin"
)

func Run() {

	logger := logger.NewLogger()

	cfg, err := config.NewConfig()
	if err != nil {
		logger.Fatal(err)
	}

	storage := storage.NewRedisStorage(cfg.DBAddr)
	handler := handler.NewHandler(storage)
	router := gin.Default()

	router.GET("/fibonacci/:x,y", handler.GetFibonacci)
	router.POST("/fibonacci/:len", handler.CreateFibonacci)

	if err := router.Run(cfg.HTTPAddr); err != nil {

		logger.Fatal(err)
	}

}
