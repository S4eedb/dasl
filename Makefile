
GOCMD=go
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet
BINARY_NAME=dasl
BUILD_DIR = build
BIN_DIR= $(BUILD_DIR)/bin

build: init
	go build -o $(BIN_DIR)/$(BINARY_NAME) cmd/cmd.go


init:
	mkdir -p $(BIN_DIR)

clean:
	rm -rf $(BUILD_DIR)