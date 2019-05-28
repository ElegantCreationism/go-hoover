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

.PHONY: gen
gen:


.PHONY: unit_test
unit_test: gen
	go test -v -cover `go list ./... | grep -v tests_system`

.PHONY: dockerise
dockerise:
	docker build -t "quay.io/90poe/external-data-gateway:${IMAGE_TAG}" .

.PHONY: system_test
system_test: dockerise system_test_default healthcheck_tests

.PHONY: system_test
system_test_default:
	docker-compose -f docker-composition/default.yml down --volumes --remove-orphans
	docker-compose -f docker-composition/default.yml rm --force --stop -v
	IMAGE_TAG=${IMAGE_TAG} \
	docker-compose \
		-f docker-composition/default.yml \
		-f docker-composition/system-test-mask.yml \
		up -d --force-recreate --remove-orphans --build
	sleep 5
	go test -v -tags=system_tests ./tests_system/...
	docker-compose -f docker-composition/default.yml down --volumes --remove-orphans
	docker-compose -f docker-composition/default.yml rm --force --stop -v

.PHONY: healthcheck_tests
healthcheck_tests:
	docker-compose -f docker-composition/default.yml down --volumes --remove-orphans
	docker-compose -f docker-composition/default.yml rm --force --stop -v
	IMAGE_TAG=${IMAGE_TAG} \
	docker-compose \
		-f docker-composition/default.yml \
		-f docker-composition/system-test-mask.yml \
		up -d --force-recreate --remove-orphans --build
	sleep 5
	docker-compose -f docker-composition/default.yml down --volumes --remove-orphans
	docker-compose -f docker-composition/default.yml rm --force --stop -v


.PHONY: lint
lint:
	golangci-lint run

