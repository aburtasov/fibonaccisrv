FROM golang:1.17 as builder
WORKDIR /go/src/fibonacci
COPY . .
RUN go build -o /build/bin/fibsrv ./cmd/main.go


FROM centos
COPY --from=builder /build/bin/fibsrv /build/bin/fibsrv
COPY . .
ENV FIB_HTTPADDR=":8080"
ENV FIB_DBADDR="redis:6379"

CMD [ "/build/bin/fibsrv" ]