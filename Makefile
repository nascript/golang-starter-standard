.PHONY: run watch critic lint tests api-spec just-api-spec mock

# Exporting bin folder to the path for makefile
export PATH   := $(PWD)/bin:$(PATH)
# Default Shell
export SHELL  := bash
# Type of OS: Linux or Darwin.
export OSTYPE := $(shell uname -s)

# --- Tooling & Variables ----------------------------------------------------------------
include ./misc/make/tools.Makefile

install-deps: gotestsum mockery golangci-lint
deps: $(GOTESTSUM) $(MOCKERY) $(GOLANGCI)
deps:
	@ echo "Required Tools Are Available"

mock: $(MOCKERY)
	mockery --all --output=mocks/genmocks --outpkg=mocks

api-spec: tests
	@ echo "Re-generate API-Spec docs"
	@ swag init --parseDependency --parseInternal \
		--parseDepth 4 -g ./cmd/api/main.go
	@ echo "generate swagger file done"

just-api-spec:
	@ echo "Re-generate API-Spec docs"
	@ swag init --parseDependency --parseInternal \
		--parseDepth 4 -g ./cmd/api/main.go
	@ echo "generate swagger file done"

tests: $(MOCKERY) $(GOTESTSUM) lint critic
	@ echo "Trying to run all tests cases"
	@ gotestsum --format pkgname-and-test-fails \
		--hide-summary=skipped \
		-- -coverprofile=cover.out ./...
	@ rm cover.out

lint: $(GOLANGCI)
	@ echo "Trying apply linter"
	@ golangci-lint cache clean
	@ golangci-lint run -c .golangci.yaml ./...

critic:
	@ echo "Trying to audit the code"
	@ gocritic check -enableAll ./...

run:
	@echo "Run App"
	go mod tidy -compat=1.19
	go run ./cmd/api/main.go

watch:
	air --build.cmd "go build -o ./tmp/ ./cmd/api/main.go"