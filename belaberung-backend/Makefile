BINARY_NAME=bin/belaberung-server
GOFLAGS :=  -race -v -work -x -compiler gc -asmflags=$(GO_ASMFLAGS) -gcflags=$(GO_GOCFLAGS) -gccgoflags=$(GO_GCCGOFLAGS) -ldflags=$(GO_LDFLAGS)
GOFLAGS_DOCKER := -compiler gc -asmflags=$(GO_ASMFLAGS) -gcflags=$(GO_GOCFLAGS) -gccgoflags=$(GO_GCCGOFLAGS) -ldflags=$(GO_LDFLAGS)
GO_ASMFLAGS :=
GO_GOCFLAGS :=
GO_GCCGOFLAGS :=
GO_LDFLAGS :=

GO_CLEANFLAGS := -x -cache -testcache -modcache

all: build

build:
	go build $(GOFLAGS) -o bin/${BINARY_NAME} .

docker: 
	go build $(GOFLAGS_DOCKER) -o bin/${BINARY_NAME} .

run: 
	go run $(GOFLAGS) .

format: 
	go fmt delfi.dev/belaberung-v2
	go fmt delfi.dev/belaberung-v2/db
	go fmt delfi.dev/belaberung-v2/crypt
	go fmt delfi.dev/belaberung-v2/model
	go fmt delfi.dev/belaberung-v2/route
	go fmt delfi.dev/belaberung-v2/ws
	go mod tidy 

clean: 
	go clean $(GO_CLEANFLAGS)
	rm --recursive --force bin vendor