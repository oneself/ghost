BINARY_NAME=ghost

build:
	go mod tidy
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME} main.go
#GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin main.go
#GOARCH=amd64 GOOS=window go build -o ${BINARY_NAME}-windows main.go

run:
	go run . --dir .. --count 10

build_run: build run

test:
	go test

init:
	go mod init ${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-windows
