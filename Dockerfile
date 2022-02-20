FROM golang:1.17

WORKDIR /go/src/fibonacci

COPY . .

RUN go build -o /build/bin/fibsrv ./cmd/main.go
RUN EXPORT FIB_HTTPADDR=":8080"
RUN EXPORT FIB_DBADDR=":6539"

CMD [ "/build/bin/fibsrv" ]