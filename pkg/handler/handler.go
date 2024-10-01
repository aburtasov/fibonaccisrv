package handler

import (
	"net/http"
	"strconv"

	"github.com/aburtasov/fibonaccisrv/pkg/storage"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// ErrorResponse представляет структуру ответа с ошибкой
type ErrorResponse struct {
	Message string `json:"message"`
}

// SuccessResponse представляет структуру успешного ответа
type SuccessResponse struct {
	Message string `json:"message"`
}

// FibonacciResponse представляет структуру ответа с срезом чисел Фибоначчи
type FibonacciResponse struct {
	Slice []string `json:"slice"`
}

// Handler обрабатывает HTTP-запросы
type Handler struct {
	storage storage.Storage
	logger  *logrus.Logger
}

// NewHandler создаёт новый экземпляр Handler
func NewHandler(storage storage.Storage, logger *logrus.Logger) *Handler {
	return &Handler{
		storage: storage,
		logger:  logger,
	}
}

// CreateFibonacci обрабатывает запрос на создание последовательности Фибоначчи
func (h *Handler) CreateFibonacci(c *gin.Context) {
	lenParam := c.Param("len")

	// Преобразование параметра "len" в целое число
	length, err := strconv.Atoi(lenParam)
	if err != nil {
		h.logger.Errorf("Invalid 'len' parameter: %v", err)
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: "Параметр 'len' должен быть целым числом",
		})
		return
	}

	// Проверка, что длина последовательности положительна
	if length < 1 {
		h.logger.Warnf("Invalid 'len' parameter: %d (must be >= 1)", length)
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: "Параметр 'len' должен быть положительным числом",
		})
		return
	}

	// Вставка последовательности Фибоначчи в хранилище
	if err := h.storage.Insert(length); err != nil {
		h.logger.Errorf("Failed to insert Fibonacci sequence: %v", err)
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: "Не удалось вставить данные",
		})
		return
	}

	h.logger.Infof("Successfully inserted Fibonacci sequence of length %d", length)
	c.JSON(http.StatusOK, SuccessResponse{
		Message: "Вставка завершена успешно",
	})
}

// GetFibonacci обрабатывает запрос на получение чисел Фибоначчи в диапазоне [x, y]
func (h *Handler) GetFibonacci(c *gin.Context) {
	// Извлечение параметров "x" и "y" из строки запроса
	xParam := c.Query("x")
	yParam := c.Query("y")

	// Проверка наличия параметров
	if xParam == "" || yParam == "" {
		h.logger.Warn("Missing query parameters 'x' and/or 'y'")
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: "Необходимо указать параметры 'x' и 'y'",
		})
		return
	}

	// Преобразование параметров "x" и "y" в целые числа
	x, err := strconv.Atoi(xParam)
	if err != nil {
		h.logger.Errorf("Invalid 'x' parameter: %v", err)
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: "Параметр 'x' должен быть целым числом",
		})
		return
	}

	y, err := strconv.Atoi(yParam)
	if err != nil {
		h.logger.Errorf("Invalid 'y' parameter: %v", err)
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: "Параметр 'y' должен быть целым числом",
		})
		return
	}

	// Проверка, что x <= y
	if x > y {
		h.logger.Warnf("Invalid range: x (%d) > y (%d)", x, y)
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: "Параметр 'x' должен быть меньше или равен 'y'",
		})
		return
	}

	// Вызов метода Get из хранилища для получения чисел Фибоначчи
	fibSlice, err := h.storage.Get(x, y)
	if err != nil {
		h.logger.Errorf("Failed to get Fibonacci sequence: %v", err)
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: "Не удалось получить данные",
		})
		return
	}

	h.logger.Infof("Successfully retrieved Fibonacci sequence from %d to %d", x, y)
	c.JSON(http.StatusOK, FibonacciResponse{
		Slice: fibSlice,
	})
}
