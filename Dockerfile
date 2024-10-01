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

COPY --from=builder /build/bin/fibsrv /usr/local/bin/fibsrv

ENTRYPOINT ["/usr/local/bin/fibsrv"]
