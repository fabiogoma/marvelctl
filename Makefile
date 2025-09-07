APP_NAME := marvelctl
PKG := ./...
GO := go

.PHONY: build

build: deps vet lint
	$(GO) build -o $(APP_NAME) main.go

clean:
	@rm -rf $(APP_NAME)
	@rm -rf coverage.out

test:
	$(GO) test -v $(PKG)

test-coverage:
	$(GO) test -coverprofile=coverage.out $(PKG)
	$(GO) tool cover -func=coverage.out

fmt:
	$(GO) fmt $(PKG)

lint:
	@if [ -x "$$(command -v golangci-lint)" ]; then \
		golangci-lint run $(PKG); \
	else \
		echo "golangci-lint not installed. Skipping lint."; \
	fi

vet:
	$(GO) vet $(PKG)

run:
	$(GO) run main.go

install:
	$(GO) install $(PKG)

deps:
	$(GO) mod tidy
