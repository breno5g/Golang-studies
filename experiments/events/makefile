.PHONY: default run-server run-producer race test

default: run-server

run-server:
	@go run cmd/consumer/main.go

run-producer:
	@go run cmd/producer/main.go

race:
	@go run --race cmd/consumer/main.go

test:
	@go test -v ./...