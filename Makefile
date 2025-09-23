
export GOBIN := $(PWD)/bin
export PATH := $(GOBIN):$(PATH)

# Version information
VERSION ?= dev
COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS := -X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(DATE)

.PHONY: run
run:
	go run -ldflags "$(LDFLAGS)" . $(ARGS)

.PHONY: lint
lint:
	go tool golangci-lint run -v ./...

.PHONY: lint-fix
lint-fix:
	go tool golangci-lint run --fix -v ./...

.PHONY: test
test:
	go test -v -race ./...

.PHONY: cover
cover:
	go test -v -race -coverpkg=./... -coverprofile=coverage.out ./...

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: build
build:
	go build -ldflags "$(LDFLAGS)" -o bin/app .
