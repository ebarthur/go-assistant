# Variables
BINARY_NAME=groq-api

# Build the project
build:
	@echo "Building the project..."
	go build -o $(BINARY_NAME)

# Run the project using compiledaemon
run:
	@echo "Running the project with compiledaemon..."
	compiledaemon --command="./$(BINARY_NAME)"

# Test the project
test:
	@echo "Running tests..."
	go test ./...

# Clean the build
clean:
	@echo "Cleaning up..."
	go clean
	rm -f $(BINARY_NAME)

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod tidy

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Lint code
lint:
	@echo "Linting code..."
	golangci-lint run

.PHONY: build run test clean deps fmt lint install-lint
