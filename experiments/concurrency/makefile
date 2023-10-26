.PHONY: default run race test

default: test

run: 
	@go run main.go

race:
	@go run --race main.go

test:
	@ab -n 10000 -c 100 http://localhost:3001/
	@curl http://localhost:3001
