BINARY=certm
VERSION ?= $(shell git describe --tags --abbrev=0)-snapshot

all: build

clean:
	@rm -rf certm certm_*

build: build-darwin build-linux

build-darwin:
	GOOS=darwin GO111MODULE=on CGO_ENABLED=0 go build -a -tags 'netgo' -ldflags "-w -X github.com/ehazlett/certm/version.Version=$(VERSION)" -o build/Darwin/${BINARY} main.go

build-linux:
	GOOS=linux GO111MODULE=on CGO_ENABLED=0 go build -a -tags 'netgo' -ldflags "-w -X github.com/ehazlett/certm/version.Version=$(VERSION)" -o build/Linux/${BINARY} main.go

image: build
	@echo Building image $(TAG)
	@docker build -t $(REPO):$(TAG) .

release: build
	rm -rf release
	glu release

test:
	@bats test/integration/cli.bats test/integration/certs.bats

.PHONY: all build clean image test release
