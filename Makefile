run:
	@go run cli/main.go

build:
	@go build && go install

sqlc:
	sqlc generate -f ./data/sqlite/sqlc.yaml
	@echo "sqlc done!"

