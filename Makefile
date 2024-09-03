run:
	@go run cli/main.go

sqlc:
	sqlc generate -f ./data/sqlite/sqlc.yaml
	@echo "sqlc done!"