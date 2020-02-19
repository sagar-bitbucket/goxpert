GOCMD=go
GOBUILD=$(GOCMD) build


BINARY_NAME=executor-service
SERVICE_NAME=executor
BUILD_PATH=./bin/$(BINARY_NAME)

build:
	$(GOBUILD) -o bin/$(BINARY_NAME) cmd/main.go

run:
	$(BUILD_PATH)