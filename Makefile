include .env
export $(shell sed 's/=.*//' .env)

MIGRATIONS_PATH = ./cmd/migrate/migrations

# this command will compile project and create binary file named bin/social
build:
	@go build -o ./bin/yollapi ./cmd/api

# if compile operation is success, this command(run) runs the binary file on terminal
run: build
	@./bin/yollapi

compose-up:
	@docker compose up

compose-down:
	@docker compose stop

connect-database:
	@docker exec -it d4d437257b9e psql -U $(DB_USER) -d $(DB_NAME)

.PHONY: migration-create
migration-create:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(name)

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) up

.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) down $(step)

migrate-force:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) force $(version)