
# See `make help` for a list of all available commands.
# Configuration
PROJECT_NAME ?= url-shortener
BUILD_TIMESTAMP := $(shell date +%Y-%m-%d-%H-%M-%S)
CI_COMMIT_SHORT_SHA := $(shell git rev-parse --short HEAD)
.ONESHELL:
.DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables

.PHONY: up
up: build start

.PHONY: down
down: stop

.PHONY: build
build:
	@echo "Building backend..."
	@go build -o ./bin/urlshortener ./cmd/api/*
	@echo "Back end built!"

.PHONY: start
start:
	@echo "Starting backend..."
	@ ./bin/urlshortener &
	@echo "Back end started!"

.PHONY: stop
stop:
	@echo "stopping backend ..."
	@-pkill -SIGTERM -f "urlshortener"
	@echo "stopped backend ..."