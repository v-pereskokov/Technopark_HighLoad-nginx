.PHONY: all test run

all: test

test:
	go test ./test

run:
	go run ./src/main.go
