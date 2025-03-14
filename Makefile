# Load environment variables from .env file
include .env
export

MIGRATE=migrate
DATABASE_URL=postgres://$$DB_USER:$$DB_PASSWORD@$$DB_HOST:$$DB_PORT/$$DB_NAME?sslmode=disable

# Nama direktori tempat file migrasi disimpan
MIGRATIONS_DIR=migrations

.PHONY : format clean install build create-migrations migrate rollback create-seeder run-seed create-mock create-mock-all

run:
	go run ./cmd/main.go

format:
	gofmt -s -w .

clean:
	go mod tidy

run-this:
	echo "hello"

install:
	go mod download
	
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags musl -o main ./cmd/webapi/main.go

start:
	./main

test:
	go test ./internal/domain/usecases/... -coverprofile=coverages/coverage.out
	go tool cover -html=coverages/coverage.out -o coverages/coverage.html

server-test:
	go test ./internal/domain/usecases/... -coverprofile=cover.out

create-migrations:
ifndef name
	$(error "name is undefined. Usage: make migrate name=YOUR_MIGRATION_NAME")
endif
	goose -dir=$(MIGRATIONS_DIR) create $(name) sql

migrate:
	goose -dir=$(MIGRATIONS_DIR) postgres $(DATABASE_URL) up

rollback:
	goose -dir=$(MIGRATIONS_DIR) postgres $(DATABASE_URL) down

create-mock-repo:
ifndef name
	$(error "name is undefined. Usage: make create-mock-repo name=YOUR_MOCK_NAME")
endif
	mockery --name=$(name) --dir=internal/domain/repositories --output=internal/domain/repositories/mocks

create-mock-repo-all:
	mockery --all --dir=internal/domain/repositories --output=internal/domain/repositories/mocks

create-mock-usecase:
ifndef name
	$(error "name is undefined. Usage: make create-mock-usecase name=YOUR_MOCK_NAME")
endif
	mockery --name=$(name) --dir=internal/domain/usecases --output=internal/domain/usecases/mocks

create-mock-helper:
ifndef name
	$(error "name is undefined. Usage: make create-mock-helper name=YOUR_MOCK_NAME")
endif
	mockery --name=$(name) --dir=internal/domain/helpers --output=internal/domain/helpers/mocks


