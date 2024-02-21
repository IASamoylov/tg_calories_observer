# Change these variables as necessary.
# source ./build/.env
include ./.build/.env
export


LOCAL_BIN?=$(CURDIR)/bin
CONTAINER_REGISTRY?=
BUILD_APP_VERSION?=dev
BUILD_SHA_SHORT?=${BUILD_APP_VERSION}
GOOSE_DRIVER=postgres
APP_POSTGRES_USER?=postgres
APP_POSTGRES_PASS?=
APP_POSTGRES_PORT?=5432
APP_POSTGRES_SSL_MODE?=prefer

export PATH := $(PATH):${LOCAL_BIN}

# ==================================================================================== #
# LDFLAGS ENVS
# ==================================================================================== #
APP_LDFLAGS_MODULE_NAME=${shell head -n 1 go.mod | cut -c 8-}
APP_LDFLAGS=-X '${APP_LDFLAGS_MODULE_NAME}/internal/config/debug.AppName=${APP_NAME}'\
			-X '${APP_LDFLAGS_MODULE_NAME}/internal/config/debug.Version=${BUILD_APP_VERSION}'\
			-X '${APP_LDFLAGS_MODULE_NAME}/internal/config/debug.GithubSHA=${BUILD_SHA}'\
			-X '${APP_LDFLAGS_MODULE_NAME}/internal/config/debug.GithubSHAShort=${BUILD_SHA_SHORT}'


# ==================================================================================== #
# TESTS ENVS
# ==================================================================================== #


GO_UNIT_TEST_DIRECTORY:=./internal/...
GO_UNIT_TEST_COVER_PKG:=${GO_TEST_DIRECTORY}
GO_UNIT_TEST_COVER_EXCLUDE:=mocks

GO_INTEGRATION_TEST_DIRECTORY:=./e2e/...
GO_INTEGRATION_TEST_COVER_PKG:=${GO_UNIT_TEST_DIRECTORY}
GO_INTEGRATION_TEST_COVER_EXCLUDE:=mocks


# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# ==================================================================================== #
# INSTALL DEPENDENCIES
# ==================================================================================== #

.PHONY: .install-lint
.install-lint:
ifeq ($(wildcard $(LOCAL_BIN)/golangci-lint),)
	GOPATH=LOCAL_BIN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh
endif


.PHONY: ci-cd-deps
ci-cd-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.10.0
	GOBIN=$(LOCAL_BIN) go install gotest.tools/gotestsum@latest

## bin-deps: installs the dependencies for the correct operation of the application
.PHONY: bin-deps
bin-deps: .install-lint ci-cd-deps
	GOBIN=$(LOCAL_BIN) go install go.uber.org/mock/mockgen@latest
	GOBIN=$(LOCAL_BIN) go install github.com/wadey/gocovmerge@latest
	GOBIN=$(LOCAL_BIN) go install github.com/onsi/ginkgo/v2/ginkgo
	go get github.com/golang/mock/mockgen


# ==================================================================================== #
# INFRA
# ==================================================================================== #

## infra: starts ./.build/docker-compose.yaml with force recreate
.PHONY: infra
infra:
	docker-compose -f ./.build/docker-compose.yaml up -d --force-recreate --wait

## migration-reset: rollback migrations
.PHONY: migration-reset
migration-reset:
	$(LOCAL_BIN)/goose -dir "./migrations" "host=localhost user=${APP_POSTGRES_USER} dbname=${APP_NAME} port=${APP_POSTGRES_PORT} sslmode=${APP_POSTGRES_SSL_MODE} password=${APP_POSTGRES_PASS} " reset

## migration-up: apply migrations
.PHONY: migration-up
migration-up:
	$(LOCAL_BIN)/goose -dir "./migrations" "host=localhost user=${APP_POSTGRES_USER} dbname=${APP_NAME} port=${APP_POSTGRES_PORT} sslmode=${APP_POSTGRES_SSL_MODE} password=${APP_POSTGRES_PASS} " up

## migration: rollback migrations and apply them after
.PHONY: migration
migration: migration-reset migration-up

## migration-infra: launches the local environment environment in docker after which there is a rollback migrations and apply them after
.PHONY: migration-infra
migration-infra: infra migration

# ==================================================================================== #
# BUILD
# ==================================================================================== #

## codegen: runs a command `go generate` to create mockups using the tool mockgen
.PHONY: codegen
codegen:
	go generate ./...

## build: builds ./cmd/telegram/main.go to output ${LOCAL_BIN}/${BINARY_NAME}
.PHONY: build
build:
	CGO_ENABLED=0 go build -ldflags="${APP_LDFLAGS}" -o=${LOCAL_BIN}/${APP_NAME} ./cmd/telegram/main.go

## build-docker: builds docker image
.PHONY: build-docker
build-docker:
	docker build \
		--build-arg APP_LDFLAGS="-s -w ${APP_LDFLAGS}" \
		--build-arg GO_VERSION=${GO_VERSION} \
		--tag ${CONTAINER_REGISTRY}${APP_NAME}:${BUILD_SHA_SHORT} \
		--file .build/Dockerfile \
		.

