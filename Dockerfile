FROM golang:1.17 as builder
WORKDIR /go/src/fibonacci
COPY . .
RUN go build -o /build/bin/fibsrv ./cmd/main.go


FROM ubuntu
COPY --from=builder /build/bin/fibsrv /build/bin/fibsrv

ENV FIB_HTTPADDR=":8080"
ENV FIB_DBADDR=":6539"

CMD [ "/build/bin/fibsrv" ]