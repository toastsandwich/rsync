GO = go
BINARY_NAME = rsync
BUILD_DIR = build
PREFIX = /usr/local

.PHONY: build clean install test lint fmt help

# Default target
all: build

# Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	$(GO) build -o $(BUILD_DIR)/$(BINARY_NAME) .

# Install the binary
install: build
	@echo "Installing $(BINARY_NAME) to $(PREFIX)/bin..."
	sudo cp $(BUILD_DIR)/$(BINARY_NAME) $(PREFIX)/bin/

# Clean build artifacts
clean:
	@echo "Cleaning up..."
	rm -f $(BUILD_DIR)/$(BINARY_NAME)

# Run tests
test:
	@echo "Running tests..."
	$(GO) test -v ./...

# Run linter
lint:
	@echo "Running linter..."
	golangci-lint run

# Format code
fmt:
	@echo "Formatting code..."
	$(GO) fmt ./...

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	$(GO) mod download
	$(GO) mod tidy

# Show help
help:
	@echo "Available targets:"
	@echo "  build    - Build the binary"
	@echo "  install  - Install the binary to $(PREFIX)/bin"
	@echo "  clean    - Remove build artifacts"
	@echo "  test     - Run tests"
	@echo "  lint     - Run linter"
	@echo "  fmt      - Format code"
	@echo "  deps     - Download and tidy dependencies"
	@echo "  help     - Show this help message"
