PWD = $(shell pwd)

MODULE = my_books
IMAGE_TAG ?= $(MODULE)
GITHUB_SHA ?= $(MODULE)

REPO_PATH = /go/src/github.com/ngkimphuongtrang/$(MODULE)
SRC = `go list -f {{.Dir}} ./... | grep -v /vendor/`

ndef = $(if $(value $(1)),,$(error $(1) not set))

lint:
	@echo "==> Running lint check..."
	@golangci-lint --config docker/.golangci.yml run
	@go vet $(SRC)

test:
	@echo "==> Running tests..."
	@go clean -testcache
	@go test `go list ./... | grep -v cmd` -race -p 1 --cover

generate:
	@echo "==> Generating code..."
	@go generate ./...

test-up:
	@docker-compose \
		-f docker/docker-compose.test.yml \
		-p $(GITHUB_SHA) up \
		--force-recreate \
		--abort-on-container-exit \
		--exit-code-from app \
		--build

test-down:
	@docker-compose \
		-f docker/docker-compose.test.yml \
		-p $(GITHUB_SHA) down \
 		-v --rmi local

dev-up:
	@docker compose \
		-f docker/docker-compose.dev.yml \
		-p $(GITHUB_SHA) up --build -d

dev-down:
	@docker compose \
		-f docker/docker-compose.dev.yml \
		-p $(GITHUB_SHA) down \
 		-v --rmi local

dev-ps:
	@docker compose \
		-f docker/docker-compose.dev.yml \
		-p $(GITHUB_SHA) ps -a

.PHONY: all fmt lint test install test-up test-down dev-up dev-down dev-ps build generate
