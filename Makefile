.PHONY: run test test-race tidy

run:
	go run ./cmd/api

test:
	go test ./...

test-race:
	go test -race -count=1 ./...

tidy:
	go mod tidy
