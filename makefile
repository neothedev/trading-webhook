# Makefile

run:
	APP_ENV=dev go run ./cmd/server

build:
	GOOS=linux GOARCH=amd64 go build -o bin/server ./cmd/server

tidy:
	go mod tidy

test:
	go test ./...

lint:
	golangci-lint run
