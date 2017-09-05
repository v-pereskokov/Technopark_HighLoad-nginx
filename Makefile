.PHONY: all test run

all: test

test:
	go test ./src/configs

run:
	go run ./src/main.go
