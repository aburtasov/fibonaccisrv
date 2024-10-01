FROM golang:1.21 AS builder

WORKDIR /go/src/fibonacci

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Сборка статически связанного бинарника
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /build/bin/fibsrv ./cmd/main.go

FROM ubuntu:22.04

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=builder /build/bin/fibsrv /usr/local/bin/fibsrv

ENTRYPOINT ["/usr/local/bin/fibsrv"]
