include .env

migrate_up:
	migrate -path=migrations -database "${POSTGRES_CONNECTION_STRING}" -verbose up

migrate_down:
	migrate -path=migrations -database "${POSTGRES_CONNECTION_STRING}" -verbose down

.PHONY: migrate_up migrate_down