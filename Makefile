.PHONY: vendor

db_up:
	docker-compose up -d

db_down:
	docker-compose down

run-migrations:
	go run migrations/run_migrations.go

vendor:
	go mod tidy && go mod vendor && go mod verify
