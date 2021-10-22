GO=go
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
GO_LDFLAGS=-ldflags="-s -w -extldflags '-dynamic'"
TARGET_DIR=bin/

.PHONY: build clean

build:
	mkdir -p bin
	CGO_ENABLED=1 GOOS=$(GOOS) GO_ARCH=$(GOARCH) $(GO) build $(GO_LDFLAGS) -o $(TARGET_DIR) ./...

clean:
	rm -rf bin
