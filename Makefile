# Exported environment variables
export GOBIN := $(shell pwd)/bin

.PHONY: default
default: help

.PHONY: all
all: modtidy lint test build ## Run lint, test and build

.PHONY: modtidy
modtidy: ## Update go.mod
	go mod tidy

.PHONY: build
build: ## Build and install binaries to $(GOBIN)
	go install -v ./gosort

.PHONY: clean
clean: ## Clean build files and artifacts.
	go clean ./...
	rm -rfv $(GOBIN)

.PHONY: lint
lint: ## Run the linter
	go vet

test_path ?= ./...
.PHONY: test
test: ## Run the tests
	go test -race -coverprofile=.coverage.out -covermode=atomic ${test_path}

.PHONY: open_coverage_report
open_coverage_report:  ## Open the last code coverage report in your default browser
	go tool cover -html=.coverage.out

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: help
help: ## Show the list of available makefile targets
	@grep -hE '^[/a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-26s\033[0m %s\n", $$1, $$2}'
