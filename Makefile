.PHONY: all test run

all: test

test:
	go test ./src/utils

run:
	go run ./src/main.go
