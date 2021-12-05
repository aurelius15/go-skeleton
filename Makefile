GOCMD=go
GOTEST=$(GOCMD) test
BINARY_NAME=example

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: all test lint

all: help

## Test:
test: unit-test integration-test ## Run all tests

unit-test: ## Run unit tests of the project
	$(GOTEST) -v -race ./... $(OUTPUT_OPTIONS)

integration-test: ## Run integration tests of the project
	$(GOTEST) --tags=integration -v -race ./... $(OUTPUT_OPTIONS)

coverage: ## Run the tests of the project and export the coverage
	$(GOTEST) -cover -covermode=count -coverprofile=profile.cov ./...
	$(GOCMD) tool cover -func profile.cov

## Lint:
lint: ## Use golintci-lint on your project
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.43.0 golangci-lint run -v $(OUTPUT_OPTIONS)

## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)