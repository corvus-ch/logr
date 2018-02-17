MAKEFLAGS += --warn-undefined-variables
SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := test
.DELETE_ON_ERROR:
.SUFFIXES:

# ---------------------
# Environment variables
# ---------------------

GOPATH := $(shell go env GOPATH)

# -------
# Targets
# -------

.PHONY: install
install:
	go get -t -d ./...
	go get -u github.com/AlekSi/gocoverutil

.PHONY: test
test: c.out

c.out: buffered/cover.out log/cover.out logrus/cover.out std/cover.out writer_adapter/cover.out zap/cover.out zerolog/cover.out
	find . -mindepth 2 -name cover.out -exec gocoverutil -coverprofile=c.out merge {} +

%/cover.out:
	go test -coverprofile $@ -covermode atomic ./$(@D)

.PHONY: clean
clean:
	rm **/cover.out c.out

