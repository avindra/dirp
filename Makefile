# Uses for this Makefile: CI/CD and local development

build:
	@go build -ldflags "-s -w" .

test:
	@go test -v ./...
