# Makefile for Go projects.
#
# This Makefile makes an effort to provide standard make targets, as described
# by https://www.gnu.org/prep/standards/html_node/Standard-Targets.html.

include Makefile.*

SHELL := /bin/sh

GOIMPORTS := $(GORUN) golang.org/x/tools/cmd/goimports
GOLANGCI_LINT := $(GORUN) github.com/golangci/golangci-lint/cmd/golangci-lint

SOURCES := $(shell \
	find . -name '*.go' | \
	grep -Ev './(proto|protogen|third_party|vendor)/' | \
	xargs)
ifdef DEBUG
$(info SOURCES = $(SOURCES))
endif

PROJECT := "freegrow"
TAG := $(shell git tag -l --points-at HEAD)
COMMIT := $(shell git describe --always --long --dirty --tags)
VERSION := $(shell [ ! -z "${TAG}" ] && echo "${TAG}" || echo "${COMMIT}")
REVISION := $(shell git rev-parse HEAD)
GITREMOTE := "github.com/luanguimaraesla/freegrow"

DOCKER_IMAGE := "luanguimaraesla/freegrow"

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

# Use linker flags to provide version settings to the target
# Also build it with as much as possible static links. It may do the build a bit slower
LDFLAGS=-ldflags "-X=${GITREMOTE}/internal/global.Version=$(VERSION) -extldflags '-static'"

################################################################################
## Standard make targets
################################################################################

.DEFAULT_GOAL := all
.PHONY: all
all: fix install

.PHONY: install
install:
	@echo ">  Installing ${PROJECT}"
	@GOBINDIR=$(GOBINDIR) $(GOINSTALL) $(LDFLAGS) ./...

.PHONY: uninstall
uninstall:
	@echo "> Uninstalling ${PROJECT}"
	$(RM) $(GOPATH)/bin/$(PROJECT)

.PHONY: clean
clean:
	@echo ">  Cleaning build cache"
	@GOBINDIR=$(GOBINDIR) $(GOCLEAN)
	$(RM) coverage.out

.PHONY: check
check: test

################################################################################
## Go-like targets
################################################################################

.PHONY: build
build:
	@echo ">  Building binary"
	@mkdir -p $(GOBINDIR)
	$(GOBUILD) $(LDFLAGS) -o $(GOBINDIR)/ ./...

.PHONY: test
test:
	@echo ">  Executing tests"
	$(GOTEST) -v -tags=unit -coverprofile=coverage.out ./... $(SILENT_CMD_SUFFIX)

.PHONY: cover
cover: cover/text

.PHONY: cover/html
cover/html:
	$(GOTOOL) cover -html=coverage.out

.PHONY: cover/text
cover/text:
	$(GOTOOL) cover -func=coverage.out


################################################################################
## Linters and formatters
################################################################################

.PHONY: fix
fix:
	@echo ">  Making sure go.mod matches the source code"
	$(GOMOD) vendor
	$(GOMOD) tidy
ifneq ($(SOURCES),)
	$(GOIMPORTS) -w $(SOURCES)
endif

.PHONY: lint
lint:
	@echo ">  Running lint"
	$(GOLANGCI_LINT) run

################################################################################
## Migrations
################################################################################

.PHONY: migrate
migrate:
	@echo ">  Migrating: UP"
	$(GORUN) -tags migrate cmd/migrate/main.go -cmd up

.PHONY: drop
drop:
	@echo ">  Migrating: DOWN"
	$(GORUN) -tags migrate cmd/migrate/main.go -cmd down

################################################################################
## Docker
################################################################################

.PHONY: build-docker
build-docker:
	@echo ">  Building docker"
	sudo docker build -t $(DOCKER_IMAGE):latest -t $(DOCKER_IMAGE):$(VERSION) .

.PHONY: release-docker
release-docker: build-docker
	@echo ">  Releasing docker"
	sudo docker push $(DOCKER_IMAGE):latest
	sudo docker push $(DOCKER_IMAGE):$(VERSION)
