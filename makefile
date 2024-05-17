.PHONY: default run run-with-docs build test docs clean

# Variables
APP_NAME=gopportunities

# Tasks
default: run-with-docs

# Run the application
run:
	@echo "Running the application..."
	@go run main.go

# Generate docs and run the application
run-with-docs:
	@echo "Generating documentation with swag..."
	@swag init
	@echo "Running the application with generated documentation..."
	@go run main.go

# Build the application
build:
	@echo "Building the application..."
	@go build -o $(APP_NAME) main.go
	@echo "Build completed: $(APP_NAME)"

# Run tests
test:
	@echo "Running tests..."
	@go test ./ ...
	@echo "Tests completed."

# Generate documentation
docs:
	@echo "Generating documentation..."
	@swag init
	@echo "Documentation generated."

# Clean build and docs
clean:
	@echo "Cleaning up..."
	@rm -f $(APP_NAME)
	@rm -rf ./docs
	@echo "Cleanup completed."
