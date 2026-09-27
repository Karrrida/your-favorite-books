include .env
export

MIGRATIONS_PATH=./migrations

migrate-up:
	migrate -database "$(DATABASE_URL)" -path $(MIGRATIONS_PATH) up

migrate-down:
	migrate -database "$(DATABASE_URL)" -path $(MIGRATIONS_PATH) down

migrate-version:
	migrate -database "$(DATABASE_URL)" -path $(MIGRATIONS_PATH) version
