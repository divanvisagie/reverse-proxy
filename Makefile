# Makefile for a Go project

# Project details
BINARY_NAME=reverse-proxy
BUILD_DIR=bin

# Go related variables
GO=go
GOTEST=$(GO) test ./...
GOBUILD=$(GO) build -o $(BUILD_DIR)/$(BINARY_NAME)
GOCLEAN=$(GO) clean
GOINSTALL=$(GO) install
GOMODTIDY=$(GO) mod tidy

# Build the project
build:
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD)

# Run tests
test:
	$(GOTEST)

# Clean the build directory
clean:
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

# Tidy up the dependencies
tidy:
	$(GOMODTIDY)

# Install the binary
install:
	$(GOINSTALL)

# Run the application (useful for local development)
run: build
	$(BUILD_DIR)/$(BINARY_NAME)

.PHONY: build test clean tidy install run
