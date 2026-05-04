.PHONY: help test test-cover lint fmt check tidy deps clean

# Default target
help: ## Show available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-16s\033[0m %s\n", $$1, $$2}'

test: ## Run all tests with race detector
	@go test -race -count=$${TEST_COUNT:-10} ./...

test-cover: ## Run all tests with coverage report
	@go test -race -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

lint: ## Run go vet and golangci-lint
	@go vet ./...
	@golangci-lint run ./...

fmt: ## Format source files
	@gofmt -w .
	@go run golang.org/x/tools/cmd/goimports@latest -local github.com/maxence2997 -w .

check: ## Run fmt-check, lint, all tests
	@echo "── fmt ──"
	@test -z "$$(gofmt -l .)" || (echo "formatting issues — run 'make fmt'"; exit 1)
	@test -z "$$(go run golang.org/x/tools/cmd/goimports@latest -local github.com/maxence2997 -l .)" || (echo "import issues — run 'make fmt'"; exit 1)
	@echo "── lint ──"
	@$(MAKE) --no-print-directory lint
	@echo "── test ──"
	@$(MAKE) --no-print-directory test
	@echo "── all passed ──"

tidy: ## Tidy module dependencies
	@go mod tidy

deps: ## Download all modules and sync go.sum, then tidy
	@go mod download
	@go mod tidy

clean: ## Remove build artifacts and test cache
	@rm -f coverage.out coverage.html
