package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/aburtasov/fibonaccisrv/pkg/storage"
	"github.com/gin-gonic/gin"
)

type ErrorResponce struct {
	Message string `json:"message"`
}

type Fibonacci struct {
	len int `json:"len"`
}

type Handler struct {
	storage storage.Storage
}

func NewHandler(storage storage.Storage) *Handler {
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
