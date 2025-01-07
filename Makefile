.PHONY: vendor

db_up:
	docker-compose up -d

db_down:
	docker-compose down

vendor:
	go mod vendor && go mod tidy && go mod verify

run-migrations:
	go run migrations/run_migrations.go


