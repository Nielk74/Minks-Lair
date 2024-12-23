
# Variables
APP_NAME := docker-gs-ping
DOCKER_IMAGE := go-app:test

# Default target
.PHONY: all
all: test

# Run tests
.PHONY: test
test:
	go test ./... -v -coverprofile=coverage.out

# Build application
.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -o $(APP_NAME)

# Run application locally
.PHONY: run
run: build
	./$(APP_NAME)

# Build Docker image
.PHONY: docker-build
docker-build:
	docker build -t $(DOCKER_IMAGE) .

# Run tests inside Docker
.PHONY: docker-test
docker-test: docker-build
	docker run --rm $(DOCKER_IMAGE) go test ./... -v -coverprofile=coverage.out

# Clean up build artifacts
.PHONY: clean
clean:
	rm -f $(APP_NAME) coverage.out

create-migration:
	echo "Creating migration file"
	migrate create -ext sql -dir migrations -seq $(name) -database postgres://localhost:5432/$(db)
	echo "Migration file created"
