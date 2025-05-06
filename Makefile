default: fmt lint install generate

GO_BUILD_FLAGS = -trimpath

.PHONY: build
build:
	go build $(GO_BUILD_FLAGS) -v ./... 

.PHONY: install
install: build
	go install -v ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: generate
generate:
	cd tools; go generate ./...

.PHONY: fmt
fmt:
	gofmt -s -w -e .

TEST_CMD := gotestsum 
ifeq (, $(shell command -v gotestsum))
	TEST_CMD := go test $(GO_BUILD_FLAGS)
endif

.PHONY: test
test:
	$(TEST_CMD) ./...
