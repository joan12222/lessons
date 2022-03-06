export tag=v1.0
root:
	export ROOT=github.com/lessons

build:
	echo "build httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

release: build
	echo "build httpserver container"
	docker build -t lessons/httpserver:${tag} .

push: release
	echo "push lessons/httpserver"
	docker push lessons/httpserver:v1.0

