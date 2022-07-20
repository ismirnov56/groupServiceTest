#!/bin/sh

migrate -source file://migrations -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DBNAME}?sslmode=${DB_SSLMODE}" up

go run ./cmd/main.go