.PHONY: sandbox
sandbox:
	@go run ./cmd/sandbox/*.go

.PHONY: test
test:
	@go test ./...

.PHONY: fmt
fmt:
	@go fmt ./...
