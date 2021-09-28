GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
BINARY_NAME=image-converter
LINTER=golangci-lint

.PHONY: build
build:
	$(GOBUILD) -x ./cmd/main.go

.PHONY: run
run:
	$(GORUN) ./cmd/main.go

.PHONY: mocks
mocks:
	@~/go/bin/mockgen -source=internal/service/service.go -package=mock

.PHONY: lint
lint:
	$(LINTER) run

.PHONY: test
test:
	$(GOTEST) ./... -v
