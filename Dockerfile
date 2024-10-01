FROM golang:1.21 AS builder

WORKDIR /go/src/fibonacci

# Копируем только go.mod и go.sum для загрузки зависимостей
COPY go.mod go.sum ./

RUN go mod download

# Копируем остальные файлы
COPY . .

RUN go build -o /build/bin/fibsrv ./cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

ENV FIB_HTTPADDR=":8080"

ENV FIB_DBADDR="redis:6379"

COPY --from=builder /build/bin/fibsrv /build/bin/fibsrv

CMD ["/build/bin/fibsrv"]
