BINARY_NAME=bin/belaberung-server

all: build

build:
	go build -o ${BINARY_NAME} .

run: 
	go run .

format: 
	go fmt delfi.dev/belaberung-v2
	go fmt delfi.dev/belaberung-v2/db
	go fmt delfi.dev/belaberung-v2/crypt
	go fmt delfi.dev/belaberung-v2/model
	go fmt delfi.dev/belaberung-v2/route
	go mod tidy 

clean: 
	rm -rf bin vendor