# Define the output binary name
BINARY_NAME := rss_aggregator
MIGRATION_PATH := ./db/migration

# Load environment variables from .env file
.PHONY: load-env
load-env:
	@export $$(grep -v '^#' .env | xargs) && \
	 echo "Loaded environment variables" && \
	 echo "DATABASE_URL: $$DATABASE_URL"


# Use environment variables
.PHONY: print-env
print-env: load-env
	@echo "DATABASE_URL: $$DATABASE_URL"
	@echo "GO_ENV: $${GO_ENV}"


# Default target: Build the project
.PHONY: all
all: build

# Build the Go project
.PHONY: build
build:
	go build -o $(BINARY_NAME) main.go

# Run the Go project
.PHONY: run
run: build
	./$(BINARY_NAME)

# Test the Go project
.PHONY: test
test:
	go test ./...

# Clean the build files
.PHONY: clean
clean:
	rm -f $(BINARY_NAME)

# Update dependencies
.PHONY: deps
deps:
	go mod tidy

# Format the code
.PHONY: fmt
fmt:
	go fmt ./...

# Generate SQL code with sqlc
.PHONY: sqlc
sqlc:
	sqlc generate

# Define migrate-create target with name argument
.PHONY: m-create
m-create:
	@export $$(grep -v '^#' .env | xargs) && goose -v -dir ${MIGRATION_PATH} create "$(n)" sql
 

# Run Goose migrations up
.PHONY: m-up
m-up:
	@export $$(grep -v '^#' .env | xargs) && \
	 echo "DATABASE_URL: $$DATABASE_URL_SSL" && \
	 goose -v -dir ${MIGRATION_PATH} postgres $$DATABASE_URL_SSL up 

# Run Goose migrations down
.PHONY: m-down
m-down:
	@export $$(grep -v '^#' .env | xargs) && \
	 echo "DATABASE_URL: $$DATABASE_URL_SSL" && \
	 goose -v -dir ${MIGRATION_PATH} postgres $$DATABASE_URL_SSL down 