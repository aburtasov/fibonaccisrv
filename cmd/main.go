package main

import (
	"log"

	"github.com/aburtasov/fibonaccisrv/pkg/config"
	"github.com/aburtasov/fibonaccisrv/pkg/handler"
	"github.com/aburtasov/fibonaccisrv/pkg/storage"

	"github.com/gin-gonic/gin"
)

func main() {

<<<<<<< HEAD
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	storage := storage.NewRedisStorage(cfg.DBAddr)
=======
	// Добавить работу с конфигом
	// Добавить логер
	// добавить graceful shutdown

	storage := storage.NewRedisStorage()
>>>>>>> 2083c60c2a390d3afe97cf52b4dd4fe6a90eb47a
	handler := handler.NewHandler(storage)
	router := gin.Default()

	router.GET("/fibonacci/:x,y", handler.GetFibonacci)
	router.POST("/fibonacci/:len", handler.CreateFibonacci)

	if err := router.Run(cfg.HTTPAddr); err != nil {
		log.Fatal(err)
	}

}
