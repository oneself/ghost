BINARY_NAME=ghost
OUT="out"
BINARY_PATH=${OUT}/${BINARY_NAME}
COVERAGE=${OUT}/coverage.out

build:
	go mod tidy
	mkdir -p ${OUT}
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_PATH} main.go
#GOARCH=amd64 GOOS=darwin go build -o ${BINARY_PATH}-darwin main.go
#GOARCH=amd64 GOOS=window go build -o ${BINARY_PATH}-windows main.go

run:
	go run . --dir ~/dropbox/Ghost --count 10

build_run: build run

test:
	go test -v ./...

cover:
	go test -coverprofile=${COVERAGE}.out ./...

show_cover: coverage.out
	go tool cover -html=${COVERAGE}

init:
	go mod init ${BINARY_NAME}

clean:
	go clean
	rm -rf ${OUT}
