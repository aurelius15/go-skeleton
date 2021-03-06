GOCMD=go
GOTEST=$(GOCMD) test
BINARY_NAME=example

DOCKER_COMPOSE_FILE = "deployments/local/docker-compose.yml"

RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
$(eval $(RUN_ARGS):;@:)

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
	$(GOTEST) -race ./... $(RUN_ARGS)

all-tests: ## Run all tests of the project
	$(GOTEST) --tags=integration -race ./...

coverage: ## Run the tests of the project and export the coverage
	$(GOTEST) -cover -coverprofile=coverage.txt -covermode=atomic ./...
	$(GOCMD) tool cover -func coverage.txt

## Check code:
lint: ## Use golintci-lint on your project
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.43.0 golangci-lint run -v

## Local environment:
up: ## Up services
	docker-compose -f ${DOCKER_COMPOSE_FILE} up --build -d $(RUN_ARGS)
	docker-compose -f ${DOCKER_COMPOSE_FILE} ps

down: ## Down all services
	docker-compose -f ${DOCKER_COMPOSE_FILE} down

restart: down up ## Restart all services

logs: ## Show logs
	docker-compose -f ${DOCKER_COMPOSE_FILE} logs $(RUN_ARGS)

## Profiling:
pprof: ## Heap
	curl -sK -v http://localhost:8080/debug/pprof/heap > heap.out
	go tool pprof -http=:8081 heap.out

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