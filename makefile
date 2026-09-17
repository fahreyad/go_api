build:
	@go build -o bin/api ./cmd/api

run: build
	@./bin/api

migrate_up:
	@go run ./cmd/migrate up

migrate_down:
	@go run ./cmd/migrate down