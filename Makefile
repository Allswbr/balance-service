.PHONY: build
build:
	# -o - записывает исполняемый файл, -v - печать имен пакетов по мере компиляции
	go build -o balance-service -v ./cmd/apiserver

run:
	./balance-service

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build