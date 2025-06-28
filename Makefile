.PHONY: setup swag run sqlc deploy down downv db-new db-up db-down db-status db-dump

DBMATE_DIR := internal/adapter/outbound/db/migrations
DB_SCHEMA_FILE := internal/adapter/outbound/db/schema.sql
DATABASE_URL ?= "postgres://postgres:postgres@localhost:5321/smartfarm?sslmode=disable"

setup:
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install github.com/amacneil/dbmate/v2@latest

swag:
	swag init -g cmd/server/main.go -o ./docs --parseDependency --parseInternal

run:
	make swag
	go run cmd/server/main.go

sqlc:
	make db-dump
	DATABASE_URL=$(DATABASE_URL) sqlc generate -f internal/adapter/outbound/db/sqlc.yml

db-new:
	DATABASE_URL=$(DATABASE_URL) dbmate -d $(DBMATE_DIR) new $(filter-out $@,$(MAKECMDGOALS))

db-up:
	DATABASE_URL=$(DATABASE_URL) dbmate -d $(DBMATE_DIR) up

db-down:
	DATABASE_URL=$(DATABASE_URL) dbmate -d $(DBMATE_DIR) down

db-dump:
	docker exec smartfarm-postgres pg_dump -U postgres -d smartfarm --schema-only > $(DB_SCHEMA_FILE)

db-status:
	DATABASE_URL=$(DATABASE_URL) dbmate -d $(DBMATE_DIR) status

deploy:
	POSTGRES_PORT=5321 docker compose -f deploy/docker-compose-postgres.yml up -d

down:
	docker compose -f deploy/docker-compose-postgres.yml down

downv:
	docker compose -f deploy/docker-compose-postgres.yml down