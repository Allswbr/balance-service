.PHONY: build
build:
	go build -o ./bin/balance-service -v ./cmd/apiserver

run:
	./bin/balance-service

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build