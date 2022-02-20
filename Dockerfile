FROM golang:1.17 as builder
WORKDIR /go/src/fibonacci
COPY . .
RUN go build -o /build/bin/fibsrv ./cmd/main.go


FROM alpine
COPY --from=builder /build/bin/fibsrv /build/bin/fibsrv
RUN EXPORT FIB_HTTPADDR=":8080"
RUN EXPORT FIB_DBADDR=":6539"

CMD [ "/build/bin/fibsrv" ]