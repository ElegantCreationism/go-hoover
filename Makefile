SHELL=/bin/bash

.PHONY: all
all: deps lint unit_test build docker_compose docker_compose_teardownm

.PHONY: ci
ci: lint unit_test docker_compose

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -a -o ./artifacts/svc .

.PHONY: deps
deps:
	go mod vendor

.PHONY: unit_test
unit_test:
	GOCHACHE=off go test -v -cover `go list ./... | grep -v tests_system` -count=1

.PHONY: dockerise
dockerise:
	docker build -t "moh90poe/go-hoover:latest" .

.PHONY: docker_compose
docker_compose: dockerise docker_compose_default

.PHONY: docker_compose_default
docker_compose_default:
	docker-compose -f docker-composition/default.yml down --volumes --remove-orphans
	docker-compose -f docker-composition/default.yml rm --force --stop -v
	docker-compose -f docker-composition/default.yml up -d --force-recreate --remove-orphans --build

.PHONY: docker_compose_teardown
docker_compose_teardown:
	docker-compose -f docker-composition/default.yml down --volumes --remove-orphans
	docker-compose -f docker-composition/default.yml rm --force --stop -v

.PHONY: lint
lint:
	golangci-lint run

