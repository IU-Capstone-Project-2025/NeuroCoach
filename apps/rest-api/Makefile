# Makefile for Fitness AI Project (Docker DB without Compose)

# Configuration
PROJECT_NAME := fitness-ai
DB_NAME := fitness_ai
DB_USER := postgres
DB_PASSWORD := postgres
DB_HOST := localhost
DB_PORT := 5432
DB_CONTAINER_NAME := fitness-ai-db
DATABASE_URL := postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable
MIGRATE_PATH := migrations

.PHONY: help
help: ## Show this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: db-up
db-up: ## Start PostgreSQL in Docker
	@if [ -z "$$(docker ps -q -f name=$(DB_CONTAINER_NAME))" ]; then \
		echo "Starting PostgreSQL container..."; \
		docker run --name $(DB_CONTAINER_NAME) -d -p $(DB_PORT):5432 \
			-e POSTGRES_DB=$(DB_NAME) \
			-e POSTGRES_USER=$(DB_USER) \
			-e POSTGRES_PASSWORD=$(DB_PASSWORD) \
			postgres:15-alpine; \
		echo "Waiting for database to be ready..."; \
		until docker exec $(DB_CONTAINER_NAME) pg_isready -U $(DB_USER); do sleep 2; done; \
		echo "Database is ready!"; \
	else \
		echo "Database container already running"; \
	fi

.PHONY: db-down
db-down: ## Stop and remove PostgreSQL container
	@echo "Stopping database container..."
	@-docker stop $(DB_CONTAINER_NAME) >/dev/null 2>&1
	@-docker rm $(DB_CONTAINER_NAME) >/dev/null 2>&1
	@echo "Database container removed"

.PHONY: migrate-up
migrate-up: ## Run database migrations using Docker
	@echo "Applying database migrations using Docker..."
	@docker run --rm \
		-v $(PWD)/$(MIGRATE_PATH):/migrations \
		--network host \
		migrate/migrate \
		-path=/migrations -database "$(DATABASE_URL)" up

.PHONY: migrate-down
migrate-down: ## Rollback database migrations using Docker
	@echo "Rolling back database migrations using Docker..."
	@docker run --rm \
		-v $(PWD)/$(MIGRATE_PATH):/migrations \
		--network host \
		migrate/migrate \
		-path=/migrations -database "$(DATABASE_URL)" down 1

.PHONY: db-shell
db-shell: ## Access database shell
	@PGPASSWORD=$(DB_PASSWORD) psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -d $(DB_NAME)

.PHONY: run-api
run-api: ## Run the Go API server
	@echo "Starting API server..."
	@go run cmd/server/main.go

.PHONY: test
test: ## Run tests
	@echo "Running tests..."
	@go test -v ./...

.PHONY: build
build: ## Build the application binary
	@echo "Building application..."
	@go build -o bin/fitness-ai cmd/server/main.go

.PHONY: clean
clean: db-down ## Clean up resources
	@echo "Cleaning up..."
	@rm -rf bin/

.PHONY: full-start
full-start: db-up migrate-up run-api ## Full start: DB, migrations, and API
	@echo "System started successfully!"

.PHONY: restart-db
restart-db: db-down db-up ## Restart database

.PHONY: mongo-up
mongo-up: ## Start MongoDB using Docker
	docker run -d --name neurocoach-mongo -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=user -e MONGO_INITDB_ROOT_PASSWORD=password -e MONGO_INITDB_DATABASE=fitness_ai mongo:6.0

.PHONY: mongo-down
mongo-down: ## Stop and remove MongoDB container
	docker stop neurocoach-mongo || true
	docker rm neurocoach-mongo || true

.PHONY: mongo-shell
mongo-shell: ## Open a shell to the running MongoDB container
	docker exec -it neurocoach-mongo mongosh -u user -p password --authenticationDatabase admin fitness_ai