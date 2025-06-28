setup:
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

swag:
	swag init -g cmd/server/main.go -o ./docs --parseDependency --parseInternal

run:
	make swag
	go run cmd/server/main.go

sqlc:
	sqlc generate -f internal/adapter/outbound/db/sqlc.yml

deploy:
	POSTGRES_PORT=5321 docker compose -f deploy/docker-compose-postgres.yml up -d

down:
	docker compose -f deploy/docker-compose-postgres.yml down

downv:
	docker compose -f deploy/docker-compose-postgres.yml down

.PHONY: setup swag run sqlc deploy down downv