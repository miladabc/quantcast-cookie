BINARY_NAME=most_active_cookie

build:
	go build -o $(BINARY_NAME) -v cmd/main.go

test:
	go test ./...

lint:
	golangci-lint run ./...
