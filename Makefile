MIGRATIONS_DIR=./app/gateway/postgres/migrations

.PHONY: migrate up down api load mocks test pprof

include .env
export $(shell sed 's/=.*//' .env)

PPROF_DIR=./tests/pprof
SERVER=http://localhost:5000/debug/pprof

migrate:
	@for f in $(MIGRATIONS_DIR)/*.sql; do \
		filename=$$(basename $$f); \
		docker-compose exec -T db psql $$DB_URL -f /migrations/$$filename; \
		echo "Migração $$filename aplicada com sucesso!"; \
	done

up:
	@docker-compose up -d --remove-orphans
	@echo "PostgreSQL iniciado!"

down:
	@docker-compose down
	@echo "PostgreSQL parado e removido!"

load:
	@k6 run --vus 100 --duration 15s tests/loadtests/load-test.js

pprof:
	mkdir -p $(PPROF_DIR)
	curl -o $(PPROF_DIR)/heap.prof $(SERVER)/heap
	curl -o $(PPROF_DIR)/goroutine.prof $(SERVER)/goroutine
	curl -o $(PPROF_DIR)/profile.prof $(SERVER)/profile
	echo "Arquivos baixados em $(PPROF_DIR)."

mocks:
	@go generate ./...

test:
	@go test ./...

coverage:
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out

api:
	@go run ./cmd/api/main.go
