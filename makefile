# Load .env file if it exists
ifneq (,$(wildcard .env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

APP_NAME := trello-services
BINARY=bin/$(APP_NAME)
MAIN=./cmd/server/main.go
DOCKER_IMAGE=$(APP_NAME):latest

GO=go
GOTEST=$(GO) test
GOFLAGS=

.PHONY: all dev build run prod clean docker-build docker-run

dev:
	@if ! command -v air >/dev/null 2>&1; then \
		echo "Air not found. Install it with: go install github.com/air-verse/air@latest"; \
		exit 1; \
	fi
	@echo "Running $(APP_NAME) in development mode..."
	air

# Build binary for production
build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p bin
	$(GO) build $(GOFLAGS) -o $(BINARY) $(MAIN)
	@echo "Binary built at $(BINARY)"

# Run the already built binary
run:
	@echo "Running $(APP_NAME) binary..."
	./$(BINARY)

# Build & run (production shortcut)
prod: build run

# Build Docker image
docker-build:
	@echo "Building Docker image $(DOCKER_IMAGE)..."
	docker build -t $(DOCKER_IMAGE) .

# Run Docker container (pakai env var dari .env)
docker-run:
	@if [ ! -f .env ]; then \
		echo ".env file not found. Please create one before running Docker."; \
		exit 1; \
	fi
	@echo "Running Docker container $(DOCKER_IMAGE) on port $(SERVER_PORT)..."
	docker run --rm -p $(SERVER_PORT):$(SERVER_PORT) --env-file .env $(DOCKER_IMAGE)

# Clean up binaries and temporary files
clean:
	@echo "Cleaning up..."
	rm -rf bin tmp
