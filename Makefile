# Change these variables as necessary.
LOCAL_BIN?=$(CURDIR)/bin
APP_NAME:=tbco-web-server
APP_VERSION?=dev
# APP_LDFLAGS?="\"-s -w\""
APP_LDFLAGS?=

GO_VERSION:=1.20
GO_TEST_DIRECTORY:=./internal/...
GO_TEST_COVER_PROFILE?=unit.coverage.out
GO_TEST_REPORT?=unit.report.xml
GO_TEST_COVER_EXCLUDE:=mocks|config

export PATH:=$(PATH):$(LOCAL_BIN)
export GO111MODULE=on

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
	go build -o=${LOCAL_BIN}/${APP_NAME} ./cmd/telegram/main.go

## build-docker: builds docker image
.PHONY: build-docker
build-docker:
	docker build \
		--build-arg APP_LDFLAGS=${APP_LDFLAGS} \
		--build-arg GO_VERSION=${GO_VERSION} \
		--tag ${APP_NAME}:${APP_VERSION} \
		--file .build/Dockerfile \
		.

## run-docker: run docker image with binding port 9090
.PHONY: run-docker
run-docker: build-docker
	docker run -d -p 9090:9090 ${APP_NAME}:${APP_VERSION}

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
		-- -covermode=count -coverprofile=$(GO_TEST_COVER_PROFILE).tmp -coverpkg=$(GO_TEST_DIRECTORY)
	grep -vE '$(GO_TEST_COVER_EXCLUDE)' $(GO_TEST_COVER_PROFILE).tmp > $(GO_TEST_COVER_PROFILE)
	rm $(GO_TEST_COVER_PROFILE).tmp


## cg-test: runs codegen before tests
.PHONY: cg-test
cg-test: codegen
	GOEXPERIMENT=nocoverageredesign $(LOCAL_BIN)/gotestsum \
		--format testname \
		--packages $(GO_TEST_DIRECTORY) \
		--junitfile $(GO_TEST_REPORT) \
		--junitfile-testcase-classname relative \
		-- -covermode=count -coverprofile=$(GO_TEST_COVER_PROFILE).tmp -coverpkg=$(GO_TEST_DIRECTORY)
	grep -vE '$(GO_TEST_COVER_EXCLUDE)' $(GO_TEST_COVER_PROFILE).tmp > $(GO_TEST_COVER_PROFILE)
	rm $(GO_TEST_COVER_PROFILE).tmp

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