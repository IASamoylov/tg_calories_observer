export GO111MODULE=on

LOCAL_BIN?=$(CURDIR)/bin

.PHONY: bin-deps
bin-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.10.0
	GOBIN=$(LOCAL_BIN) go install github.com/golang/mock/mockgen@v1.6.0
	GOBIN=$(LOCAL_BIN) go install gotest.tools/gotestsum@latest

GO_TEST_DIRECTORY:=./internal/...
GO_TEST_COVER_PROFILE?=unit.coverage.out
GO_TEST_REPORT?=unit.report.xml

.PHONY: test
test:
	$(LOCAL_BIN)/gotestsum \
		--format testname \
		--packages $(GO_TEST_DIRECTORY) \
		--junitfile $(GO_TEST_REPORT) \
		--junitfile-testcase-classname relative \
		-- -covermode=count -coverprofile=$(GO_TEST_COVER_PROFILE) -coverpkg=$(GO_TEST_DIRECTORY)

.PHONY: cover
cover:
	go tool cover -html=$(GO_TEST_COVER_PROFILE)

.PHONY: build
build:
	go build ./...

.PHONY: .install-lint
.install-lint:
ifeq ($(wildcard $(LOCAL_BIN)/golangci-lint),)
	$(info Downloading golangci-lint v$(GOLANGCI_TAG))
	GOPATH=LOCAL_BIN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh
endif

.PHONY: lint
lint: .install-lint
	echo $(LOCAL_BIN)
	$(LOCAL_BIN)/golangci-lint run --new-from-rev=origin/main --config=.golangci.yaml ./...

.PHONY: lint-full
lint-full: .install-lint
	$(LOCAL_BIN)/golangci-lint run --config=.golangci.yaml ./...
