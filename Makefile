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