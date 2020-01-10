# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=version_check
BINARY_UNIX=$(BINARY_NAME)_unix

build:
        $(GOBUILD) -o $(BINARY_NAME) -v
clean:
        $(GOCLEAN)
        rm -f $(BINARY_NAME)
        rm -f $(BINARY_UNIX)
run:
        $(GOBUILD) -o $(BINARY_NAME) -v ./...
        ./$(BINARY_NAME)
deps:
        $(GOGET) github.com/mikemintang/go-curl
        $(GOGET) "github.com/hashicorp/consul/api"


# Cross compilation
build-linux:
        CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
docker-build:
        docker run --rm -it -v "$(GOPATH)":/go -w /go/src golang:1.12.13 go build -o "$(BINARY_UNIX)" -v