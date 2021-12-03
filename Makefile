SHELL := /bin/bash

# Gather Host Information
# {
UNAME := $(shell uname -s)
OS    := UNKNOWN

ifeq ($(UNAME),Linux)
    OS := linux
endif

ifeq ($(UNAME),Darwin)
    OS := macos
endif
# }

# Some useful variables...

ERROR   := "   \033[41;1m error \033[0m "
INFO    := "    \033[34;1m info \033[0m "
OK      := "      \033[32;1m ok \033[0m "
WARNING := " \033[33;1m warning \033[0m "

CWD    := $(shell pwd)

# Task: ling | Run linter
# {
.PHONY: lint
lint:
	docker run --rm -v $(CWD):/app -w /app golangci/golangci-lint:v1.43.0 golangci-lint run -v
# }

# Task: test | Run all go tests with race condition.
# {
.PHONY: test
unit-test:
	go test -v -race -timeout 30s ./...
# }
