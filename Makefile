.PHONY: build
build:
	docker-compose build

run:
	docker-compose up

stop:
	docker-compose down

migrate-up:
	migrate -path ./schema -database 'postgres://postgres:0000@localhost:5432/postgres?sslmode=disable' up

migrate-down:
	migrate -path ./schema -database 'postgres://postgres:0000@localhost:5432/postgres?sslmode=disable' down

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build