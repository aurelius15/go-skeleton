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
test: all-tests coverage ## Run all tests with coverage

unit-tests: ## Run unit tests of the project
	$(GOTEST) -race ./... $(OUTPUT_OPTIONS)

all-tests: ## Run all tests of the project
	$(GOTEST) --tags=integration -race ./... $(OUTPUT_OPTIONS)

coverage: ## Run the tests of the project and export the coverage
	$(GOTEST) -cover -coverprofile=coverage.txt -covermode=atomic ./...
	$(GOCMD) tool cover -func coverage.txt

## Check code:
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