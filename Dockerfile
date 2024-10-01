# Стадия сборки
FROM golang:1.21 AS builder

WORKDIR /go/src/fibonacci

# Копируем файлы проекта
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Сборка приложения
RUN go build -o /build/bin/fibsrv ./cmd/main.go

# Финальная стадия
FROM centos:latest

# Установка необходимых пакетов
RUN apk --no-cache add ca-certificates

# Копируем исполняемый файл из стадии сборки
COPY --from=builder /build/bin/fibsrv /usr/local/bin/fibsrv

# Определяем точку входа
ENTRYPOINT ["/usr/local/bin/fibsrv"]
