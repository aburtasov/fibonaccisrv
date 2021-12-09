package main

import (
	"context"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"

	"github.com/go-redis/redis/v8"
)

func main() {

	storage := NewRedisStorage()
	handler := NewHandler(storage)
	router := gin.Default()

	router.GET("/fibonacci/:x,y", handler.GetFibonacci)
	router.POST("/fibonacci/:len", handler.CreateFibonacci)

	router.Run()
}

type RedisStorage struct {
	rdb   *redis.Client
	mutex sync.Mutex
}

type Storage interface {
	Insert(length int)
	Get(x, y int) []int
}

func NewRedisStorage() *RedisStorage {
	return &RedisStorage{
		rdb: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DV
		}),
	}
}

func (r *RedisStorage) Insert(length int) {

	var ctx = context.Background()

	fib := make(map[int]int)
	fib[1] = 0
	fib[2] = 1

	for i := 3; i <= length; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	for i, _ := range fib {
		r.mutex.Lock()
		err := r.rdb.Set(ctx, strconv.Itoa(i), fib[i], 0).Err()
		if err != nil {
			fmt.Printf("can't set data in Redis:%s\n", err.Error())
		}
		r.mutex.Unlock()
	}

}

func (r *RedisStorage) Get(x, y int) []int {

	var ctx = context.Background()
	var fibSlice []int

	for i := 1; i <= y; i++ {

		if i >= x && i <= y {
			key := strconv.Itoa(i)
			val, err := r.rdb.Get(ctx, key).Result()

			if err != nil {
				fmt.Printf("failed to get value from Redis:%s\n", err.Error())
			}

			num, _ := strconv.Atoi(val)

			fibSlice = append(fibSlice, num)
		} else {
			continue
		}
	}

	sort.Ints(fibSlice)

	return fibSlice

}

type ErrorResponce struct {
	Message string `json:"message"`
}

type Fibonacci struct {
	len int `json:"len"`
}

type Handler struct {
	storage Storage
}

func NewHandler(storage Storage) *Handler {
	return &Handler{storage: storage}
}

func (h *Handler) CreateFibonacci(c *gin.Context) {

	l, err := strconv.Atoi(c.Param("len"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return
	}
	h.storage.Insert(l)

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "inserting done!",
	})

}

func (h *Handler) GetFibonacci(c *gin.Context) {

	str := c.Param("x,y")
	newstr := strings.Split(str, ",")

	x, err := strconv.Atoi(newstr[0])
	if err != nil {
		fmt.Printf("failed to convert x param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return
	}

	y, err := strconv.Atoi(newstr[1])
	if err != nil {
		fmt.Printf("failed to convert y param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return
	}

	fibSlice := h.storage.Get(x, y)

	c.JSON(http.StatusOK, map[string]interface{}{
		"slice": fibSlice,
	})

}
