BINARY_NAME := cmdgenie
BIN_DIR := bin

.PHONY: all
all: build

.PHONY: build
build: build-linux-amd64 build-linux-arm64 build-darwin-amd64 build-darwin-arm64 build-windows-amd64

.PHONY: build-linux-amd64
build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/$(BINARY_NAME)-linux-amd64 .

.PHONY: build-linux-arm64
build-linux-arm64:
	GOOS=linux GOARCH=arm64 go build -o $(BIN_DIR)/$(BINARY_NAME)-linux-arm64 .

.PHONY: build-darwin-amd64
build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -o $(BIN_DIR)/$(BINARY_NAME)-darwin-amd64 .

.PHONY: build-darwin-arm64
build-darwin-arm64:
	GOOS=darwin GOARCH=arm64 go build -o $(BIN_DIR)/$(BINARY_NAME)-darwin-arm64 .

.PHONY: build-windows-amd64
build-windows-amd64:
	GOOS=windows GOARCH=amd64 go build -o $(BIN_DIR)/$(BINARY_NAME)-windows-amd64.exe .

.PHONY: clean
clean:
	rm -f $(BIN_DIR)/$(BINARY_NAME)-*

.PHONY: lint
lint:
	go fmt ./...
	golangci-lint run

.PHONY: test
test:
	go test -v ./...

.PHONY: deps
deps:
	go mod tidy
