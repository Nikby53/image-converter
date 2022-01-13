GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
BINARY_NAME=image-converter
LINTER=golangci-lint
DOCKER=docker-compose


.PHONY: build
build:
	$(GOBUILD) -x ./cmd/=main.go

.PHONY: run
run:
	$(GORUN) ./cmd/api/main.go

.PHONY: mocks
mocks:
	@~/go/bin/mockgen -source=internal/service/service.go -package=mock

.PHONY: lint
lint:
	$(LINTER) run

.PHONY: test
test:
	$(GOTEST) -v --short ./...

.PHONY: integration
integration:
	docker-compose -f docker-compose.local.yml up --build -d
	go test  -v ./internal/repository -timeout 30s -run ^TestRepository_Transactional$
	docker stop postgresql
