# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

# Copyright 2016 Ante Vojvodic <ante@atcorner.hr>

NAME     := asc
VERSION  ?= v1.0.0
GIT_REV  := $(shell git describe --always --abbrev --long)
LD_FLAGS := -ldflags="-X main.Version=$(VERSION) -X main.GitRevision=$(GIT_REV)"
GOOS     ?= linux
GOARCH   ?= amd64
CGO      ?= 0
BUILD    := GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=$(CGO) go build $(LD_FLAGS)
OUTDIR   := build
SSH_HOST ?= centos7-test

build: clean $(NAME)

$(NAME):
	@echo "==> Building"
	@(cd src && $(BUILD) -o ../$@)

.PHONY: clean
clean:
	@echo "==> Cleaning"
	@rm -f $(NAME)

.PHONY: run
run: $(NAME)
	scp $(NAME) "$(SSH_HOST):"
	scp config.json "$(SSH_HOST):"
	ssh "$(SSH_HOST)" "./$(NAME)"

.PHONY: release
release:
	mkdir -p $(OUTDIR)/opt/asc
	mkdir -p $(OUTDIR)/usr/lib/systemd/system
	cp $(NAME) $(OUTDIR)/opt/asc
	cp config.json $(OUTDIR)/opt/asc
	cp tools/$(NAME).service $(OUTDIR)/usr/lib/systemd/system

.PHONY: test
test:
	@echo "==> Running go vet"
	@(cd src && go vet)
	@echo "==> Running static check"
	@(cd src && staticcheck -f stylish .)
	@echo "==> Running linter"
	@(cd src && golangci-lint run .)
	@echo "==> Running tests"
	@(cd src && go test -v)
