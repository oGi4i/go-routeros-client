LOCAL_BIN:=$(CURDIR)/bin
GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
GOLANGCI_TAG:=v1.52.2

export CGO_ENABLED=0
export GOSUMDB=sum.golang.org
export GONOPROXY=
export GONOSUMDB=
export GOPRIVATE=
export GOPROXY=

.PHONY: bin-deps
bin-deps:
ifeq (,$(wildcard $(GOLANGCI_BIN)))
	$(info Installing golangci-lint dependency...)
	tmp=$$(mktemp -d) && cd $$tmp && pwd && go mod init temp && \
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_TAG)
endif

.PHONY: lint
lint: bin-deps
	$(info Running lint...)
	$(GOLANGCI_BIN) run --new-from-rev=origin/master --config=.golangci.yaml ./...

.PHONY: lint-full
lint-full: bin-deps
	$(info Running lint-full...)
	$(GOLANGCI_BIN) run --config=.golangci.yaml ./...

.PHONY: test
test:
	$(info Running tests...)
	@go test -v -coverprofile=cover.out ./...

.PHONY: cover
cover: test
	$(info Generating coverage...)
	@go tool cover -html=cover.out -o=cover.html
