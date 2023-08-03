.PHONY: build run test

build:
	@go build -o ./bin/api ./cmd/api

run: build
	@./bin/api

test:
	@go test -v ./...
