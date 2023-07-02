# Change these variables as necessary.
# source ./build/.env
include ./.build/.env
export



LOCAL_BIN?=$(CURDIR)/bin
CONTAINER_REGISTRY?=
BUILD_APP_VERSION?=dev
BUILD_SHA_SHORT?=${BUILD_APP_VERSION}

export PATH := $(PATH):${LOCAL_BIN}

# ==================================================================================== #
# LDFLAGS ENVS
# ==================================================================================== #
APP_LDFLAGS_MODULE_NAME=${shell head -n 1 go.mod | cut -c 8-}
APP_LDFLAGS=-X '${APP_LDFLAGS_MODULE_NAME}/internal/config/debug.AppName=${APP_NAME}'\
			-X '${APP_LDFLAGS_MODULE_NAME}/internal/config/debug.Version=${BUILD_APP_VERSION}'\
			-X '${APP_LDFLAGS_MODULE_NAME}/internal/config/debug.GithubSHA=${BUILD_SHA}'\
			-X '${APP_LDFLAGS_MODULE_NAME}/internal/config/debug.GithubSHAShort=${BUILD_SHA_SHORT}'\
			-X '${APP_LDFLAGS_MODULE_NAME}/internal/config/debug.BuildedAt=$(shell date -u)'


# ==================================================================================== #
# TESTS ENVS
# ==================================================================================== #


GO_TEST_DIRECTORY:=./internal/...
GO_TEST_COVER_PKG:=${GO_TEST_DIRECTORY}
GO_TEST_COVER_PROFILE?=unit.coverage.out
GO_TEST_REPORT?=unit.report.xml
GO_TEST_COVER_EXCLUDE:=mocks

GO_INTEGRATION_TEST_DIRECTORY:=./integration_test/...
GO_INTEGRATION_TEST_COVER_PKG:=${GO_TEST_DIRECTORY}
GO_INTEGRATION_TEST_COVER_PROFILE?=integration.coverage.out
GO_INTEGRATION_TEST_REPORT?=integration.report.xml
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

## bin-deps: installs the dependencies for the correct operation of the application
.PHONY: .install-lint
.install-lint:
ifeq ($(wildcard $(LOCAL_BIN)/golangci-lint),)
	GOPATH=LOCAL_BIN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh
endif

.PHONY: bin-deps
bin-deps: .install-lint
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.10.0
	GOBIN=$(LOCAL_BIN) go install github.com/golang/mock/mockgen@v1.6.0
	GOBIN=$(LOCAL_BIN) go install gotest.tools/gotestsum@latest
	go get github.com/golang/mock/mockgen

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
		--packages $(GO_TEST_DIRECTORY) \
		--junitfile $(GO_TEST_REPORT) \
		--junitfile-testcase-classname relative \
		-- -cover -covermode=count -coverprofile=$(GO_TEST_COVER_PROFILE).tmp -coverpkg=$(GO_TEST_COVER_PKG)
	grep -vE '$(GO_TEST_COVER_EXCLUDE)' $(GO_TEST_COVER_PROFILE).tmp > $(GO_TEST_COVER_PROFILE)
	rm $(GO_TEST_COVER_PROFILE).tmp


## integration-test: runs integration tests via gotestsum with coverage
.PHONY: integration-test
integration-test: 
	GOEXPERIMENT=nocoverageredesign $(LOCAL_BIN)/gotestsum \
		--format testname \
		--packages $(GO_INTEGRATION_TEST_DIRECTORY) \
		--junitfile $(GO_INTEGRATION_TEST_REPORT) \
		--junitfile-testcase-classname relative \
		-- -tags=integration_test -cover -covermode=count -coverprofile=$(GO_INTEGRATION_TEST_COVER_PROFILE).tmp -coverpkg=$(GO_INTEGRATION_TEST_COVER_PKG)
	grep -vE '$(GO_INTEGRATION_TEST_COVER_EXCLUDE)' $(GO_INTEGRATION_TEST_COVER_PROFILE).tmp > $(GO_INTEGRATION_TEST_COVER_PROFILE)
	rm $(GO_INTEGRATION_TEST_COVER_PROFILE).tmp

## cg-test: runs codegen before tests
.PHONY: cg-test
cg-test: codegen test

## cg-integration-test: runs codegen before integration-test
.PHONY: cg-integration-test
cg-integration-test: codegen integration-test

## cover: runs web display of coverage report
.PHONY: cover
cover: test
	go tool cover -html=$(GO_TEST_COVER_PROFILE)

## lint: runs lint for changes using config .golangci.yaml
.PHONY: lint
lint: .install-lint
	$(LOCAL_BIN)/golangci-lint run --new-from-rev=origin/main --config=.golangci.yaml ./...

## lint-full: runs lint for all project using config .golangci.yaml
.PHONY: lint-full
lint-full: .install-lint
	$(LOCAL_BIN)/golangci-lint run --config=.golangci.yaml ./...