## push-docker: push image to registry
.PHONY: push-docker
push-docker:
	docker push ${CONTAINER_REGISTRY}${APP_NAME}:${BUILD_SHA_SHORT}

## run-docker: run docker image with binding port 9090
.PHONY: run-docker
run-docker: build-docker
	docker ps -aq --filter "name=${APP_NAME}" | xargs -r docker rm -f
	docker run -p 9090:9090 --name ${APP_NAME} -d ${APP_NAME}:${BUILD_SHA_SHORT}

## run: runs web server
.PHONY: run
run: build
	${LOCAL_BIN}/${APP_NAME}

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## test: runs tests via gotestsum with coverage
.PHONY: test
test: 
	GOEXPERIMENT=nocoverageredesign $(LOCAL_BIN)/gotestsum \
		--format testname \
		--packages $(GO_UNIT_TEST_DIRECTORY) \
		--junitfile $(GO_UNIT_TEST_REPORT) \
		--junitfile-testcase-classname relative \
		-- -cover -covermode=count -coverprofile=$(GO_UNIT_TEST_COVER_PROFILE).tmp -coverpkg=$(GO_UNIT_TEST_COVER_PKG)
	grep -vE '$(GO_UNIT_TEST_COVER_EXCLUDE)' $(GO_UNIT_TEST_COVER_PROFILE).tmp > $(GO_UNIT_TEST_COVER_PROFILE)
	rm $(GO_UNIT_TEST_COVER_PROFILE).tmp


## e2e: runs integration tests via ginkgo with coverage
.PHONY: e2e-debug
e2e-debug: infra migration-up
	GOEXPERIMENT=nocoverageredesign $(LOCAL_BIN)/ginkgo -tags=e2e -v ./e2e/...
	#GOEXPERIMENT=nocoverageredesign $(LOCAL_BIN)/gotestsum \
#		--format testname \
#		--packages $(GO_INTEGRATION_TEST_DIRECTORY) \
#		--junitfile $(GO_INTEGRATION_TEST_REPORT) \
#		--junitfile-testcase-classname relative \
#		-- -tags=e2e -cover -covermode=count -coverprofile=$(GO_INTEGRATION_TEST_COVER_PROFILE).tmp -coverpkg=$(GO_INTEGRATION_TEST_COVER_PKG)
	#grep -vE '$(GO_INTEGRATION_TEST_COVER_EXCLUDE)' $(GO_INTEGRATION_TEST_COVER_PROFILE).tmp > $(GO_INTEGRATION_TEST_COVER_PROFILE)
	#rm $(GO_INTEGRATION_TEST_COVER_PROFILE).tmp

## e2e: runs integration tests via ginkgo with coverage
.PHONY: e2e
e2e: infra migration-up
	GOEXPERIMENT=nocoverageredesign $(LOCAL_BIN)/ginkgo -tags=e2e --succinct ./e2e/...
	#GOEXPERIMENT=nocoverageredesign $(LOCAL_BIN)/gotestsum \
#		--format testname \
#		--packages $(GO_INTEGRATION_TEST_DIRECTORY) \
#		--junitfile $(GO_INTEGRATION_TEST_REPORT) \
#		--junitfile-testcase-classname relative \
#		-- -tags=e2e -cover -covermode=count -coverprofile=$(GO_INTEGRATION_TEST_COVER_PROFILE).tmp -coverpkg=$(GO_INTEGRATION_TEST_COVER_PKG)
	#grep -vE '$(GO_INTEGRATION_TEST_COVER_EXCLUDE)' $(GO_INTEGRATION_TEST_COVER_PROFILE).tmp > $(GO_INTEGRATION_TEST_COVER_PROFILE)
	#rm $(GO_INTEGRATION_TEST_COVER_PROFILE).tmp

## cg-test: runs codegen before tests
.PHONY: cg-test
cg-test: codegen test

## cg-integration-test: runs codegen before integration-test
.PHONY: cg-e2e
cg-integration-test: codegen e2e

## cover: runs web display of coverage report for unit tests
.PHONY: cover
cover: test
	go tool cover -html=$(GO_UNIT_TEST_COVER_PROFILE)

## integration-cover: runs web display of coverage report for integration tests
.PHONY: e2e-cover
e2e-cover: e2e
	go tool cover -html=$(GO_INTEGRATION_TEST_COVER_PROFILE)


## merge-cover: integration tests and unit tests coverage report
.PHONY: merge-cover
merge-cover:
	$(LOCAL_BIN)/gocovmerge ./$(GO_UNIT_TEST_COVER_PROFILE) ./$(GO_INTEGRATION_TEST_COVER_PROFILE) > $(GO_TEST_COVER_PROFILE)

## lint: runs lint for changes using config .golangci.yaml
.PHONY: lint
lint: .install-lint
	$(LOCAL_BIN)/golangci-lint run --new-from-rev=origin/main --config=.golangci.yaml ./...

## lint-full: runs lint for all project using config .golangci.yaml
.PHONY: lint-full
lint-full: .install-lint
	$(LOCAL_BIN)/golangci-lint run --config=.golangci.yaml ./...