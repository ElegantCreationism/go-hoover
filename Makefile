SHELL=/bin/bash
IMAGE_TAG := $(shell git rev-parse HEAD)

.PHONY: all
all: deps lint unit_test build system_test

.PHONY: ci
ci: lint unit_test system_test

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -a -o ./artifacts/svc .

.PHONY: deps
deps:
	go mod vendor

.PHONY: unit_test
unit_test:
	go test -v -cover `go list ./... | grep -v tests_system` -count=1

.PHONY: dockerise
dockerise:
	docker build -t "moh90poe/go-hoover:${IMAGE_TAG}" .

.PHONY: system_test
system_test: dockerise



.PHONY: lint
lint:
	golangci-lint run

