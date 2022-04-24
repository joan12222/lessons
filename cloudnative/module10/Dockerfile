#FROM golang:1.17 AS build
#WORKDIR /demo8
#COPY . .
#ENV CGO_ENABLED=0
#ENV GO11MODULE=on
#ENV GOOS=linux go build -installsuffix cgo -o bin/httpserver main.go

FROM busybox
COPY ./demo8-server /demo8-server
EXPOSE 8080
ENTRYPOINT ["/demo8-server"]
