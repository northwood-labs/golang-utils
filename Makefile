#-------------------------------------------------------------------------------
# Global variables.

GO=/usr/local/bin/go

#-------------------------------------------------------------------------------
# Running `make` will show the list of subcommands that will run.

all: help

.PHONY: help
## help: [help] Prints this help message.
help:
	@echo "Usage:"
	@echo ""
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

#-------------------------------------------------------------------------------
# Clean Go code

.PHONY: clean
## clean: [clean] Clean Go's module cache.
clean: clean-app clean-go

.PHONY: clean-app
## clean-app: [clean] Cleans local build directories
clean-app:
	rm -Rf ./bin ./dist

.PHONY: clean-go
## clean-go: [clean] clean Go's module cache
clean-go:
	$(GO) clean -i -r

.PHONY: clean-go-deep
## clean-go-deep: [clean] deep-clean Go's module cache
clean-go-deep:
	$(GO) clean -i -r -x -testcache -modcache -cache

#-------------------------------------------------------------------------------
# Golang docs

.PHONY: docs
docs:
	@ MODPATH="$$(cat go.mod | grep ^module | awk '{print $$2}')" && echo "http://localhost:6060/pkg/$$MODPATH" && open "http://localhost:6060/pkg/$$MODPATH" && godoc -http=:6060

#-------------------------------------------------------------------------------
# Testing

.PHONY: test
## test: [testing] Run ALL testing tasks.
test:
	@ echo " "
	@ echo "=====> Running go test..."
	$(GO) test ./...

#-------------------------------------------------------------------------------
# Linting

.PHONY: godeps
## godeps: [lint] go.mod lint, update, and download.
godeps:
	@ echo " "
	@ echo "=====> Running go get -v..."
	$(GO) mod tidy -go=1.17 && $(GO) get -v ./...

.PHONY: gofmt
## gofmt: [lint] Rewrite the source files to match canonical Go format.
gofmt:
	@ echo " "
	@ echo "=====> Running gofumpt and goimports..."
	gofumpt -s -w . && goimports -e -w .

.PHONY: golint
## golint: [lint] Lints all Golang code.
golint: godeps gofmt
	@ echo " "
	@ echo "=====> Running golangci-lint..."
	golangci-lint run --fix **/*.go

.PHONY: pylint
## pylint: [lint] Lints all Python code.
pylint:
	@ echo " "
	@ echo "=====> Running yapf..."
	find . -type f -name "*.py"  -not -path "*/venv/*" | xargs -I% yapf --in-place "%"

.PHONY: markdownlint
## markdownlint: [lint] Lints all Markdown files.
markdownlint:
	@ echo " "
	@ echo "=====> Running Markdownlint..."
	npx markdownlint-cli --fix '**/*.md' --ignore 'node_modules'

.PHONY: lint
## lint: [lint] Runs ALL linting tasks.
lint: markdownlint pylint golint

#-------------------------------------------------------------------------------
# Git Tasks

.PHONY: tag
## tag: [release] Git tags (and GPG-signs) the release.
tag:
	@ if [ $$(git status -s -uall | wc -l) != 1 ]; then echo 'ERROR: Git workspace must be clean.'; exit 1; fi;

	@echo "This release will be tagged as: $$(cat ./VERSION)"
	@echo "This version should match your release. If it doesn't, re-run 'make version'."
	@echo "---------------------------------------------------------------------"
	@read -p "Press any key to continue, or press Control+C to cancel. " x;

	@echo " "
	@chag update $$(cat ./VERSION)
	@echo " "

	@echo "These are the contents of the CHANGELOG for this release. Are these correct?"
	@echo "---------------------------------------------------------------------"
	@chag contents
	@echo "---------------------------------------------------------------------"
	@echo "Are these release notes correct? If not, cancel and update CHANGELOG.md."
	@read -p "Press any key to continue, or press Control+C to cancel. " x;

	@echo " "

	git add .
	git commit -a -m "Preparing the $$(cat ./VERSION) release."
	chag tag --sign

.PHONY: version
## version: [release] Sets the version for the next release. Pre-requisite for a release tag.
version:
	@echo "Current version: $$(cat ./VERSION)"
	@read -p "Enter new version number: " nv; \
	printf "$$nv" > ./VERSION
