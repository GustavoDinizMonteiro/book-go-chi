MIGRATIONS_DIR=./app/gateway/postgres/migrations

.PHONY: migrate up down api

include .env
export $(shell sed 's/=.*//' .env)

migrate:
	@for f in $(MIGRATIONS_DIR)/*.sql; do \
		filename=$$(basename $$f); \
		docker-compose exec -T db psql $$DB_URL -f /migrations/$$filename; \
		echo "Migração $$filename aplicada com sucesso!"; \
	done

up:
	@docker-compose up -d
	@echo "PostgreSQL iniciado!"

down:
	@docker-compose down
	@echo "PostgreSQL parado e removido!"


api:
	@go run ./cmd/api/main.go
