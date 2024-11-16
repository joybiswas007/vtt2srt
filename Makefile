# Simple Makefile for a Go project

# Build the application

build:
	@echo "Building..."
	@go build -o main main.go

# Run the application
run:
	@go run main.go

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

.PHONY: build run clean
