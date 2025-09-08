# ${{values.app_name}} - Makefile
# Based on lab-go-cli patterns

# Build variables
BINARY_NAME=${{values.app_name}}
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
DATE ?= $(shell date +%FT%T%z)
BUILT_BY ?= $(shell git config user.email 2>/dev/null || echo "unknown")

# Go variables
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=gofmt

# Build flags
LDFLAGS=-ldflags "-X 'github.com/Portfolio-jaime/${{values.app_name}}/cmd.version=$(VERSION)' \
                  -X 'github.com/Portfolio-jaime/${{values.app_name}}/cmd.commit=$(COMMIT)' \
                  -X 'github.com/Portfolio-jaime/${{values.app_name}}/cmd.date=$(DATE)' \
                  -X 'github.com/Portfolio-jaime/${{values.app_name}}/cmd.builtBy=$(BUILT_BY)'"

# Directories
BUILD_DIR=build
DOCS_DIR=docs

.PHONY: all build clean test coverage lint fmt vet mod-tidy mod-download dev install uninstall docs help

# Default target
all: clean fmt vet test build

# Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) ./main.go

# Build for multiple platforms
build-all: build-linux build-darwin build-windows

build-linux:
	@echo "Building for Linux..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 ./main.go

build-darwin:
	@echo "Building for macOS..."
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 ./main.go

build-windows:
	@echo "Building for Windows..."
	@mkdir -p $(BUILD_DIR)
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe ./main.go

# Clean build artifacts
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	@rm -rf $(BUILD_DIR)

# Run tests
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

# Run tests with coverage
coverage:
	@echo "Running tests with coverage..."
	$(GOTEST) -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Lint code
lint:
	@echo "Running golint..."
	@command -v golint >/dev/null 2>&1 || { echo "golint not installed. Installing..."; $(GOGET) golang.org/x/lint/golint; }
	golint ./...

# Format code
fmt:
	@echo "Formatting code..."
	$(GOFMT) -s -w .

# Vet code
vet:
	@echo "Vetting code..."
	$(GOCMD) vet ./...

# Tidy go modules
mod-tidy:
	@echo "Tidying go modules..."
	$(GOMOD) tidy

# Download go modules
mod-download:
	@echo "Downloading go modules..."
	$(GOMOD) download

# Development mode (build and run)
dev: build
	@echo "Running in development mode..."
	./$(BUILD_DIR)/$(BINARY_NAME)

# Install binary to GOPATH/bin
install:
	@echo "Installing $(BINARY_NAME)..."
	$(GOBUILD) $(LDFLAGS) -o $(GOPATH)/bin/$(BINARY_NAME) ./main.go

# Uninstall binary from GOPATH/bin
uninstall:
	@echo "Uninstalling $(BINARY_NAME)..."
	@rm -f $(GOPATH)/bin/$(BINARY_NAME)

# Generate documentation
docs:
	@echo "Generating documentation..."
	@mkdir -p $(DOCS_DIR)
	./$(BUILD_DIR)/$(BINARY_NAME) --help > $(DOCS_DIR)/help.txt
	@echo "Documentation generated in $(DOCS_DIR)/"

# Display help
help:
	@echo "Available targets:"
	@echo "  all          - Clean, format, vet, test, and build"
	@echo "  build        - Build the binary"
	@echo "  build-all    - Build for all platforms"
	@echo "  clean        - Clean build artifacts"
	@echo "  test         - Run tests"
	@echo "  coverage     - Run tests with coverage report"
	@echo "  lint         - Lint the code"
	@echo "  fmt          - Format the code"
	@echo "  vet          - Vet the code"
	@echo "  mod-tidy     - Tidy go modules"
	@echo "  mod-download - Download go modules"
	@echo "  dev          - Build and run in development mode"
	@echo "  install      - Install binary to GOPATH/bin"
	@echo "  uninstall    - Remove binary from GOPATH/bin"
	@echo "  docs         - Generate documentation"
	@echo "  help         - Show this help message"