MAKEFILE_URL = https://raw.githubusercontent.com/mercadolibre/fury_go-core/master/Makefile
GITHUB_TOKEN ?= $(shell gh auth token 2>/dev/null)

# Sets the default goal to be used if no targets were specified on the command line.
.DEFAULT_GOAL := all

# Catch all targets and delegate to the Makefile in the repository MAKEFILE_URL.
# This allows us to use the same Makefile for all projects.
# If you want to override a target, just define it at the end of this file.
# If the cURL command fails, it will fallback to the last version of the Makefile that was downloaded.
# If a previous version does not exist, an error is returned.
%:
	@if [ -z "${GITHUB_TOKEN}" ]; then echo "GITHUB_TOKEN is not set"; exit 1; fi
	@curl -H "Authorization: token $(GITHUB_TOKEN)" -s -f -o Makefile.common $(MAKEFILE_URL) || (test -f Makefile.common || (echo "ERROR: Unable to download Makefile" && exit 1))
	@$(MAKE) -f Makefile.common $@

.PHONY: run
run:
	@echo "=> Running application locally"
	@go run cmd/server/*.go

.PHONY: all
all: generate tidy format vet staticcheck test

.PHONY: generate
generate: init
	@echo "=> Running go generate"
	@go generate ./...

.PHONY: init
init:
	@echo "=> Initializing project"
	@git submodule init
	@git pull --recurse-submodule
	@go generate ./...