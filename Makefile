Version := $(shell git describe --tags --dirty)
GitCommit := $(shell git rev-parse HEAD)
PKG_NAME := $(shell cat go.mod | grep module | sed -e "s/^module\s//")
LDFLAGS := "-s -w -X $(PKG_NAME)/internal.Version=$(Version) -X $(PKG_NAME)/internal.GitCommit=$(GitCommit)"

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o bin/qpp

.PHONY: dist
dist:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -o bin/qpp-amd64 -a -installsuffix cgo -ldflags $(LDFLAGS)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -mod=vendor -o bin/qpp-arm64 -a -installsuffix cgo -ldflags $(LDFLAGS)