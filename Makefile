# GNUmakefile
SHELL := bash
.ONESHELL:
.DELETE_ON_ERROR:
.SHELLFLAGS := -eu -o pipefail -c
MAKEFLAGS += --warn-undefined-variables

.PHONY: help
help: ## Display this help section
	@echo -e "Usage: make <command>\n\nAvailable commands are:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {printf "  %-38s %s\n", $$1, $$2}' ${MAKEFILE_LIST}
.DEFAULT_GOAL := help
################################################################################

# TARGETS
.PHONY: test cover fmt gofmt govet deps

test: fmt ## Generate mocks and run tests
	go test -cover ./opentok

cover: ## Runs test for the given package name, and shows the coverage report
	go test -coverprofile coverage.out ./opentok && go tool cover -html coverage.out

fmt: gofmt govet deps ## Clean dependencies and run code linting / formatting

gofmt:
	@echo gofmt -l -w
	@gofmt -l -w $(shell find . -type f -name "*.go" -not -path "./vendor/*")

govet:
	go vet ./opentok

deps:
	go mod tidy
