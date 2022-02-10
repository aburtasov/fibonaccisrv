package main

import (
	"log"

	"github.com/aburtasov/fibonaccisrv/pkg/handler"
	"github.com/aburtasov/fibonaccisrv/pkg/storage"

	"github.com/gin-gonic/gin"
)

func main() {

	// Добавить работу с конфигом
	// Добавить логер
	// добавить graceful shutdown

	storage := storage.NewRedisStorage()
	handler := handler.NewHandler(storage)
	router := gin.Default()

	router.GET("/fibonacci/:x,y", handler.GetFibonacci)
	router.POST("/fibonacci/:len", handler.CreateFibonacci)

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}

}
