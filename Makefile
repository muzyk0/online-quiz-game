# Makefile for Online Quiz Game
SHELL := /bin/bash

.PHONY: help
help: ## Show this help message
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# ── Build ──────────────────────────────────────────────────────────────────────

.PHONY: build
build: ## Build the server binary → bin/server
	go build -o bin/server ./cmd/server/

.PHONY: run
run: ## Run the server (loads .env if present)
	go run ./cmd/server/

.PHONY: clean
clean: ## Remove build artifacts
	rm -rf bin/

# ── Test & Quality ─────────────────────────────────────────────────────────────

.PHONY: test
test: ## Run all tests
	go test ./...

.PHONY: test-v
test-v: ## Run all tests with verbose output
	go test -v ./...

.PHONY: test-race
test-race: ## Run all tests with race detector
	go test -race ./...

.PHONY: test-cover
test-cover: ## Run tests with coverage report
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

.PHONY: test-e2e-simple
test-e2e-simple: ## Run e2e tests (requires DATABASE_URL)
	go test -tags e2e -v -timeout 120s ./test/e2e/

.PHONY: test-e2e
test-e2e: ## Run e2e tests with race detection, coverage and summary (requires DATABASE_URL)
	@go test -tags e2e -race -count=1 -v -timeout 120s \
		-coverprofile=e2e_coverage.out \
		-coverpkg=./internal/... \
		./test/e2e/ 2>&1 | tee /tmp/.e2e_out.txt; \
	STATUS=$${PIPESTATUS[0]}; \
	PASS=$$(awk '/--- PASS/{c++} END{print c+0}' /tmp/.e2e_out.txt); \
	FAIL=$$(awk '/--- FAIL/{c++} END{print c+0}' /tmp/.e2e_out.txt); \
	RACE=$$(awk '/DATA RACE/{c++} END{print c+0}' /tmp/.e2e_out.txt); \
	echo ""; \
	printf '━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n'; \
	printf '  Total %-3d   Pass %-3d   Fail %-3d' $$((PASS+FAIL)) $$PASS $$FAIL; \
	[ "$$RACE" -gt "0" ] && printf '   Race conditions: %d' $$RACE; printf '\n'; \
	[ -f e2e_coverage.out ] && \
		go tool cover -func=e2e_coverage.out | tail -1 | awk '{printf "  Coverage: %s\n", $$NF}'; \
	printf '━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n'; \
	rm -f /tmp/.e2e_out.txt; \
	exit $$STATUS

.PHONY: lint
lint: ## Run golangci-lint
	golangci-lint run

.PHONY: fmt
fmt: ## Format code with go fmt
	go fmt ./...

.PHONY: tidy
tidy: ## Tidy go modules
	go mod tidy

.PHONY: generate
generate: ## Run go generate (regenerate mocks)
	go generate ./...

# ── Migrations ─────────────────────────────────────────────────────────────────

.PHONY: migrate-up
migrate-up: ## Apply all pending migrations
	go run ./cmd/migrate/ -action up

.PHONY: migrate-down
migrate-down: ## Rollback all migrations
	go run ./cmd/migrate/ -action down

.PHONY: migrate-version
migrate-version: ## Show current migration version
	go run ./cmd/migrate/ -action version

.PHONY: migrate-create
migrate-create: ## Create new migration: make migrate-create name=add_users
	@test -n "$(name)" || (echo "Usage: make migrate-create name=<migration_name>"; exit 1)
	migrate create -ext sql -dir internal/app/database/migrations $(name)

# ── Tunnel ─────────────────────────────────────────────────────────────────────

.PHONY: tunnel
tunnel: ## Expose local server via nport (HTTPS tunnel → localhost:8080)
	npx nport 8080 -s quiz-game-9art

# ── Docker ─────────────────────────────────────────────────────────────────────

COMPOSE_FILE := deployments/docker-compose.yml
DOCKER_COMPOSE := $(shell docker compose version >/dev/null 2>&1 && echo "docker compose" || echo "docker-compose")
DC := $(DOCKER_COMPOSE) -f $(COMPOSE_FILE)

.PHONY: docker-build
docker-build: ## Build Docker image
	docker build -f build/package/Dockerfile -t quiz-server .

.PHONY: docker-up
docker-up: ## Start all services (postgres + server) in background
	$(DC) up -d

.PHONY: docker-up-db
docker-up-db: ## Start only postgres in background
	$(DC) up -d postgres

.PHONY: docker-down
docker-down: ## Stop and remove containers
	$(DC) down

.PHONY: docker-down-v
docker-down-v: ## Stop containers and remove volumes (wipes DB data)
	$(DC) down -v

.PHONY: docker-logs
docker-logs: ## Follow logs from all containers
	$(DC) logs -f

.PHONY: docker-ps
docker-ps: ## Show running containers status
	$(DC) ps

.PHONY: docker-migrate
docker-migrate: ## Run DB migrations inside the running server container
	$(DC) exec server ./server -migrate up

# ── Swagger ────────────────────────────────────────────────────────────────────

.PHONY: swag
swag: ## Generate Swagger docs from annotations
	@which swag >/dev/null 2>&1 || go install github.com/swaggo/swag/cmd/swag@latest
	swag init -g cmd/server/main.go -o internal/app/swagger/docs --parseDependency --parseInternal
