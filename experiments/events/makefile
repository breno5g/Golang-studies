.PHONY: default run race test

default: run

run:
	@go run cmd/consumer/main.go

race:
	@go run --race cmd/consumer/main.go

test:
	@go test -v ./...