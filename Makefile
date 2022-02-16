.PHONY:build
build:
	go build -o fibsrv cmd/main.go

.DEFAULT_GOAL := build