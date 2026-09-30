include .env
export

MIGRATIONS_PATH=./migrations

migrate-up:
	migrate -database "$(DATABASE_URL)" -path $(MIGRATIONS_PATH) up

migrate-up-force:
	migrate -database "$(DATABASE_URL)" -path $(MIGRATIONS_PATH) force 2

migrate-down:
	migrate -database "$(DATABASE_URL)" -path $(MIGRATIONS_PATH) down

migrate-version:
	migrate -database "$(DATABASE_URL)" -path $(MIGRATIONS_PATH) version
