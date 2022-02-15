.PHONY: build
 build:
       go build -o fibsrv ./cmd

.DEFAULT_GOAL := build