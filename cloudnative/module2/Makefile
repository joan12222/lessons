export tag=v1.0
BINARY_NAME=httpserver


build:
	echo "build httpserver binary"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME} .

clean:
	go clean
	rm ${BINARY_NAME}

release: build
	echo "build httpserver container"
	docker build -t lessons/httpserver:${tag} .

push: release
	echo "push lessons/httpserver"
	docker push flyinpenguin/httpserver:v1.0

