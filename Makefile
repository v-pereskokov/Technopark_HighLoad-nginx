.PHONY: all test run

all: test

test:
	go test ./...

run:
	go run ./src/main.go
