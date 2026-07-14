.PHONY: \
	help \
	setup \
	dev \
	build \
	test \
	test-verbose \
	fmt \
	migration-create \
	migration-up \
	migration-down \
	migration-down-all \
	migration-version

-include .env
export

MIGRATIONS_PATH := ./cmd/migrations

MIGRATE := docker run --rm \
	--network host \
	-v "$(CURDIR)/cmd/migrations:/migrations" \
	migrate/migrate

help:
	@echo "Available commands:"
	@echo ""
	@echo "  make setup"
	@echo "      Apply all pending database migrations."
	@echo ""
	@echo "  make dev"
	@echo "      Start the API locally."
	@echo ""
	@echo "  make build"
	@echo "      Build the API binary."
	@echo ""
	@echo "  make test"
	@echo "      Run all unit tests."
	@echo ""
	@echo "  make test-verbose"
	@echo "      Run all unit tests with verbose output."
	@echo ""
	@echo "  make fmt"
	@echo "      Format all Go files."
	@echo ""
	@echo "  make migration-create NAME=create_users_table"
	@echo "      Create a new sequential migration using the local migrate CLI."
	@echo ""
	@echo "  make migration-up"
	@echo "      Apply all pending migrations."
	@echo ""
	@echo "  make migration-down"
	@echo "      Roll back the latest migration."
	@echo ""
	@echo "  make migration-down-all"
	@echo "      Roll back all migrations."
	@echo ""
	@echo "  make migration-version"
	@echo "      Show the current migration version."

setup: migration-up
	@echo "Development environment is ready."

dev:
	go run ./cmd/api

build:
	@mkdir -p ./bin
	go build -o ./bin/api ./cmd/api

test:
	go test ./...

test-verbose:
	go test -v ./...

fmt:
	go fmt ./...

migration-create:
	@if [ -z "$(NAME)" ]; then \
		echo "NAME is required."; \
		echo "Example: make migration-create NAME=add_created_at_to_users"; \
		exit 1; \
	fi
	migrate create \
		-ext sql \
		-dir "$(MIGRATIONS_PATH)" \
		-seq \
		"$(NAME)"

migration-up: check-database-url
	$(MIGRATE) \
		-path=/migrations \
		-database "$(DATABASE_URL)" \
		up

migration-down: check-database-url
	$(MIGRATE) \
		-path=/migrations \
		-database "$(DATABASE_URL)" \
		down 1

migration-down-all: check-database-url
	$(MIGRATE) \
		-path=/migrations \
		-database "$(DATABASE_URL)" \
		down -all

migration-version: check-database-url
	$(MIGRATE) \
		-path=/migrations \
		-database "$(DATABASE_URL)" \
		version

.PHONY: check-database-url

check-database-url:
	@if [ -z "$(DATABASE_URL)" ]; then \
		echo "DATABASE_URL is not defined."; \
		echo "Add it to the .env file or export it in your terminal."; \
		exit 1; \
	fi
