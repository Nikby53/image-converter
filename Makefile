GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
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
	mockery --case underscore --dir ./internal/service/ --output ./internal/service/mocks --all --disable-version-string
	mockery --case underscore --dir ./internal/broker/ --output ./internal/broker/mocks --all --disable-version-string

.PHONY: lint
lint:
	$(LINTER) run

.PHONY: test
test:
	$(GOTEST) ./... -v
