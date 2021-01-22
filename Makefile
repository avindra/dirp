build:
	@go build -ldflags "-s -w" .

run:
	@go run .

test:
	@go test ./...
