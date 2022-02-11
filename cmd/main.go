package main

import (
	"fmt"
	"log"

	"github.com/aburtasov/fibonaccisrv/pkg/config"
	"github.com/aburtasov/fibonaccisrv/pkg/handler"
	"github.com/aburtasov/fibonaccisrv/pkg/storage"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Hello,%s!HTTP server is starting!\n", cfg.User)

	storage := storage.NewRedisStorage(cfg.DBAddr)

	handler := handler.NewHandler(storage)
	router := gin.Default()

	router.GET("/fibonacci/:x,y", handler.GetFibonacci)
	router.POST("/fibonacci/:len", handler.CreateFibonacci)

	if err := router.Run(cfg.HTTPAddr); err != nil {

		log.Fatal(err)
	}

}
