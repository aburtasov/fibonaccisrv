package main

import (
	handler "github/aburtasov/fibonaccisrv/pkg/handler"
	storage "github/aburtasov/fibonaccisrv/pkg/storage"

	"github.com/gin-gonic/gin"
)

func main() {

	storage := storage.NewRedisStorage()
	handler := handler.NewHandler(storage)
	router := gin.Default()

	router.GET("/fibonacci/:x,y", handler.GetFibonacci)
	router.POST("/fibonacci/:len", handler.CreateFibonacci)

	router.Run()
}
