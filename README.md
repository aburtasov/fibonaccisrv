# Fibonacci service

## Build & Run(locally)

#### Prerequisites

+ go 1.17.3 or latest
+ Redis latest
+ GC compiler
+ Docker

#### Build

You should clone this repository to your PC and then build application, for example : `go build -o app main.go`
Or you can make `Docker` image using `Dockerfile` and run application in `Docker container`. See manulas about `Docker` and `Dockerfiles`.

For Redis installation you can use Docker. Just download Redis image `docker pull redis` and then enter next command:`docker run --name redisDB -p 6379:6379 -d --rm  redis`


