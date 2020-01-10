# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOGET=$(GOCMD) get

BINARY_NAME=version_check

build:
        $(GOBUILD) -o $(BINARY_NAME) -v
run:
        chmod 777 $(BINARY_NAME)
        nohup ./$(BINARY_NAME) &
deps:
        $(GOGET) github.com/mikemintang/go-curl
        $(GOGET) github.com/hashicorp/consul/api


# Cross compilation
build-linux:
        CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v